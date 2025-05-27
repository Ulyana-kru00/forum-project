package main

import (
	"database/sql"
	"log"

	_ "github.com/Mandarinka0707/newRepoGOODarhit/chat/docs"
	"github.com/Mandarinka0707/newRepoGOODarhit/chat/internal/handler"
	"github.com/Mandarinka0707/newRepoGOODarhit/chat/internal/repository"
	"github.com/Mandarinka0707/newRepoGOODarhit/chat/internal/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Chat Microservice API
// @version 1.0
// @description This is a chat microservice with WebSocket support
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8082
// @BasePath /
// @schemes http

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	connStr := "postgres://postgres:Password123@localhost:5432/forum_db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repository.NewMessageRepository(db)
	uc := usecase.NewMessageUseCase(repo)
	h := handler.NewMessageHandler(uc)

	go h.HandleMessages()

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// WebSocket endpoint
	// @Summary WebSocket connection
	// @Description Establishes a WebSocket connection for real-time chat
	// @Tags chat
	// @Accept json
	// @Produce json
	// @Success 101 {string} string "Switching Protocols"
	// @Router /ws [get]
	r.GET("/ws", h.HandleConnections)

	// Get messages endpoint
	// @Summary Get messages
	// @Description Get all chat messages
	// @Tags chat
	// @Accept json
	// @Produce json
	// @Success 200 {array} models.Message
	// @Router /messages [get]
	r.GET("/messages", h.GetMessages)

	log.Println("Listening on :8082...")
	log.Fatal(r.Run(":8082"))
}

// package main

// import (
// 	"chat-microservice-go/internal/handler"
// 	"chat-microservice-go/internal/repository"
// 	"chat-microservice-go/internal/usecase"

// 	"github.com/gin-contrib/cors"

// 	"database/sql"
// 	"log"

// 	"github.com/gin-gonic/gin"
// 	_ "github.com/lib/pq"
// )

// func main() {
// 	connStr := "postgres://user:password@localhost:5432/database?sslmode=disable"
// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	repo := repository.NewMessageRepository(db)
// 	uc := usecase.NewMessageUseCase(repo)
// 	h := handler.NewMessageHandler(uc)

// 	go h.HandleMessages()

// 	r := gin.Default()
// 	r.Use(cors.Default())
// 	r.GET("/ws", h.HandleConnections)
// 	r.GET("/messages", h.GetMessages)

// 	log.Println("Listening on :8082...")
// 	log.Fatal(r.Run(":8082"))
// }
