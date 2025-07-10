package repository

import (
	"hris-datawarehouse/config"
)

type PerformanceScoreData struct {
	Name  string  `json:"name"`
	Total float64 `json:"total"`
}

type PerformanceScoreCount struct {
	Name  string  `json:"name"`
	Total float64 `json:"total"`
	Count int     `json:"count,omitempty"`
}

func GetAveragePerformanceScorePerDepartmentWithCount(
	startDate, endDate string,
	empStatusID, managerID, positionID, deptID int,
	state, gender string,
) (result []PerformanceScoreCount, err error) {
	query := `
		SELECT 
			d.Department AS name, 
			COUNT(f.EmpID) AS count
		FROM fact_employment f
		INNER JOIN dim_employee e ON f.EmpID = e.EmpID
		INNER JOIN dim_department d ON f.DeptID = d.DeptID
		INNER JOIN dim_performance p ON f.PerfScoreID = p.PerfScoreID
		WHERE f.Is_Terminated = 0
		AND f.PerfScoreID IS NOT NULL
	`

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
		AND (? = '' OR e.Gender = ?)
	`

	query += " GROUP BY d.Department, f.DeptID ORDER BY count DESC"

	args = append(args,
		deptID, deptID,
		empStatusID, empStatusID,
		managerID, managerID,
		positionID, positionID,
		state, state,
		gender, gender,
	)

	err = config.DB.Raw(query, args...).Scan(&result).Error
	return
}

func GetAverageEmpSatisfactionPerPositionRounded(
	startDate, endDate string,
	empStatusID, managerID, positionID, deptID int,
	state, gender string,
) (result []PerformanceScoreCount, err error) {
	query := `
		SELECT pos.Position AS name, ROUND(AVG(p.EmpSatisfaction), 2) AS total
		FROM fact_employment f
		INNER JOIN dim_employee e ON f.EmpID = e.EmpID
		INNER JOIN dim_position pos ON f.PositionID = pos.PositionID
		INNER JOIN dim_performance p ON f.PerfScoreID = p.PerfScoreID
		WHERE f.Is_Terminated = 0
		AND p.EmpSatisfaction IS NOT NULL
		AND p.EmpSatisfaction > 0
	`

	var args []interface{}
	if startDate != "" && endDate != "" {
		query += " AND f.DateofHire BETWEEN ? AND ?"
		args = append(args, startDate, endDate)
	}

	query += `
		AND (? = 0 OR f.DeptID = ?)
		AND (? = 0 OR f.EmpStatusID = ?)
		AND (? = 0 OR f.PositionID = ?)
		AND (? = '' OR e.State = ?)
		AND (? = '' OR e.Gender = ?)
	`

	query += " GROUP BY pos.Position, f.PositionID ORDER BY total DESC"

	args = append(args,
		deptID, deptID,
		empStatusID, empStatusID,
		positionID, positionID,
		state, state,
		gender, gender,
	)

	err = config.DB.Raw(query, args...).Scan(&result).Error
	return
}
