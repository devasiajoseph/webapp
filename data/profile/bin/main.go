package main

import (
	"encoding/csv"
	"log"
	"os"
	"strings"

	"github.com/devasiajoseph/webapp/core"
	"github.com/devasiajoseph/webapp/data/location"
	"github.com/devasiajoseph/webapp/data/profile"
	"github.com/devasiajoseph/webapp/db/postgres"
)

func cleanName(name string) string {
	return strings.ReplaceAll(name, " & family", "")
}

func main() {
	core.Start()
	postgres.InitDb()
	cf, err := os.Open("billionaires.csv")
	if err != nil {
		log.Println(err)
		return
	}

	defer cf.Close()
	reader := csv.NewReader(cf)
	lines, err := reader.ReadAll()

	if err != nil {
		log.Println(err)
	}
	for _, each := range lines {
		c := location.Country{CountryName: each[1]}
		err := c.NamedQuery()
		if err != nil {
			log.Println(c.CountryName)
			log.Println(err)
		}

		if c.CountryID == 0 {
			log.Println("Country not found")
			log.Println(c.CountryName)
		}

		pro := profile.Profile{
			FullName:  cleanName(each[0]),
			About:     each[2] + "," + each[3],
			CountryID: c.CountryID,
			Slug:      profile.Slugify(each[0])}

		err = pro.Create()
		if err != nil {
			log.Println(err)
			log.Println(pro.FullName)
		}

	}

}
