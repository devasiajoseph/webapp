package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/devasiajoseph/webapp/core"
	"github.com/devasiajoseph/webapp/data/location"
	"github.com/devasiajoseph/webapp/db/postgres"
)

func main() {
	core.Start()
	postgres.InitDb()
	cf, err := os.Open("countries.csv")
	if err != nil {
		log.Println(err)
		return
	}
	defer cf.Close()
	reader := csv.NewReader(cf)
	lines, err := reader.ReadAll()
	if err != nil {
		log.Println(err)
		return
	}

	for _, each := range lines {
		c := location.Country{
			CountryName: each[1],
			CountryCode: each[0],
		}
		err := c.Create()
		if err != nil {
			log.Println(err)
		}
	}
}
