package uauth

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/devasiajoseph/webapp/crypt"
	"github.com/devasiajoseph/webapp/db/postgres"
)

type UserAccount struct {
	UserAccountID int       `db:"user_account_id" json:"user_account_id"`
	Email         string    `db:"email" json:"email"`
	Phone         string    `db:"phone" json:"phone"`
	Password      string    `db:"password" json:"-"`
	FullName      string    `db:"full_name" json:"full_name"`
	Active        bool      `db:"active" json:"active"`
	CreatedOn     time.Time `db:"created_on" json:"created_on"`
	LastLogin     time.Time `db:"last_login" json:"last_login"`
	Role          string    `db:"role" json:"role"`
}

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
	Role          string    `db:"role" json:"Role"`
}

//Create creates new user account
func (ua *UserAccount) Create() error {
	db := postgres.Db
	var sqlInsertUser = "INSERT INTO user_account" +
		"(phone,email, password, full_name, active, created_on, last_login)" +
		"VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING user_account_id;"

	err := db.QueryRow(sqlInsertUser,
		ua.Phone,
		ua.Email,
		crypt.HashPassword(ua.Password),
		ua.FullName,
		ua.Active,
		time.Now(),
		time.Now()).Scan(&ua.UserAccountID)
	if err != nil {
		log.Println("Error creating user")
		log.Println(err)
		return err
	}

	return err
}

//Data fetches user account data based on id or phone or email
func (ua *UserAccount) Data() error {
	db := postgres.Db
	bsql := "select * from user_account where "
	var err error
	if ua.UserAccountID > 0 {
		bsql += " user_account_id=$1;"
		err = db.Get(ua, bsql, ua.UserAccountID)
		if err != nil {
			fmt.Println(err)
			log.Println("Unable to fetch user data with id")
		}
		return err
	}
	if ua.Phone != "" {
		bsql += " phone=$1;"
		err := db.Get(ua, bsql, ua.Phone)
		if err != nil {
			log.Println("Unable to fetch user data with phone ")
		}
		return err
	}

	if ua.Email != "" {
		bsql += " email=$1;"
		err := db.Get(ua, bsql, ua.Email)
		if err != nil {
			log.Println("Unable to fetch user data with email ")
		}
		return err
	}

	return errors.New("No fetch parameter provided")
}

func keepAlive(cookie string) {
	db := postgres.Db
	expiry := time.Now().Add(30 * time.Minute)
	_, err := db.Exec(sqlKeepAlive, expiry, cookie)
	if err != nil {
		log.Println("Error keeping alive")
	}
}

//GetUserByAuth gets user by auth token
func GetUserByAuth(uauth string) (AuthUser, error) {
	db := postgres.Db
	var au AuthUser
	err := db.Get(&au, sqlFetchAuthUser, uauth)
	if err != nil {
		log.Println(err)
		return au, err
	}
	if !au.Active {
		log.Println("Inactive user => " + au.Email + ":" + au.Phone)
		return au, errors.New("Inactive user")
	}
	return au, err
}
