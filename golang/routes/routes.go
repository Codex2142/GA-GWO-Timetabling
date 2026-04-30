package routes

import (
	"timetabling/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App,
	guruHandler *handlers.GuruHandler,
	kelasHandler *handlers.KelasHandler,
	mapelHandler *handlers.MapelHandler,
	slotHandler *handlers.SlotHandler,
	relasiHandler *handlers.RelasiGuruMapelHandler,
	waliHandler *handlers.WaliKelasHandler,
) {
	api := app.Group("/api")

	api.Post("/guru", guruHandler.Create)
	api.Get("/guru", guruHandler.GetAll)
	api.Get("/guru/:id", guruHandler.GetByID)
	api.Put("/guru/:id", guruHandler.Update)
	api.Delete("/guru/:id", guruHandler.Delete)

	api.Post("/kelas", kelasHandler.Create)
	api.Get("/kelas", kelasHandler.GetAll)
	api.Get("/kelas/:id", kelasHandler.GetByID)
	api.Put("/kelas/:id", kelasHandler.Update)
	api.Delete("/kelas/:id", kelasHandler.Delete)

	api.Post("/mapel", mapelHandler.Create)
	api.Get("/mapel", mapelHandler.GetAll)
	api.Get("/mapel/:id", mapelHandler.GetByID)
	api.Put("/mapel/:id", mapelHandler.Update)
	api.Delete("/mapel/:id", mapelHandler.Delete)

	api.Post("/slot", slotHandler.Create)
	api.Get("/slot", slotHandler.GetAll)
	api.Get("/slot/:id", slotHandler.GetByID)
	api.Put("/slot/:id", slotHandler.Update)
	api.Delete("/slot/:id", slotHandler.Delete)

	api.Post("/relasi-guru-mapel", relasiHandler.Create)
	api.Get("/relasi-guru-mapel", relasiHandler.GetAll)
	api.Get("/relasi-guru-mapel/:id", relasiHandler.GetByID)
	api.Put("/relasi-guru-mapel/:id", relasiHandler.Update)
	api.Delete("/relasi-guru-mapel/:id", relasiHandler.Delete)

	api.Post("/wali-kelas", waliHandler.Create)
	api.Get("/wali-kelas", waliHandler.GetAll)
	api.Get("/wali-kelas/:id", waliHandler.GetByID)
	api.Put("/wali-kelas/:id", waliHandler.Update)
	api.Delete("/wali-kelas/:id", waliHandler.Delete)

}
