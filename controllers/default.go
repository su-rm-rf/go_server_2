package controllers

import (
	"fmt"
	// "log"
	"net/http"
	"encoding/json"
	"reflect"
	// "strconv"
	"io/ioutil"

	"gorm.io/gorm"
	"github.com/julienschmidt/httprouter"

	. "com.yuhualing/todo/models"
	. "com.yuhualing/todo/db"
	. "com.yuhualing/todo/utils"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, reflect.TypeOf("hello"))
}

func TodoList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db := DBConnect()

	// db.AutoMigrate(&Todo{})
	// db.Create(&Todo{ Text: "yyyy", Completed: 1 })

	query := r.URL.Query()
	completed := query.Get("completed")

	var todos []map[string]interface{}
	if completed != "" {
		db.Model(Todo{}).Where("completed = ?", completed).Find(&todos)
	} else {
		db.Model(Todo{}).Find(&todos)
	}

	rr := ResultArray{Status: "ok", Data: todos}
	res, _ := json.Marshal(rr)
	w.Write(res)
}

func TodoDetail(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db := DBConnect()
	
	id := ps.ByName("id")
	
	var todo map[string]interface{}
	db.Model(Todo{}).First(&todo, id)

	ret := ResultObject{Status: "ok", Data: todo}
	res, _ := json.Marshal(ret)
	w.Write(res)
}

func TodoAdd(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db := DBConnect()

	content := make([]byte, r.ContentLength)
	r.Body.Read(content)

	var todo Todo
	json.Unmarshal([]byte(string(content)), &todo)
	
	db.Create(&todo)

	ret := ResultObject{Status: "ok", Data: todo.Id}
	res, _ := json.Marshal(ret)
	w.Write(res)
}

func TodoUpdate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db := DBConnect()

	content := make([]byte, r.ContentLength)
	r.Body.Read(content)

	var todo map[string]interface{}
	json.Unmarshal([]byte(string(content)), &todo)

	var todo2 Todo
	db.First(&todo2, todo["id"])

	text := todo["text"]
	if text != nil {
		todo2.Text = text.(string)
	}
	todo2.Completed = int(todo["completed"].(float64))
	db.Save(&todo2)

	ret := ResultObject{Status: "ok", Data: todo2}
	res, _ := json.Marshal(ret)
	w.Write(res)
}

func TodoDelete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db := DBConnect()
	
	body, _ := ioutil.ReadAll(r.Body)
	var todo map[string]interface{}
	json.Unmarshal([]byte(body), &todo)

	db.Model(Todo{}).Delete(&todo, todo["id"])

	ret := ResultObject{Status: "ok", Data: todo}
	res, _ := json.Marshal(ret)
	w.Write(res)
}

func TodoClear(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db := DBConnect()

	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Todo{})
	
	ret := ResultObject{Status: "ok", Data: ""}
	res, _ := json.Marshal(ret)
	w.Write(res)
}