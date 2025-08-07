package dto

type SubscriptionResponse struct {
	ID          string  `json:"id"`
	ServiceName string  `json:"service_name"`
	Price       int     `json:"price"`
	UserID      string  `json:"user_id"`
	StartDate   string  `json:"start_date"`
	EndDate     *string `json:"end_date,omitempty"`
}

type SubscriptionListResponse struct {
	Subs []*SubscriptionResponse `json:"subscriptions"`
}

type TotalCostResponse struct {
	TotalCost int `json:"total_cost"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
