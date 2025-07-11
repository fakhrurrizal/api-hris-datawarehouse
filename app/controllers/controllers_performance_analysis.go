package controllers

import (
	repository "hris-datawarehouse/app/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetDashboardAveragePerformanceScorePerDepartmentWithCount godoc
// @Summary Get Dashboard Average Performance Score Per Department with Count
// @Description Get Dashboard Average Performance Score Per Department with Employee Count
// @Tags PerformanceAnalysis
// @Param department_id query integer false "department_id (int)"
// @Param emp_status_id query integer false "emp_status_id (int)"
// @Param position_id query integer false "position_id (int)"
// @Param manager_id query integer false "manager_id (int)"
// @Param gender query string false "gender (string)"
// @Param state query string false "state (string)"
// @Param start_date query string false "start_date (format: 2006-01-02)"
// @Param end_date query string false "end_date (format: 2006-01-02)"
// @Produce json
// @Success 200
// @Router /v1/dashboard/average-performance-score-per-department [get]
// @Security ApiKeyAuth
// @Security JwtToken
func GetDashboardAveragePerformanceScorePerDepartmentWithCount(c echo.Context) error {
	State := c.QueryParam("state")
	Gender := c.QueryParam("gender")
	EmpStatusID, _ := strconv.Atoi(c.QueryParam("emp_status_id"))
	PositionID, _ := strconv.Atoi(c.QueryParam("position_id"))
	DepartmentID, _ := strconv.Atoi(c.QueryParam("department_id"))
	ManagerID, _ := strconv.Atoi(c.QueryParam("manager_id"))
	startDate := c.QueryParam("start_date")
	endDate := c.QueryParam("end_date")

	startDate, endDate = CheckDate(startDate, endDate)
	data, err := repository.GetAveragePerformanceScorePerDepartmentWithCount(startDate, endDate, EmpStatusID, ManagerID, PositionID, DepartmentID, State, Gender)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, data)
}

// GetDashboardBarChartAverageEmpSatisfactionPerPosition godoc
// @Summary Get Dashboard Bar Chart Average Employee Satisfaction Per Position
// @Description Get Dashboard Bar Chart Average Employee Satisfaction Per Position
// @Tags PerformanceAnalysis
// @Param department_id query integer false "department_id (int)"
// @Param emp_status_id query integer false "emp_status_id (int)"
// @Param position_id query integer false "position_id (int)"
// @Param gender query string false "gender (string)"
// @Param state query string false "state (string)"
// @Param start_date query string false "start_date (format: 2006-01-02)"
// @Param end_date query string false "end_date (format: 2006-01-02)"
// @Produce json
// @Success 200
// @Router /v1/dashboard/heatmap-average-emp-satisfaction-per-position [get]
// @Security ApiKeyAuth
// @Security JwtToken
func GetDashboardBarchartAverageEmpSatisfactionPerPosition(c echo.Context) error {
	State := c.QueryParam("state")
	Gender := c.QueryParam("gender")
	EmpStatusID, _ := strconv.Atoi(c.QueryParam("emp_status_id"))
	PositionID, _ := strconv.Atoi(c.QueryParam("position_id"))
	DepartmentID, _ := strconv.Atoi(c.QueryParam("department_id"))
	startDate := c.QueryParam("start_date")
	endDate := c.QueryParam("end_date")

	startDate, endDate = CheckDate(startDate, endDate)
	data, err := repository.GetSatisfactionHeatmapByPosition(startDate, endDate, EmpStatusID, 0, PositionID, DepartmentID, State, Gender)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, data)
}

// GetEmployeeByMaritalStatus godoc
// @Summary Get Total Employees per Marital Status
// @Description Get jumlah karyawan berdasarkan status pernikahan
// @Tags Dashboard
// @Param department_id query int false "Department ID"
// @Param emp_status_id query int false "Employee Status ID"
// @Param position_id query int false "Position ID"
// @Param manager_id query int false "Manager ID"
// @Param marital_status_id query int false "Marital Status ID"
// @Param state query string false "State"
// @Param start_date query string false "Start Date (format: 2006-01-02)"
// @Param end_date query string false "End Date (format: 2006-01-02)"
// @Produce json
// @Success 200 {array} map[string]interface{}
// @Router /v1/dashboard/employees-per-marital-status [get]
// @Security ApiKeyAuth
// @Security JwtToken
func GetEmployeeByMaritalStatus(c echo.Context) error {
	State := c.QueryParam("state")
	EmpStatusID, _ := strconv.Atoi(c.QueryParam("emp_status_id"))
	PositionID, _ := strconv.Atoi(c.QueryParam("position_id"))
	DepartmentID, _ := strconv.Atoi(c.QueryParam("department_id"))
	ManagerID, _ := strconv.Atoi(c.QueryParam("manager_id"))
	MaritalStatusID, _ := strconv.Atoi(c.QueryParam("marital_status_id"))
	startDate := c.QueryParam("start_date")
	endDate := c.QueryParam("end_date")

	startDate, endDate = CheckDate(startDate, endDate)

	data, err := repository.GetEmployeeCountByMaritalStatus(
		startDate, endDate,
		EmpStatusID, ManagerID, PositionID, DepartmentID, MaritalStatusID, State,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, data)
}
