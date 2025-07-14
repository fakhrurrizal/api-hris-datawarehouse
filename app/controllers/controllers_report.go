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
// @Param state query string false "state (string)"
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
	State := c.QueryParam("state")
	EmpStatusID, _ := strconv.Atoi(c.QueryParam("emp_status_id"))
	PositionID, _ := strconv.Atoi(c.QueryParam("position_id"))
	ManagerID, _ := strconv.Atoi(c.QueryParam("manager_id"))
	startDate := c.QueryParam("start_date")
	endDate := c.QueryParam("end_date")

	startDate, endDate = CheckDate(startDate, endDate)

	var data reqres.DashboardScoreCard
	data.TotalEmployee, _ = repository.GetTotalActiveEmployees(startDate, endDate, departmentID, Gender, EmpStatusID, ManagerID, PositionID, State)
	data.TurnoverPercentage, _ = repository.GetTurnoverRate(startDate, endDate, departmentID, Gender, EmpStatusID, ManagerID, PositionID, State)
	data.AveragePerformance, _ = repository.GetAveragePerformanceScore(departmentID, Gender, EmpStatusID, ManagerID, PositionID, State)
	data.AverageDaysLateLast30, _ = repository.GetAverageDaysLateLast30(departmentID, Gender, EmpStatusID, ManagerID, PositionID, State)
	data.GetTotalSalaryExpenditure, _ = repository.GetTotalSalaryExpenditure(departmentID, Gender, EmpStatusID, ManagerID, PositionID, State)
	return c.JSON(http.StatusOK, data)
}

// GetDashboardBarChart godoc
// @Summary Get Dashboard Bar Chart Per Department
// @Description Get Dashboard Bar Chart Per Department
// @Tags Dashboard
// @Param gender query string false "gender (string) (m/f)"
// @Param emp_status_id query integer false "emp_status_id (int)"
// @Param position_id query integer false "position_id (int)"
// @Param manager_id query integer false "manager_id (int)"
// @Param start_date query string false "start_date (format: 2006-01-02)"
// @Param end_date query string false "end_date (format: 2006-01-02)"
// @Produce json
// @Success 200
// @Router /v1/dashboard/barchart-employee-per-department [get]
// @Security ApiKeyAuth
// @Security JwtToken
func GetDashboardBarChartEmployeePerDepartment(c echo.Context) error {

	Gender := c.QueryParam("gender")
	State := c.QueryParam("state")
	EmpStatusID, _ := strconv.Atoi(c.QueryParam("emp_status_id"))
	PositionID, _ := strconv.Atoi(c.QueryParam("position_id"))
	ManagerID, _ := strconv.Atoi(c.QueryParam("manager_id"))
	startDate := c.QueryParam("start_date")
	endDate := c.QueryParam("end_date")

	startDate, endDate = CheckDate(startDate, endDate)

	data, _ := repository.GetEmployeeCountPerDepartment(startDate, endDate, EmpStatusID, ManagerID, PositionID, State, Gender)

	return c.JSON(http.StatusOK, data)
}

// GetDashboardBarChart godoc
// @Summary Get Dashboard Bar Chart Per Department
// @Description Get Dashboard Bar Chart Per Department
// @Tags Dashboard
// @Param department_id query integer false "department_id (int)"
// @Param emp_status_id query integer false "emp_status_id (int)"
// @Param position_id query integer false "position_id (int)"
// @Param manager_id query integer false "manager_id (int)"
// @Param start_date query string false "start_date (format: 2006-01-02)"
// @Param end_date query string false "end_date (format: 2006-01-02)"
// @Produce json
// @Success 200
// @Router /v1/dashboard/barchart-employee-per-gender [get]
// @Security ApiKeyAuth
// @Security JwtToken
func GetDashboardBarChartEmployeePerGender(c echo.Context) error {

	State := c.QueryParam("state")
	EmpStatusID, _ := strconv.Atoi(c.QueryParam("emp_status_id"))
	PositionID, _ := strconv.Atoi(c.QueryParam("position_id"))
	DepartmentID, _ := strconv.Atoi(c.QueryParam("department_id"))
	ManagerID, _ := strconv.Atoi(c.QueryParam("manager_id"))
	startDate := c.QueryParam("start_date")
	endDate := c.QueryParam("end_date")

	startDate, endDate = CheckDate(startDate, endDate)

	data, _ := repository.GetEmployeeCountPerGender(startDate, endDate, EmpStatusID, ManagerID, PositionID, DepartmentID, State)

	return c.JSON(http.StatusOK, data)
}

// GetDashboardBarChartEmployeePerRecruitmentSource godoc
// @Summary Get Dashboard Bar Chart Per Recruitment Source
// @Description Get Dashboard Bar Chart Per Recruitment Source
// @Tags Dashboard
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
// @Router /v1/dashboard/barchart-employee-per-recruitment-source [get]
// @Security ApiKeyAuth
// @Security JwtToken
func GetDashboardBarChartEmployeePerRecruitmentSource(c echo.Context) error {

	State := c.QueryParam("state")
	Gender := c.QueryParam("gender")
	EmpStatusID, _ := strconv.Atoi(c.QueryParam("emp_status_id"))
	PositionID, _ := strconv.Atoi(c.QueryParam("position_id"))
	DepartmentID, _ := strconv.Atoi(c.QueryParam("department_id"))
	ManagerID, _ := strconv.Atoi(c.QueryParam("manager_id"))
	startDate := c.QueryParam("start_date")
	endDate := c.QueryParam("end_date")

	startDate, endDate = CheckDate(startDate, endDate)

	data, _ := repository.GetEmployeeCountPerRecruitmentSource(startDate, endDate, EmpStatusID, ManagerID, PositionID, DepartmentID, State, Gender)

	return c.JSON(http.StatusOK, data)
}

// GetDashboardBarChartEmployeePerCitizenDesc godoc
// @Summary Get Dashboard Bar Chart Per Citizen Description
// @Description Get Dashboard Bar Chart Per Citizen Description
// @Tags Dashboard
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
// @Router /v1/dashboard/barchart-employee-per-citizen-desc [get]
// @Security ApiKeyAuth
// @Security JwtToken
func GetDashboardBarChartEmployeePerCitizenDesc(c echo.Context) error {

	State := c.QueryParam("state")
	Gender := c.QueryParam("gender")
	EmpStatusID, _ := strconv.Atoi(c.QueryParam("emp_status_id"))
	PositionID, _ := strconv.Atoi(c.QueryParam("position_id"))
	DepartmentID, _ := strconv.Atoi(c.QueryParam("department_id"))
	ManagerID, _ := strconv.Atoi(c.QueryParam("manager_id"))
	startDate := c.QueryParam("start_date")
	endDate := c.QueryParam("end_date")

	startDate, endDate = CheckDate(startDate, endDate)

	data, _ := repository.GetEmployeeCountPerCitizenDesc(startDate, endDate, EmpStatusID, ManagerID, PositionID, DepartmentID, State, Gender)

	return c.JSON(http.StatusOK, data)
}

// GetDashboardBarChartEmployeePerRaceDesc godoc
// @Summary Get Dashboard Bar Chart Per Race Description
// @Description Get Dashboard Bar Chart Per Race Description
// @Tags Dashboard
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
// @Router /v1/dashboard/barchart-employee-per-race-desc [get]
// @Security ApiKeyAuth
// @Security JwtToken
func GetDashboardBarChartEmployeePerRaceDesc(c echo.Context) error {
	State := c.QueryParam("state")
	Gender := c.QueryParam("gender")
	EmpStatusID, _ := strconv.Atoi(c.QueryParam("emp_status_id"))
	PositionID, _ := strconv.Atoi(c.QueryParam("position_id"))
	DepartmentID, _ := strconv.Atoi(c.QueryParam("department_id"))
	ManagerID, _ := strconv.Atoi(c.QueryParam("manager_id"))
	startDate := c.QueryParam("start_date")
	endDate := c.QueryParam("end_date")

	startDate, endDate = CheckDate(startDate, endDate)
	data, _ := repository.GetEmployeeCountPerRaceDesc(startDate, endDate, EmpStatusID, ManagerID, PositionID, DepartmentID, State, Gender)

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

// GetDashboardPieChartMaritalStatusRatio godoc
// @Summary Get Dashboard Pie Chart Employee Marital Status Ratio
// @Tags MaritalAnalysis
// @Param department_id query integer false "department_id"
// @Param emp_status_id query integer false "emp_status_id"
// @Param position_id query integer false "position_id"
// @Param manager_id query integer false "manager_id"
// @Param gender query string false "gender"
// @Param state query string false "state"
// @Produce json
// @Success 200
// @Router /v1/dashboard/piechart-employee-marital-ratio [get]
// @Security ApiKeyAuth
// @Security JwtToken
func GetDashboardPieChartMaritalStatusRatio(c echo.Context) error {
	state := c.QueryParam("state")
	gender := c.QueryParam("gender")
	empStatusID, _ := strconv.Atoi(c.QueryParam("emp_status_id"))
	positionID, _ := strconv.Atoi(c.QueryParam("position_id"))
	managerID, _ := strconv.Atoi(c.QueryParam("manager_id"))
	departmentID, _ := strconv.Atoi(c.QueryParam("department_id"))

	data, err := repository.GetEmployeeMaritalRatio(departmentID, empStatusID, managerID, positionID, state, gender)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, data)
}

// GetDashboardPieChartAgeRatio godoc
// @Summary Get Dashboard Pie Chart Employee Age Ratio
// @Tags AgeAnalysis
// @Param department_id query integer false "department_id"
// @Param emp_status_id query integer false "emp_status_id"
// @Param position_id query integer false "position_id"
// @Param manager_id query integer false "manager_id"
// @Param gender query string false "gender"
// @Param state query string false "state"
// @Produce json
// @Success 200
// @Router /v1/dashboard/piechart-employee-age-ratio [get]
// @Security ApiKeyAuth
// @Security JwtToken
func GetDashboardPieChartAgeRatio(c echo.Context) error {
	state := c.QueryParam("state")
	gender := c.QueryParam("gender")
	empStatusID, _ := strconv.Atoi(c.QueryParam("emp_status_id"))
	positionID, _ := strconv.Atoi(c.QueryParam("position_id"))
	managerID, _ := strconv.Atoi(c.QueryParam("manager_id"))
	departmentID, _ := strconv.Atoi(c.QueryParam("department_id"))

	data, err := repository.GetEmployeeAgeRatio(departmentID, empStatusID, managerID, positionID, state, gender)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, data)
}
