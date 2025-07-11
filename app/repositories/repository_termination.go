package repository

import "hris-datawarehouse/config"

func GetTop10TerminatedDepartments(
	startDate, endDate string,
	empStatusID, managerID, positionID int,
	state string,
) (result []struct {
	Name  string `json:"name"`
	Total int    `json:"total"`
}, err error) {
	query := `
		SELECT 
			d.Department AS name,
			COUNT(*) AS total
		FROM fact_employment f
		JOIN dim_department d ON f.DeptID = d.DeptID
		JOIN dim_employee e ON f.EmpID = e.EmpID
		WHERE f.Is_Terminated = "Yes"
	`

	var args []interface{}

	if startDate != "" && endDate != "" {
		query += ` AND f.DateofTermination BETWEEN ? AND ?`
		args = append(args, startDate, endDate)
	}

	query += `
		AND (? = 0 OR f.EmpStatusID = ?)
		AND (? = 0 OR f.ManagerID = ?)
		AND (? = 0 OR f.PositionID = ?)
		AND (? = '' OR e.State = ?)
		GROUP BY d.Department
		ORDER BY total DESC
		LIMIT 10
	`

	args = append(args,
		empStatusID, empStatusID,
		managerID, managerID,
		positionID, positionID,
		state, state,
	)

	err = config.DB.Raw(query, args...).Scan(&result).Error
	return
}
