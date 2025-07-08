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
			dashboard.GET("/barchart-employee-per-department", controllers.GetDashboardBarChartEmployeePerDepartment, middlewares.Auth())
			dashboard.GET("/barchart-employee-per-gender", controllers.GetDashboardBarChartEmployeePerGender, middlewares.Auth())
			dashboard.GET("/barchart-employee-per-recruitment-source", controllers.GetDashboardBarChartEmployeePerRecruitmentSource, middlewares.Auth())
			dashboard.GET("/barchart-employee-per-citizen-desc", controllers.GetDashboardBarChartEmployeePerCitizenDesc, middlewares.Auth())
			dashboard.GET("/barchart-employee-per-race-desc", controllers.GetDashboardBarChartEmployeePerRaceDesc, middlewares.Auth())
			dashboard.GET("/barchart-average-salary-per-department", controllers.GetDashboardBarChartAverageSalaryPerDepartment, middlewares.Auth())
			dashboard.GET("/barchart-average-salary-per-position", controllers.GetDashboardBarChartAverageSalaryPerPosition, middlewares.Auth())
			dashboard.GET("/highest-lowest-salary", controllers.GetDashboardHighestLowestSalary, middlewares.Auth())
			dashboard.GET("/average-performance-score-per-department", controllers.GetDashboardAveragePerformanceScorePerDepartmentWithCount, middlewares.Auth())
			dashboard.GET("/barchart-average-emp-satisfaction-per-position", controllers.GetDashboardBarChartAverageEmpSatisfactionPerPosition, middlewares.Auth())
			dashboard.GET("/barchart-employee-termination-by-reason", controllers.GetDashboardBarChartEmployeeTerminationByReason, middlewares.Auth())
			dashboard.GET("/barchart-employee-termination-by-department", controllers.GetDashboardBarChartEmployeeTerminationByDepartment, middlewares.Auth())
			dashboard.GET("/piechart-employee-termination-ratio", controllers.GetDashboardPieChartEmployeeTerminationRatio, middlewares.Auth())
			dashboard.GET("/linechart-recruitment-trend", controllers.GetDashboardLineChartRecruitmentTrend, middlewares.Auth())
			dashboard.GET("/linechart-performance-trend", controllers.GetDashboardLineChartPerformanceTrend, middlewares.Auth())
			dashboard.GET("/linechart-turnover-trend", controllers.GetDashboardLineChartTurnoverTrend, middlewares.Auth())
			dashboard.GET("/linechart-late-absence-trend", controllers.GetDashboardLineChartLateAbsenceTrend, middlewares.Auth())
			dashboard.GET("/linechart-late-absence-trend", controllers.GetDashboardLineChartLateAbsenceTrend, middlewares.Auth())

		}
		dim := api.Group("/dim", middlewares.CheckAPIKey())
		{
			dim.GET("/department", controllers.GetDimDepartment, middlewares.Auth())
			dim.GET("/position", controllers.GetDimPosition, middlewares.Auth())
			dim.GET("/employee", controllers.GetDimEmployee, middlewares.Auth())
			dim.GET("/manager", controllers.GetDimManager, middlewares.Auth())
			dim.GET("/performance", controllers.GetDimPerformance, middlewares.Auth())
			dim.GET("/employment-status", controllers.GetDimEmploymentStatus, middlewares.Auth())
			dim.GET("/marital-status", controllers.GetDimMaritalStatus, middlewares.Auth())
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
