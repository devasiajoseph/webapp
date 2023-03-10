package profile

import (
	"log"
	"regexp"
	"strings"

	"github.com/devasiajoseph/webapp/db/postgres"
)

type Profile struct {
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

func (p *Profile) Create() error {
	db := postgres.Db
	_, err := db.NamedExec(sqlCreate, p)
	if err != nil {
		log.Println(err)
		log.Println("Error creating new profile")
	}
	return err
}
