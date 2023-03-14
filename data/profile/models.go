package profile

import (
	"database/sql"
	"log"
	"regexp"
	"strings"

	"github.com/devasiajoseph/webapp/db/postgres"
	"github.com/devasiajoseph/webapp/uauth"
)

type Object struct {
	ProfileID   int    `json:"profile_id" db:"profile_id"`
	FullName    string `json:"full_name" db:"full_name"`
	Designation string `json:"designation" db:"designation"`
	About       string `json:"about" db:"about"`
	ProfilePic  string `json:"profile_pic" db:"profile_pic"`
	Instagram   string `json:"instagram" db:"instagram"`
	Facebook    string `json:"facebook" db:"facebook"`
	Twitter     string `json:"twitter" db:"twitter"`
	Youtube     string `json:"youtube" db:"youtube"`
	Tiktok      string `json:"tiktok" db:"tiktok"`
	Slug        string `json:"slug" db:"slug"`
	CountryID   int    `json:"country_id" db:"country_id"`
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

var sqlCreate = "insert into profile (full_name,about,profile_pic,country_id) " +
	"values (:full_name,:about,:profile_pic,:country_id);"

func (obj *Object) Create() error {
	db := postgres.Db
	_, err := db.NamedExec(sqlCreate, obj)
	if err != nil {
		log.Println(err)
		log.Println("Error creating new profile")
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
