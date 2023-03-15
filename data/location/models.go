package location

import (
	"database/sql"
	"log"

	"github.com/devasiajoseph/webapp/db/postgres"
)

// Country data
type Country struct {
	CountryID   int    `db:"country_id" json:"country_id"`
	CountryName string `db:"country_name" json:"country_name"`
	CountryCode string `db:"country_code" json:"country_code"`
}

func (c *Country) Create() error {
	db := postgres.Db
	_, err := db.NamedExec("insert into country (country_name,country_code) values (:country_name,:country_code)", c)
	if err != nil {
		log.Println("Error saving country")
	}
	return err
}

func (c *Country) NamedQuery() error {
	db := postgres.Db
	err := db.Get(c, "select country_name,country_id from country where country_name=$1", c.CountryName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		log.Println("Error getting country by name")
		return err
	}
	return err
}

func GetCountryList() ([]Country, error) {
	db := postgres.Db
	cl := []Country{}
	err := db.Select(&cl, "select country_id,country_name,country_code from country;")
	if err != nil {
		log.Println("Error fetching country list")
	}
	return cl, err
}
