package file

import (
	"log"

	"github.com/devasiajoseph/webapp/db/postgres"
)

var sqlInsertImage = "insert into image (file_name,path,original_image) values (:file_name,:path,:original_image) returning image_id;"

func (imgD *ImageData) Save() error {
	db := postgres.Db
	rows, err := db.NamedQuery(sqlInsertImage, imgD)
	if err != nil {
		log.Println(err)
		log.Println("Error creating new profile")
	}

	if rows.Next() {
		rows.Scan(&imgD.ImageID)
	}
	return err
}

func (imgD *ImageData) Delete() error {
	db := postgres.Db
	_, err := db.Exec("delete from image where image_id=$1", imgD.ImageID)
	if err != nil {
		log.Println("error deleting image")
	}
	return err
}
