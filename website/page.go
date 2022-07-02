/*
 * Copyright (C) Centipair Technologies Private Limited - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written By Devasia Joseph <devasiajoseph@centipair.com>, January 2019
 */

package website

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/devasiajoseph/webapp/core"
	"github.com/devasiajoseph/webapp/db/postgres"

	//"gopkg.in/russross/blackfriday.v2"
	"html/template"
	//"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/csrf"
	//"os"
	//	"strings"
)

type Paths struct {
	DirPath   string
	StaticUrl string
}

type MetaData struct {
	Id      int    `storm:"id,increment" json:"id" `
	PageId  int    `storm:"index" json:"PageId"`
	Name    string `storm:"index" json:"Name"`
	Content string `json:"Content"`
}

type PageData struct {
	Id             int           `storm:"id,increment" json:"Id" `
	Slug           string        `storm:"index,unique" json:"Slug"`
	Title          string        `json:"Title"`
	HtmlContent    template.HTML `json: "HtmlContent"`
	Markdown       string        `json:"Markdown"`
	PageType       string        `json:"PageType"`
	Meta           []MetaData    `json:"Meta"`
	LastEdited     time.Time     `json:"LastEdited"`
	UserId         int           `storm:"index" json:"UserId"`
	Author         string        `storm:"index" json:"Author"`
	BasePageId     int           `storm:"index" json:"BasePageId"`
	TemplatePageId int           `storm:"index" json:"TemplatePageId"`
	Csrf           string
	StaticUrl      string
	KeyValue       string
	UAuthLoggedIn  bool
	LoggedInUser   string
	RecapSiteKey   string
	PageFile       string
	BasePageFile   string
}

func HandlePageError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

var tmpDir = core.AbsolutePath("html/")

//AuthUser current authenticated user
type AuthUser struct {
	UserAccountID int       `db:"user_account_id" json:"UserAccountID"`
	FullName      string    `db:"full_name" json:"FullName"`
	Email         string    `db:"email" json:"Email"`
	Phone         string    `db:"phone" json:"Phone"`
	Active        bool      `db:"active" json:"Active"`
	CreatedOn     time.Time `db:"created_on" json:"CreatedOn"`
	LastLogin     time.Time `db:"last_login" json:"LastLogin"`
	SessionExpiry time.Time `db:"session_expiry" json:"SessionExpiry"`
}

var sqlFetchAuthUser = "SELECT user_session.session_expiry, " +
	"user_account.full_name,user_account.email,user_account.phone,user_account.user_account_id,user_account.active " +
	"from user_session, user_account WHERE auth_token=$1 " +
	"AND user_session.user_account_id=user_account.user_account_id;"

//GetUserByAuth gets user by auth token
func GetUserByAuth(uauth string) (AuthUser, error) {
	db := postgres.Db
	var au AuthUser
	err := db.Get(&au, sqlFetchAuthUser, uauth)
	if err != nil {
		//log.Println(err)
		//log.Println("No auth user found")
		return au, err
	}
	if !au.Active {
		log.Println("Inactive user => " + au.Email + ":" + au.Phone)
		return au, errors.New("inactive user")
	}
	return au, err
}

func GetAuthenticatedUser(r *http.Request) (AuthUser, error) {
	var authUser AuthUser
	cookie, err := r.Cookie("uauth-token")
	if err != nil {
		return authUser, errors.New("Unauthorized")
	}
	if cookie == nil {
		return authUser, errors.New("Unauthorized")
	}
	authUser, err = GetUserByAuth(cookie.Value)
	if err != nil {
		return authUser, err
	}
	return authUser, err
}

func PagePath(dirPath string, spage string) string {
	return core.FixPathSlash(dirPath + "/html/" + spage)
}

func DbPath(dirPath string, dbFile string) string {
	return core.FixPathSlash(dirPath + "/dbf/" + dbFile)
}

func PageError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	tmpl500, err500 := template.ParseFiles(tmpDir+"base.html", tmpDir+"500.html")
	err500 = tmpl500.Execute(w, nil)
	HandlePageError(err500)
}

func PageNotFound(w http.ResponseWriter, r *http.Request, p Paths) {
	w.WriteHeader(http.StatusNotFound)
	tmpl404, err404 := template.ParseFiles(core.FixPathSlash(p.DirPath+"/html/base.html"), core.FixPathSlash(p.DirPath+"/html/404.html"))
	pd := PageData{
		//Title: "hello",
		Csrf:      csrf.Token(r),
		StaticUrl: p.StaticUrl,
	}
	err404 = tmpl404.Execute(w, pd)
	HandlePageError(err404)
}

func RenderPage(w http.ResponseWriter, r *http.Request, pd PageData) {
	//pageData := pd
	paths := Paths{DirPath: core.BinPath, StaticUrl: core.StaticUrl}
	pagePath := PagePath(paths.DirPath, pd.PageFile)
	pd.StaticUrl = "/static"
	pd.RecapSiteKey = core.Config.RecapSiteKey
	//fmt.Println(PagePath(paths.DirPath, pd.BasePageFile))
	//fmt.Println(pagePath)
	authUser, err := GetAuthenticatedUser(r)
	if err == nil {
		pd.UAuthLoggedIn = true
		pd.LoggedInUser = authUser.FullName
	}
	pd.Csrf = csrf.Token(r)
	tmpl, err := template.ParseFiles(PagePath(paths.DirPath, pd.BasePageFile), pagePath)
	if err != nil {
		log.Println("Template error")
	}

	err = tmpl.Execute(w, pd)
	if err != nil {
		log.Println(err)
		log.Println("Template exe error")
	}
}

type RecaptchaResponse struct {
	Success     bool      `json:"success"`
	Score       float64   `json:"score"`
	Action      string    `json:"action"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

var verifyURL = "https://www.google.com/recaptcha/api/siteverify"

func ValidRecaptcha(token string) bool {

	resp, err := http.PostForm(verifyURL,
		url.Values{"secret": {core.Config.RecapSecretKey}, "response": {token}})
	if err != nil {
		log.Printf("Post error: %s\n", err)
		return false
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Read error: could not read body")
		log.Println(err)
		return false
	}
	var r RecaptchaResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		log.Println("Cannot umarshall recaptcha json")
		log.Println(err)
		return false
	}
	if r.Success && (r.Score > 0.5) {
		return true
	}
	fmt.Println(r)
	return false

}
