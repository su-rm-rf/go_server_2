package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"com.yuhualing/todo/middlewares"
	"com.yuhualing/todo/controllers"
)

func main() {
	router := httprouter.New()
	
	router.GET("/", controllers.Index)
	router.GET("/todo/list", middlewares.Cors(controllers.TodoList))
	// router.GET("/todo/:id", middlewares.Cors(controllers.TodoDetail))

	// 非常的无语，Golang的生态，实在是太鸡肋了

	log.Fatal(http.ListenAndServe(":8702", router))
}