package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tieu-minh/pkg/log"
	"tieu-minh/pkg/router"
	"tieu-minh/service/review/dao"
)

func GetListTodo(w http.ResponseWriter, r *http.Request) {
	repo := &dao.Repo{}

	err := repo.GetList()
	if err != nil {
		log.Println(log.LogLevelError, "GetListTodo: repo.GetList()", err)
		router.ResponseInternalError(w, "Failed")
		return
	}

	router.ResponseSuccessWithData(w, "B.200", "Get list todo", repo)
}

func AddTodo(w http.ResponseWriter, r *http.Request) {
	var todo dao.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Println(log.LogLevelError, "AddTodo: Decode(&todo)", err)
		router.ResponseInternalError(w, err.Error())
		return
	}

	err = todo.Insert()
	if err != nil {
		log.Println(log.LogLevelError, "AddTodo: Insert", err)
		router.ResponseInternalError(w, err.Error())
		return
	}
	router.ResponseSuccess(w, "B.200", "")
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	paramId := r.URL.Query().Get("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		router.ResponseInternalError(w, err.Error())
		log.Println(log.LogLevelDebug, "GetReviews: strconv.Atoi(paramSkip)", err)
		return
	}

	todo := &dao.Todo{
		Id: id,
	}

	err = todo.Delete()
	if err != nil {
		log.Println(log.LogLevelError, "DeleteTodo: Delete", err)
		router.ResponseInternalError(w, err.Error())
		return
	}
	router.ResponseSuccess(w, "B.200", "")
}
