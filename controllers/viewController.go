package controllers

import (
	"beegons/models"

	"beegons/utils"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"

	"beegons/models/database"
)

func IndexGET(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "index", nil)
}
