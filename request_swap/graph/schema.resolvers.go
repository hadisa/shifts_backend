package graph

import (
	"context"
	"errors"
	"fmt"
	"request_swaps/graph/generated"
	"request_swaps/graph/model"
	"request_swaps/util"
	"time"

	sentry "github.com/getsentry/sentry-go"
	"github.com/google/uuid"
)

// CreateRequestSwap is the resolver for the createRequestSwap field.
func (r *mutationResolver) CreateRequestSwap(ctx context.Context, input model.RequestSwapInput, authUserID *string) (*model.RequestSwapResponse, error) {
	var shiftError []*model.ShiftError
	fieldError := "Create Request Swap"
	errorMessage := "Something went wrong while adding the Request Swap." // default error message

	if authUserID == nil || *authUserID == string("") {
		errorMessage = "Authenticated user id is required"
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeRequired,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	var err error
	permission := false

	// validate permission
	permission, err = util.CheckPermission("request_swap", "WRITE_ALL", *authUserID)
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	if !permission {
		permission, err = util.CheckPermission("request_swap", "WRITE", *authUserID)
		if err != nil {
			sentry.CaptureException(err)
			defer sentry.Flush(2 * time.Second)

			errorMessage = err.Error()
			shiftError = append(shiftError, &model.ShiftError{
				Code:    model.ShiftErrorCodeInvalid,
				Field:   &fieldError,
				Message: &errorMessage,
			})

			return &model.RequestSwapResponse{
				Errors:  shiftError,
				Request: nil,
			}, nil
		}
	}

	if !permission {
		errorMessage = "Permission denied: request_swap.WRITE, request_swap.WRITE_ALL"
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeRequired,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	requestId, err := util.CreateRequest(&input.ChannelID, &input.UserID)
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	if requestId == "" {
		errorMessage = "Something went wrong while creating Request."
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	requestSwap := &model.RequestSwap{
		ID:                        uuid.New().String(),
		ChannelID:                 input.ChannelID,
		RequestID:                 &requestId,
		UserID:                    input.UserID,
		AssignedUserShiftID:       &input.AssignedUserShiftID,
		AssignedUserShiftIDToSwap: input.AssignedUserShiftIDToSwap,
		Status:                    model.RequestStatusPending,
		RequestNote:               input.RequestNote,
		ResponseNote:              input.ResponseNote,
		ResponseByUserID:          input.ResponseByUserID,
		ResponseAt:                input.ResponseAt,
		CreatedAt:                 time.Now().UTC(),
	}

	err = r.DB.Create(&requestSwap).Error
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{

			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	// get the user from the user service and return it
	user, err := util.GetUser(input.UserID)

	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{

			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	if user.ID == nil || *user.ID == string("") {
		errorMessage = "User not found"
		shiftError = append(shiftError, &model.ShiftError{

			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	/* TODO: get offered shift and assigned to user shift */

	requestResponse := model.RequestResponse{
		ID:           requestSwap.ID,
		ChannelID:    requestSwap.ChannelID,
		RequestID:    *requestSwap.RequestID,
		RequestNote:  requestSwap.RequestNote,
		Status:       &requestSwap.Status,
		ResponseNote: requestSwap.ResponseNote,
		ResponseAt:   requestSwap.ResponseAt,
		CreatedAt:    &requestSwap.CreatedAt,
		User:         user,
	}

	return &model.RequestSwapResponse{
		Errors:  nil,
		Request: &requestResponse,
	}, nil
}

// UpdateRequestSwap is the resolver for the updateRequestSwap field.
func (r *mutationResolver) UpdateRequestSwap(ctx context.Context, id string, input model.RequestSwapInput, authUserID *string) (*model.RequestSwapResponse, error) {
	var shiftError []*model.ShiftError
	fieldError := "Update Request Swap"
	errorMessage := "Something went wrong while updating the Request Swap." // default error message

	if authUserID == nil || *authUserID == string("") {
		errorMessage = "Authenticated user id is required"
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeRequired,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	var err error
	permission := false

	// validate permission
	permission, err = util.CheckPermission("request_swap", "WRITE_ALL", *authUserID)
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil

	}

	if !permission {
		permission, err = util.CheckPermission("request_swap", "WRITE", *authUserID)
		if err != nil {
			sentry.CaptureException(err)
			defer sentry.Flush(2 * time.Second)

			errorMessage = err.Error()
			shiftError = append(shiftError, &model.ShiftError{
				Code:    model.ShiftErrorCodeRequired,
				Field:   &fieldError,
				Message: &errorMessage,
			})

			return &model.RequestSwapResponse{
				Errors:  shiftError,
				Request: nil,
			}, nil
		}
	}

	if !permission {
		errorMessage = "Permission denied: request_swap.WRITE, request_swap.WRITE_ALL"
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeRequired,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	if id == "" {
		errorMessage = "Request Swap id is required"
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeRequired,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	err = r.DB.Model(&model.RequestSwap{}).Where("id = ?", id).Updates(model.RequestSwap{
		ChannelID:                 input.ChannelID,
		UserID:                    input.UserID,
		AssignedUserShiftID:       &input.AssignedUserShiftID,
		AssignedUserShiftIDToSwap: input.AssignedUserShiftIDToSwap,
		Status:                    model.RequestStatusPending,
		RequestNote:               input.RequestNote,
		ResponseNote:              input.ResponseNote,
		ResponseByUserID:          input.ResponseByUserID,
		ResponseAt:                input.ResponseAt,
	}).Error

	if err != nil {
		errorMessage = "Error updating Request Swap: " + err.Error()
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	// get the updated request swap
	var requestSwap model.RequestSwap
	err = r.DB.Where("id = ?", id).First(&requestSwap).Error
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeNotFound,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	user, err := util.GetUser(input.UserID)

	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{

			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	if user.ID == nil || *user.ID == string("") {
		errorMessage = "User not found"
		shiftError = append(shiftError, &model.ShiftError{

			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	/* TODO: get offered shift and assigned to user shift */

	requestResponse := model.RequestResponse{
		ID:           requestSwap.ID,
		ChannelID:    requestSwap.ChannelID,
		RequestID:    *requestSwap.RequestID,
		RequestNote:  requestSwap.RequestNote,
		Status:       &requestSwap.Status,
		ResponseNote: requestSwap.ResponseNote,
		ResponseAt:   requestSwap.ResponseAt,
		CreatedAt:    &requestSwap.CreatedAt,
		User:         user,
	}

	return &model.RequestSwapResponse{
		Errors:  nil,
		Request: &requestResponse,
	}, nil
}

// DeleteRequestSwap is the resolver for the deleteRequestSwap field.
func (r *mutationResolver) DeleteRequestSwap(ctx context.Context, id string, authUserID *string) (*model.RequestSwapResponse, error) {
	var shiftError []*model.ShiftError
	fieldError := "Delete Request Swap"
	errorMessage := "Something went wrong while deleting the Request Swap." // default error message

	if authUserID == nil || *authUserID == string("") {
		errorMessage = "Authenticated user id is required"
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeRequired,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	var err error
	permission := false

	// validate permission
	permission, err = util.CheckPermission("request_swap", "WRITE_ALL", *authUserID)
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	if !permission {
		permission, err = util.CheckPermission("request_swap", "WRITE", *authUserID)
		if err != nil {
			sentry.CaptureException(err)
			defer sentry.Flush(2 * time.Second)

			errorMessage = err.Error()
			shiftError = append(shiftError, &model.ShiftError{
				Code:    model.ShiftErrorCodeInvalid,
				Field:   &fieldError,
				Message: &errorMessage,
			})

			return &model.RequestSwapResponse{
				Errors:  shiftError,
				Request: nil,
			}, nil
		}
	}

	if !permission {
		errorMessage = "Permission denied: request_swap.WRITE, request_swap.WRITE_ALL"
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeRequired,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}
	if id == "" {
		return nil, errors.New("id is required")
	}

	var requestSwap model.RequestSwap
	// get request swap	by id
	err = r.DB.First(&requestSwap, "id= ?", id).Error
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeNotFound,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	// delete request swap
	err = r.DB.Delete(&requestSwap, "id = ?", id).Error

	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	// delete request
	deleted, err := util.DeleteRequest(requestSwap.RequestID)
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		return nil, err
	}
	if !deleted {
		errorMessage = "There was an error deleting the request in Request API"
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	user, err := util.GetUser(requestSwap.UserID)

	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{

			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	if user.ID == nil || *user.ID == string("") {
		errorMessage = "User not found"
		shiftError = append(shiftError, &model.ShiftError{

			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	/* TODO: get offered shift and assigned to user shift */

	requestResponse := model.RequestResponse{
		ID:           requestSwap.ID,
		ChannelID:    requestSwap.ChannelID,
		RequestID:    *requestSwap.RequestID,
		RequestNote:  requestSwap.RequestNote,
		Status:       &requestSwap.Status,
		ResponseNote: requestSwap.ResponseNote,
		ResponseAt:   requestSwap.ResponseAt,
		CreatedAt:    &requestSwap.CreatedAt,
		User:         user,
	}

	return &model.RequestSwapResponse{
		Errors:  nil,
		Request: &requestResponse,
	}, nil
}

// CancelRequestSwap is the resolver for the cancelRequestSwap field.
func (r *mutationResolver) CancelRequestSwap(ctx context.Context, channelID string, requestID string, authUserID *string) (*model.RequestSwapResponse, error) {
	// in the frontend, the user can only cancel their own request swap so the other request swap must not be displayed to the user

	var shiftError []*model.ShiftError
	fieldError := "Cancel Request Swap"
	errorMessage := "Something went wrong while canceling the Request Swap." // default error message

	if authUserID == nil || *authUserID == string("") {
		errorMessage = "Authenticated user id is required"
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeRequired,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	var err error
	permission := false

	// validate permission
	permission, err = util.CheckPermission("request_swap", "READ_ALL", *authUserID)
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	if !permission {
		permission, err = util.CheckPermission("request_swap", "READ", *authUserID)
		if err != nil {
			sentry.CaptureException(err)
			defer sentry.Flush(2 * time.Second)

			errorMessage = err.Error()
			shiftError = append(shiftError, &model.ShiftError{
				Code:    model.ShiftErrorCodeInvalid,
				Field:   &fieldError,
				Message: &errorMessage,
			})

			return &model.RequestSwapResponse{
				Errors:  shiftError,
				Request: nil,
			}, nil
		}
	}

	if !permission {
		errorMessage = "Permission denied: request_swap.READ_ALL, request_swap.READ"
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeRequired,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	if channelID == "" || requestID == "" {
		return nil, errors.New("channelID and requestID are required")
	}

	var requestSwap model.RequestSwap
	status := model.RequestStatusCancelled

	// update the request swap and set the status to cancelled
	err = r.DB.Model(&model.RequestSwap{}).Where("channel_id = ? AND request_id = ?", channelID, requestID).Update("status", status).Error

	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	// get the updated request swap
	err = r.DB.Where("channel_id = ? AND request_id = ?", channelID, requestID).First(&requestSwap).Error
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeNotFound,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	// get the user from the user service and return it
	user, err := util.GetUser(*authUserID)

	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{

			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	if user.ID == nil || *user.ID == string("") {
		errorMessage = "User not found"
		shiftError = append(shiftError, &model.ShiftError{

			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	/* TODO: get offered shift and assigned to user shift */

	requestResponse := model.RequestResponse{
		ID:           requestSwap.ID,
		ChannelID:    requestSwap.ChannelID,
		RequestID:    *requestSwap.RequestID,
		RequestNote:  requestSwap.RequestNote,
		Status:       &requestSwap.Status,
		ResponseNote: requestSwap.ResponseNote,
		ResponseAt:   requestSwap.ResponseAt,
		CreatedAt:    &requestSwap.CreatedAt,
		User:         user,
	}

	return &model.RequestSwapResponse{
		Errors:  nil,
		Request: &requestResponse,
	}, nil
}

// ApproveRequestSwap is the resolver for the approveRequestSwap field.
func (r *mutationResolver) ApproveRequestSwap(ctx context.Context, id string, responseNote *string, authUserID *string) (*model.RequestSwapResponse, error) {
	var shiftError []*model.ShiftError
	fieldError := "Approve Request Swap"
	errorMessage := "Something went wrong while approving the Request Swap." // default error message

	if authUserID == nil || *authUserID == string("") {
		errorMessage = "Authenticated user id is required"
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeRequired,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	var err error
	permission := false

	// validate permission
	permission, err = util.CheckPermission("request_swap", "MANAGE_ALL", *authUserID)
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	if !permission {
		permission, err = util.CheckPermission("request_swap", "MANAGE", *authUserID)
		if err != nil {
			sentry.CaptureException(err)
			defer sentry.Flush(2 * time.Second)

			errorMessage = err.Error()
			shiftError = append(shiftError, &model.ShiftError{
				Code:    model.ShiftErrorCodeInvalid,
				Field:   &fieldError,
				Message: &errorMessage,
			})

			return &model.RequestSwapResponse{
				Errors:  shiftError,
				Request: nil,
			}, nil
		}
	}

	if !permission {
		errorMessage = "Permission denied: request_swap.MANAGE, request_swap.MANAGE_ALL"
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeRequired,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	if id == "" {
		return nil, errors.New("id is required")
	}

	// approve the request swap
	var requestSwap model.RequestSwap
	status := model.RequestStatusApproved
	approve := &model.RequestSwap{
		Status:       status,
		ResponseNote: responseNote,
	}
	// update the request swap and set the status to approved
	err = r.DB.Model(&model.RequestSwap{}).Where("id = ?", id).Updates(approve).Error
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	// get the updated request swap
	err = r.DB.Where("id = ?", id).First(&requestSwap).Error
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeNotFound,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	user, err := util.GetUser(requestSwap.UserID)

	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{

			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	if user.ID == nil || *user.ID == string("") {
		errorMessage = "User not found"
		shiftError = append(shiftError, &model.ShiftError{

			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	/* TODO: get offered shift and assigned to user shift */

	requestResponse := model.RequestResponse{
		ID:           requestSwap.ID,
		ChannelID:    requestSwap.ChannelID,
		RequestID:    *requestSwap.RequestID,
		RequestNote:  requestSwap.RequestNote,
		Status:       &requestSwap.Status,
		ResponseNote: requestSwap.ResponseNote,
		ResponseAt:   requestSwap.ResponseAt,
		CreatedAt:    &requestSwap.CreatedAt,
		User:         user,
	}

	return &model.RequestSwapResponse{
		Errors:  nil,
		Request: &requestResponse,
	}, nil
}

// DenyRequestSwap is the resolver for the denyRequestSwap field.
func (r *mutationResolver) DenyRequestSwap(ctx context.Context, id string, responseNote *string, authUserID *string) (*model.RequestSwapResponse, error) {
	var shiftError []*model.ShiftError
	fieldError := "Deny Request Swap"
	errorMessage := "Something went wrong while denying the Request Swap." // default error message

	if authUserID == nil || *authUserID == string("") {
		errorMessage = "Authenticated user id is required"
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeRequired,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	var err error
	permission := false

	// validate permission
	permission, err = util.CheckPermission("request_swap", "MANAGE_ALL", *authUserID)
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	if !permission {
		permission, err = util.CheckPermission("request_swap", "MANAGE", *authUserID)
		if err != nil {

			sentry.CaptureException(err)
			defer sentry.Flush(2 * time.Second)

			errorMessage = err.Error()
			shiftError = append(shiftError, &model.ShiftError{
				Code:    model.ShiftErrorCodeInvalid,
				Field:   &fieldError,
				Message: &errorMessage,
			})

			return &model.RequestSwapResponse{
				Errors:  shiftError,
				Request: nil,
			}, nil
		}
	}

	if !permission {
		errorMessage = "Permission denied: request_swap.MANAGE, request_swap.MANAGE_ALL"
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	if id == "" {
		return nil, errors.New("id is required")
	}

	// approve the request swap
	var requestSwap model.RequestSwap
	status := model.RequestStatusDenied
	
	deny := &model.RequestSwap{
		Status:       status,
		ResponseNote: responseNote,
	}

	// update the request swap and set the status to approved
	err = r.DB.Model(&model.RequestSwap{}).Where("id = ?", id).Updates(deny).Error
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	// get the updated request swap
	err = r.DB.Where("id = ?", id).First(&requestSwap).Error
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeNotFound,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	user, err := util.GetUser(requestSwap.UserID)

	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{

			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	if user.ID == nil || *user.ID == string("") {
		errorMessage = "User not found"
		shiftError = append(shiftError, &model.ShiftError{

			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.RequestSwapResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	/* TODO: get offered shift and assigned to user shift */

	requestResponse := model.RequestResponse{
		ID:           requestSwap.ID,
		ChannelID:    requestSwap.ChannelID,
		RequestID:    *requestSwap.RequestID,
		RequestNote:  requestSwap.RequestNote,
		Status:       &requestSwap.Status,
		ResponseNote: requestSwap.ResponseNote,
		ResponseAt:   requestSwap.ResponseAt,
		CreatedAt:    &requestSwap.CreatedAt,
		User:         user,
	}

	return &model.RequestSwapResponse{
		Errors:  nil,
		Request: &requestResponse,
	}, nil
}

// GetRequestsSwaps is the resolver for the getRequestsSwaps field.
func (r *queryResolver) GetRequestsSwaps(ctx context.Context, channelID string, authUserID *string) ([]*model.RequestSwap, error) {
	if authUserID == nil || *authUserID == string("") {
		return nil, fmt.Errorf("Authenticated user id is required")
	}

	var err error
	permission := false

	// validate permission
	permission, err = util.CheckPermission("request_swap", "READ_ALL", *authUserID)
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		return nil, err
	}

	// if !permission {
	// 	permission, err = util.CheckPermission("request_swap", "READ", *authUserID)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	if !permission {
		return nil, fmt.Errorf("Permission denied: request_swap.READ_ALL")
	}

	if channelID == "" {
		return nil, errors.New("channelID is required")
	}

	var requestSwaps []*model.RequestSwap

	err = r.DB.Where("channel_id = ?", channelID).Find(&requestSwaps).Error
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		return nil, err
	}

	return requestSwaps, nil
}

// GetRequestsSwapsByChannelIDRequestID is the resolver for the getRequestsSwapsByChannelIdRequestId field.
func (r *queryResolver) GetRequestsSwapsByChannelIDRequestID(ctx context.Context, channelID string, requestID string) (*model.RequestSwap, error) {
	if channelID == "" {
		return nil, errors.New("channel ID is required")
	}

	if requestID == "" {
		return nil, errors.New("request ID is required")
	}

	var requestSwaps *model.RequestSwap

	err := r.DB.Where("channel_id = ? AND request_id = ?", channelID, requestID).First(&requestSwaps).Error
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		return nil, err
	}

	return requestSwaps, nil
}

// GetRequestSwap is the resolver for the getRequestSwap field.
func (r *queryResolver) GetRequestSwap(ctx context.Context, id string, authUserID *string) (*model.RequestSwap, error) {
	if authUserID == nil || *authUserID == string("") {
		return nil, fmt.Errorf("Authenticated user id is required")
	}

	var err error
	permission := false

	// validate permission
	permission, err = util.CheckPermission("request_swap", "READ", *authUserID)
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		return nil, err
	}

	if !permission {
		permission, err = util.CheckPermission("request_swap", "READ_ALL", *authUserID)
		if err != nil {
			sentry.CaptureException(err)
			defer sentry.Flush(2 * time.Second)

			return nil, err
		}
	}

	if !permission {
		return nil, fmt.Errorf("Permission denied: request_swap.READ, request_swap.READ_ALL")
	}

	if id == "" {
		return nil, errors.New("id is required")
	}

	var requestSwap model.RequestSwap
	err = r.DB.Where("id = ?", id).First(&requestSwap).Error
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		return nil, err
	}

	return &requestSwap, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
