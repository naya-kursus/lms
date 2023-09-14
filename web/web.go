package web

import (
	"embed"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/gowok/gowok"
	"github.com/hadihammurabi/belajar-go-rest-api/driver"
	"github.com/hadihammurabi/belajar-go-rest-api/web/middleware"
)

func ConfigureRoute(api *Web) {
	api.HTTP.Mount("", Index(api))
	// api.HTTP.Mount("/auth", Auth(api).router)
}

var a *Web
var viewsFS embed.FS

func Get() *Web {
	if a != nil {
		return a
	}

	a = NewWeb()
	ConfigureRoute(a)
	return a
}

func Run() {
	go (Get()).Run()
}

// Web struct
type Web struct {
	HTTP        *fiber.App
	Middlewares middleware.Middlewares
}

func NewWeb() *Web {
	app := gowok.NewHTTP(&driver.Get().Config.App.Rest)

	appWithView := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})
	appWithView.Mount("", app)

	api := &Web{
		HTTP:        appWithView,
		Middlewares: middleware.NewMiddleware(),
	}
	return api
}

func (d *Web) Run() {
	restConf := driver.Get().Config.App.Rest
	if !restConf.Enabled {
		return
	}

	log.Println("API REST started at", restConf.Host)
	if err := d.HTTP.Listen(restConf.Host); err != nil {
		log.Printf("Server is not running! Reason: %v", err)
	}
}

func (d *Web) Stop() {
	d.HTTP.Shutdown()
	log.Println("Server was stopped")
}
