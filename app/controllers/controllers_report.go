package controllers

import (
	repository "hris-datawarehouse/app/repositories"
	"hris-datawarehouse/app/reqres"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

// GetDashboardScoreCard godoc
// @Summary Get Dashboard Score Card
// @Description Get Dashboard Score Card
// @Tags Dashboard
// @Param department_id query integer false "department_id (int)"
// @Param gender query string false "gender (string) (m/f)"
// @Param emp_status_id query integer false "emp_status_id (int)"
// @Param position_id query integer false "position_id (int)"
// @Param manager_id query integer false "manager_id (int)"
// @Param start_date query string false "start_date (format: 2006-01-02)"
// @Param end_date query string false "end_date (format: 2006-01-02)"
// @Produce json
// @Success 200
// @Router /v1/dashboard/score-card [get]
// @Security ApiKeyAuth
// @Security JwtToken
func GetDashboardScoreCard(c echo.Context) error {

	departmentID, _ := strconv.Atoi(c.QueryParam("department_id"))
	Gender := c.QueryParam("gender")
	EmpStatusID, _ := strconv.Atoi(c.QueryParam("emp_status_id"))
	PositionID, _ := strconv.Atoi(c.QueryParam("position_id"))
	ManagerID, _ := strconv.Atoi(c.QueryParam("manager_id"))
	startDate := c.QueryParam("start_date")
	endDate := c.QueryParam("end_date")

	startDate, endDate = CheckDate(startDate, endDate)

	var data reqres.DashboardScoreCard
	data.TotalEmployee, _ = repository.GetTotalActiveEmployees(startDate, endDate, departmentID, Gender, EmpStatusID, ManagerID, PositionID, "")
	data.TurnoverPercentage, _ = repository.GetTurnoverRate(startDate, endDate, departmentID, Gender, EmpStatusID, ManagerID, PositionID, "")
	data.AveragePerformance, _ = repository.GetAveragePerformanceScore(departmentID, Gender, EmpStatusID, ManagerID, PositionID, "")
	data.AverageDaysLateLast30, _ = repository.GetAverageDaysLateLast30(departmentID, Gender, EmpStatusID, ManagerID, PositionID, "")
	return c.JSON(http.StatusOK, data)
}

func CheckDate(startDate, endDate string) (string, string) {
	if startDate == "" && endDate == "" {
		return "", ""
	}

	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		location = time.Local
	}

	if startDate == "" {
		startDate = time.Now().In(location).AddDate(0, -1, 0).Format("2006-01-02 15:04:05")
	}
	if endDate == "" {
		endDate = time.Now().In(location).Format("2006-01-02 15:04:05")
	}

	return startDate, endDate
}
