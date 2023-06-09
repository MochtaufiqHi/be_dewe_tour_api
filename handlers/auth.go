package handlers

import (
	authdto "dumbmerch/dto/auth"
	dto "dumbmerch/dto/result"
	"dumbmerch/models"
	"dumbmerch/pkg/bcrypt"
	jwtToken "dumbmerch/pkg/jwt"
	"dumbmerch/repository"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlersAuth struct {
	AuthRepository repository.AuthRepository
}

func HandlersAuth(AuthRepository repository.AuthRepository) *handlersAuth {
	return &handlersAuth{AuthRepository}
}

func (h *handlersAuth) Register(c echo.Context) error {
	request := new(authdto.AuthRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "binding gagal"})
	}

	// fmt.Println(request)

	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "hashing password gagal"})
	}

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "validate gagal"})
	}

	user := models.User{
		Fullname: request.Fullname,
		Email:    request.Email,
		Password: password,
		Phone:    request.Phone,
		Address:  request.Address,
		Role:     "user",
	}
	// fmt.Println(user)

	data, err := h.AuthRepository.Register(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: "create gagal"})
	}

	registerResponse := authdto.RegisterResponse{
		//Fullname: user.Fullname,
		Email:    data.Email,
		Password: data.Password,
		//Token: token,
	}

	// fmt.Println(registerResponse)
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: registerResponse})
}

func (h *handlersAuth) Login(c echo.Context) error {
	request := new(authdto.LoginRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	user := models.User{
		Email:    request.Email,
		Password: request.Password,
	}

	// fmt.Println(user)
	//check mail
	user, err := h.AuthRepository.Login(user.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	// check password
	isValid := bcrypt.CheckPasswordHash(request.Password, user.Password)
	if !isValid {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	// generate token
	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	loginResponse := authdto.LoginResponse{
		Fullname: user.Fullname,
		Email:    user.Email,
		Role:     user.Role,
		Address:  user.Address,
		Phone:    user.Phone,
		ID:       user.ID,
		// TransactionID: user.TransactionID,
		// Transaction: user.Transaction,

		// Password: user.Password,
		Token: token,
	}
	// fmt.Println(user.Role)
	// fmt.Println(loginResponse)

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: loginResponse})

}

func (h *handlersAuth) CheckAuth(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userID := userLogin.(jwt.MapClaims)["id"].(float64)

	user, _ := h.AuthRepository.CheckAuth(int(userID))

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: user})
}
