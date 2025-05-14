package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ofojichigozie/iphms-go-backend/dtos"
	"github.com/ofojichigozie/iphms-go-backend/middleware"
	"github.com/ofojichigozie/iphms-go-backend/responses"
	"github.com/ofojichigozie/iphms-go-backend/services"
)

type VitalsController struct {
	vitalsService services.VitalsService
}

func NewVitalsController(vitalsService services.VitalsService) *VitalsController {
	return &VitalsController{vitalsService}
}

func (vc *VitalsController) CreateVitals(ctx *gin.Context) {
	currentUser, err := middleware.GetCurrentUser(ctx)
	if err != nil {
		responses.Error(ctx, http.StatusUnauthorized, "Authentication required", err.Error())
		return
	}

	var body dtos.CreateVitalsInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		responses.Error(ctx, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	body.UserId = currentUser.UserId

	vitals, err := vc.vitalsService.CreateVitals(body)
	if err != nil {
		responses.Error(ctx, http.StatusBadRequest, "Failed to record vitals", err.Error())
		return
	}

	responses.Success(ctx, http.StatusCreated, "Vitals recorded successfully", vitals)
}

func (vc *VitalsController) GetAllVitals(ctx *gin.Context) {
	currentUser, err := middleware.GetCurrentUser(ctx)
	if err != nil {
		responses.Error(ctx, http.StatusUnauthorized, "Authentication required", err.Error())
		return
	}

	filters := make(map[string]interface{})

	if currentUser.Role != "admin" {
		filters["user_id"] = currentUser.UserId
	} else {
		if userIdStr := ctx.Query("userId"); userIdStr != "" {
			userId, err := strconv.ParseUint(userIdStr, 10, 32)
			if err != nil {
				responses.Error(ctx, http.StatusBadRequest,
					"Invalid user ID", "User ID must be a positive integer")
				return
			}
			filters["user_id"] = uint(userId)
		}
	}

	if temperature := ctx.Query("temperature"); temperature != "" {
		if temp, err := strconv.ParseFloat(temperature, 32); err == nil {
			filters["temperature"] = float32(temp)
		}
	}

	if startDate := ctx.Query("startDate"); startDate != "" {
		filters["created_at >= ?"] = startDate
	}

	if endDate := ctx.Query("endDate"); endDate != "" {
		filters["created_at <= ?"] = endDate
	}

	vitals, err := vc.vitalsService.GetAllVitals(filters)
	if err != nil {
		responses.Error(ctx, http.StatusInternalServerError, "Couldn't fetch vitals", nil)
		return
	}

	responses.Success(ctx, http.StatusOK, "Vitals retrieved successfully", vitals)
}

func (vc *VitalsController) GetVitals(ctx *gin.Context) {
	currentUser, err := middleware.GetCurrentUser(ctx)
	if err != nil {
		responses.Error(ctx, http.StatusUnauthorized, "Authentication required", err.Error())
		return
	}

	id := ctx.Param("id")
	vitalsId, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		responses.Error(ctx, http.StatusBadRequest,
			"Invalid vitals ID", "Vitals ID must be a positive integer")
		return
	}

	vitals, err := vc.vitalsService.GetVitalsById(uint(vitalsId))
	if err != nil {
		responses.Error(ctx, http.StatusNotFound, "Vitals not found", nil)
		return
	}

	if currentUser.Role != "admin" && currentUser.UserId != vitals.UserID {
		responses.Error(ctx, http.StatusForbidden,
			"Access denied", "You can only view your own vitals")
		return
	}

	responses.Success(ctx, http.StatusOK, "Vitals retrieved successfully", vitals)
}

func (vc *VitalsController) DeleteVital(ctx *gin.Context) {
	currentUser, err := middleware.GetCurrentUser(ctx)
	if err != nil {
		responses.Error(ctx, http.StatusUnauthorized, "Authentication required", err.Error())
		return
	}

	id := ctx.Param("id")
	vitalsId, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		responses.Error(ctx, http.StatusBadRequest,
			"Invalid vitals ID", "Vitals ID must be a positive integer")
		return
	}

	vitals, err := vc.vitalsService.GetVitalsById(uint(vitalsId))
	if err != nil {
		responses.Error(ctx, http.StatusNotFound, "Vitals not found", nil)
		return
	}

	if currentUser.Role != "admin" && currentUser.UserId != vitals.UserID {
		responses.Error(ctx, http.StatusForbidden,
			"Not authorized", "You can only delete your own vitals records")
		return
	}

	err = vc.vitalsService.DeleteVitalsById(uint(vitalsId))
	if err != nil {
		responses.Error(ctx, http.StatusNotFound, "Vitals not found", nil)
		return
	}

	responses.Success(ctx, http.StatusOK, "Vitals deleted successfully", nil)
}
