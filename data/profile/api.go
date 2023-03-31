package profile

import (
	"log"
	"net/http"

	"github.com/devasiajoseph/webapp/api"
	"github.com/devasiajoseph/webapp/file"
	"github.com/devasiajoseph/webapp/uauth"
	"github.com/devasiajoseph/webapp/validator"
	"github.com/gorilla/mux"
)

var apiObj = "profile"

func (obj *Object) hasAuth(w http.ResponseWriter, r *http.Request) bool {
	obj.ProfileID = api.ObjID(r, "profile_id")
	if obj.ProfileID == 0 {
		return true
	}

	ua, err := uauth.GetAuthenticatedUser(r)
	if err != nil {
		log.Println(err)
		api.ServerError(w)
		return false
	}

	obj.UserAccount = ua

	auth, err := obj.IsManager(ua)

	if err != nil {
		log.Println(err)
		return false
	}

	if !auth {
		api.AuthError(w)
		return false
	}

	err = obj.Get()
	if err != nil {
		log.Println("error while getting object in hasAuth")
		return false
	}

	return true
}

func ValidateSave(r *http.Request) validator.ValidatorResponse {
	res := validator.InitResponse()
	validator.RequiredStringValidation(r.FormValue("full_name"), "full_name", &res)
	validator.RequiredStringValidation(r.FormValue("about"), "about", &res)
	return res
}

func saveApi(w http.ResponseWriter, r *http.Request) {
	vRes := ValidateSave(r)
	if !vRes.Valid {
		api.ValidationError(w, vRes)
		return
	}
	obj := Object{}
	if !obj.hasAuth(w, r) {
		return
	}
	obj.FullName = r.FormValue("full_name")
	obj.About = r.FormValue("about")
	obj.Instagram = r.FormValue("instagram")
	obj.Linkedin = r.FormValue("linkedin")
	obj.Facebook = r.FormValue("facebook")
	obj.Twitter = r.FormValue("twitter")
	obj.Youtube = r.FormValue("youtube")
	obj.Tiktok = r.FormValue("tiktok")
	obj.CountryID = api.PostInt(r, "country_id")
	obj.ProfileID = api.PostInt(r, "profile_id")
	err := obj.Save()
	if err != nil {
		api.ServerError(w)
	}
	api.ObjectResponse(w, obj)
}

func listApi(w http.ResponseWriter, r *http.Request) {
	ua, err := uauth.GetAuthenticatedUser(r)
	if err != nil {
		log.Println(err)
		api.ServerError(w)
		return
	}
	ol := ObjectList{Page: api.QueryInt(r, "page"),
		Limit:         50,
		UserAccountID: ua.UserAccountID}
	err = ol.Fetch()

	if err != nil {
		api.ServerError(w)
		return
	}

	api.ObjectResponse(w, ol)
}

func deleteApi(w http.ResponseWriter, r *http.Request) {

}

func getApi(w http.ResponseWriter, r *http.Request) {
	objID := api.ObjID(r, "profile_id")
	if objID == 0 {
		api.AuthError(w)
		return
	}

	obj := Object{ProfileID: objID}
	err := obj.Get()
	if err != nil {
		log.Println(err)
		api.ServerError(w)
		return
	}
	if obj.ProfileID == 0 {
		api.ObjectNotFound(w)
		return
	}
	api.ObjectResponse(w, obj)
}

func uploadDP(w http.ResponseWriter, r *http.Request) {
	obj := Object{}
	if !obj.hasAuth(w, r) {
		api.AuthError(w)
		return
	}
	imgData := file.ImageData{MaxUploadSize: 15, Width: 200}
	err := imgData.ProcessUpload(w, r, "profile_pic")
	if err != nil {
		api.ServerError(w)
	}
	err = obj.AddProfilePic(imgData)
	if err != nil {
		api.ServerError(w)
	}
	api.ObjectResponse(w, imgData)
}

func AddRoutes(r *mux.Router) {
	r.HandleFunc("/api/"+apiObj, saveApi).Methods("POST")
	r.HandleFunc("/api/"+apiObj, listApi).Methods("GET")
	r.HandleFunc("/api/"+apiObj+"/{profile_id}", getApi).Methods("GET")
	r.HandleFunc("/api/"+apiObj+"/upload-dp/{profile_id}", uploadDP).Methods("POST")

}

// Start initializes bitcoin based functions
func Start(r *mux.Router) {
	log.Println("Starting " + apiObj + " api")
	AddRoutes(r)
}
