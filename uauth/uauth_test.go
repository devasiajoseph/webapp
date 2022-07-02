package uauth

import (
	"testing"

	"github.com/devasiajoseph/webapp/core"
	"github.com/devasiajoseph/webapp/db/postgres"
)

func TestCreate(t *testing.T) {
	core.Start()
	postgres.InitDb()
	ua := UserAccount{
		Phone:    "9539100781",
		Password: "password",
		Email:    "devasiajoseph@gmail.com",
	}

	err := ua.Create()
	if err != nil {
		t.Errorf("Error creating ")
	}
}
