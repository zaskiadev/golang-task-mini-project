package main

import (
	"github.com/julienschmidt/httprouter"
	"golang-task-mini-project/controllers"
	"golang-task-mini-project/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	db, err := gorm.Open(sqlite.Open("databasetask.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	err = db.AutoMigrate(&models.UserTasks{}, &models.Task{}, &models.TaskComment{})
	if err != nil {
		//fmt.Println(err.Error())
	}

	webController := &controllers.WebControllers{}

	router := httprouter.New()
	router.GET("/", webController.Login)
	router.POST("/", webController.Login)
	router.GET("/home", webController.Home)
	router.GET("/add_task", webController.AddTask)
	router.POST("/add_task", webController.AddTask)
	router.GET("/edit_task/:codetask", webController.EditTask)
	router.POST("/update/:codetask", webController.UpdateTask)

	router.GET("/add_comment_task/", webController.AddCommentTask)
	router.POST("/add_comment_task", webController.AddCommentTask)

	router.GET("/show_all_comment_task", webController.ShowAllCommentTask)

	router.POST("/done/:codetask", webController.DoneTask)

	log.Fatal(http.ListenAndServe(":6490", router))

}
