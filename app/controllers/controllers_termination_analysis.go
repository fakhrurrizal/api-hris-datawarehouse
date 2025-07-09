package controllers

import (
	repository "hris-datawarehouse/app/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Controller functions for employee termination analysis

// GetDashboardBarChartEmployeeTerminationByReason godoc
// @Summary Get Dashboard Bar Chart Employee Termination By Reason
// @Description Get Dashboard Bar Chart Employee Termination By Reason
// @Tags TerminationAnalysis
// @Param department_id query integer false "department_id (int)"
// @Param emp_status_id query integer false "emp_status_id (int)"
// @Param position_id query integer false "position_id (int)"
// @Param gender query string false "gender (string)"
// @Param state query string false "state (string)"
// @Param start_date query string false "start_date (format: 2006-01-02)"
// @Param end_date query string false "end_date (format: 2006-01-02)"
// @Produce json
// @Success 200
// @Router /v1/dashboard/barchart-employee-termination-by-reason [get]
// @Security ApiKeyAuth
// @Security JwtToken
func GetDashboardBarChartEmployeeTerminationByReason(c echo.Context) error {
	State := c.QueryParam("state")
	Gender := c.QueryParam("gender")
	EmpStatusID, _ := strconv.Atoi(c.QueryParam("emp_status_id"))
	PositionID, _ := strconv.Atoi(c.QueryParam("position_id"))
	DepartmentID, _ := strconv.Atoi(c.QueryParam("department_id"))
	startDate := c.QueryParam("start_date")
	endDate := c.QueryParam("end_date")

	startDate, endDate = CheckDate(startDate, endDate)
	data, err := repository.GetEmployeeTerminationByReason(startDate, endDate, EmpStatusID, 0, PositionID, DepartmentID, State, Gender)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, data)
}

// GetDashboardBarChartEmployeeTerminationByDepartment godoc
// @Summary Get Dashboard Bar Chart Employee Termination By Department
// @Description Get Dashboard Bar Chart Employee Termination By Department
// @Tags TerminationAnalysis
// @Param department_id query integer false "department_id (int)"
// @Param emp_status_id query integer false "emp_status_id (int)"
// @Param position_id query integer false "position_id (int)"
// @Param gender query string false "gender (string)"
// @Param state query string false "state (string)"
// @Param start_date query string false "start_date (format: 2006-01-02)"
// @Param end_date query string false "end_date (format: 2006-01-02)"
// @Produce json
// @Success 200
// @Router /v1/dashboard/barchart-employee-termination-by-department [get]
// @Security ApiKeyAuth
// @Security JwtToken
func GetDashboardBarChartEmployeeTerminationByDepartment(c echo.Context) error {
	State := c.QueryParam("state")
	Gender := c.QueryParam("gender")
	EmpStatusID, _ := strconv.Atoi(c.QueryParam("emp_status_id"))
	PositionID, _ := strconv.Atoi(c.QueryParam("position_id"))
	DepartmentID, _ := strconv.Atoi(c.QueryParam("department_id"))
	startDate := c.QueryParam("start_date")
	endDate := c.QueryParam("end_date")

	startDate, endDate = CheckDate(startDate, endDate)
	data, err := repository.GetEmployeeTerminationByDepartment(startDate, endDate, EmpStatusID, 0, PositionID, DepartmentID, State, Gender)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, data)
}

// 3. Controller untuk Rasio Voluntarily Terminated vs Terminated for Cause
// GetDashboardPieChartEmployeeTerminationRatio godoc
// @Summary Get Dashboard Pie Chart Employee Termination Ratio
// @Description Get Dashboard Pie Chart Employee Termination Ratio (Voluntarily vs Terminated for Cause)
// @Tags TerminationAnalysis
// @Param department_id query integer false "department_id (int)"
// @Param emp_status_id query integer false "emp_status_id (int)"
// @Param position_id query integer false "position_id (int)"
// @Param gender query string false "gender (string)"
// @Param state query string false "state (string)"
// @Param start_date query string false "start_date (format: 2006-01-02)"
// @Param end_date query string false "end_date (format: 2006-01-02)"
// @Produce json
// @Success 200
// @Router /v1/dashboard/piechart-employee-termination-ratio [get]
// @Security ApiKeyAuth
// @Security JwtToken
func GetDashboardPieChartEmployeeTerminationRatio(c echo.Context) error {
	State := c.QueryParam("state")
	Gender := c.QueryParam("gender")
	EmpStatusID, _ := strconv.Atoi(c.QueryParam("emp_status_id"))
	PositionID, _ := strconv.Atoi(c.QueryParam("position_id"))
	managerID, _ := strconv.Atoi(c.QueryParam("manager_id"))
	DepartmentID, _ := strconv.Atoi(c.QueryParam("department_id"))
	startDate := c.QueryParam("start_date")
	endDate := c.QueryParam("end_date")

	startDate, endDate = CheckDate(startDate, endDate)
	data, err := repository.GetEmployeeTerminationRatio(startDate, endDate, EmpStatusID, managerID, PositionID, DepartmentID, State, Gender)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, data)
}
