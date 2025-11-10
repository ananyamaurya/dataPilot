package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"

	"datapilot/internal/api"
	"datapilot/internal/db"
	"datapilot/internal/storage"
)

func main() {
	// 1️⃣ Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ No .env file found, using system env variables")
	}

	port := os.Getenv("GO_API_PORT")
	if port == "" {
		port = "8080"
	}

	// 2️⃣ Initialize database connection
	pg, err := db.InitPostgres()
	if err != nil {
		log.Fatalf("❌ Failed to connect to Postgres: %v", err)
	}
	defer pg.Close()

	// 3️⃣ Initialize MinIO (object storage)
	minioClient, err := storage.InitMinio()
	if err != nil {
		log.Fatalf("❌ Failed to connect to MinIO: %v", err)
	}

	// 4️⃣ Create Fiber app
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	// 5️⃣ Health check route
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// 6️⃣ Register routes
	api.SetupRoutes(app, pg, minioClient)

	// 7️⃣ Start server
	log.Printf("🚀 Go API running on port %s", port)
	if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("❌ Server failed: %v", err)
	}
}
