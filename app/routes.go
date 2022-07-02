package app

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/devasiajoseph/webapp/core"
	"github.com/devasiajoseph/webapp/website"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
)

type AppPageData struct {
	Id             int                `storm:"id,increment" json:"Id" `
	Slug           string             `storm:"index,unique" json:"Slug"`
	Title          string             `json:"Title"`
	FileName       string             `storm:"index,unique" json:"FileName"`
	BaseFileName   string             `json:"BaseFileName"`
	HtmlContent    template.HTML      `json: "HtmlContent"`
	Markdown       string             `json:"Markdown"`
	PageType       string             `json:"PageType"`
	Meta           []website.MetaData `json:"Meta"`
	LastEdited     time.Time          `json:"LastEdited"`
	UserId         int                `storm:"index" json:"UserId"`
	Author         string             `storm:"index" json:"Author"`
	BasePageId     int                `storm:"index" json:"BasePageId"`
	TemplatePageId int                `storm:"index" json:"TemplatePageId"`
	Csrf           string
	StaticUrl      string
	KeyValue       string
	UAuthLoggedIn  bool
	LoggedInUser   string
	Message        string
	PageAuth       bool
	ItemCount      string
	RecapSiteKey   string
}

func home(w http.ResponseWriter, r *http.Request) {
	loggedIn := false
	paths := website.Paths{DirPath: core.BinPath, StaticUrl: core.StaticUrl}
	pagePath := website.PagePath(paths.DirPath, "home.html")
	tmpl, err := template.ParseFiles(website.PagePath(paths.DirPath, "base.html"), pagePath)
	if err != nil {
		fmt.Println(err)
	}
	authUser, err := website.GetAuthenticatedUser(r)
	if err == nil {
		loggedIn = true
	}

	//itemCount := printMessage.Sprintf("%d", GetItemCount())
	p := AppPageData{
		//Title: "hello",
		Csrf:          csrf.Token(r),
		StaticUrl:     paths.StaticUrl,
		UAuthLoggedIn: loggedIn,
		LoggedInUser:  authUser.FullName,
		RecapSiteKey:  core.Config.RecapSiteKey,
		//ItemCount:     itemCount,
	}
	err = tmpl.Execute(w, p)

	if err != nil {
		log.Println(err)
		log.Println("Unable to process template file")
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}

}

func AddRoutes(r *mux.Router) {
	r.HandleFunc("/", home).Methods("GET")
}

func Start(r *mux.Router) {
	AddRoutes(r)
}
