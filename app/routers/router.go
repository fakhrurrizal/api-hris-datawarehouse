package router

import (
	"fmt"
	"hris-datawarehouse/app/controllers"
	"hris-datawarehouse/app/middlewares"
	"hris-datawarehouse/config"
	"html/template"
	"io"
	"log"
	"net/http"

	_ "hris-datawarehouse/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Init(app *echo.Echo) {
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("*.html")),
	}
	app.Renderer = renderer
	app.Use(middlewares.Cors())
	app.Use(middlewares.Secure())
	app.Use(middlewares.Gzip())
	app.Use(middlewares.Recover())
	app.Use(middlewares.Logger())

	app.GET("/", controllers.Index)
	app.GET("/test", controllers.Test)
	app.GET("/version", controllers.Version)
	app.GET("/swagger/*", echoSwagger.WrapHandler)
	app.GET("/docs", func(c echo.Context) error {
		err := c.Render(http.StatusOK, "docs.html", map[string]interface{}{
			"BaseUrl": config.LoadConfig().BaseUrl,
			"Title":   "Api Documentation of " + config.LoadConfig().AppName,
		})
		fmt.Println("err:", err)
		return err
	})
	app.Static("/assets", "assets")

	api := app.Group("/v1", middlewares.StripHTMLMiddleware, middlewares.CheckAPIKey())
	{
		auth := api.Group("/auth")
		{
			auth.POST("/signin", controllers.SignIn)
		}
		dashboard := api.Group("/dashboard", middlewares.CheckAPIKey())
		{
			dashboard.GET("/score-card", controllers.GetDashboardScoreCard, middlewares.Auth())
		}
	}
	log.Printf("Server started...")
}

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}
