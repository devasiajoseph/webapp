/*
 * Copyright (C) Centipair Technologies Private Limited - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written By Devasia Joseph <devasiajoseph@centipair.com>, January 2019
 */

package website

import (
	"net/http"

	"github.com/NYTimes/gziphandler"
	"github.com/devasiajoseph/webapp/core"
	"github.com/gorilla/mux"
)

func AddRoutes(r *mux.Router) {

	if core.Config.Compress {
		ccljs := gziphandler.GzipHandler(http.FileServer(http.Dir(adminUIPath)))
		cljs := http.StripPrefix("/admin-ui/", ccljs)
		r.PathPrefix("/admin-ui/").Handler(cljs)

		cas := gziphandler.GzipHandler(http.FileServer(http.Dir(adminStaticPath)))
		astatic := http.StripPrefix("/admin-static/", cas)
		r.PathPrefix("/admin-static/").Handler(astatic)

		fs := gziphandler.GzipHandler(http.FileServer(http.Dir(staticPath)))
		s := http.StripPrefix("/static/", fs)
		r.PathPrefix("/static/").Handler(s)

		ffcljs := gziphandler.GzipHandler(http.FileServer(http.Dir(userUIPath)))
		fcljs := http.StripPrefix("/user-ui/", ffcljs)
		r.PathPrefix("/user-ui/").Handler(fcljs)

		ffcljsp := gziphandler.GzipHandler(http.FileServer(http.Dir(userUIPath)))
		fcljsp := http.StripPrefix("/p/user-ui/", ffcljsp)
		r.PathPrefix("/p/user-ui/").Handler(fcljsp)

		ffcljsapp := gziphandler.GzipHandler(http.FileServer(http.Dir(userUIPath)))
		fcljsapp := http.StripPrefix("/app/user-ui/", ffcljsapp)
		r.PathPrefix("/app/user-ui/").Handler(fcljsapp)

	} else {
		cljs := http.StripPrefix("/admin-ui/", http.FileServer(http.Dir(adminUIPath)))
		r.PathPrefix("/admin-ui/").Handler(cljs)

		astatic := http.StripPrefix("/admin-static/", http.FileServer(http.Dir(adminStaticPath)))
		r.PathPrefix("/admin-static/").Handler(astatic)

		s := http.StripPrefix("/static/", http.FileServer(http.Dir(staticPath)))
		r.PathPrefix("/static/").Handler(s)

		fcljs := http.StripPrefix("/user-ui/", http.FileServer(http.Dir(userUIPath)))
		r.PathPrefix("/user-ui/").Handler(fcljs)

		figcljs := http.StripPrefix("/cljs-out/", http.FileServer(http.Dir(fuserUIPath)))
		r.PathPrefix("/cljs-out/").Handler(figcljs)

		fcljsp := http.StripPrefix("/p/user-ui/", http.FileServer(http.Dir(userUIPath)))
		r.PathPrefix("/p/user-ui/").Handler(fcljsp)

		fcljsapp := http.StripPrefix("/app/user-ui/", http.FileServer(http.Dir(userUIPath)))
		r.PathPrefix("/app/user-ui/").Handler(fcljsapp)

	}

}
