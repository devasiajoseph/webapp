package profile

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/devasiajoseph/webapp/cpmath"
	"github.com/devasiajoseph/webapp/db/postgres"
	"github.com/devasiajoseph/webapp/file"
	"github.com/devasiajoseph/webapp/uauth"
)

// ObjectList has list of companies
type ObjectList struct {
	Data          []ObjectListData `json:"data"`
	Total         int              `json:"total"`
	Page          int              `json:"page"`
	Limit         int              `json:"limit"`
	Offset        int              `json:"offset"`
	UserAccountID int              `json:"-"`
}

type ObjectListData struct {
	ProfileID  int    `json:"profile_id" db:"profile_id"`
	FullName   string `json:"full_name" db:"full_name"`
	ProfilePic string `json:"profile_pic" db:"profile_pic"`
}

type Object struct {
	ProfileID   int            `json:"profile_id" db:"profile_id"`
	FullName    string         `json:"full_name" db:"full_name"`
	Designation string         `json:"designation" db:"designation"`
	About       string         `json:"about" db:"about"`
	ProfilePic  string         `json:"profile_pic" db:"profile_pic"`
	ImageID     int            `json:"image_id" db:"image_id"`
	Instagram   string         `json:"instagram" db:"instagram"`
	Linkedin    string         `json:"linkedin" db:"linkedin"`
	Facebook    string         `json:"facebook" db:"facebook"`
	Twitter     string         `json:"twitter" db:"twitter"`
	Youtube     string         `json:"youtube" db:"youtube"`
	Tiktok      string         `json:"tiktok" db:"tiktok"`
	CountryID   int            `json:"country_id" db:"country_id"`
	Slug        string         `json:"slug" db:"slug"`
	UserAccount uauth.AuthUser `json:"-"`
}

type ProfileManager struct {
	ProfileManagerID int `json:"profile_manager_id" db:"profile_manager_id"`
	UserAccountID    int `json:"user_account_id" db:"user_account_id"`
	ProfileID        int `json:"profile_id" db:"profile_id"`
}

func Slugify(str string) string {
	// Convert to lowercase
	str = strings.ToLower(str)

	// Remove all non-word characters and replace with "-"
	reg, err := regexp.Compile(`[\W]+`)
	if err != nil {
		panic(err)
	}
	str = reg.ReplaceAllString(str, "-")

	// Remove any leading or trailing "-"
	str = strings.Trim(str, "-")

	return str
}

var sqlCreate = "insert into profile (full_name,about,profile_pic,image_id,instagram,linkedin,facebook,twitter,youtube,tiktok,country_id,slug) " +
	"values (:full_name,:about,:profile_pic,:image_id,:instagram,:linkedin,:facebook,:twitter,:youtube,:tiktok,:country_id,:slug) returning profile_id;"

func (obj *Object) Create() error {
	db := postgres.Db
	rows, err := db.NamedQuery(sqlCreate, obj)

	if err != nil {
		log.Println(err)
		fmt.Println(obj)
		log.Println("Error creating new profile")
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&obj.ProfileID)
	}

	return err
}

func (obj *Object) Save() error {
	obj.Slug = Slugify(obj.FullName)
	if obj.ProfileID == 0 {
		dp := file.GetBlankImage()
		obj.ImageID = dp.ImageID
		obj.ProfilePic = dp.Src
		err := obj.Create()
		if err != nil {
			return err
		}
		err = obj.AddManager(obj.UserAccount)
		if err != nil {
			return err
		}
	}

	return obj.Update()
}

var sqlGet = "select * from profile where profile_id = $1"

func (obj *Object) Get() error {
	db := postgres.Db
	err := db.Get(obj, sqlGet, obj.ProfileID)
	if err != nil {
		if err == sql.ErrNoRows {
			obj.ProfileID = 0
			return nil
		}
		log.Println("Error getting profile")
	}

	return err
}

var sqlUpdate = "update profile set " +
	" full_name=:full_name,about=:about,instagram=:instagram,linkedin=:linkedin," +
	"facebook=:facebook,twitter=:twitter,youtube=:youtube,tiktok=:tiktok,country_id=:country_id," +
	"slug=:slug where profile_id=:profile_id;"

func (obj *Object) Update() error {
	db := postgres.Db
	_, err := db.NamedExec(sqlUpdate, obj)
	if err != nil {
		log.Println(err)
		log.Println("Error updating profile")
	}
	return err
}

var sqlManager = "select profile_id, user_account_id from profile_manager where profile_id=$1 and user_account_id=$2;"

func (obj *Object) IsManager(ua uauth.AuthUser) (bool, error) {
	if !ua.Active {
		return false, nil
	}
	db := postgres.Db
	pm := ProfileManager{}
	err := db.Get(&pm, sqlManager, obj.ProfileID, ua.UserAccountID)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		log.Println("Error while fetching profile manager")
		return false, err
	}
	return true, nil
}

func (obj *Object) AddManager(ua uauth.AuthUser) error {
	if !ua.Active {
		return errors.New("user not active")
	}
	db := postgres.Db
	sqlInsertManager := "insert into profile_manager (profile_id,user_account_id) values ($1,$2);"
	_, err := db.Exec(sqlInsertManager, obj.ProfileID, ua.UserAccountID)

	if err != nil {
		log.Println(err)
	}
	return err
}

func (obj *Object) AddBlankProfilePic() error {
	sqlAddImage := "update profile set profile_pic=$1,image_id=$2 where profile_id=$3;"
	db := postgres.Db
	_, err := db.Exec(sqlAddImage, file.BlankProfileImage, 0, obj.ProfileID)
	if err != nil {
		log.Println("error adding blank profile pic")

	}
	return err
}

func (obj *Object) DeleteProfilePic() error {
	if obj.ImageID == 0 {
		return nil
	}
	err := file.DeleteImage(obj.ImageID)
	if err != nil {
		log.Println("error deleting profile pic")
	}
	return err
}

func (obj *Object) AddProfilePic(imgData file.ImageData) error {
	err := obj.DeleteProfilePic()
	if err != nil {
		log.Println(err)
		log.Println("error while deleting profile pic for adding new profile pic")
	}
	sqlAddImage := "update profile set profile_pic=$1,image_id=$2 where profile_id=$3;"

	db := postgres.Db
	_, err = db.Exec(sqlAddImage, imgData.Src, imgData.ImageID, obj.ProfileID)
	if err != nil {
		log.Println("error adding profile pic")
		return err
	}
	return nil
}

func (ol *ObjectList) Fetch() error {
	ol.Offset = cpmath.Offset(ol.Page, ol.Limit)
	sqlList := "select profile.profile_id,profile.full_name,profile.profile_pic from profile left join profile_manager " +
		"on profile.profile_id=profile_manager.profile_id where " +
		"profile_manager.user_account_id=$1 limit $2 offset $3;"
	sqlTotal := "select count(*) from profile left join profile_manager " +
		"on profile.profile_id=profile_manager.profile_id where " +
		"profile_manager.user_account_id=$1;"
	db := postgres.Db

	err := db.Get(&ol.Total, sqlTotal, ol.UserAccountID)
	if err != nil {
		log.Println("error getting total in profile list")
		log.Println(err)
		return err
	}
	err = db.Select(&ol.Data, sqlList, ol.UserAccountID, ol.Limit, ol.Offset)

	if err != nil {
		log.Println("error fecthing profile list")
		log.Println(err)
	}
	return err
}

func (ol *ObjectList) Search(q string) error {
	ol.Offset = cpmath.Offset(ol.Page, ol.Limit)
	sq := "%" + q + "%"
	sqlList := "select profile.profile_id,profile.full_name,profile.profile_pic from profile left join profile_manager " +
		"on profile.profile_id=profile_manager.profile_id where " +
		" profile.full_name ilike $1 and " +
		"profile_manager.user_account_id=$2 limit $3 offset $4;"
	sqlTotal := "select count(*) from profile left join profile_manager " +
		"on profile.profile_id=profile_manager.profile_id where " +
		" profile.full_name ilike $1 and " +
		"profile_manager.user_account_id=$2;"
	db := postgres.Db

	err := db.Get(&ol.Total, sqlTotal, sq, ol.UserAccountID)
	if err != nil {
		log.Println("error getting total in profile list")
		log.Println(err)
		return err
	}
	err = db.Select(&ol.Data, sqlList, sq, ol.UserAccountID, ol.Limit, ol.Offset)

	if err != nil {
		log.Println("error fecthing profile list")
		log.Println(err)
	}
	return err
}
