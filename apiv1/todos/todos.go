package todos

import (
	// "fmt"
	models "jubo-hw/models"
	"jubo-hw/schemas"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	// "strconv"
)

// @Summary get all todo items
// @Schemes
// @Description get all todo items
// @Tags todo
// @Accept json
// @Produce json
// @Success 200 {array} schemas.TodoItem
// @Success 500 {object} schemas.HTTPError
// @Router /todos [get]
func GetTODOItems(g *gin.Context) {
	data, err := models.GetAllTODOItems()
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	g.JSON(http.StatusOK, data)
}

// @Summary get todo item
// @Schemes
// @Description get todo item info
// @Tags todo
// @Accept json
// @Produce json
// @Param id path int true "todo item id"
// @Success 200 {object} schemas.TodoItem
// @Success 500 {object} schemas.HTTPError
// @Router /todos/{id} [get]
func GetTODOItem(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "id must be an integer"})
		return
	}
	data, err := models.GetTODOItem(id)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "id not found"})
		return
	}
	g.JSON(http.StatusOK, data)
}

// @Summary create todo item
// @Schemes
// @Description create todo item
// @Tags todo
// @Accept json
// @Produce json
// @Param todoitem body schemas.CreateTodoItemBody true "todo item"
// @Success 500 {object} schemas.HTTPError
// @Router /todos [post]
func CreateTODOItem(g *gin.Context) {
	var todoItem schemas.CreateTodoItemBody
	if err := g.ShouldBindJSON(&todoItem); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if err := models.CreateTODOItem(todoItem); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
}

// @Summary update todo item
// @Schemes
// @Description change todo item status to completed
// @Tags todo
// @Accept json
// @Produce json
// @Param id path int true "todo item id"
// @Success 500 {object} schemas.HTTPError
// @Router /todos/{id} [put]
func UpdateTODOItem(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "id must be an integer"})
		return
	}
	if err := models.UpdateTODOItem(id); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "id not found"})
		return
	}
}

// @Summary delete todo item
// @Schemes
// @Description delete todo item
// @Tags todo
// @Accept json
// @Produce json
// @Param id path int true "todo item id"
// @Success 500 {object} schemas.HTTPError
// @Router /todos/{id} [delete]
func DeleteTODOItem(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "id must be an integer"})
		return
	}
	if err := models.DeleteTODOItem(id); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "id not found"})
		return
	}
}
