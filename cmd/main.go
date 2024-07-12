package main

import (
	"blog/database"
	"blog/internal"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	engine := html.New("html", ".html")
	engine.Reload(true)
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	db, err := database.ConnectToDB()
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	internal.Route(app, db)
	go func() {
		log.Fatal(app.Listen(":8888"))
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGQUIT, syscall.SIGTERM)
	<-quit
	log.Printf("Shutting down server...")
	app.ShutdownWithContext(context.Background())
}
