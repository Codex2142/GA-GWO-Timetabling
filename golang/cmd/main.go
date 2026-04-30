// @title Timetabling Backend API
// @version 1.0
// @description Backend API untuk manajemen penjadwalan sekolah dengan Guided GA-GWO algorithm
// @contact.name API Support
// @host localhost:12000
// @BasePath /api
// @schemes http
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"timetabling/config"
	"timetabling/handlers"
	"timetabling/repositories"
	"timetabling/routes"
	"timetabling/seeds"
	"timetabling/services"

	_ "timetabling/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func main() {
	// Ambil subcommand dari argumen pertama (jika ada)
	subcommand := "serve"
	if len(os.Args) >= 2 {
		subcommand = os.Args[1]
	}

	switch subcommand {
	case "serve":
		// go run ./cmd/...
		// go run ./cmd/... serve
		db := mustConnectDB()
		defer db.Close()
		// Seed jika data belum ada
		if err := seeds.SeedDatabase(db); err != nil {
			log.Printf("warning: %v (data mungkin sudah ada)", err)
		}
		runServer(db)

	case "migrate:fresh":
		// go run ./cmd/... migrate:fresh
		// go run ./cmd/... migrate:fresh --seed
		withSeed := hasFlag("--seed")

		db := mustConnectDB()
		defer db.Close()

		log.Println("migrate:fresh → Mereset database (drop + recreate tables)...")
		if err := config.ResetDatabase(db); err != nil {
			log.Fatalf("migrate:fresh gagal: %v", err)
		}
		log.Println("migrate:fresh → Selesai. Database bersih.")

		if withSeed {
			log.Println("migrate:fresh --seed → Seeding dari CSV...")
			if err := seeds.SeedDatabase(db); err != nil {
				log.Fatalf("seeding gagal: %v", err)
			}
			log.Println("migrate:fresh --seed → Selesai. Data berhasil diisi ulang.")
		}

	default:
		fmt.Fprintln(os.Stderr, "Perintah tidak dikenal:", subcommand)
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "Perintah yang tersedia:")
		fmt.Fprintln(os.Stderr, "  go run ./cmd/...                       Start server (default)")
		fmt.Fprintln(os.Stderr, "  go run ./cmd/... serve                 Start server")
		fmt.Fprintln(os.Stderr, "  go run ./cmd/... migrate:fresh         Reset database")
		fmt.Fprintln(os.Stderr, "  go run ./cmd/... migrate:fresh --seed  Reset database + seed dari CSV")
		os.Exit(1)
	}
}

// hasFlag mengecek apakah flag tertentu ada di os.Args (mulai indeks ke-2).
func hasFlag(flag string) bool {
	for _, arg := range os.Args[2:] {
		if arg == flag {
			return true
		}
	}
	return false
}

// mustConnectDB menghubungkan ke database atau langsung fatal jika gagal.
func mustConnectDB() *sql.DB {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	return db
}

// runServer menyiapkan handler, routes, dan menjalankan Fiber.
func runServer(db *sql.DB) {
	guruRepo := repositories.NewGuruRepository(db)
	kelasRepo := repositories.NewKelasRepository(db)
	mapelRepo := repositories.NewMapelRepository(db)
	slotRepo := repositories.NewSlotRepository(db)
	relasiRepo := repositories.NewRelasiGuruMapelRepository(db)
	waliRepo := repositories.NewWaliKelasRepository(db)

	guruService := services.NewGuruService(guruRepo)
	kelasService := services.NewKelasService(kelasRepo)
	mapelService := services.NewMapelService(mapelRepo)
	slotService := services.NewSlotService(slotRepo)
	relasiService := services.NewRelasiGuruMapelService(relasiRepo, guruRepo, mapelRepo)
	waliService := services.NewWaliKelasService(waliRepo, guruRepo, kelasRepo)

	guruHandler := handlers.NewGuruHandler(guruService)
	kelasHandler := handlers.NewKelasHandler(kelasService)
	mapelHandler := handlers.NewMapelHandler(mapelService)
	slotHandler := handlers.NewSlotHandler(slotService)
	relasiHandler := handlers.NewRelasiGuruMapelHandler(relasiService)
	waliHandler := handlers.NewWaliKelasHandler(waliService)

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin,Content-Type,Accept,Authorization",
	}))

	// Swagger endpoint
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	routes.RegisterRoutes(app, guruHandler, kelasHandler, mapelHandler, slotHandler, relasiHandler, waliHandler)

	log.Println("server running on http://localhost:12000")
	if err := app.Listen(":12000"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
