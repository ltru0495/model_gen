package consts

const ControllerConst = `
package controllers
import (
	"../models"

	"log"
	"net/http"

	"github.com/gorilla/schema"
)

func %sCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		} else {
			%s := new(models.%[1]s)
			
			decoder := schema.NewDecoder()
			err = decoder.Decode(%[2]s, r.PostForm)
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


func %[1]sTable(w http.ResponseWriter, r *http.Request) {
	users := models.AllUsers()

	context := make(map[string]interface{})
	context["Users"] = users

	utils.RenderTemplate(w, "info", context)
}

`
