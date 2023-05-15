package main

import (
	"fmt"

	"github.com/FurkanSamaraz/Dependency-Injection/app/controllers"
	"github.com/FurkanSamaraz/Dependency-Injection/app/pkg/middleware"
	"github.com/FurkanSamaraz/Dependency-Injection/app/pkg/model"
	"github.com/FurkanSamaraz/Dependency-Injection/app/pkg/services"
	"github.com/FurkanSamaraz/Dependency-Injection/platform/database"
	"github.com/FurkanSamaraz/Dependency-Injection/platform/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func main() {
	app := fiber.New()
	jwtMiddleware := middleware.NewJWTMiddleware("gizli-anahtar")

	// redis
	redisConn := redis.Connect()
	defer redisConn.Close()

	db := database.Connect()

	// Logger'ı oluştur
	redisModel := model.RedisModel{Conn: redisConn}
	modelRed := model.NewRedis(redisConn)

	chatS := services.NewChatService(db, &redisModel)
	userS := services.NewUserService(db)

	ds := model.UserModel{Scv: redisModel}

	chatC := controllers.NewChatController(chatS)
	userC := controllers.NewUserController(userS, &ds)

	app.Post("/register", userC.RegisterHandler)
	app.Post("/login", userC.LoginHandler)

	// Middleware'leri ekle
	app.Use(jwtMiddleware.Middleware())

	app.Get("/verify-contact", chatC.VerifyContactHandler)
	app.Get("/chat-history", chatC.ChatHistoryHandler)
	app.Get("/contact-list", chatC.ContactListHandler)

	ws := controllers.NewWsController(modelRed)
	app.Get("/ws", websocket.New(ws.WsHandler))

	// Rotaları ayarla
	// routes.SetupRoute1(app, service1, logger, model1)
	// routes.SetupRoute2(app, service2, validator, model2)

	// Uygulamayı başlat
	err := app.Listen(":8080")
	if err != nil {
		fmt.Printf("Uygulama başlatılırken bir hata oluştu: %v", err)
	}
}
