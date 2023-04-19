package controllers

import (
	"fmt"
	// "log"
	"net/http"
	"encoding/json"
	// "reflect"

	"github.com/julienschmidt/httprouter"

	. "com.yuhualing/todo/models"
	. "com.yuhualing/todo/db"
)

type ResultArray struct {
	Status string	`json:"status"`
	Data []map[string]interface{} `json:"data"`
}
type ResultObject struct {
	Status string	`json:"status"`
	Data map[string]interface{} `json:"data"`
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "hello")
}

func TodoList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db := DBConnect()

	// db.AutoMigrate(&Todo{})
	// db.Create(&Todo{ Text: "yyyy", Completed: 1 })

	var todos []map[string]interface{}
	db.Model(Todo{}).Find(&todos)

	rr := ResultArray{Status: "okk", Data: todos}
	res, _ := json.Marshal(rr)

	w.Write(res)
}

func TodoDetail(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db := DBConnect()

	id := ps.ByName("id")

	var todo map[string]interface{}
	db.Model(Todo{}).First(&todo, id)

	rr := ResultObject{Status: "ok", Data: todo}
	res, _ := json.Marshal(rr)
	
	w.Write(res)
}

// router.GET("/todo/list", func(ctx *gin.Context) {
// 	completed := ctx.Query("completed")

// 	list := "select * from todo"
	
// 	var rows *sql.Rows
// 	if completed != "" {
// 		list += " where completed = ?"
// 		rows, _ = db.Query(list, completed)
// 	}
// 	if completed == "" {
// 		rows, _ = db.Query(list)
// 	}

// 	data := []Todo{}
// 	for rows.Next() {
// 		var todo Todo
// 		err := rows.Scan(&todo.Id, &todo.Text, &todo.Completed)

// 		if err != nil {
// 			fmt.Println("scan err: ", err)
// 		}
		
// 		data = append(data, todo)
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{ "data": data})
// })

// router.GET("/todo/:id", func(ctx *gin.Context) {
// 	// id := ctx.Params.ByName("id")
// 	id := ctx.Param("id")
// 	detail := "select * from todo where id = ?"
// 	var todo Todo
// 	err := db.QueryRow(detail, id).Scan(&todo.Id, &todo.Text, &todo.Completed)

// 	if err != nil {
// 		fmt.Println("query err: ", err)
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{ "data": todo })
// })

// router.POST("/todo/add", func (ctx *gin.Context) {
// 	var params Todo
// 	_ = ctx.Bind(&params)
// 	text := params.Text
// 	completed := params.Completed

// 	add := "insert into todo (text, completed) values (?, ?)"
// 	ret, err := db.Exec(add, text, completed)
// 	if err != nil {
// 		fmt.Println("add err: ", err)
// 	}

// 	id, err := ret.LastInsertId()
// 	if err != nil {
// 		fmt.Println("get id err: ", err)
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{ "data": id })
// })

// router.POST("/todo/update", func (ctx *gin.Context) {
// 	var params Todo
// 	_ = ctx.Bind(&params)
// 	id := params.Id
// 	text := params.Text
// 	completed := params.Completed

// 	if text != "" {
// 		update := "update todo set text = ?, completed = ? where id = ?"
// 		ret, _ := db.Exec(update, text, completed, id)
// 		num, _ := ret.RowsAffected()
// 		ctx.JSON(http.StatusOK, gin.H{ "data": num })
// 	}
// 	if text == "" {
// 		update := "update todo set completed = ? where id = ?"
// 		ret, _ := db.Exec(update, completed, id)
// 		num, _ := ret.RowsAffected()
// 		ctx.JSON(http.StatusOK, gin.H{ "data": num })
// 	}
// })

// router.POST("/todo/delete", func (ctx *gin.Context) {
// 	var params map[string]int
// 	json.NewDecoder(ctx.Request.Body).Decode(&params)
// 	id := params["id"]

// 	del := "delete from todo where id = ?"
// 	ret, err := db.Exec(del, id)
// 	if err != nil {
// 		fmt.Println("delete err: ", err)
// 	}
// 	num, _ := ret.RowsAffected()

// 	ctx.JSON(http.StatusOK, gin.H{ "data": num })
// })

// router.POST("/todo/clear", func (ctx *gin.Context) {
// 	clear := "delete from todo"
// 	ret, _ := db.Exec(clear)
// 	num, _ := ret.RowsAffected()
// 	ctx.JSON(http.StatusOK, gin.H{ "data": num })
// })

// router.NoRoute(func (ctx *gin.Context) {

// })