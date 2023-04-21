package news

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/devasiajoseph/webapp/db/postgres"
	"github.com/devasiajoseph/webapp/libs/format"
)

type Object struct {
	NewsID     int    `json:"news_id" db:"news_id"`
	DomainID   int    `json:"domain_id" db:"domain_id"`
	Title      string `json:"title" db:"title"`
	MinContent string `json:"min_content" db:"min_content"`
	Content    string `json:"content" db:"content"`
	CoverPhoto string `json:"cover_photo" db:"cover_photo"`
	Thumbnail  string `json:"thumbnail" db:"thumbnail"`
}

var sqlCreate = "insert into news (domain_id,title,min_content,content) values (domain_id,title,min_content,content) returning news_id;"
var sqlDelete = "delete from news where news_id=$1;"
var sqlUpdate = "update news set domain_id=:domain_id,title=:title,min_content=:min_content,content=:content where news_id=:news_id;"
var sqlSlug = "update news set slug=$1 where news_id=$2;"
var sqlPublish = "update news set published=$1 where news_id=$2;"

func (obj *Object) UpdateSlug() error {
	if obj.NewsID == 0 {
		return errors.New("news does not exist to add slug")
	}
	slug := format.Slugify(obj.Title + strconv.Itoa(obj.NewsID))
	db := postgres.Db
	_, err := db.Exec(sqlSlug, slug, obj.NewsID)
	if err != nil {
		log.Println(err)
		log.Println("error updating news slug")
	}
	return err
}

func (obj *Object) Create() error {
	db := postgres.Db
	rows, err := db.NamedQuery(sqlCreate, obj)

	if err != nil {
		log.Println(err)
		fmt.Println(obj)
		log.Println("error creating news")
		return err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&obj.NewsID)
	}

	return obj.UpdateSlug()
}

func (obj *Object) Update() error {
	db := postgres.Db
	_, err := db.NamedExec(sqlUpdate, obj)
	if err != nil {
		log.Println(err)
		log.Println("error updating news")
	}
	return err
}

func (obj *Object) Save() error {
	if obj.NewsID == 0 {
		return obj.Create()
	}

	return obj.Update()
}

func (obj *Object) Delete() error {
	return nil
}
