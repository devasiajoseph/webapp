package file

import (
	"log"

	"github.com/devasiajoseph/webapp/db/postgres"
)

var sqlInsertImage = "insert into image (file_name,path,original_image,reverse_id) values (:file_name,:path,:original_image,:reverse_id) returning image_id;"

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

func GetBlankImage() (ImageData, error) {
	db := postgres.Db
	sqlGetBlank := "select image_id from image where tag='_blank';"
	var imgs []ImageData
	var imgData ImageData
	err := db.Select(&imgs, sqlGetBlank)
	if err != nil {
		log.Println("error getting blank image")
	}
	if len(imgs) > 0 {
		imgData = imgs[0]
	}
	return imgData, err
}
