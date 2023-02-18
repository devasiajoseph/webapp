package app

import (
	"log"

	"github.com/devasiajoseph/webapp/db/postgres"
)

var sqlGetPage = "select page_file, base_page_file from page where page_slug=$1"

type Page struct {
	PageID       int    `db:"page_id" json:"page_id"`
	PageSlug     string `db:"page_slug" json:"page_slug"`
	PageFile     string `db:"page_file" json:"page_file"`
	BasePageFile string `db:"base_page_file" json:"base_page_file"`
}

func (apd *AppPageData) GetPage() error {
	db := postgres.Db

	err := db.Get(apd, sqlGetPage, apd.PageSlug)
	if err != nil {
		log.Println("Error getting page")
		log.Println(err)
	}

	return err
}
