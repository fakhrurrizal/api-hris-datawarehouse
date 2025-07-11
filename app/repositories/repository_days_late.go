package repository

import (
	"hris-datawarehouse/config"
)

func GetLateHeatmapMonthlyByDepartment(
	startDate, endDate string,
	empStatusID, managerID, positionID, deptID int,
	state string,
) (result []struct {
	Y     string `json:"y"`     // Bulan (format: YYYY-MM)
	X     string `json:"x"`     // Nama Departemen
	Value int    `json:"value"` // Jumlah DaysLateLast30
}, err error) {

	query := `
		SELECT 
			DATE_FORMAT(f.CurrentDate, '%Y-%m') AS y,
			d.Department AS x,
			SUM(f.DaysLateLast30) AS value
		FROM fact_employment f
		JOIN dim_employee e ON f.EmpID = e.EmpID
		JOIN dim_department d ON f.DeptID = d.DeptID
		WHERE f.DaysLateLast30 > 0
			AND f.Is_Terminated = 0
	`

	var args []interface{}

	if startDate != "" && endDate != "" {
		query += " AND f.CurrentDate BETWEEN ? AND ?"
		args = append(args, startDate, endDate)
	}

	query += `
		AND (? = 0 OR f.DeptID = ?)
		AND (? = 0 OR f.EmpStatusID = ?)
		AND (? = 0 OR f.ManagerID = ?)
		AND (? = 0 OR f.PositionID = ?)
		AND (? = '' OR e.State = ?)
		GROUP BY y, x
		ORDER BY y ASC, x ASC
	`

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
