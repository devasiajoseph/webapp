package file

import (
	"database/sql"
	"log"

	"github.com/devasiajoseph/webapp/db/postgres"
)

var sqlInsertImage = "insert into image (file_name,src,original_image) " +
	"values (:file_name,:src,:original_image) returning image_id;"

func (imgD *ImageData) Save() error {
	db := postgres.Db
	rows, err := db.NamedQuery(sqlInsertImage, imgD)
	if err != nil {
		log.Println(err)
		log.Println("Error creating new image")
	}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&imgD.ImageID)
	}
	return err
}

func (imgd *ImageData) Delete() error {
	db := postgres.Db
	sqlDeleteImage := "delete from image where image_id=$1;"
	_, err := db.Exec(sqlDeleteImage, imgd.ImageID)

	if err != nil {
		log.Println("error deleting image")
		return err
	}
	err = DeleteFile(ImageuploadPath + imgd.Filename)
	if err != nil {
		log.Println("error deleting file from path")
	}

	return err
}

func GetImage(imageID int) (ImageData, error) {
	db := postgres.Db
	var img ImageData
	err := db.Get(&img, "select * from image where image_id=$1")
	if err != nil {
		log.Println("error getiing image with image id")
	}

	return img, err
}

func DeleteImage(imageID int) error {
	if imageID == 0 {
		return nil
	}
	img, err := GetImage(imageID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		log.Println("error getting image to delete")
		return err
	}
	return img.Delete()
}
