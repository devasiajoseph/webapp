/*
 * Copyright (C) Centipair Technologies Private Limited - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written By Devasia Joseph <devasiajoseph@centipair.com>, June 2018
 */
package uauth

import (
	"encoding/json"
	"errors"

	"github.com/devasiajoseph/webapp/core"
	"github.com/devasiajoseph/webapp/crypt"
	"github.com/devasiajoseph/webapp/libs/api"
	"github.com/devasiajoseph/webapp/website"

	"log"
	"net/http"
	"time"

	"github.com/devasiajoseph/webapp/validator"
	"github.com/gorilla/mux"
)

func IsSecureCookie() bool {
	if !core.Config.Dev && core.Config.Ssl {
		return true
	} else {
		return false
	}
}

func IsAdmin(r *http.Request) bool {
	ua, err := GetAuthenticatedUser(r)
	if err != nil {
		return false
	}
	if ua.Role == RoleAdmin {
		return true
	}
	return false
}

func ValidateReg(r *http.Request) validator.ValidatorResponse {
	res := validator.InitResponse()
	//validator.RequiredStringValidation(r.FormValue("Email"), "Email", &res)
	validator.RequiredEmailValidation(r.FormValue("email"), "email", &res)
	validator.RequiredStringValidation(r.FormValue("phone"), "phone", &res)
	validator.RequiredStringValidation(r.FormValue("password"), "password", &res)
	validator.RequiredStringValidation(r.FormValue("full-name"), "full-name", &res)
	UniqueEmailValidation(r.FormValue("email"), "email", &res)
	UniquePhoneValidation(r.FormValue("phone"), "phone", &res)
	return res
}

func ValidateUpdate(r *http.Request) validator.ValidatorResponse {
	res := validator.InitResponse()
	email := r.FormValue("email")
	validator.RequiredStringValidation(email, "email", &res)
	validator.RequiredStringValidation(r.FormValue("full-name"), "full-name", &res)
	ua := UserAccount{Email: email}
	err := ua.Data()
	if err == nil {
		UpdatePhoneValidation(ua, email, "phone", &res)
	}
	return res
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
	go keepAlive(cookie.Value)
	return authUser, err
}

func ValidSession(r *http.Request) bool {
	au, err := GetAuthenticatedUser(r)
	if err != nil {
		return false
	}
	return au.Active
}

func ValidObjSession(r *http.Request, UserAccountID int) bool {
	au, err := GetAuthenticatedUser(r)
	if err != nil {
		return false
	}
	if !au.Active {
		return false
	}
	if !(au.UserAccountID == UserAccountID) {
		return false
	}
	return true
}

func ValidateResendOTP(r *http.Request) validator.ValidatorResponse {
	res := validator.InitResponse()
	validator.RequiredStringValidation(r.FormValue("Phone"), "Phone", &res)
	PhoneExistValidation(r.FormValue("Phone"), "Phone", &res)
	return res
}

func resendOTP(w http.ResponseWriter, r *http.Request) {
	if !website.ValidRecaptcha(r.FormValue("recap-token")) {
		api.AuthError(w)
		return
	}
	vRes := ValidateResendOTP(r)
	if !vRes.Valid {
		api.ValidationError(w, vRes)
		return
	}

	var ua UserAccount
	err := ua.Account(r.FormValue("Phone"))
	if err != nil {
		api.ServerError(w)
		return
	}

	uk, err := ua.sendRegistrationOTP()

	if err != nil {
		api.ServerError(w)
		return
	}
	api.ObjectResponse(w, uk)
	return
}

type UserRegistrationResponse struct {
	Message       string `json:"message"`
	Success       bool   `json:"success"`
	ActivationKey string `json:"activation-key"`
}

func RegRequest(w http.ResponseWriter, r *http.Request) {
	if !website.ValidRecaptcha(r.FormValue("recap-token")) {
		api.AuthError(w)
		return
	}

	vRes := ValidateReg(r)

	if !vRes.Valid {
		w.Header().Set("Content-Type", "application/json")
		responseJSON, _ := json.Marshal(vRes)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(responseJSON)
		return
	}
	ua := UserAccount{
		Email:    r.FormValue("email"),
		Phone:    r.FormValue("phone"),
		Password: r.FormValue("password"),
		FullName: r.FormValue("full-name"),
	}

	uk, err := ua.Register()
	if err != nil {
		api.ServerError(w)
		return
	}

	regR := UserRegistrationResponse{Message: "success", Success: true, ActivationKey: uk.KeyValue}
	api.ObjectResponse(w, regR)

}

// LoginAPI handles user login request
func LoginAPI(w http.ResponseWriter, r *http.Request) {
	userStatus := UserStatus{LoggedIn: false, Role: ""}
	//r.ParseForm()
	var ua UserAccount
	err := ua.Account(r.FormValue("email"))

	if err != nil {
		userStatus.Message = "Invalid Account"
		api.ObjectResponse(w, userStatus)
		return
	}

	if !ua.Active {
		userStatus.Message = "User not active"
		api.ObjectResponse(w, userStatus)
		return
	}

	userSession, err := ua.Login(r.FormValue("password"))
	if err != nil {
		userStatus.Message = "Invalid Login"
		api.ObjectResponse(w, userStatus)
		return

	} else {
		//UpdateLoginTime(userSession.UserAccountID)
		cookie := http.Cookie{
			Path:     "/",
			Name:     "uauth-token",
			Value:    userSession.UAuthToken,
			Expires:  userSession.SessionExpiry,
			HttpOnly: true,
			Secure:   IsSecureCookie()}
		http.SetCookie(w, &cookie)
		w.WriteHeader(http.StatusOK)
		userStatus.LoggedIn = true
	}

	api.ObjectResponse(w, userStatus)
}

func logoutAPI(w http.ResponseWriter, r *http.Request) {
	userStatus := UserStatus{LoggedIn: false, Role: ""}
	cookie, err := r.Cookie("uauth-token")
	if err == nil {
		go Logout(cookie.Value)
	}
	ec := http.Cookie{
		Path:     "/",
		Name:     "uauth-token",
		Value:    "",
		MaxAge:   -1,
		Expires:  time.Now().Add(-100 * time.Hour),
		HttpOnly: true,
		Secure:   IsSecureCookie()}
	http.SetCookie(w, &ec)
	w.WriteHeader(http.StatusOK)
	statusJSON, err := json.Marshal(userStatus)
	if err != nil {
		log.Println("Error while json conversion")
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(statusJSON)
}

func userStatus(w http.ResponseWriter, r *http.Request) {
	userStatus := UserStatus{LoggedIn: false, Role: ""}
	authUser, err := GetAuthenticatedUser(r)
	if err == nil {
		userStatus.LoggedIn = true
		userStatus.Username = authUser.Email
		userStatus.Role = authUser.Role
	} else {
		log.Println("Error in user status")
		log.Println(err)
	}
	statusJSON, err := json.Marshal(userStatus)

	if err != nil {
		log.Println("Error while json conversion")
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(statusJSON)
	return
}

func logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("uauth-token")
	if err == nil {
		go Logout(cookie.Value)
	}
	ec := http.Cookie{
		Path:     "/",
		Name:     "uauth-token",
		Value:    "",
		MaxAge:   -1,
		Expires:  time.Now().Add(-100 * time.Hour),
		HttpOnly: true,
		Secure:   IsSecureCookie()}
	http.SetCookie(w, &ec)
	http.Redirect(w, r, "/", 302)

}

func validatePasswordReset(r *http.Request) validator.ValidatorResponse {
	res := validator.InitResponse()
	validator.RequiredStringValidation(r.FormValue("Email"), "Email", &res)
	EmailExistValidation(r.FormValue("Email"), "Email", &res)
	return res
}

func passwordResetStart(w http.ResponseWriter, r *http.Request) {
	pData := website.PageData{}
	pData.PageFile = "forgot-password.html"
	pData.BasePageFile = "base.html"

	website.RenderPage(w, r, pData)
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	pData := website.PageData{}
	pData.PageFile = "login.html"
	pData.BasePageFile = "base.html"

	website.RenderPage(w, r, pData)
}

func registerPage(w http.ResponseWriter, r *http.Request) {

	pData := website.PageData{}
	pData.PageFile = "register.html"
	pData.BasePageFile = "base.html"

	website.RenderPage(w, r, pData)
}

func passwordResetRequest(w http.ResponseWriter, r *http.Request) {
	vRes := validatePasswordReset(r)

	if !vRes.Valid {
		validator.ValidationErrorResponse(w, vRes)
		return
	}

	ua := UserAccount{Email: r.FormValue("Email")}
	err := ua.Data()
	if err != nil {
		log.Println(err)
		api.Error(w, errors.New("Unable to get user data"))
		return
	}

	err = ua.ResetPasswordRequest()

	if err != nil {
		api.Error(w, errors.New("Unable to get request user password reset"))
		return
	}
	api.RequestComplete(w, "password request completed")
}

func passwordResetConfirm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	keyValue := vars["key"]

	uk := UserKeys{KeyName: passwordResetKey, KeyValue: keyValue}
	err := uk.Data()
	pData := website.PageData{}
	if err != nil {
		pData.PageFile = "key-not-found.html"
		pData.BasePageFile = "base.html"
	} else {
		pData.PageFile = "password-reset.html"
		pData.BasePageFile = "base.html"
		pData.KeyValue = keyValue
	}

	website.RenderPage(w, r, pData)
}

func validatePasswordResetSave(r *http.Request) validator.ValidatorResponse {
	res := validator.InitResponse()
	validator.RequiredStringValidation(r.FormValue("Password"), "Password", &res)
	return res
}

func passwordResetSave(w http.ResponseWriter, r *http.Request) {
	vRes := validatePasswordResetSave(r)

	if !vRes.Valid {
		validator.ValidationErrorResponse(w, vRes)
		return
	}
	keyValue := r.FormValue("PasswordResetKey")

	uk := UserKeys{KeyName: passwordResetKey, KeyValue: keyValue}
	err := uk.Data()
	if err != nil {
		log.Println("key error => ", keyValue)
		api.Error(w, errors.New("Unable to get request user password reset"))
		return
	}

	newPassword := r.FormValue("Password")
	err = uk.CompletePasswordReset(newPassword)
	if err != nil {
		log.Println("error in saving password")
		api.Error(w, errors.New("Unable to save new password"))
		return
	}

	api.RequestComplete(w, "password reset completed")
}

func otpActivationPage(w http.ResponseWriter, r *http.Request) {
	regKey := api.QueryParam(r, "registration-key")
	pData := website.PageData{KeyValue: regKey}
	uk := UserKeys{
		KeyName:  registrationKey,
		KeyValue: regKey,
	}
	err := uk.Data()
	if err != nil {
		pData.PageFile = "404.html"
	} else {
		pData.PageFile = "otp.html"
	}

	pData.BasePageFile = "base.html"

	website.RenderPage(w, r, pData)
}

func validateOTP(r *http.Request) validator.ValidatorResponse {
	res := validator.InitResponse()
	uk := UserKeys{
		KeyName:  registrationKey,
		KeyValue: r.FormValue("registration-key"),
	}
	if !uk.VerifyOTP(r.FormValue("otp")) {
		validator.AppendError("Invalid OTP", "otp", &res)
	}

	return res
}

func otpActivationAPI(w http.ResponseWriter, r *http.Request) {
	vRes := validateOTP(r)

	if !vRes.Valid {
		validator.ValidationErrorResponse(w, vRes)
		return
	}

	uk := UserKeys{
		KeyName:  registrationKey,
		KeyValue: r.FormValue("registration-key"),
	}
	err := uk.Data()
	if err != nil {
		log.Println(err)
		api.ServerError(w)
		return
	}
	var resObj api.ResponseObj
	ua := UserAccount{UserAccountID: uk.UserAccountID}
	err = ua.Data()
	if err != nil {
		log.Println(err)
		api.GenericError(w)
		return
	}
	err = ua.Activate()
	if err != nil {
		log.Println(err)
		api.GenericError(w)
		return
	}
	uk.Delete()
	resObj.Code = 200
	resObj.Valid = true
	api.GenericResponse(w, resObj)
	return

}

func validatePasswordChange(ua AuthUser, r *http.Request) validator.ValidatorResponse {
	res := validator.InitResponse()
	ChangePasswordValidation(ua, r.FormValue("Password"), "Password", &res)
	return res
}

func changePassword(w http.ResponseWriter, r *http.Request) {
	ua, err := GetAuthenticatedUser(r)
	if err != nil {
		api.AuthError(w)
	}

	vRes := validatePasswordChange(ua, r)
	if !vRes.Valid {
		api.ValidationError(w, vRes)
		return
	}
	userAccount := UserAccount{UserAccountID: ua.UserAccountID}
	err = userAccount.Data()

	if err != nil {
		api.ServerError(w)
		return
	}

	err = userAccount.SetPassword(r.FormValue("NewPassword"))
	if err != nil {
		api.ServerError(w)
		return
	}

}

func validateSaveProfile(ua UserAccount, r *http.Request) validator.ValidatorResponse {
	res := validator.InitResponse()
	//validator.RequiredStringValidation(r.FormValue("Email"), "Email", &res)
	validator.EmailValue(r.FormValue("Email"), "Email", &res)
	validator.RequiredStringValidation(r.FormValue("Phone"), "Phone", &res)
	validator.RequiredStringValidation(r.FormValue("FullName"), "FullName", &res)
	UpdatePhoneValidation(ua, r.FormValue("Phone"), "Phone", &res)
	return res
}

func updateProfile(w http.ResponseWriter, r *http.Request) {
	au, err := GetAuthenticatedUser(r)
	if err != nil {
		api.AuthError(w)
	}
	ua := UserAccount{Phone: au.Phone}
	err = ua.Data()
	if err != nil {
		api.AuthError(w)
	}
	vRes := validateSaveProfile(ua, r)

	if !vRes.Valid {
		api.ValidationError(w, vRes)
		return
	}

	userAccount := UserAccount{UserAccountID: ua.UserAccountID}
	userAccount.Email = r.FormValue("Email")
	userAccount.Phone = r.FormValue("Phone")
	userAccount.FullName = r.FormValue("FullName")

	err = userAccount.UpdateProfile()
	if err != nil {
		api.ServerError(w)
		return
	}

}

func myProfile(w http.ResponseWriter, r *http.Request) {
	ua, err := GetAuthenticatedUser(r)
	if err != nil {
		api.ServerError(w)
		return
	}

	api.ObjectResponse(w, ua)
}

func validateForgotPassword(r *http.Request) validator.ValidatorResponse {
	res := validator.InitResponse()
	validator.RequiredStringValidation(r.FormValue("Phone"), "Phone", &res)
	validator.IndianPhoneValidation(r.FormValue("Phone"), "Phone", &res)
	PhoneExistValidation(r.FormValue("Phone"), "Phone", &res)
	return res
}

func forgotPassword(w http.ResponseWriter, r *http.Request) {
	if !website.ValidRecaptcha(r.FormValue("recap-token")) {
		api.AuthError(w)
		return
	}
	vRes := validateForgotPassword(r)

	if !vRes.Valid {
		validator.ValidationErrorResponse(w, vRes)
		return
	}

	ua := UserAccount{Phone: r.FormValue("Phone")}
	err := ua.Data()
	if err != nil {
		log.Println(err)
		api.Error(w, errors.New("Unable to get user data"))
		return
	}

	uk, err := ua.sendForgotPasswordOTP()

	if err != nil {
		api.Error(w, errors.New("Unable to get request user password reset"))
		return
	}
	otpSession := OTPSession{KeyName: uk.KeyName, KeyValue: uk.KeyValue}
	api.ObjectResponse(w, otpSession)
}

func forgotPasswordOTPValidation(otp string, otpSession string, id string, r *validator.ValidatorResponse) UserKeys {

	uk := UserKeys{KeyName: passwordResetKey, KeyValue: otpSession}
	if !uk.verifyOTP(otp) {
		validator.AppendError("Invalid OTP", id, r)
	}
	return uk
}

func validateForgotPasswordOTP(r *http.Request) validator.ValidatorResponse {
	res := validator.InitResponse()
	validator.RequiredStringValidation(r.FormValue("otp"), "otp", &res)
	validator.RequiredStringValidation(r.FormValue("otp-session"), "otp-session", &res)

	return res
}

func forgotPasswordOTP(w http.ResponseWriter, r *http.Request) {
	vRes := validateForgotPasswordOTP(r)
	uk := forgotPasswordOTPValidation(r.FormValue("otp"), r.FormValue("otp-session"), "otp", &vRes)

	if !vRes.Valid {
		validator.ValidationErrorResponse(w, vRes)
		return
	}

	otpSession := OTPSession{KeyName: uk.KeyName, KeyValue: uk.KeyValue}
	api.ObjectResponse(w, otpSession)
}

func validatePasswordResetSubmit(r *http.Request) validator.ValidatorResponse {
	res := validator.InitResponse()
	validator.RequiredStringValidation(r.FormValue("otp-session"), "otp-session", &res)
	validator.RequiredStringValidation(r.FormValue("Password"), "Password", &res)
	return res
}

func passwordResetSubmit(w http.ResponseWriter, r *http.Request) {
	vRes := validatePasswordResetSubmit(r)

	if !vRes.Valid {
		validator.ValidationErrorResponse(w, vRes)
		return
	}

	uk := UserKeys{KeyName: passwordResetKey, KeyValue: r.FormValue("otp-session")}

	err := uk.Data()
	if err != nil {
		api.AuthError(w)
		return
	}

	var ua UserAccount

	ua.UserAccountID = uk.UserAccountID
	err = ua.SetPassword(r.FormValue("Password"))

	if err != nil {
		api.ServerError(w)
		return
	}
	api.RequestComplete(w, "Password reset completed")
}

func adminUserList(w http.ResponseWriter, r *http.Request) {
	if !IsAdmin(r) {
		api.AuthError(w)
		return
	}
	var ol ObjList
	page, limit := api.PageLimit(r)

	ol.Page = page
	ol.Limit = limit
	ol.Query = api.QueryParam(r, "q")
	if ol.Query != "" {
		err := ol.Search()

		if err != nil {
			api.ServerError(w)
			return
		}

	} else {
		err := ol.List()

		if err != nil {
			api.ServerError(w)
			return
		}
	}

	api.ObjectResponse(w, ol)
	return
}

func adminSearchUser(w http.ResponseWriter, r *http.Request) {
	if !IsAdmin(r) {
		api.AuthError(w)
		return
	}

	var ol ObjList
	page, limit := api.PageLimit(r)

	ol.Page = page
	ol.Limit = limit
	ol.Query = api.QueryParam(r, "q")
	err := ol.Search()

	if err != nil {
		api.ServerError(w)
		return
	}

	api.ObjectResponse(w, ol)
	return

}

func adminUserSave(w http.ResponseWriter, r *http.Request) {
	if !IsAdmin(r) {
		api.AuthError(w)
		return
	}
	update := false
	var vRes validator.ValidatorResponse
	uID := api.PostInt(r, "UserAccountID")
	if uID == 0 {
		vRes = ValidateReg(r)
	} else {
		update = true
		vRes = ValidateUpdate(r)
	}

	if !vRes.Valid {
		api.ValidationError(w, vRes)
		return
	}
	ua := UserAccount{
		UserAccountID: uID,
		Email:         r.FormValue("Email"),
		Phone:         r.FormValue("Phone"),
		Password:      crypt.HashPassword(r.FormValue("Password")),
		FullName:      r.FormValue("FullName"),
		Active:        true}

	if update {
		err := ua.UpdateProfile()
		if err != nil {
			log.Println("Error updating admin user")
			log.Println(err)
			api.ServerError(w)
			return
		}

	} else {
		err := ua.CreateRaw()
		if err != nil {
			log.Println("Error creating user from admin")
			log.Println(err)
			api.ServerError(w)
			return
		}
	}
	api.ObjectResponse(w, ua)

}

func adminUserFetch(w http.ResponseWriter, r *http.Request) {
	if !IsAdmin(r) {
		api.AuthError(w)
		return
	}

	id := api.ObjID(r, "id")
	if id == 0 {
		api.ObjIDError(w)
		return
	}

	ua := UserAccount{UserAccountID: id}

	err := ua.Data()
	if err != nil {
		api.ObjectNotFound(w)
		return
	}
	api.ObjectResponse(w, ua)

}

func adminChangeUserPassword(w http.ResponseWriter, r *http.Request) {

	if !IsAdmin(r) {
		api.AuthError(w)
		return
	}

	id := api.PostInt(r, "UserAccountID")
	if id == 0 {
		api.ObjIDError(w)
		return
	}

	ua := UserAccount{UserAccountID: id}

	err := ua.ChangePassword(r.FormValue("Password"))
	if err != nil {
		api.ServerError(w)
		return
	}
	api.ObjectResponse(w, ua)
}

func adminDeleteUser(w http.ResponseWriter, r *http.Request) {
	if !IsAdmin(r) {
		api.AuthError(w)
		return
	}
	id := api.ObjID(r, "id")
	obj := UserAccount{
		UserAccountID: id}
	err := obj.Delete()
	if err != nil {
		log.Println("Error while deleting user")
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	return
}

// AddRoutes adds uauth routes to main application
func AddRoutes(r *mux.Router) {
	r.HandleFunc("/api/uauth/logout", logoutAPI).Methods("POST")
	r.HandleFunc("/api/uauth/status", userStatus).Methods("GET")
	r.HandleFunc("/api/uauth/register", RegRequest).Methods("POST")
	r.HandleFunc("/api/uauth/login", LoginAPI).Methods("POST")
	r.HandleFunc("/api/uauth/forgot-password", forgotPassword).Methods("POST")
	r.HandleFunc("/api/uauth/forgot-password-otp", forgotPasswordOTP).Methods("POST")
	r.HandleFunc("/api/uauth/password-reset-submit", passwordResetSubmit).Methods("POST")
	r.HandleFunc("/api/uauth/password-reset-save", passwordResetSave).Methods("POST")
	r.HandleFunc("/app/uauth/password-reset-confirm/{key}", passwordResetConfirm).Methods("GET")
	r.HandleFunc("/app/uauth/password-reset-start", passwordResetStart).Methods("GET")
	r.HandleFunc("/app/uauth-login", loginPage).Methods("GET")
	r.HandleFunc("/app/uauth-logout", logout).Methods("GET")
	r.HandleFunc("/app/uauth-register", registerPage).Methods("GET")
	r.HandleFunc("/app/uauth-otp", otpActivationPage).Methods("GET")
	r.HandleFunc("/api/uauth/otp-activate", otpActivationAPI).Methods("POST")
	r.HandleFunc("/api/uauth/resend-otp", resendOTP).Methods("POST")
	r.HandleFunc("/api/uauth/update-profile", updateProfile).Methods("POST")
	r.HandleFunc("/api/uauth/change-password", changePassword).Methods("POST")
	r.HandleFunc("/api/uauth/my-profile", myProfile).Methods("GET")

	r.HandleFunc("/api/admin/uauth/search", adminSearchUser).Methods("GET")

	r.HandleFunc("/api/admin/uauth", adminUserList).Methods("GET")
	r.HandleFunc("/api/admin/uauth", adminUserSave).Methods("POST")
	r.HandleFunc("/api/admin/uauth/{id}", adminUserFetch).Methods("GET")
	r.HandleFunc("/api/admin/uauth/{id}", adminDeleteUser).Methods("DELETE")
	r.HandleFunc("/api/admin/uauth/password", adminChangeUserPassword).Methods("POST")
}

// Start initializes uauth based functions
func Start(r *mux.Router) {
	AddRoutes(r)
}
