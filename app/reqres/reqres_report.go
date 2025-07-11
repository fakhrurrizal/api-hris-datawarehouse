package reqres

type DashboardScoreCard struct {
	TotalEmployee         int     `json:"total_employee"`          
	TurnoverPercentage    float64 `json:"turnover_percentage"`      
	AveragePerformance    float64 `json:"average_performance"`    
	AverageDaysLateLast30 float64 `json:"average_days_late_last30"` 
}

type EmploymentResponse struct {
    EmpID        int     `json:"emp_id"`
    EmployeeName string  `json:"employee_name"`
    Position     string  `json:"position"`
    Department   string  `json:"department"`
    Manager      string  `json:"manager"`
    Salary       float64 `json:"salary"`
    DateOfHire   string  `json:"date_of_hire"`
}

type EmploymentPagingResponse struct {
    Page  int                 `json:"page"`
    Limit int                 `json:"limit"`
    Total int64               `json:"total"`
    Data  []EmploymentResponse `json:"data"`
}