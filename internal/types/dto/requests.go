package dto

type CreateSubscriptionRequest struct {
	ServiceName string  `json:"service_name"`
	Price       int     `json:"price" validate:"required,gte=0"`
	UserID      string  `json:"user_id" validate:"required,uuid4"`
	StartDate   string  `json:"start_date" validate:"required,datetime=01-2006"`
	EndDate     *string `json:"end_date" validate:"omitempty,datetime=01-2006"`
}

type GetSubscriptionRequest struct {
	SubID string `validate:"required,uuid4"` // FROM PATH.
}

type UpdateSubscriptionRequest struct {
	SubID   string  `validate:"required,uuid4"` // FROM PATH.
	Price   *int    `json:"price" validate:"required_without=EndDate,gte=0"`
	EndDate *string `json:"end_date" validate:"required_without=Price,datetime=01-2006"`
}

type DeleteSubscriptionRequest struct {
	SubID string `validate:"required,uuid4"` // FROM PATH.
}

type ListSubscriptionsRequest struct {
	UserID      *string `schema:"user_id" validate:"required_without=ServiceName,uuid4"`
	ServiceName *string `schema:"service" validate:"required_without=UserID"`
}

type GetTotalCostRequest struct {
	FromDate    string  `schema:"from" validate:"required,datetime=01-2006"`
	ToDate      string  `schema:"to" validate:"required,datetime=01-2006"`
	UserID      *string `schema:"user_id" validate:"required_without=ServiceName,uuid4"`
	ServiceName *string `schema:"service" validate:"required_without=UserID"`
}
