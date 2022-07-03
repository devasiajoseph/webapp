package website

import (
	//"context"

	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/devasiajoseph/webapp/core"
	"github.com/devasiajoseph/webapp/file"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
)

var adminUIPath = core.AbsolutePath("cljs/admin-ui")
var userUIPath = core.AbsolutePath("cljs/user-ui")
var fuserUIPath = core.AbsolutePath("cljs/target/public/cljs-out")
var adminStaticPath = core.AbsolutePath("cljs/admin/static")
var staticPath = core.AbsolutePath("static")

func ServerPort() string {
	port := "8080"
	if core.Config.Port > 0 {
		port = strconv.Itoa(core.Config.Port)
	}

	if len(os.Args) > 1 {
		if _, err := strconv.Atoi(os.Args[1]); err != nil {
			log.Fatal("Inavlid port")
		}
		port = os.Args[1]
	}
	port = ":" + port
	return port
}

func Start(r *mux.Router) {

	port := ServerPort()
	log.Println("Starting webserver at port" + port)
	AddRoutes(r)

	protectionMiddleware := func(handler http.Handler) http.Handler {
		protectionFn := csrf.Protect(
			[]byte(core.SKey),
			csrf.Secure(core.Secure),
			csrf.Path("/"),
		)

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Use some kind of condition here to see if the router should use
			// the CSRF protection. For the sake of this example, we'll check
			// the path prefix.
			if !strings.HasPrefix(r.URL.Path, "/webhooks") {
				protectionFn(handler).ServeHTTP(w, r)
				return
			}

			handler.ServeHTTP(w, r)
		})
	}
	err := http.ListenAndServe(port,
		protectionMiddleware(r),
	)
	/*err := http.ListenAndServe(port,
	csrf.Protect(
		[]byte(core.SKey),
		csrf.Secure(core.Secure), // Pass it *to* this constructor
	)(r))*/
	if err != nil {
		log.Println(err)
	}
}

func ServiceWorker(w http.ResponseWriter, r *http.Request) {
	dat, err := file.ReadBytes(core.AbsolutePath("static/js/sw.js"))

	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/javascript")
	w.Write(dat)
}
