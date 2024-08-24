// internal/handler/todo_handler.go
package handler

import (
    "net/http"
    "strconv"
    "practice02/internal/domain"
    "practice02/internal/usecase"

    "github.com/labstack/echo/v4"
)

type TodoHandler struct {
    TodoUsecase usecase.TodoUsecase
}

func NewTodoHandler(e *echo.Echo, u usecase.TodoUsecase) {
    handler := &TodoHandler{
        TodoUsecase: u,
    }

    e.POST("/todos", handler.CreateTodo)
    e.GET("/todos", handler.GetAllTodos)
    e.GET("/todos/:id", handler.GetTodoByID)
    e.PUT("/todos/:id", handler.UpdateTodo)
    e.DELETE("/todos/:id", handler.DeleteTodo)
}

func (h *TodoHandler) CreateTodo(c echo.Context) error {
    todo := new(domain.Todo)
    if err := c.Bind(todo); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }

    if err := h.TodoUsecase.Create(todo); err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusCreated, todo)
}

func (h *TodoHandler) GetAllTodos(c echo.Context) error {
    todos, err := h.TodoUsecase.GetAll()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, todos)
}

func (h *TodoHandler) GetTodoByID(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    todo, err := h.TodoUsecase.GetByID(uint(id))
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, todo)
}

func (h *TodoHandler) UpdateTodo(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    todo := new(domain.Todo)
    if err := c.Bind(todo); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }

    todo.ID = uint(id)
    if err := h.TodoUsecase.Update(todo); err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusOK, todo)
}

func (h *TodoHandler) DeleteTodo(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    if err := h.TodoUsecase.Delete(uint(id)); err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.NoContent(http.StatusNoContent)
}
