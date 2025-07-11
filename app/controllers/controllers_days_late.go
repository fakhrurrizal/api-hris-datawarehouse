package controllers

import (
	repository "hris-datawarehouse/app/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetHeatmapLateByDepartment godoc
// @Summary Get Heatmap Days Late per Department
// @Description Heatmap keterlambatan masuk berdasarkan hari dan departemen
// @Tags Dashboard
// @Param department_id query int false "Department ID"
// @Param emp_status_id query int false "Employee Status ID"
// @Param position_id query int false "Position ID"
// @Param manager_id query int false "Manager ID"
// @Param state query string false "State"
// @Param start_date query string false "Start Date (format: 2006-01-02)"
// @Param end_date query string false "End Date (format: 2006-01-02)"
// @Produce json
// @Success 200 {array} map[string]interface{}
// @Router /v1/dashboard/heatmap-days-late [get]
// @Security ApiKeyAuth
// @Security JwtToken
func GetMonthlyHeatmapLateByDepartment(c echo.Context) error {
	State := c.QueryParam("state")
	EmpStatusID, _ := strconv.Atoi(c.QueryParam("emp_status_id"))
	PositionID, _ := strconv.Atoi(c.QueryParam("position_id"))
	DepartmentID, _ := strconv.Atoi(c.QueryParam("department_id"))
	ManagerID, _ := strconv.Atoi(c.QueryParam("manager_id"))
	startDate := c.QueryParam("start_date")
	endDate := c.QueryParam("end_date")

	// Kalau kamu ingin, bisa tetap pakai CheckDate atau biarkan kosong
	startDate, endDate = CheckDate(startDate, endDate)

	data, err := repository.GetLateHeatmapMonthlyByDepartment(
		startDate, endDate,
		EmpStatusID, ManagerID, PositionID, DepartmentID, State,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, data)
}
