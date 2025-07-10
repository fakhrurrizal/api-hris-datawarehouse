package repository

import (
	"fmt"
	"hris-datawarehouse/config"
)

type LineChartData struct {
	X     string  `json:"x"`
	Y     float64 `json:"y"`
	Label string  `json:"label,omitempty"`
}

func GetRecruitmentTrend(
	startDate, endDate string,
	empStatusID, managerID, positionID, deptID int,
	state, periodType string,
) (result []LineChartData, err error) {
	var dateColumn string
	if periodType == "year" {
		dateColumn = "YEAR(f.DateofHire)"
	} else {
		dateColumn = "DATE_FORMAT(f.DateofHire, '%Y-%m')"
	}

	query := fmt.Sprintf(`
		SELECT %s AS x, COUNT(DISTINCT f.EmpID) AS y, 'Recruitment' AS label
		FROM fact_employment f
		JOIN dim_employee e USING(EmpID)
		JOIN dim_department d USING(DeptID)
		WHERE f.DateofHire IS NOT NULL AND f.Is_Terminated = 0
	`, dateColumn)

	var args []interface{}
	if startDate != "" && endDate != "" {
		query += " AND f.DateofHire BETWEEN ? AND ?"
		args = append(args, startDate, endDate)
	}
	query += `
		AND (? = 0 OR f.DeptID = ?)
		AND (? = 0 OR f.EmpStatusID = ?)
		AND (? = 0 OR f.ManagerID = ?)
		AND (? = 0 OR f.PositionID = ?)
		AND (? = '' OR e.State = ?)
	`
	query += fmt.Sprintf(" GROUP BY %s ORDER BY x", dateColumn)
	args = append(args, deptID, deptID, empStatusID, empStatusID, managerID, managerID, positionID, positionID, state, state)

	err = config.DB.Raw(query, args...).Scan(&result).Error
	return
}

func GetPerformanceTrend(
	startDate, endDate string,
	empStatusID, managerID, positionID, deptID int,
	state, periodType string,
) (result []struct {
	X     string `json:"x"`
	Y     int    `json:"y"`
	Label string `json:"label"`
}, err error) {
	var dateColumn string
	if periodType == "year" {
		dateColumn = "YEAR(f.DateofHire)"
	} else {
		dateColumn = "DATE_FORMAT(f.DateofHire, '%Y-%m')"
	}

	query := fmt.Sprintf(`
		SELECT %s AS x, COUNT(*) AS y, p.PerformanceScore AS label
		FROM fact_employment f
		JOIN dim_employee e USING(EmpID)
		JOIN dim_department d USING(DeptID)
		JOIN dim_performance p ON f.PerfScoreID = p.PerfScoreID
		WHERE f.PerfScoreID IS NOT NULL AND f.Is_Terminated = 0
	`, dateColumn)

	var args []interface{}
	if startDate != "" && endDate != "" {
		query += " AND f.DateofHire BETWEEN ? AND ?"
		args = append(args, startDate, endDate)
	}
	query += `
		AND (? = 0 OR f.DeptID = ?)
		AND (? = 0 OR f.EmpStatusID = ?)
		AND (? = 0 OR f.ManagerID = ?)
		AND (? = 0 OR f.PositionID = ?)
		AND (? = '' OR e.State = ?)
	`
	query += fmt.Sprintf(" GROUP BY %s, p.PerformanceScore ORDER BY x", dateColumn)
	args = append(args, deptID, deptID, empStatusID, empStatusID, managerID, managerID, positionID, positionID, state, state)

	err = config.DB.Raw(query, args...).Scan(&result).Error
	return
}

func GetTurnoverTrend(
	startDate, endDate string,
	empStatusID, managerID, positionID, deptID int,
	state, periodType string,
) (result []LineChartData, err error) {
	var dateColumn string
	if periodType == "year" {
		dateColumn = "YEAR(f.DateofTermination)"
	} else {
		dateColumn = "DATE_FORMAT(f.DateofTermination, '%Y-%m')"
	}

	query := fmt.Sprintf(`
		SELECT %s AS x, COUNT(DISTINCT f.EmpID) AS y, 'Turnover' AS label
		FROM fact_employment f
		JOIN dim_employee e USING(EmpID)
		JOIN dim_department d USING(DeptID)
		WHERE f.DateofTermination IS NOT NULL
	`, dateColumn)

	var args []interface{}
	if startDate != "" && endDate != "" {
		query += " AND f.DateofTermination BETWEEN ? AND ?"
		args = append(args, startDate, endDate)
	}
	query += `
		AND (? = 0 OR f.DeptID = ?)
		AND (? = 0 OR f.EmpStatusID = ?)
		AND (? = 0 OR f.ManagerID = ?)
		AND (? = 0 OR f.PositionID = ?)
		AND (? = '' OR e.State = ?)
	`
	query += fmt.Sprintf(" GROUP BY %s ORDER BY x", dateColumn)
	args = append(args, deptID, deptID, empStatusID, empStatusID, managerID, managerID, positionID, positionID, state, state)

	err = config.DB.Raw(query, args...).Scan(&result).Error
	return
}

func GetLateAbsenceTrend(
	startDate, endDate string,
	empStatusID, managerID, positionID, deptID int,
	state, periodType, trendType string,
) (result []LineChartData, err error) {
	var dateColumn string
	if periodType == "year" {
		dateColumn = "YEAR(f.DateofHire)"
	} else {
		dateColumn = "DATE_FORMAT(f.DateofHire, '%Y-%m')"
	}

	var selectClause, label string
	switch trendType {
	case "late":
		selectClause = "AVG(f.DaysLateLast30) AS y"
		label = "Average Days Late"
	case "absence":
		selectClause = "AVG(f.Absences) AS y"
		label = "Average Absences"
	default:
		selectClause = "AVG(f.DaysLateLast30 + f.Absences) AS y"
		label = "Average Late + Absences"
	}

	query := fmt.Sprintf(`
		SELECT %s AS x, %s, '%s' AS label
		FROM fact_employment f
		JOIN dim_employee e USING(EmpID)
		JOIN dim_department d USING(DeptID)
		WHERE f.Is_Terminated = 0 AND f.DateofHire IS NOT NULL
	`, dateColumn, selectClause, label)

	var args []interface{}
	if startDate != "" && endDate != "" {
		query += " AND f.DateofHire BETWEEN ? AND ?"
		args = append(args, startDate, endDate)
	}
	query += `
		AND (? = 0 OR f.DeptID = ?)
		AND (? = 0 OR f.EmpStatusID = ?)
		AND (? = 0 OR f.ManagerID = ?)
		AND (? = 0 OR f.PositionID = ?)
		AND (? = '' OR e.State = ?)
	`
	query += fmt.Sprintf(" GROUP BY %s ORDER BY x", dateColumn)
	args = append(args, deptID, deptID, empStatusID, empStatusID, managerID, managerID, positionID, positionID, state, state)

	err = config.DB.Raw(query, args...).Scan(&result).Error
	return
}
