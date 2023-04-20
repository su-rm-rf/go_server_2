package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"com.yuhualing/todo/middlewares"
	"com.yuhualing/todo/controllers"
)

func main() {
	router := httprouter.New()

	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			header := w.Header()
			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", "*")
			header.Set("Access-Control-Allow-Credentials", "true")
			header.Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With")
		} else {
			fmt.Println(r.Header.Get("Access-Control-Request-Method"), w.Header().Get("Allow"))
		}
    // Adjust status code to 204
    w.WriteHeader(http.StatusNoContent)
	})
	
	// 非常的无语，Golang的生态，实在是太鸡肋了
	router.GET("/", controllers.Index)
	router.GET("/todolist", middlewares.Cors(controllers.TodoList))
	router.GET("/tododetail/:id", middlewares.Cors(controllers.TodoDetail))
	router.POST("/todoadd", middlewares.Cors(controllers.TodoAdd))
	router.POST("/todoupdate", middlewares.Cors(controllers.TodoUpdate))
	router.POST("/tododelete", middlewares.Cors(controllers.TodoDelete))
	router.POST("/todoclear", middlewares.Cors(controllers.TodoClear))

	log.Fatal(http.ListenAndServe(":8702", router))
}