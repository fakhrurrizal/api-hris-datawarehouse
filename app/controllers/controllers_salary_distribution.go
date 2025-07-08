package controllers

import (
	repository "hris-datawarehouse/app/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetDashboardBarChartAverageSalaryPerDepartment godoc
// @Summary Get SalaryDistribution Bar Chart Average Salary Per Department
// @Description Get SalaryDistribution Bar Chart Average Salary Per Department
// @Tags SalaryDistribution
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
// @Router /v1/dashboard/barchart-average-salary-per-department [get]
// @Security ApiKeyAuth
// @Security JwtToken
func GetDashboardBarChartAverageSalaryPerDepartment(c echo.Context) error {
	State := c.QueryParam("state")
	Gender := c.QueryParam("gender")
	EmpStatusID, _ := strconv.Atoi(c.QueryParam("emp_status_id"))
	PositionID, _ := strconv.Atoi(c.QueryParam("position_id"))
	DepartmentID, _ := strconv.Atoi(c.QueryParam("department_id"))
	ManagerID, _ := strconv.Atoi(c.QueryParam("manager_id"))
	startDate := c.QueryParam("start_date")
	endDate := c.QueryParam("end_date")

	startDate, endDate = CheckDate(startDate, endDate)
	data, _ := repository.GetAverageSalaryPerDepartment(startDate, endDate, EmpStatusID, ManagerID, PositionID, DepartmentID, State, Gender)

	return c.JSON(http.StatusOK, data)
}

// GetDashboardBarChartAverageSalaryPerPositionWithCount godoc
// @Summary Get SalaryDistribution Bar Chart Average Salary Per Position with Employee Count
// @Description Get SalaryDistribution Bar Chart Average Salary Per Position with Employee Count
// @Tags SalaryDistribution
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
// @Router /v1/dashboard/barchart-average-salary-per-position [get]
// @Security ApiKeyAuth
// @Security JwtToken
func GetDashboardBarChartAverageSalaryPerPosition(c echo.Context) error {
	State := c.QueryParam("state")
	Gender := c.QueryParam("gender")
	EmpStatusID, _ := strconv.Atoi(c.QueryParam("emp_status_id"))
	PositionID, _ := strconv.Atoi(c.QueryParam("position_id"))
	DepartmentID, _ := strconv.Atoi(c.QueryParam("department_id"))
	ManagerID, _ := strconv.Atoi(c.QueryParam("manager_id"))
	startDate := c.QueryParam("start_date")
	endDate := c.QueryParam("end_date")

	startDate, endDate = CheckDate(startDate, endDate)
	data, _ := repository.GetAverageSalaryPerPositionWithCount(startDate, endDate, EmpStatusID, ManagerID, PositionID, DepartmentID, State, Gender)

	return c.JSON(http.StatusOK, data)
}

// GetDashboardHighestLowestSalary godoc
// @Summary Get Dashboard Highest and Lowest Salary
// @Description Get Dashboard Highest and Lowest Salary
// @Tags SalaryDistribution
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
// @Router /v1/dashboard/highest-lowest-salary [get]
// @Security ApiKeyAuth
// @Security JwtToken
func GetDashboardHighestLowestSalary(c echo.Context) error {
	State := c.QueryParam("state")
	Gender := c.QueryParam("gender")
	EmpStatusID, _ := strconv.Atoi(c.QueryParam("emp_status_id"))
	PositionID, _ := strconv.Atoi(c.QueryParam("position_id"))
	DepartmentID, _ := strconv.Atoi(c.QueryParam("department_id"))
	ManagerID, _ := strconv.Atoi(c.QueryParam("manager_id"))
	startDate := c.QueryParam("start_date")
	endDate := c.QueryParam("end_date")

	startDate, endDate = CheckDate(startDate, endDate)
	data, err := repository.GetHighestAndLowestSalary(startDate, endDate, EmpStatusID, ManagerID, PositionID, DepartmentID, State, Gender)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, data)
}
