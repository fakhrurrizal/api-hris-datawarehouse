package controllers

import (
	repository "hris-datawarehouse/app/repositories"
	"hris-datawarehouse/app/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type EmploymentPagingResponse struct {
	Page  int                  `json:"page"`
	Limit int                  `json:"limit"`
	Total int64                `json:"total"`
	Data  []EmploymentResponse `json:"data"`
}

type EmploymentResponse struct {
	EmpID        int     `json:"emp_id"`
	EmployeeName string  `json:"employee_name"`
	Position     string  `json:"position"`
	Department   string  `json:"department"`
	Manager      string  `json:"manager"`
	Salary       float64 `json:"salary"`
	DateOfHire   string  `json:"date_of_hire"`
}

// GetDimDepartment godoc
// @Summary Get All Department With Pagination
// @Description Get All Department With Pagination
// @Tags Dim Table
// @Param search query string false "search (string)"
// @Param page query integer false "page (int)"
// @Param limit query integer false "limit (int)"
// @Param sort query string false "sort (ASC/DESC)"
// @Param order query string false "order by (default: id)"
// @Param created_at_margin_top query string false "created_at_margin_top (format: 2006-01-02)"
// @Param created_at_margin_bottom query string false "created_at_margin_top (format: 2006-01-02)"
// @Produce json
// @Success 200
// @Router /v1/dim/department [get]
// @Security ApiKeyAuth
// @Security JwtToken
func GetDimDepartment(c echo.Context) error {
	// companyID, _ := strconv.Atoi(c.QueryParam("company_id"))
	// plantID, _ := strconv.Atoi(c.QueryParam("plant_id"))
	// userID, _ := strconv.Atoi(c.QueryParam("user_id"))
	param := utils.PopulatePaging(c, "status")

	data, err := repository.GetDepartments(param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get departments",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, data)
}

// GetDimEmployee godoc
// @Summary Get All Employees With Pagination
// @Description Get All Employees With Pagination
// @Tags Dim Table
// @Param search query string false "search (string)"
// @Param page query integer false "page (int)"
// @Param limit query integer false "limit (int)"
// @Param sort query string false "sort (ASC/DESC)"
// @Param order query string false "order by (default: EmpID)"
// @Produce json
// @Success 200
// @Router /v1/dim/employee [get]
// @Security ApiKeyAuth
// @Security JwtToken
func GetDimEmployee(c echo.Context) error {
	param := utils.PopulatePaging(c, "status")

	data, err := repository.GetEmployees(param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get employees",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, data)
}

// GetDimEmploymentStatus godoc
// @Summary Get All Employment Statuses With Pagination
// @Description Get All Employment Statuses With Pagination
// @Tags Dim Table
// @Param search query string false "search (string)"
// @Param page query integer false "page (int)"
// @Param limit query integer false "limit (int)"
// @Param sort query string false "sort (ASC/DESC)"
// @Param order query string false "order by (default: EmpStatusID)"
// @Produce json
// @Success 200
// @Router /v1/dim/employment-status [get]
// @Security ApiKeyAuth
// @Security JwtToken
func GetDimEmploymentStatus(c echo.Context) error {
	param := utils.PopulatePaging(c, "status")

	data, err := repository.GetEmploymentStatuses(param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get employment statuses",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, data)
}

// GetDimManager godoc
// @Summary Get All Managers With Pagination
// @Description Get All Managers With Pagination
// @Tags Dim Table
// @Param search query string false "search (string)"
// @Param page query integer false "page (int)"
// @Param limit query integer false "limit (int)"
// @Param sort query string false "sort (ASC/DESC)"
// @Param order query string false "order by (default: ManagerID)"
// @Produce json
// @Success 200
// @Router /v1/dim/manager [get]
// @Security ApiKeyAuth
// @Security JwtToken
func GetDimManager(c echo.Context) error {
	param := utils.PopulatePaging(c, "status")

	data, err := repository.GetManagers(param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get managers",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, data)
}

// GetDimMaritalStatus godoc
// @Summary Get All Marital Statuses With Pagination
// @Description Get All Marital Statuses With Pagination
// @Tags Dim Table
// @Param search query string false "search (string)"
// @Param page query integer false "page (int)"
// @Param limit query integer false "limit (int)"
// @Param sort query string false "sort (ASC/DESC)"
// @Param order query string false "order by (default: MaritalStatusID)"
// @Produce json
// @Success 200
// @Router /v1/dim/marital-status [get]
// @Security ApiKeyAuth
// @Security JwtToken
func GetDimMaritalStatus(c echo.Context) error {
	param := utils.PopulatePaging(c, "status")

	data, err := repository.GetMaritalStatuses(param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get marital statuses",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, data)
}

// GetDimPerformance godoc
// @Summary Get All Performance Scores With Pagination
// @Description Get All Performance Scores With Pagination
// @Tags Dim Table
// @Param search query string false "search (string)"
// @Param page query integer false "page (int)"
// @Param limit query integer false "limit (int)"
// @Param sort query string false "sort (ASC/DESC)"
// @Param order query string false "order by (default: PerfScoreID)"
// @Produce json
// @Success 200
// @Router /v1/dim/performance [get]
// @Security ApiKeyAuth
// @Security JwtToken
func GetDimPerformance(c echo.Context) error {
	param := utils.PopulatePaging(c, "status")

	data, err := repository.GetPerformances(param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get performance scores",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, data)
}

// GetDimPosition godoc
// @Summary Get All Positions With Pagination
// @Description Get All Positions With Pagination
// @Tags Dim Table
// @Param search query string false "search (string)"
// @Param page query integer false "page (int)"
// @Param limit query integer false "limit (int)"
// @Param sort query string false "sort (ASC/DESC)"
// @Param order query string false "order by (default: PositionID)"
// @Produce json
// @Success 200
// @Router /v1/dim/position [get]
// @Security ApiKeyAuth
// @Security JwtToken
func GetDimPosition(c echo.Context) error {
	param := utils.PopulatePaging(c, "status")

	data, err := repository.GetPositions(param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get positions",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, data)
}

// GetEmploymentWithFilters godoc
// @Summary Get All Employment Data With Filters
// @Description Get Employment Data with Filters and Pagination
// @Tags Fact Table
// @Param search query string false "Search by Employee Name"
// @Param gender query string false "Gender"
// @Param state query string false "State"
// @Param department_id query integer false "Department ID"
// @Param position_id query integer false "Position ID"
// @Param emp_status_id query integer false "Employee Status ID"
// @Param manager_id query integer false "Manager ID"
// @Param start_date query string false "Start Date (YYYY-MM-DD)"
// @Param end_date query string false "End Date (YYYY-MM-DD)"
// @Param page query integer false "Page (int)"
// @Param limit query integer false "Limit (int)"
// @Param sort query string false "Sort (ASC/DESC)"
// @Param order query string false "Order by field (default: e.Employee_Name)"
// @Produce json
// @Success 200
// @Router /v1/dim/fact-employment [get]
// @Security ApiKeyAuth
// @Security JwtToken
func GetEmploymentWithFilters(c echo.Context) error {
	param := utils.PopulatePaging(c, "e.Employee_Name")

	// Ambil semua query param
	gender := c.QueryParam("gender")
	state := c.QueryParam("state")
	startDate := c.QueryParam("start_date")
	endDate := c.QueryParam("end_date")

	// Parsing integer, jika error dianggap 0
	departmentID, _ := strconv.Atoi(c.QueryParam("department_id"))
	positionID, _ := strconv.Atoi(c.QueryParam("position_id"))
	empStatusID, _ := strconv.Atoi(c.QueryParam("emp_status_id"))
	managerID, _ := strconv.Atoi(c.QueryParam("manager_id"))

	param.Custom = map[string]interface{}{
		"gender":        gender,
		"state":         state,
		"start_date":    startDate,
		"end_date":      endDate,
		"position_id":   positionID,
		"emp_status_id": empStatusID,
		"manager_id":    managerID,
	}

	data, err := repository.GetEmploymentWithFilters(departmentID, param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get employment data",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, data)
}
