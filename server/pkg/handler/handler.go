package handler

import (
	service "github.com/Argento-prg/NoteBook/server/pkg/service"
	gin "github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(cors.AllowAll())

	// регистрация и авторизация
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp) //регистрация
		auth.POST("/sign-in", h.signIn) //авторизация
	}

	// взаимодействие с остальными сущностями
	api := router.Group("/api", h.userIdentity)
	{
		// взаимодействие со списками
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)      // добавление списка
			lists.GET("/", h.getAllLists)      // получение списков
			lists.GET("/:id", h.getListById)   // получение определённого списка
			lists.PUT("/:id", h.updateList)    // редактирование списка
			lists.DELETE("/:id", h.deleteList) // удаление списка
			// взаимодействие с задачами
			todos := lists.Group(":id/todos")
			{
				todos.POST("/", h.createTodo)           // добавление задачи
				todos.GET("/", h.getAllTodos)           // получение всех задач из списка
				todos.GET("/:todo_id", h.getTodoById)   // получение определённой задачи
				todos.PUT("/:todo_id", h.updateTodo)    // редактирование задачи
				todos.DELETE("/:todo_id", h.deleteTodo) // удаление задачи
			}
		}
	}
	return router
}
