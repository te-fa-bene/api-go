package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/te-fa-bene/api-go/internal/app/http/middleware"
	"github.com/te-fa-bene/api-go/internal/app/repository"
	"github.com/te-fa-bene/api-go/internal/app/service"
)

type AuthHandler struct {
	auth       *service.AuthService
	repository *repository.EmployeeRepository
}

func NewAuthHandler(auth *service.AuthService, repository *repository.EmployeeRepository) *AuthHandler {
	return &AuthHandler{
		auth:       auth,
		repository: repository,
	}
}

type LoginRequest struct {
	StoreID  string `json:"store_id" binding:"required,uuid"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type MeResponse struct {
	ID      string `json:"id"`
	StoreID string `json:"store_id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Role    string `json:"role"`
}

// Login
// @Summary      Login
// @Description  Login a user
// @Tags         auth
// @Accept    	 json
// @Produce      json
// @Success      200  {object}  LoginResponse
// @Failure      401  {object}  map[string]string
// @Router       /auth/login [post]
func (h *AuthHandler) Login(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := h.auth.Login(req.StoreID, req.Email, req.Password)
	if err != nil {
		if err == service.ErrInvalidCredentials {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid credentials",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal error",
		})
		return
	}

	ctx.JSON(http.StatusOK, LoginResponse{
		AccessToken: res.AccessToken,
		TokenType:   "Bearer",
		ExpiresIn:   res.ExpiresIn,
	})
}

// Me
// @Summary      Me
// @Description  Get user information
// @Tags         auth
// @Security 		 BearerAuth
// @Produce      json
// @Success      200  {object}  MeResponse
// @Failure      401  {object}  map[string]string
// @Router       /me [get]
func (h *AuthHandler) Me(ctx *gin.Context) {
	employeeID := ctx.GetString(middleware.CtxEmployeeID)
	storeID := ctx.GetString(middleware.CtxStoreID)

	employee, err := h.repository.FindActiveByIDAndStore(employeeID, storeID)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, MeResponse{
		ID:      employee.ID,
		StoreID: employee.StoreID,
		Name:    employee.Name,
		Email:   employee.Email,
		Role:    employee.Role,
	})
}
