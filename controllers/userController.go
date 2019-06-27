package controllers

import (
	"../models"

	"log"
	"net/http"

	"github.com/gorilla/schema"
)

func UserCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		} else {
			user := new(models.User)

			decoder := schema.NewDecoder()
			err = decoder.Decode(user, r.PostForm)
			if err != nil {
				log.Println(err)
			} else {
				err = user.Insert()
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
	utils.RenderTemplate(w, "user_create", nil)
}

func UserTable(w http.ResponseWriter, r *http.Request) {
	users := models.AllUsers()

	context := make(map[string]interface{})
	context["Users"] = users

	utils.RenderTemplate(w, "info", context)
}
