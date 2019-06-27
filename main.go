package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/model_gen/consts"
)

func AppendToFile(filename, text string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err = f.WriteString(text); err != nil {
		panic(err)
	}
}

func CreateFile(filename, text string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	_, err = f.WriteString(text)
	if err != nil {
		f.Close()
		return err
	}
	fmt.Println("Created file", filename)
	return nil
}

func AddRoute(filename, route string) error {
	f, err := os.Open("./routers/app.go")
	if err != nil {
		f.Close()
		return err
	}
	scanner := bufio.NewScanner(f)
	text := ""
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), ").Handler(appRouter)") {
			text += "\t" + route + "\n" + scanner.Text() + "\n"
		} else {
			text += scanner.Text() + "\n"
		}
	}
	f.Close()
	err = os.Remove(filename)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = CreateFile(filename, text)

	fmt.Println(text)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func CheckSyntax(name string) error {
	strs := strings.Split(name, "_")
	if len(strs) == 1 {
		return errors.New("Bad syntax")
	}
	return nil
}

func Add(name string) {
	templateName := name
	controllerName := ""
	route := ""

	name = strings.ToLower(name)
	strs := strings.Split(name, "_")
	for k := 0; k < len(strs); k++ {
		route += "/" + strs[k]
		controllerName += strings.Title(strs[k])
	}

	route = fmt.Sprintf(`appRouter.HandleFunc("%s", controllers.%s).Methods("GET")`, route, controllerName)

	controllerFunc := fmt.Sprintf(consts.FuncControllerConst, controllerName, templateName)
	AppendToFile("./controllers/viewController.go", controllerFunc)

	template := fmt.Sprintf(consts.HTMLConst, templateName)
	CreateFile("./templates/"+templateName+".html", template)

	AddRoute("./routers/app.go", route)

}

type ModelField struct {
	Name string
	Type string
}

type Model struct {
	Name   string
	Fields []ModelField
}

func modelForm(model Model) string {
	formGroups := ""
	for _, field := range model.Fields {
		formGroups += fmt.Sprintf(consts.FormGroup, strings.Title(field.Name))
	}
	formContent := fmt.Sprintf(consts.Form, formGroups)
	return formContent
}

func modelTable(model Model) string {
	header := ""
	body := ""
	for _, field := range model.Fields {
		header += fmt.Sprintf("\t\t\t\t\t\t\t<th scope=\"col\">%s</th>\n", strings.Title(field.Name))
		body += fmt.Sprintf("\t\t\t\t\t\t\t\t<td>{{ .%s }}</td>\n", strings.Title(field.Name))
	}

	return fmt.Sprintf(consts.TableConst, header, model.Name, body)
}

func ModelGen(model Model) {
	model.Fields = append([]ModelField{{"id", "primitive.ObjectID"}}, model.Fields...)
	fieldsContent := ""
	updateContent := ""
	for _, field := range model.Fields {
		fieldname := strings.ToLower(field.Name)
		fieldRow := fmt.Sprintf("%s\t%s\t `json:\"%s\" bson:\"%[3]s\"`", strings.Title(field.Name), strings.ToLower(field.Type), fieldname)
		fieldsContent += fieldRow + "\n\t"

		nameWithQuotes := "\t\t\"" + fieldname + "\" : " + strings.ToLower(model.Name[:1]) + "." + strings.Title(field.Name) + ", "
		updateContent += nameWithQuotes + "\n"
	}
	fileContent := fmt.Sprintf(consts.ModelConst, model.Name, fieldsContent)
	fileContent += fmt.Sprintf(consts.ModelDAOConst, strings.ToUpper(model.Name)+"_COL", strings.ToLower(model.Name)+"s",
		strings.ToLower(model.Name[:1]), model.Name, updateContent)
	// fmt.Printf(fileContent)
	CreateFile("./models/"+strings.ToLower(model.Name)+".go", fileContent)

	htmlCreateContent := fmt.Sprintf(consts.HTMLConst, strings.ToLower(model.Name)+"_create", modelForm(model))
	CreateFile("./templates/"+strings.ToLower(model.Name)+"_create.html", htmlCreateContent)

	// fmt.Printf(htmlCreateContent)
	controllerContent := fmt.Sprintf(consts.ControllerConst, strings.Title(model.Name), strings.ToLower(model.Name))
	CreateFile("./controllers/"+strings.ToLower(model.Name)+"Controller.go", controllerContent)
	// fmt.Println(controllerContent)

	// htmlTableContent := fmt.Sprintf(consts.HTMLConst, strings.ToLower(model.Name)+"_table", modelTable(model))
	// fmt.Println(htmlTableContent)

}

func main() {
	model := Model{}
	model.Name = "User"
	model.Fields = []ModelField{{"name", "string"}, {"username", "string"}, {"password", "string"}}
	ModelGen(model)
}

// Concept
// Ela Admin
