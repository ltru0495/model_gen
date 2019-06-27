package consts

const FuncControllerConst = `func %s(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "%s", nil)
}
`

const HTMLConst = `{{ define "%s"}}
{{template "header"}}
</head>
<body>
	<div class="container">
		<div class="row">
			<div class="col-md-12">%s
			</div>
		</div>
	</div>
{{template "footer"}}
</body>
{{ end }}`

const ModelConst = `package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type %s struct {
	%s
}
`
const ModelDAOConst = `
const %s = "%s"

func (%s *%s) Insert() error {
	%[3]s.Id = primitive.NewObjectID()
	res, err := getDatabase().Collection(%[1]s).InsertOne(context.Background(), t)
	%[3]s.Id = res.InsertedID.(primitive.ObjectID)
	return err
}

func (%[3]s *%[4]s) Delete() error {
	res, err := getDatabase().Collection(%[1]s).DeleteOne(context.Background(), bson.D{{"_id", %[3]s.Id}})
	%[3]s.Id = res.InsertedID.(primitive.ObjectID)
	return err
}

func (%[3]s *%[4]s) Update() error {
	update := bson.M{"$set": bson.M{
%s
	}}
	_, err := getDatabase().Collection(%[1]s).UpdateOne(context.Background(), bson.D{{"_id", %[3]s.Id}}, update)
	return err
}

func Find%[4]s(id primitive.ObjectID) (%[3]s %[4]s, err error) {
	filter := bson.D{{"_id", id}}
	err = getDatabase().Collection(%[1]s).FindOne(context.Background(), filter).Decode(&%[3]s)
	return
}

func All%[4]ss() (%[2]s []%[4]s, err error) {
	cursor, err := getDatabase().Collection(%[1]s).Find(context.Background(), bson.D{{}})
	if err != nil {
		return
	}
	var %[3]s %[4]s
	for cursor.Next(context.TODO()) {
		err = cursor.Decode(&%[3]s)
		if err != nil {
			return
		}
		tests = append(%[2]s, %[3]s)
	}
	if err = cursor.Err(); err != nil {
		return
	}
	cursor.Close(context.TODO())
	return
}
`

const Form = `
			<form>%s
			  	<button type="submit" class="btn btn-primary">Submit</button>
			</form>
`
const FormGroup = `
			<div class="form-group">
			    <label for="input%s">%[1]s</label>
			    <input type="text" class="form-control" id="input%[1]s" placeholder="Enter %[1]s">
			</div>
`

const TableConst = `
				<table class="table">
					<thead>
						<tr>
%s
						</tr>
					</thead>

					<tbody>
						{{ range .%ss }}
							<tr>
%s
							</tr>
						{{end}}
					</tbody>
				</table>`

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
