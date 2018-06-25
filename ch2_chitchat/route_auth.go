package main

import (
	"net/http"
	"time"
	"github.com/sausheong/gwp/Chapter_2_Go_ChitChat/chitchat/data"
)

func authenticate(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	user, _ := data.UserByEmail(r.PostFormValue("email"))
	if user.Password == data.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			danger(err, "Cannot create session")
		}
		cookie := http.Cookie{
			Name: "_cookie",
			Value: session.Uuid,
			HttpOnly: true,
		}

		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

type Session struct {
	Id int
	Uuid string
	Email string
	UserId int
	CreatedAt time.Time
}