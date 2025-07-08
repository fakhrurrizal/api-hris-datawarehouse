package repository

import (
	"fmt"
	"hris-datawarehouse/config"
)

// Repository untuk Line Chart Dashboard

type LineChartData struct {
	X     string  `json:"x"`
	Y     float64 `json:"y"`
	Label string  `json:"label,omitempty"`
}

type LineChartResponse struct {
	Data []LineChartData `json:"data"`
}

func GetRecruitmentTrend(
	startDate, endDate string,
	empStatusID, managerID, positionID, deptID int,
	state, periodType string,
) (result []LineChartData, err error) {
	var dateColumn string

	if periodType == "year" {
		dateColumn = "YEAR(e.DateofHire)"
	} else {
		dateColumn = "DATE_FORMAT(e.DateofHire, '%Y-%m')"
	}

	query := fmt.Sprintf(`
		SELECT %s AS x, COUNT(e.EmpID) AS y, 'Recruitment' AS label
		FROM dim_employee e
		INNER JOIN dim_department d ON e.DeptID = d.DeptID
		WHERE e.DateofHire IS NOT NULL
	`, dateColumn)

	var args []interface{}

	if startDate != "" && endDate != "" {
		query += " AND e.DateofHire BETWEEN ? AND ?"
		args = append(args, startDate, endDate)
	}

	query += `
		AND (? = 0 OR e.DeptID = ?)
		AND (? = 0 OR e.EmpStatusID = ?)
		AND (? = 0 OR d.ManagerID = ?)
		AND (? = 0 OR e.PositionID = ?)
		AND (? = '' OR e.State = ?)
	`

	query += fmt.Sprintf(" GROUP BY %s ORDER BY x", dateColumn)

	args = append(args,
		deptID, deptID,
		empStatusID, empStatusID,
		managerID, managerID,
		positionID, positionID,
		state, state,
	)

	err = config.DB.Raw(query, args...).Scan(&result).Error
	return
}

func GetPerformanceTrend(
	startDate, endDate string,
	empStatusID, managerID, positionID, deptID int,
	state, periodType string,
) (result []LineChartData, err error) {
	var dateColumn string

	if periodType == "year" {
		dateColumn = "YEAR(e.LastPerformanceReview_Date)"
	} else {
		dateColumn = "DATE_FORMAT(e.LastPerformanceReview_Date, '%Y-%m')"
	}

	query := fmt.Sprintf(`
		SELECT %s AS x, AVG(p.PerformanceScore) AS y, 'Performance Score' AS label
		FROM dim_employee e
		INNER JOIN dim_department d ON e.DeptID = d.DeptID
		INNER JOIN dim_performance p ON e.PerfScoreID = p.PerfScoreID
		WHERE e.LastPerformanceReview_Date IS NOT NULL
		AND e.DateofTermination IS NULL
	`, dateColumn)

	var args []interface{}

	if startDate != "" && endDate != "" {
		query += " AND e.LastPerformanceReview_Date BETWEEN ? AND ?"
		args = append(args, startDate, endDate)
	}

	query += `
		AND (? = 0 OR e.DeptID = ?)
		AND (? = 0 OR e.EmpStatusID = ?)
		AND (? = 0 OR d.ManagerID = ?)
		AND (? = 0 OR e.PositionID = ?)
		AND (? = '' OR e.State = ?)
	`

	query += fmt.Sprintf(" GROUP BY %s ORDER BY x", dateColumn)

	args = append(args,
		deptID, deptID,
		empStatusID, empStatusID,
		managerID, managerID,
		positionID, positionID,
		state, state,
	)

	err = config.DB.Raw(query, args...).Scan(&result).Error
	return
}

// GetTurnoverTrend - Tren Turnover
func GetTurnoverTrend(
	startDate, endDate string,
	empStatusID, managerID, positionID, deptID int,
	state, periodType string,
) (result []LineChartData, err error) {
	var dateColumn string

	if periodType == "year" {
		dateColumn = "YEAR(e.DateofTermination)"
	} else {
		dateColumn = "DATE_FORMAT(e.DateofTermination, '%Y-%m')"
	}

	query := fmt.Sprintf(`
		SELECT %s AS x, COUNT(e.EmpID) AS y, 'Turnover' AS label
		FROM dim_employee e
		INNER JOIN dim_department d ON e.DeptID = d.DeptID
		WHERE e.DateofTermination IS NOT NULL
	`, dateColumn)

	var args []interface{}

	if startDate != "" && endDate != "" {
		query += " AND e.DateofTermination BETWEEN ? AND ?"
		args = append(args, startDate, endDate)
	}

	query += `
		AND (? = 0 OR e.DeptID = ?)
		AND (? = 0 OR e.EmpStatusID = ?)
		AND (? = 0 OR d.ManagerID = ?)
		AND (? = 0 OR e.PositionID = ?)
		AND (? = '' OR e.State = ?)
	`

	query += fmt.Sprintf(" GROUP BY %s ORDER BY x", dateColumn)

	args = append(args,
		deptID, deptID,
		empStatusID, empStatusID,
		managerID, managerID,
		positionID, positionID,
		state, state,
	)

	err = config.DB.Raw(query, args...).Scan(&result).Error
	return
}

// GetLateAbsenceTrend - Tren Late/Absence
func GetLateAbsenceTrend(
	startDate, endDate string,
	empStatusID, managerID, positionID, deptID int,
	state, periodType, trendType string,
) (result []LineChartData, err error) {
	var dateColumn string

	if periodType == "year" {
		dateColumn = "YEAR(e.DateofHire)"
	} else {
		dateColumn = "DATE_FORMAT(e.DateofHire, '%Y-%m')"
	}

	var selectClause string
	var label string

	switch trendType {
	case "late":
		selectClause = "AVG(e.DaysLateLast30) AS y"
		label = "Average Days Late"
	case "absence":
		selectClause = "AVG(e.Absences) AS y"
		label = "Average Absences"
	default: // both
		selectClause = "AVG(e.DaysLateLast30 + e.Absences) AS y"
		label = "Average Late + Absences"
	}

	query := fmt.Sprintf(`
		SELECT %s AS x, %s, '%s' AS label
		FROM dim_employee e
		INNER JOIN dim_department d ON e.DeptID = d.DeptID
		WHERE e.DateofTermination IS NULL
		AND e.DateofHire IS NOT NULL
	`, dateColumn, selectClause, label)

	var args []interface{}

	if startDate != "" && endDate != "" {
		query += " AND e.DateofHire BETWEEN ? AND ?"
		args = append(args, startDate, endDate)
	}

	query += `
		AND (? = 0 OR e.DeptID = ?)
		AND (? = 0 OR e.EmpStatusID = ?)
		AND (? = 0 OR d.ManagerID = ?)
		AND (? = 0 OR e.PositionID = ?)
		AND (? = '' OR e.State = ?)
	`

	query += fmt.Sprintf(" GROUP BY %s ORDER BY x", dateColumn)

	args = append(args,
		deptID, deptID,
		empStatusID, empStatusID,
		managerID, managerID,
		positionID, positionID,
		state, state,
	)

	err = config.DB.Raw(query, args...).Scan(&result).Error
	return
}
