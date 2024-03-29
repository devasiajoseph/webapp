package profile

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
	PageId         int    `db:"page_id" json:"page_id" `
	PageSlug       string `db:"page_slug" json:"page_slug"`
	Title          string `json:"title"`
	PageFile       string `db:"page_file" json:"page_file"`
	BasePageFile   string `db:"base_page_file" json:"base_page_file"`
	HtmlContent    template.HTML
	Markdown       string
	PageType       string
	Meta           []website.MetaData
	LastEdited     time.Time
	UserId         int
	Author         string
	BasePageId     int
	TemplatePageId int
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

func ProfilePage(w http.ResponseWriter, r *http.Request) {
	loggedIn := false
	paths := website.Paths{DirPath: core.BinPath, StaticUrl: core.StaticUrl}
	pagePath := website.PagePath(paths.DirPath, "profile.html")
	tmpl, err := template.ParseFiles(website.PagePath(paths.DirPath, "base.html"), pagePath)
	if err != nil {
		fmt.Println(err)
	}

	authUser, err := website.GetAuthenticatedUser(r)
	if err == nil {
		loggedIn = true
	}

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

func AddPageRoutes(r *mux.Router) {
	r.HandleFunc("/"+apiObj+"/{slug}", ProfilePage).Methods("GET")
}
