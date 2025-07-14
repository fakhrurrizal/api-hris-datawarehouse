package repository

import (
	"fmt"
	"hris-datawarehouse/config"
)

func buildFilterQuery(
	departmentID int, gender string, empStatusID int, managerID int, positionID int, state string,
) (string, []interface{}) {
	var query string
	var args []interface{}

	if departmentID != 0 {
		query += " AND f.DeptID = ?"
		args = append(args, departmentID)
	}
	if gender != "" {
		query += " AND e.Gender = ?"
		args = append(args, gender)
	}
	if empStatusID != 0 {
		query += " AND f.EmpStatusID = ?"
		args = append(args, empStatusID)
	}
	if managerID != 0 {
		query += " AND f.ManagerID = ?"
		args = append(args, managerID)
	}
	if positionID != 0 {
		query += " AND f.PositionID = ?"
		args = append(args, positionID)
	}
	if state != "" {
		query += " AND e.State = ?"
		args = append(args, state)
	}

	return query, args
}

func GetTotalActiveEmployees(
	startDate, endDate string,
	departmentID int, gender string, empStatusID int, managerID int, positionID int, state string,
) (total int, err error) {
	query := `
		SELECT COUNT(DISTINCT f.EmpID) AS total 
		FROM fact_employment f
		JOIN dim_employee e USING(EmpID)
		JOIN dim_department d USING(DeptID)
		WHERE f.DateofTermination IS NULL
	`
	var args []interface{}

	if startDate != "" && endDate != "" {
		query += " AND f.DateofHire BETWEEN ? AND ?"
		args = append(args, startDate, endDate)
	}

	filter, filterArgs := buildFilterQuery(departmentID, gender, empStatusID, managerID, positionID, state)
	query += filter
	args = append(args, filterArgs...)

	err = config.DB.Raw(query, args...).Scan(&total).Error
	return
}

func GetTurnoverRate(
	startDate, endDate string,
	departmentID int, gender string, empStatusID int, managerID int, positionID int, state string,
) (percentage float64, err error) {
	var args []interface{}
	var resignedFilter, activeFilter string

	// Filter waktu resignation
	if startDate != "" && endDate != "" {
		resignedFilter += " AND f.DateofTermination BETWEEN ? AND ?"
		args = append(args, startDate, endDate)
	}

	// Filter dinamis
	filter1, filterArgs1 := buildFilterQuery(departmentID, gender, empStatusID, managerID, positionID, state)
	resignedFilter += filter1
	args = append(args, filterArgs1...)

	// Filter waktu active
	if startDate != "" && endDate != "" {
		activeFilter += " AND f.DateofHire BETWEEN ? AND ?"
		args = append(args, startDate, endDate)
	}

	// Filter dinamis lagi
	filter2, filterArgs2 := buildFilterQuery(departmentID, gender, empStatusID, managerID, positionID, state)
	activeFilter += filter2
	args = append(args, filterArgs2...)

	query := fmt.Sprintf(`
		WITH resigned AS (
			SELECT COUNT(*) AS total
			FROM fact_employment f
			JOIN dim_employee e USING(EmpID)
			JOIN dim_department d USING(DeptID)
			WHERE f.DateofTermination IS NOT NULL
			%s
		),
		active AS (
			SELECT COUNT(*) AS total
			FROM fact_employment f
			JOIN dim_employee e USING(EmpID)
			JOIN dim_department d USING(DeptID)
			WHERE f.DateofTermination IS NULL
			%s
		)
		SELECT COALESCE(ROUND((CAST(resigned.total AS DECIMAL) / NULLIF(active.total, 0)) * 100, 2), 0.0) AS turnover_rate_percentage
		FROM resigned, active
	`, resignedFilter, activeFilter)

	err = config.DB.Raw(query, args...).Scan(&percentage).Error
	return
}

func GetAveragePerformanceScore(
	departmentID int, gender string, empStatusID int, managerID int, positionID int, state string,
) (avgScore float64, err error) {
	query := `
		SELECT COALESCE(ROUND(AVG(f.PerfScoreID), 2), 0.0) AS avg_score 
		FROM fact_employment f
		JOIN dim_employee e USING(EmpID)
		JOIN dim_department d USING(DeptID)
		WHERE f.PerfScoreID IS NOT NULL 
		AND f.DateofTermination IS NULL
	`

	filter, filterArgs := buildFilterQuery(departmentID, gender, empStatusID, managerID, positionID, state)
	query += filter

	err = config.DB.Raw(query, filterArgs...).Scan(&avgScore).Error
	return
}

func GetAverageDaysLateLast30(
	departmentID int, gender string, empStatusID int, managerID int, positionID int, state string,
) (avgDays float64, err error) {
	query := `
		SELECT COALESCE(ROUND(AVG(f.DaysLateLast30), 2), 0.0) AS avg_days 
		FROM fact_employment f
		JOIN dim_employee e USING(EmpID)
		JOIN dim_department d USING(DeptID)
		WHERE f.DaysLateLast30 IS NOT NULL 
		AND f.DateofTermination IS NULL
	`

	filter, filterArgs := buildFilterQuery(departmentID, gender, empStatusID, managerID, positionID, state)
	query += filter

	err = config.DB.Raw(query, filterArgs...).Scan(&avgDays).Error
	return
}

func GetTotalSalaryExpenditure(
	departmentID int, gender string, empStatusID int, managerID int, positionID int, state string,
) (totalSalary float64, err error) {
	query := `
		SELECT SUM(f.Salary) AS total_salary
		FROM fact_employment f
		JOIN dim_employee e USING(EmpID)
		JOIN dim_department d USING(DeptID)
		WHERE f.DateofTermination IS NULL
	`

	filter, filterArgs := buildFilterQuery(departmentID, gender, empStatusID, managerID, positionID, state)
	query += filter

	err = config.DB.Raw(query, filterArgs...).Scan(&totalSalary).Error
	return
}
