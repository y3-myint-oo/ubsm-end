package routers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	m "ymo.me/sbum-end/models"
	u "ymo.me/sbum-end/utils"
)

func createTodo(context *gin.Context) {
	context.Header("Content-Type", "application/json")

	title := context.PostForm("Title")
	completed, _ := strconv.ParseBool(context.PostForm("Completed"))
	var todo = m.Todo{bson.NewObjectId(), title, completed}
	fmt.Println("" + todo.Title + " completed: " + strconv.FormatBool(todo.Completed))
	err := u.TodosCollection.Insert(&todo)
	if err != nil {
		log.Fatal(err)
	}

	context.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "todo item created successfully",
	})
}

func fetchAllTodo(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	context.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   m.MockTodos,
	})
}

func fetchSingleTodo(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   "Hello",
	})
}

func updateTodo(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Todo updated successfully!",
	})
}

func deleteTodo(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Todo deleted successfully!",
	})
}
