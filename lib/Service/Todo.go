package service_todo

import (
	"net/http"
	// "strings"
	// "encoding/json"
	// "strconv"
	"github.com/kawanishik/todo-go-vue/lib/Utility"
	"github.com/gin-gonic/gin"
)

type Todo struct {
	Id    int
	Title string
	Description string
	Completed bool
	Created_at string
	Updated_at string
}

type Params struct {
	Id    int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
}

func Index(c *gin.Context) {
	selected, err := utility.Db.Query("SELECT * FROM todos")
	if err != nil {
		panic(err.Error())
	}
	data := []Todo{}
	for selected.Next() {
		todo := Todo{}
		err = selected.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Completed, &todo.Created_at, &todo.Updated_at)
		if err != nil {
			panic(err.Error())
		}
		data = append(data, todo)
	}
	defer selected.Close()
	c.JSON(http.StatusOK, data)
}

func Create(c *gin.Context) {
	var params Params
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if params.Title == "" || params.Description == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title or description is empty"})
		return
	}

	insert, err := utility.Db.Prepare("INSERT INTO todos(title, description) VALUES(?,?)")
	if err != nil {
		panic(err.Error())
	}
	insert.Exec(params.Title, params.Description)
	c.JSON(http.StatusOK, gin.H{"message": "suncess created!!"})
}

func Delete(c *gin.Context) {
	var params Params
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// paramsが存在するのかチェック
	if params.Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is empty"})
		return
	}

	// 消したい要素があるかチェック
	selected, err := utility.Db.Query("SELECT * FROM todos WHERE id = ?", params.Id)
	if !selected.Next() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is not found"})
		return
	}
	defer selected.Close()

	delete, err := utility.Db.Prepare("DELETE FROM todos WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	delete.Exec(params.Id)
	c.JSON(http.StatusOK, gin.H{"message": "suncess deleted!!"})
}

func Edit(c *gin.Context) {
	var params Params
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// params.idが存在するのかチェック
	if params.Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is empty"})
		return
	}

	// 要素が存在するのかをチェック
	selected, err := utility.Db.Query("SELECT * FROM todos WHERE id = ?", params.Id)
	if !selected.Next() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is not found"})
		return
	}
	defer selected.Close()

	update, err := utility.Db.Prepare("UPDATE todos SET title = ?, description = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	update.Exec(params.Title, params.Description, params.Id)
	c.JSON(http.StatusOK, gin.H{"message": "suncess updated!!"})
}
