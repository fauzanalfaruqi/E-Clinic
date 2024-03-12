package bookingDelivery

import (
	"avengers-clinic/model/dto"
	"avengers-clinic/model/dto/json"
	"avengers-clinic/pkg/constants"
	"avengers-clinic/pkg/utils"
	"avengers-clinic/src/booking"
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type bookingDelivery struct {
	bookingUC booking.BookingUsecase
}

func NewBookingDelivery(v1Group *gin.RouterGroup, bookingUC booking.BookingUsecase) {
	handler := bookingDelivery{
		bookingUC,
	}

	bookingGroup := v1Group.Group("/booking")
	{
		bookingGroup.GET("", handler.GetAll)
		bookingGroup.GET("/:id", handler.GetByID)
		bookingGroup.POST("", handler.Create)
		bookingGroup.PUT("/:id", handler.EditSchedule)
		bookingGroup.PUT("/done/:id", handler.Done)
		bookingGroup.PUT("/cancel/:id", handler.Cancel)
	}
}

func (bd bookingDelivery) GetAll(ctx *gin.Context) {

	//Get query date
	//If not exist, the date is now()
	date := ctx.Query("date")
	if date == "" {
		date = time.Now().Format("2006-01-02")
	} else {

		if _, err := time.Parse("2006-01-02", date); err != nil {
			json.NewResponseBadRequest(ctx, nil, "invalid date type", constants.BookingService, "01")
			return
		}
	}

	//Get the datas
	data, err := bd.bookingUC.GetAll()
	if err != nil {
		json.NewResponseError(ctx, err.Error(), constants.BookingService, "01")
		return
	}

	json.NewResponseSuccess(ctx, data, "success", constants.BookingService, "01")
}

func (bd bookingDelivery) GetByID(ctx *gin.Context) {

	//Get id from url param, parse to uuid
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		json.NewResponseBadRequest(ctx, nil, err.Error(), constants.BookingService, "01")
		return
	}

	//Get data
	data, err := bd.bookingUC.GetOneByID(id)

	//validating error
	if err != nil && err == sql.ErrNoRows {
		json.NewResponseBadRequest(ctx, nil, "data not found", constants.BookingService, "01")
		return
	} else if err != nil {
		json.NewResponseError(ctx, err.Error(), constants.BookingService, "01")
		return
	}

	json.NewResponseSuccess(ctx, data, "success", constants.BookingService, "01")

}

func (bd bookingDelivery) GetByDoctorID(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		json.NewResponseBadRequest(ctx, nil, err.Error(), constants.BookingService, "01")
		return
	}

	//get data
	data, err := bd.bookingUC.GetAllByDoctorID(id)
	if err != nil && err == sql.ErrNoRows {
		json.NewResponseBadRequest(ctx, nil, "data not found", constants.BookingService, "01")
		return
	} else if err != nil {
		json.NewResponseError(ctx, err.Error(), constants.BookingService, "01")
		return
	}

	json.NewResponseSuccess(ctx, data, "success", constants.BookingService, "01")

}

func (bd bookingDelivery) Create(ctx *gin.Context) {
	var input dto.CreateBooking

	//Binding req body
	if err := ctx.ShouldBindJSON(&input); err != nil {
		json.NewResponseError(ctx, err.Error(), constants.BookingService, "01")
		return
	}

	//validate tag
	if err := utils.Validated(input); err != nil {
		json.NewResponseBadRequest(ctx, err, "Bad request", constants.BookingService, "01")
		return
	}

	//Create booking
	data, err := bd.bookingUC.Create(input)
	//if create failed, it return err no rows 
	//because we do use validation create where not exist
	//and returnin ID
	if err != nil && err == sql.ErrNoRows {
		json.NewResponseBadRequest(ctx, nil, constants.ErrScheduleTaken, constants.BookingService, "01")
		return
	//we also use validate match doctor_schedules.day_of_week == day_of_week(bookings.booking_date)
	} else if err != nil && (err.Error() == constants.ErrDocSchedNotExist || err.Error() == constants.ErrScheduleNotMatch) {
		json.NewResponseBadRequest(ctx, nil, err.Error(), constants.BookingService, "01")
		return
	} else if err != nil {
		json.NewResponseError(ctx, err.Error(), constants.BookingService, "01")
		return
	}

	json.NewResponseCreated(ctx, data, "success", constants.BookingService, "01")
}

func (bd bookingDelivery) EditSchedule(ctx *gin.Context) {

	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		json.NewResponseBadRequest(ctx, nil, err.Error(), constants.BookingService, "01")
		return
	}

	var input dto.UpdateBookingSchedule

	if err := ctx.ShouldBindJSON(&input); err != nil {
		json.NewResponseError(ctx, err.Error(), constants.BookingService, "01")
		return
	}

	//Validating struct tag
	if err := utils.Validated(input); err != nil {
		json.NewResponseBadRequest(ctx, err, "Bad request", constants.BookingService, "01")
		return
	}

	data, err := bd.bookingUC.EditSchedule(id, input)
	if err != nil && (err == sql.ErrNoRows || err.Error() == constants.ErrScheduleTaken) {
		json.NewResponseBadRequest(ctx, nil, constants.ErrScheduleTaken, constants.BookingService, "01")
		return
	} else if err != nil {
		json.NewResponseError(ctx, err.Error(), constants.BookingService, "01")
		return
	}

	json.NewResponseCreated(ctx, data, "success", constants.BookingService, "01")
}

func (bd bookingDelivery) Done(ctx *gin.Context) {
	// var input dto.UpdateBooking

	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		json.NewResponseBadRequest(ctx, nil, err.Error(), constants.BookingService, "01")
		return
	}

	// file, _ := ctx.FormFile("img")
	// file.Filename = "coba"
	// file.
	// ctx.SaveUploadedFile(file, )

	// if err := ctx.ShouldBindJSON(&input); err != nil {
	// 	json.NewResponseError(ctx, err.Error(), constants.BookingService, "01")
	// 	return
	// }

	// if err := utils.Validated(input); err != nil {
	// 	json.NewResponseBadRequest(ctx, err, "Bad request", constants.BookingService, "01")
	// 	return
	// }

	data, err := bd.bookingUC.FinishBooking(id)
	if err != nil && err == sql.ErrNoRows {
		json.NewResponseBadRequest(ctx, nil, "data not found", constants.BookingService, "01")
		return
	} else if err != nil {
		json.NewResponseError(ctx, err.Error(), constants.BookingService, "01")
		return
	}

	json.NewResponseCreated(ctx, data, "success", constants.BookingService, "01")
}

func (bd bookingDelivery) Cancel(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		json.NewResponseBadRequest(ctx, nil, err.Error(), constants.BookingService, "01")
		return
	}

	data, err := bd.bookingUC.Cancel(id)
	if err != nil && err == sql.ErrNoRows {
		json.NewResponseBadRequest(ctx, nil, "data not found", constants.BookingService, "01")
		return
	} else if err != nil {
		json.NewResponseError(ctx, err.Error(), constants.BookingService, "01")
		return
	}
	json.NewResponseCreated(ctx, data, "canceled", constants.BookingService, "01")
}
