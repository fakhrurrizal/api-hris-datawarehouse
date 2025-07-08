package controllers

import (
	repository "hris-datawarehouse/app/repositories"
	"hris-datawarehouse/app/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

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
