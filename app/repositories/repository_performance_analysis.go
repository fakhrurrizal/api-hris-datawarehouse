package repository

import (
	"fmt"
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

func GetSatisfactionHeatmapByPosition(
	startDate, endDate string,
	empStatusID, managerID, positionID, deptID int,
	state, gender string,
) (result []struct {
	X     string  `json:"x"`     // Posisi
	Y     string  `json:"y"`     // Tahun-Bulan
	Value float64 `json:"value"` // Rata-rata EmpSatisfaction
}, err error) {

	dateFormat := "DATE_FORMAT(f.DateofHire, '%Y-%m')"

	query := fmt.Sprintf(`
		SELECT 
			pos.Position AS x, 
			%s AS y, 
			ROUND(AVG(p.EmpSatisfaction), 2) AS value
		FROM fact_employment f
		INNER JOIN dim_employee e ON f.EmpID = e.EmpID
		INNER JOIN dim_position pos ON f.PositionID = pos.PositionID
		INNER JOIN dim_performance p ON f.PerfScoreID = p.PerfScoreID
		WHERE f.Is_Terminated = 0
			AND p.EmpSatisfaction IS NOT NULL
			AND p.EmpSatisfaction > 0
	`, dateFormat)

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

	query += fmt.Sprintf(` GROUP BY pos.Position, %s ORDER BY %s ASC, pos.Position ASC`, dateFormat, dateFormat)

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


func GetEmployeeCountByMaritalStatus(
	startDate, endDate string,
	empStatusID, managerID, positionID, deptID, maritalStatusID int,
	state string,
) (result []struct {
	Name  string `json:"name"`
	Total int    `json:"total"`
}, err error) {

	query := `
		SELECT 
			ms.MaritalDesc AS name,
			COUNT(*) AS total
		FROM fact_employment f
		JOIN dim_employee e ON f.EmpID = e.EmpID
		JOIN dim_marital_status ms ON f.MaritalStatusID_validate = ms.MaritalStatusID
		WHERE f.Is_Terminated = "No"
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
		AND (? = 0 OR f.MaritalStatusID_validate = ?)
		AND (? = '' OR e.State = ?)
		GROUP BY ms.MaritalDesc
		ORDER BY total DESC
	`

	args = append(args,
		deptID, deptID,
		empStatusID, empStatusID,
		managerID, managerID,
		positionID, positionID,
		maritalStatusID, maritalStatusID,
		state, state,
	)

	err = config.DB.Raw(query, args...).Scan(&result).Error
	return
}
