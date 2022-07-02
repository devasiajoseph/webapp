package file

import (
	//"github.com/gorilla/mux"

	"time"

	"github.com/devasiajoseph/webapp/core"
)

type MiscFile struct {
	Id               int `storm:"id,increment" json:"Id" `
	OriginalFileName string
	Filename         string    `storm:"index" json:"Filename"`
	UploadedTime     time.Time `json:"UploadedTime"`
	FileSize         int64
}

var fileUploadPath = core.AbsolutePath("static/uploads/files")
