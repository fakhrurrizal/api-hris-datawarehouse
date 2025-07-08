package controllers

import (
	repository "hris-datawarehouse/app/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Controllers untuk Line Chart Dashboard

// GetDashboardLineChartRecruitmentTrend godoc
// @Summary Get Dashboard Line Chart Recruitment Trend
// @Description Get Dashboard Line Chart Recruitment Trend per Month/Year
// @Tags Dashboard
// @Param department_id query integer false "department_id (int)"
// @Param emp_status_id query integer false "emp_status_id (int)"
// @Param position_id query integer false "position_id (int)"
// @Param manager_id query integer false "manager_id (int)"
// @Param start_date query string false "start_date (format: 2006-01-02)"
// @Param end_date query string false "end_date (format: 2006-01-02)"
// @Param period_type query string false "period_type (month/year)" default(month)
// @Produce json
// @Success 200
// @Router /v1/dashboard/linechart-recruitment-trend [get]
// @Security ApiKeyAuth
// @Security JwtToken
func GetDashboardLineChartRecruitmentTrend(c echo.Context) error {
	State := c.QueryParam("state")
	EmpStatusID, _ := strconv.Atoi(c.QueryParam("emp_status_id"))
	PositionID, _ := strconv.Atoi(c.QueryParam("position_id"))
	DepartmentID, _ := strconv.Atoi(c.QueryParam("department_id"))
	ManagerID, _ := strconv.Atoi(c.QueryParam("manager_id"))
	startDate := c.QueryParam("start_date")
	endDate := c.QueryParam("end_date")
	periodType := c.QueryParam("period_type")

	if periodType == "" {
		periodType = "month"
	}

	startDate, endDate = CheckDate(startDate, endDate)

	data, _ := repository.GetRecruitmentTrend(startDate, endDate, EmpStatusID, ManagerID, PositionID, DepartmentID, State, periodType)

	return c.JSON(http.StatusOK, data)
}

// GetDashboardLineChartPerformanceTrend godoc
// @Summary Get Dashboard Line Chart Performance Score Trend
// @Description Get Dashboard Line Chart Average Performance Score Trend per Month/Year
// @Tags Dashboard
// @Param department_id query integer false "department_id (int)"
// @Param emp_status_id query integer false "emp_status_id (int)"
// @Param position_id query integer false "position_id (int)"
// @Param manager_id query integer false "manager_id (int)"
// @Param start_date query string false "start_date (format: 2006-01-02)"
// @Param end_date query string false "end_date (format: 2006-01-02)"
// @Param period_type query string false "period_type (month/year)" default(month)
// @Produce json
// @Success 200
// @Router /v1/dashboard/linechart-performance-trend [get]
// @Security ApiKeyAuth
// @Security JwtToken
func GetDashboardLineChartPerformanceTrend(c echo.Context) error {
	State := c.QueryParam("state")
	EmpStatusID, _ := strconv.Atoi(c.QueryParam("emp_status_id"))
	PositionID, _ := strconv.Atoi(c.QueryParam("position_id"))
	DepartmentID, _ := strconv.Atoi(c.QueryParam("department_id"))
	ManagerID, _ := strconv.Atoi(c.QueryParam("manager_id"))
	startDate := c.QueryParam("start_date")
	endDate := c.QueryParam("end_date")
	periodType := c.QueryParam("period_type")

	if periodType == "" {
		periodType = "month"
	}

	startDate, endDate = CheckDate(startDate, endDate)

	data, _ := repository.GetPerformanceTrend(startDate, endDate, EmpStatusID, ManagerID, PositionID, DepartmentID, State, periodType)

	return c.JSON(http.StatusOK, data)
}

// GetDashboardLineChartTurnoverTrend godoc
// @Summary Get Dashboard Line Chart Turnover Trend
// @Description Get Dashboard Line Chart Turnover Trend per Month/Year
// @Tags Dashboard
// @Param department_id query integer false "department_id (int)"
// @Param emp_status_id query integer false "emp_status_id (int)"
// @Param position_id query integer false "position_id (int)"
// @Param manager_id query integer false "manager_id (int)"
// @Param start_date query string false "start_date (format: 2006-01-02)"
// @Param end_date query string false "end_date (format: 2006-01-02)"
// @Param period_type query string false "period_type (month/year)" default(month)
// @Produce json
// @Success 200
// @Router /v1/dashboard/linechart-turnover-trend [get]
// @Security ApiKeyAuth
// @Security JwtToken
func GetDashboardLineChartTurnoverTrend(c echo.Context) error {
	State := c.QueryParam("state")
	EmpStatusID, _ := strconv.Atoi(c.QueryParam("emp_status_id"))
	PositionID, _ := strconv.Atoi(c.QueryParam("position_id"))
	DepartmentID, _ := strconv.Atoi(c.QueryParam("department_id"))
	ManagerID, _ := strconv.Atoi(c.QueryParam("manager_id"))
	startDate := c.QueryParam("start_date")
	endDate := c.QueryParam("end_date")
	periodType := c.QueryParam("period_type")

	if periodType == "" {
		periodType = "month"
	}

	startDate, endDate = CheckDate(startDate, endDate)

	data, _ := repository.GetTurnoverTrend(startDate, endDate, EmpStatusID, ManagerID, PositionID, DepartmentID, State, periodType)

	return c.JSON(http.StatusOK, data)
}

// GetDashboardLineChartLateAbsenceTrend godoc
// @Summary Get Dashboard Line Chart Late/Absence Trend
// @Description Get Dashboard Line Chart Late/Absence Trend per Month/Year
// @Tags Dashboard
// @Param department_id query integer false "department_id (int)"
// @Param emp_status_id query integer false "emp_status_id (int)"
// @Param position_id query integer false "position_id (int)"
// @Param manager_id query integer false "manager_id (int)"
// @Param start_date query string false "start_date (format: 2006-01-02)"
// @Param end_date query string false "end_date (format: 2006-01-02)"
// @Param period_type query string false "period_type (month/year)" default(month)
// @Param trend_type query string false "trend_type (late/absence/both)" default(both)
// @Produce json
// @Success 200
// @Router /v1/dashboard/linechart-late-absence-trend [get]
// @Security ApiKeyAuth
// @Security JwtToken
func GetDashboardLineChartLateAbsenceTrend(c echo.Context) error {
	State := c.QueryParam("state")
	EmpStatusID, _ := strconv.Atoi(c.QueryParam("emp_status_id"))
	PositionID, _ := strconv.Atoi(c.QueryParam("position_id"))
	DepartmentID, _ := strconv.Atoi(c.QueryParam("department_id"))
	ManagerID, _ := strconv.Atoi(c.QueryParam("manager_id"))
	startDate := c.QueryParam("start_date")
	endDate := c.QueryParam("end_date")
	periodType := c.QueryParam("period_type")
	trendType := c.QueryParam("trend_type")

	if periodType == "" {
		periodType = "month"
	}
	if trendType == "" {
		trendType = "both"
	}

	startDate, endDate = CheckDate(startDate, endDate)

	data, _ := repository.GetLateAbsenceTrend(startDate, endDate, EmpStatusID, ManagerID, PositionID, DepartmentID, State, periodType, trendType)

	return c.JSON(http.StatusOK, data)
}
