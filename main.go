package main

import (
	"github.com/devasiajoseph/webapp/app"
	"github.com/devasiajoseph/webapp/core"
	"github.com/devasiajoseph/webapp/db/postgres"
	"github.com/devasiajoseph/webapp/uauth"
	"github.com/devasiajoseph/webapp/website"
	"github.com/gorilla/mux"
)

func main() {
	core.Start()
	postgres.InitDb()
	r := mux.NewRouter()
	//r.HandleFunc("/p/{page}", website.StaticPageHandler)
	//r.HandleFunc("/r/{page}", website.RawPageHandler)
	//r.HandleFunc("/sp/{page}", website.StandAlonePageHandler)
	//r.HandleFunc("/page/{page}", website.MDPageHandler)
	r.HandleFunc("/usw.js", website.ServiceWorker)
	//bitcoin.Start(r)
	uauth.Start(r)
	app.Start(r)
	//location.Start(r)
	//profile.Start(r)
	website.Start(r)

}
