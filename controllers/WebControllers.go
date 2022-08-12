package controllers

import (
	"github.com/julienschmidt/httprouter"
	"golang-task-mini-project/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var UserCurrentLogin string = ""

type WebControllers struct{}

func (controller *WebControllers) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	_, err := gorm.Open(sqlite.Open("databasetask.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	files := []string{
		"./views/base.html",
		"./views/login.html",
	}

	htmlTemplate, err := template.ParseFiles(files...)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	//var DataTasks []models.Task
	//db.Find(&DataTasks)
	datas := map[string]interface{}{}
	err = htmlTemplate.ExecuteTemplate(w, "base", datas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (controller *WebControllers) Home(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	db, err := gorm.Open(sqlite.Open("databasetask.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	files := []string{
		"./views/base.html",
		"./views/home.html",
	}

	htmlTemplate, err := template.ParseFiles(files...)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	var DataTasks []models.Task
	db.Find(&DataTasks)
	datas := map[string]interface{}{
		"Tasks": DataTasks,
	}
	err = htmlTemplate.ExecuteTemplate(w, "base", datas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (controller *WebControllers) AddTask(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	db, err := gorm.Open(sqlite.Open("databasetask.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())

	}

	if r.Method == "POST" {
		var selectedValue = r.FormValue("dataUserTaskTo")

		var data = models.UserTasks{}
		db.First(&data, "user_name=? ", strings.Replace(selectedValue, " ", "", -1))

		//db.Where("user_name = ?", selectedValue).Select("CodeUserTask").Find(&data)
		//var getSelectedUserTask = data.CodeUserTask
		//fmt.Println("selected user task : %s", getSelectedUserTask)

		task := models.Task{
			CodeUserCreateTask:      "0",
			CodeUserDestinationTask: data.CodeUserTask,
			Task:                    r.FormValue("taskDetail"),
			DateDeadLineTask:        r.FormValue("deadLineTask"),
			StatusTask:              "Create",
			TaskComment:             nil,
		}

		result := db.Create(&task)

		if result.Error != nil {
			log.Println(result.Error)
		}

		http.Redirect(w, r, "/home", http.StatusFound)
	} else {
		files := []string{
			"./views/base.html",
			"./views/add_task.html",
		}

		htmlTemplate, err := template.ParseFiles(files...)

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		var data []models.UserTasks
		db.Find(&data)
		datas := map[string]interface{}{
			"DataUser": data,
		}
		err = htmlTemplate.ExecuteTemplate(w, "base", datas)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}

}

func (controller *WebControllers) AddCommentTask(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	db, err := gorm.Open(sqlite.Open("databasetask.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	files := []string{
		"./views/base.html",
		"./views/login.html",
		"./views/home.html",
		"./views/add_task.html",
		"./views/add_comment_task.html",
		"./views/show_all_comment_task ",
	}

	htmlTemplate, err := template.ParseFiles(files...)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	var DataTasks []models.Task
	db.Find(&DataTasks)
	datas := map[string]interface{}{
		"Tasks": DataTasks,
	}
	err = htmlTemplate.ExecuteTemplate(w, "base", datas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (controller *WebControllers) ShowAllCommentTask(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	db, err := gorm.Open(sqlite.Open("databasetask.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	files := []string{
		"./views/base.html",
		"./views/login.html",
		"./views/home.html",
		"./views/add_task.html",
		"./views/add_comment_task.html",
		"./views/show_all_comment_task ",
	}

	htmlTemplate, err := template.ParseFiles(files...)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	var ds []models.Task
	db.Find(&ds)
	datas := map[string]interface{}{
		"Tasks": ds,
	}

	println(datas)
	err = htmlTemplate.ExecuteTemplate(w, "base", datas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
