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
			COUNT(e.EmpID) AS count
		FROM dim_employee e
		INNER JOIN dim_department d ON e.DeptID = d.DeptID
		INNER JOIN dim_performance p ON e.PerfScoreID = p.PerfScoreID
		WHERE e.DateofTermination IS NULL
		AND p.PerformanceScore IS NOT NULL
	`

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
		AND (? = '' OR e.Gender = ?)
	`

	query += " GROUP BY d.Department, d.DeptID ORDER BY count DESC"

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
		SELECT pos.Position AS name, ROUND(AVG(e.EmpSatisfaction), 2) AS total
		FROM dim_employee e
		INNER JOIN dim_position pos ON e.PositionID = pos.PositionID
		WHERE e.DateofTermination IS NULL
		AND e.EmpSatisfaction IS NOT NULL
		AND e.EmpSatisfaction > 0
	`
	var args []interface{}
	if startDate != "" && endDate != "" {
		query += " AND e.DateofHire BETWEEN ? AND ?"
		args = append(args, startDate, endDate)
	}
	query += `
		AND (? = 0 OR e.DeptID = ?)
		AND (? = 0 OR e.EmpStatusID = ?)
		AND (? = 0 OR e.PositionID = ?)
		AND (? = '' OR e.State = ?)
		AND (? = '' OR e.Gender = ?)
	`
	query += " GROUP BY pos.Position, pos.PositionID ORDER BY total DESC"
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
