package reqres

type DashboardScoreCard struct {
	TotalEmployee         int     `json:"total_employee"`          
	TurnoverPercentage    float64 `json:"turnover_percentage"`      
	AveragePerformance    float64 `json:"average_performance"`    
	AverageDaysLateLast30 float64 `json:"average_days_late_last30"` 
}
