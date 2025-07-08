package repository

import (
	"fmt"
	"hris-datawarehouse/config"
)

func buildFilterQuery() string {
	return `
		AND (? = 0 OR e.DeptID = ?)
		AND (? = '' OR e.Gender = ?)
		AND (? = 0 OR e.EmpStatusID = ?)
		AND (? = 0 OR d.ManagerID = ?)
		AND (? = 0 OR e.PositionID = ?)
		AND (? = '' OR e.State = ?)
	`
}
func GetTotalActiveEmployees(
	startDate, endDate string,
	departmentID int, gender string, empStatusID int, managerID int, positionID int, state string,
) (total int, err error) {
	query := `
		SELECT COUNT(*) AS total 
		FROM dim_employee e
		JOIN dim_department d USING(DeptID)
		WHERE e.DateofTermination IS NULL
	`

	var args []interface{}

	if startDate != "" && endDate != "" {
		query += " AND e.DateofHire BETWEEN ? AND ?"
		args = append(args, startDate, endDate)
	}

	query += buildFilterQuery()

	args = append(args,
		departmentID, departmentID,
		gender, gender,
		empStatusID, empStatusID,
		managerID, managerID,
		positionID, positionID,
		state, state,
	)

	err = config.DB.Raw(query, args...).Scan(&total).Error
	return
}

func GetTurnoverRate(
	startDate, endDate string,
	departmentID int, gender string, empStatusID int, managerID int, positionID int, state string,
) (percentage float64, err error) {
	filter := buildFilterQuery()

	// Base query
	query := `
		WITH resigned AS (
			SELECT COUNT(*) AS total
			FROM dim_employee e
			JOIN dim_department d USING(DeptID)
			WHERE e.DateofTermination IS NOT NULL
			%s
		),
		active AS (
			SELECT COUNT(*) AS total
			FROM dim_employee e
			JOIN dim_department d USING(DeptID)
			WHERE e.DateofTermination IS NULL
			%s
		)
		SELECT COALESCE(ROUND((CAST(resigned.total AS DECIMAL) / NULLIF(active.total, 0)) * 100, 2), 0.0) AS turnover_rate_percentage
		FROM resigned, active
	`

	var resignedFilter, activeFilter string
	var args []interface{}

	// Tanggal resignation
	if startDate != "" && endDate != "" {
		resignedFilter += " AND e.DateofTermination BETWEEN ? AND ?"
		args = append(args, startDate, endDate)
	}

	// Tambah filter dinamis lainnya
	resignedFilter += filter
	args = append(args,
		departmentID, departmentID,
		gender, gender,
		empStatusID, empStatusID,
		managerID, managerID,
		positionID, positionID,
		state, state,
	)

	// Tanggal hire
	if startDate != "" && endDate != "" {
		activeFilter += " AND e.DateofHire BETWEEN ? AND ?"
		args = append(args, startDate, endDate)
	}

	// Tambah filter dinamis lainnya
	activeFilter += filter
	args = append(args,
		departmentID, departmentID,
		gender, gender,
		empStatusID, empStatusID,
		managerID, managerID,
		positionID, positionID,
		state, state,
	)

	finalQuery := fmt.Sprintf(query, resignedFilter, activeFilter)

	// Eksekusi
	err = config.DB.Raw(finalQuery, args...).Scan(&percentage).Error
	return
}

func GetAveragePerformanceScore(
	departmentID int, gender string, empStatusID int, managerID int, positionID int, state string,
) (avgScore float64, err error) {
	query := `
		SELECT COALESCE(ROUND(AVG(e.PerfScoreID), 2), 0.0) AS avg_score 
		FROM dim_employee e
		JOIN dim_department d USING(DeptID)
		WHERE e.PerfScoreID IS NOT NULL 
		AND e.DateofTermination IS NULL
	` + buildFilterQuery()

	err = config.DB.Raw(query,
		departmentID, departmentID,
		gender, gender,
		empStatusID, empStatusID,
		managerID, managerID,
		positionID, positionID,
		state, state,
	).Scan(&avgScore).Error
	return
}

func GetAverageDaysLateLast30(
	departmentID int, gender string, empStatusID int, managerID int, positionID int, state string,
) (avgDays float64, err error) {
	query := `
		SELECT COALESCE(ROUND(AVG(e.DaysLateLast30), 2), 0.0) AS avg_days 
		FROM dim_employee e
		JOIN dim_department d USING(DeptID)
		WHERE e.DaysLateLast30 IS NOT NULL 
		AND e.DateofTermination IS NULL
	` + buildFilterQuery()

	err = config.DB.Raw(query,
		departmentID, departmentID,
		gender, gender,
		empStatusID, empStatusID,
		managerID, managerID,
		positionID, positionID,
		state, state,
	).Scan(&avgDays).Error
	return
}
