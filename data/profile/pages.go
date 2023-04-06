package profile

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/devasiajoseph/webapp/core"
	"github.com/devasiajoseph/webapp/website"
	"github.com/gorilla/mux"
)

func ProfilePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("here")
	paths := website.Paths{DirPath: core.BinPath, StaticUrl: core.StaticUrl}
	pagePath := website.PagePath(paths.DirPath, "profile.html")
	tmpl, err := template.ParseFiles(website.PagePath(paths.DirPath, "base.html"), pagePath)
	if err != nil {
		fmt.Println(err)
	}

	//itemCount := printMessage.Sprintf("%d", GetItemCount())
	obj := Object{}
	err = tmpl.Execute(w, obj)

	if err != nil {
		log.Println(err)
		log.Println("Unable to process template file")
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
}

func AddPageRoutes(r *mux.Router) {
	fmt.Println("adding page routes")
	r.HandleFunc(apiObj+"/{slug}", ProfilePage).Methods("GET")

}
