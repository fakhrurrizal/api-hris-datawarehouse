package controllers

import (
	repository "hris-datawarehouse/app/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetTop10TerminatedDepartments godoc
// @Summary Get Top 10 Terminated Departments
// @Description Top 10 departemen dengan jumlah karyawan yang paling banyak diberhentikan
// @Tags Dashboard
// @Param emp_status_id query int false "Employee Status ID"
// @Param manager_id query int false "Manager ID"
// @Param position_id query int false "Position ID"
// @Param state query string false "State"
// @Param start_date query string false "Start Date (format: 2006-01-02)"
// @Param end_date query string false "End Date (format: 2006-01-02)"
// @Produce json
// @Success 200 {array} map[string]interface{}
// @Router /v1/dashboard/top-10-termination-departments [get]
// @Security ApiKeyAuth
// @Security JwtToken
func GetTop10TerminatedDepartmentsController(c echo.Context) error {
	startDate := c.QueryParam("start_date")
	endDate := c.QueryParam("end_date")
	state := c.QueryParam("state")

	empStatusID, _ := strconv.Atoi(c.QueryParam("emp_status_id"))
	managerID, _ := strconv.Atoi(c.QueryParam("manager_id"))
	positionID, _ := strconv.Atoi(c.QueryParam("position_id"))

	startDate, endDate = CheckDate(startDate, endDate)

	data, err := repository.GetTop10TerminatedDepartments(
		startDate, endDate,
		empStatusID, managerID, positionID,
		state,
	)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, data)
}
