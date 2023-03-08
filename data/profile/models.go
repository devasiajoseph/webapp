package profile

import (
	"log"

	"github.com/devasiajoseph/webapp/db/postgres"
)

type Profile struct {
	ProfileID  int    `json:"profile_id" db:"profile_id"`
	FullName   string `json:"full_name" db:"full_name"`
	About      string `json:"about" db:"about"`
	ProfilePic string `json:"profile_pic" db:"profile_pic"`
	CountryID  int    `json:"country_id" db:"country_id"`
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
