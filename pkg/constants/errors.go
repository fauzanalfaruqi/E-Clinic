package constants

const (
	ErrScheduleTaken            = "doctors schedule at that time has been taken"
	ErrScheduleNotMatch         = "your booking isn't match with doctor's schedule"
	ErrScheduleDateExist        = "the schedule for that date already exists"
	ErrDateFormat               = "invalid date format"
	ErrDocSchedNotExist         = "doctor schedule is not exist"
	ErrPaymentAlreadyTrue       = "the payment has already been set to true"
	ErrQuantityGreaterThanStock = "quantity amount is greater than the stock available"
	ErrNoStockAvailable         = "no stock available for this item"
)
