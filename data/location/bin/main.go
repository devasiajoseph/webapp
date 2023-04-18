package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/devasiajoseph/webapp/core"
	"github.com/devasiajoseph/webapp/data/location"
	"github.com/devasiajoseph/webapp/db/postgres"
)

var sqlInsert = "insert into city (city_id,city_name,city_type,location,longitude,latitude,district_id,district_centre) " +
	"values (:city_id,:city_name,:city_type,ST_SetSRID(ST_MakePoint(:longitude,:latitude),4326),:longitude,:latitude,:district_id,:district_centre)"

func ParseInt(num string) int {
	i, err := strconv.Atoi(num)
	if err != nil {
		return 0
	}
	return i
}

func cleanTitle(v string) string {
	return strings.Title(strings.ToLower(strings.TrimSpace(v)))
}

func districtCentre(val string) bool {
	if strings.ToLower(strings.TrimSpace(val)) == "true" {
		return true
	}
	return false
}
func loadCountries() {
	cf, err := os.Open("files/countries.csv")
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
			log.Println(c)
		}
	}
}

func loadTowns() {
	csvfile, err := os.Open("files/cities.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	r := csv.NewReader(csvfile)
	towns := []location.City{}
	for {
		// Read each record from csv

		record, err := r.Read()
		if err == io.EOF {
			break

		}
		if err != nil {
			log.Fatal(err)
		}

		c := location.City{}

		c.DistrictID = ParseInt(record[0])
		c.CityName = cleanTitle(record[2])
		c.CityType = cleanTitle(record[3])
		c.Latitude = record[4]
		c.Longitude = record[5]
		c.CityID = ParseInt(record[6])
		c.DistrictCentre = districtCentre(record[7])

		towns = append(towns, c)

	}

	db := postgres.Db
	for _, each := range towns {
		_, err := db.NamedExec(sqlInsert, each)

		if err != nil {

			log.Println(err)
			log.Println(each)
		}
	}
}

func main() {
	core.Start()
	postgres.InitDb()
	//loadCountries()
	loadTowns()
}
