package handlers

import (
	dto "dumbmerch/dto/result"
	trandto "dumbmerch/dto/transaction"
	"fmt"
	"strconv"

	"dumbmerch/models"
	"dumbmerch/repository"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type transactionHandlers struct {
	TransactionRepository repository.TransactionRepository
}

func HandlerTransaction(TransactionRepository repository.TransactionRepository) *transactionHandlers {
	return &transactionHandlers{TransactionRepository}
}

func (h *transactionHandlers) GetAllTransaction(c echo.Context) error {
	var transaction []models.Transaction
	// fmt.Println(transaction)
	transaction, err := h.TransactionRepository.GetAllTransaction()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	// fmt.Println(transaction)
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: transaction})
}

func (h *transactionHandlers) GetTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var transaction models.Transaction
	transaction, err := h.TransactionRepository.GetTransaction(id)

	fmt.Println(transaction)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTransaction(transaction)})
}

func (h *transactionHandlers) CreateTransaction(c echo.Context) error {
	request := new(trandto.CreateTransactionRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	trip, _ := h.TransactionRepository.GetTripByID(request.TripID)

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	// fmt.Println(request)
	transaction := models.Transaction{
		CounterQty: request.CounterQty,
		Total:      request.Total,
		Status:     request.Status,
		Attachment: request.Attachment,
		TripID:     request.TripID,
		Trip:       trip,
		UserID:     int(userId),
	}

	data, err := h.TransactionRepository.CreateTransaction(transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTransaction(data)})
}

func (h *transactionHandlers) UpdateTransaction(c echo.Context) error {
	request := new(trandto.CreateTransactionRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))

	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	trip, _ := h.TransactionRepository.GetTripByID(request.TripID)
	user, _ := h.TransactionRepository.GetUserByID(request.UserID)

	if request.CounterQty != 0 {
		transaction.CounterQty = request.CounterQty
	}

	if request.Total != 0 {
		transaction.Total = request.Total
	}

	if request.Status != "" {
		transaction.Status = request.Status
	}

	if request.Attachment != "" {
		transaction.Attachment = request.Attachment
	}

	transaction.TripID = request.TripID
	transaction.Trip = trip

	transaction.UserID = request.UserID
	transaction.User = user

	fmt.Println(user)
	// fmt.Println(trip)
	// fmt.Println(transaction.Trip)

	data, err := h.TransactionRepository.UpdateTransaction(transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTransaction(data)})
}

func (h *transactionHandlers) DeleteTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.TransactionRepository.DeleteTransaction(transaction, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTransaction(data)})

}

func (h *transactionHandlers) GetTransactionByUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	// var transaction models.Transaction
	transaction, err := h.TransactionRepository.GetTransactionByUser(id)

	fmt.Println(transaction)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: transaction})
}

func convertResponseTransaction(u models.Transaction) trandto.TransactionResponse {
	return trandto.TransactionResponse{
		ID:         u.ID,
		CounterQty: u.CounterQty,
		Total:      u.Total,
		Status:     u.Status,
		Attachment: u.Attachment,
		TripID:     u.TripID,
		Trip:       models.TripResponse(u.Trip),
		User:       models.User{},
	}
}

func convertResponseTransactionUpdate(u models.Transaction) trandto.TransactionUpdateResponse {
	return trandto.TransactionUpdateResponse{
		CounterQty: u.CounterQty,
		Total:      u.Total,
		Status:     u.Status,
		Attachment: u.Attachment,
		TripID:     u.TripID,
		Trip:       models.Trip(u.Trip),
	}
}
