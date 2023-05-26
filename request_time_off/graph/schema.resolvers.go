package graph

import (
	"context"
	"errors"
	"fmt"
	"request_time_offs/graph/generated"
	"request_time_offs/graph/model"
	"request_time_offs/util"
	"time"

	sentry "github.com/getsentry/sentry-go"
	"github.com/google/uuid"
)

// CreateRequestTimeOff is the resolver for the createRequestTimeOff field.
func (r *mutationResolver) CreateRequestTimeOff(ctx context.Context, input model.RequestTimeOffInput, authUserID *string) (*model.TimeOffResponse, error) {
	var shiftError []*model.ShiftError
	fieldError := "Create Request Time Off"
	errorMessage := "Something went wrong while creating the Request Time Off." // default error message

	if authUserID == nil || *authUserID == string("") {
		errorMessage = "Authenticated user id is required"
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeRequired,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	var err error
	permission := false

	// validate permission
	permission, err = util.CheckPermission("request_time_off", "WRITE_ALL", *authUserID)
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		return nil, err
	}

	if !permission {
		permission, err = util.CheckPermission("request_time_off", "WRITE", *authUserID)
		if err != nil {
			sentry.CaptureException(err)
			defer sentry.Flush(2 * time.Second)

			return nil, err
		}
	}

	if !permission {
		errorMessage = "Permission denied: request_time_off.WRITE, request_time_off.WRITE_ALL"
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeRequired,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
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
			Code:    model.ShiftErrorCodeRequired,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	if requestId == "" {
		errorMessage = "There was an error creating the request in Request API"
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	// create RequestTimeOff
	timeNow := time.Now().UTC()
	requestTimeOff := &model.RequestTimeOff{
		ID:               uuid.New().String(),
		RequestID:        &requestId,
		ChannelID:        &input.ChannelID,
		UserID:           input.UserID,
		Reason:           input.Reason,
		StartTime:        input.StartTime,
		EndTime:          input.EndTime,
		Is24Hours:        &input.Is24Hours,
		RequestNote:      input.RequestNote,
		Status:           model.RequestStatusPending,
		ResponseNote:     input.ResponseNote,
		ResponseByUserID: input.ResponseByUserID,
		ResponseAt:       input.ResponseAt,
		CreatedAt:        timeNow,
	}

	err = r.DB.Create(requestTimeOff).Error
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
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

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	if user.ID == nil || *user.ID == string("") {
		errorMessage = "User not found"
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeNotFound,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	requestResponse := model.RequestResponse{
		ID:           requestTimeOff.ID,
		ChannelID:    *requestTimeOff.ChannelID,
		RequestID:    *requestTimeOff.RequestID,
		RequestNote:  requestTimeOff.RequestNote,
		Status:       &requestTimeOff.Status,
		ResponseNote: requestTimeOff.ResponseNote,
		ResponseAt:   requestTimeOff.ResponseAt,
		CreatedAt:    &requestTimeOff.CreatedAt,
		User:         user,
	}

	return &model.TimeOffResponse{
		Errors:  nil,
		Request: &requestResponse,
	}, nil
}

// UpdateRequestTimeOff is the resolver for the updateRequestTimeOff field.
func (r *mutationResolver) UpdateRequestTimeOff(ctx context.Context, id string, input model.RequestTimeOffInput, authUserID *string) (*model.TimeOffResponse, error) {
	var shiftError []*model.ShiftError
	fieldError := "Update Request Time Off"
	errorMessage := "Something went wrong while updating the Request Time Off." // default error message

	if authUserID == nil || *authUserID == string("") {
		errorMessage = "Authenticated user id is required"
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeRequired,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	var err error
	permission := false

	// validate permission
	permission, err = util.CheckPermission("request_time_off", "WRITE_ALL", *authUserID)
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	if !permission {
		permission, err = util.CheckPermission("request_time_off", "WRITE", *authUserID)
		if err != nil {
			sentry.CaptureException(err)
			defer sentry.Flush(2 * time.Second)

			errorMessage = err.Error()
			shiftError = append(shiftError, &model.ShiftError{
				Code:    model.ShiftErrorCodeInvalid,
				Field:   &fieldError,
				Message: &errorMessage,
			})

			return &model.TimeOffResponse{
				Errors:  shiftError,
				Request: nil,
			}, nil
		}
	}

	if !permission {
		errorMessage = "Permission denied: request_time_off.WRITE, request_time_off.WRITE_ALL"
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeRequired,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	if id == "" {
		errorMessage = "Request Time Off ID is required"
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeRequired,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	// update RequestTimeOff
	err = r.DB.Model(&model.RequestTimeOff{}).Where("id = ?", id).Updates(
		model.RequestTimeOff{
			UserID:           input.UserID,
			Reason:           input.Reason,
			StartTime:        input.StartTime,
			EndTime:          input.EndTime,
			Is24Hours:        &input.Is24Hours,
			RequestNote:      input.RequestNote,
			Status:           model.RequestStatusDenied,
			ResponseNote:     input.ResponseNote,
			ResponseByUserID: input.ResponseByUserID,
			ResponseAt:       input.ResponseAt,
			ChannelID:        &input.ChannelID},
	).Error

	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	// get RequestTimeOff
	var requestTimeOff model.RequestTimeOff
	err = r.DB.Where("id = ?", id).First(&requestTimeOff).Error
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeNotFound,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	// get the user from the user service and return it
	user, err := util.GetUser(requestTimeOff.UserID)

	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	if user.ID == nil || *user.ID == string("") {
		errorMessage = "Something went wrong while fetching the user."
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	requestResponse := model.RequestResponse{
		ID:           requestTimeOff.ID,
		ChannelID:    *requestTimeOff.ChannelID,
		RequestID:    *requestTimeOff.RequestID,
		RequestNote:  requestTimeOff.RequestNote,
		Status:       &requestTimeOff.Status,
		ResponseNote: requestTimeOff.ResponseNote,
		ResponseAt:   requestTimeOff.ResponseAt,
		CreatedAt:    &requestTimeOff.CreatedAt,
		User:         user,
	}

	return &model.TimeOffResponse{
		Errors:  nil,
		Request: &requestResponse,
	}, nil
}

// DeleteRequestTimeOff is the resolver for the deleteRequestTimeOff field.
func (r *mutationResolver) DeleteRequestTimeOff(ctx context.Context, id string, authUserID *string) (*model.TimeOffResponse, error) {
	var shiftError []*model.ShiftError
	fieldError := "Delete Request Time Off"
	errorMessage := "Something went wrong while deleting the Request Time Off." // default error message

	if authUserID == nil || *authUserID == string("") {
		errorMessage = "Authenticated user id is required"
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeRequired,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	var err error
	permission := false

	// validate permission
	permission, err = util.CheckPermission("request_time_off", "WRITE_ALL", *authUserID)
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	if !permission {
		permission, err = util.CheckPermission("request_time_off", "WRITE", *authUserID)
		if err != nil {
			sentry.CaptureException(err)
			defer sentry.Flush(2 * time.Second)

			errorMessage = err.Error()
			shiftError = append(shiftError, &model.ShiftError{
				Code:    model.ShiftErrorCodeInvalid,
				Field:   &fieldError,
				Message: &errorMessage,
			})

			return &model.TimeOffResponse{
				Errors:  shiftError,
				Request: nil,
			}, nil
		}
	}

	if !permission {
		errorMessage = "Permission denied: request_time_off.WRITE, request_time_off.WRITE_ALL"
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeRequired,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	if id == "" {
		errorMessage = "Request Time Off ID is required"
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeRequired,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	var requestTimeOff model.RequestTimeOff
	// get request time off by id
	err = r.DB.First(&requestTimeOff, "id= ?", id).Error
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeNotFound,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	// delete request time off
	err = r.DB.Delete(&requestTimeOff, "id = ?", id).Error

	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	// delete request
	deleted, err := util.DeleteRequest(requestTimeOff.RequestID)
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeNotFound,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}
	if !deleted {
		errorMessage = "Request not found or already deleted"
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeNotFound,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	// get the user from the user service and return it
	user, err := util.GetUser(requestTimeOff.UserID)

	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	if user.ID == nil || *user.ID == string("") {
		errorMessage = "Something went wrong while fetching the user."
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	requestResponse := model.RequestResponse{
		ID:           requestTimeOff.ID,
		ChannelID:    *requestTimeOff.ChannelID,
		RequestID:    *requestTimeOff.RequestID,
		RequestNote:  requestTimeOff.RequestNote,
		Status:       &requestTimeOff.Status,
		ResponseNote: requestTimeOff.ResponseNote,
		ResponseAt:   requestTimeOff.ResponseAt,
		CreatedAt:    &requestTimeOff.CreatedAt,
		User:         user,
	}

	return &model.TimeOffResponse{
		Errors:  nil,
		Request: &requestResponse,
	}, nil
}

// CancelRequestTimeOff is the resolver for the cancelRequestTimeOff field.
func (r *mutationResolver) CancelRequestTimeOff(ctx context.Context, channelID string, requestID string, authUserID *string) (*model.TimeOffResponse, error) {
	var shiftError []*model.ShiftError
	fieldError := "Cancel Request Time Off"
	errorMessage := "Something went wrong while canceling the Request Time Off." // default error message

	if authUserID == nil || *authUserID == string("") {
		errorMessage = "Authenticated user id is required"
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeRequired,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	var err error
	permission := false

	// validate permission
	permission, err = util.CheckPermission("request_time_off", "READ_ALL", *authUserID)
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	if !permission {
		permission, err = util.CheckPermission("request_time_off", "READ", *authUserID)
		if err != nil {
			sentry.CaptureException(err)
			defer sentry.Flush(2 * time.Second)

			errorMessage = err.Error()
			shiftError = append(shiftError, &model.ShiftError{
				Code:    model.ShiftErrorCodeInvalid,
				Field:   &fieldError,
				Message: &errorMessage,
			})

			return &model.TimeOffResponse{
				Errors:  shiftError,
				Request: nil,
			}, nil
		}
	}

	if !permission {
		errorMessage = "Permission denied: request_time_off.READ, request_time_off.READ_ALL"
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeRequired,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	if channelID == "" || requestID == "" {
		errorMessage = "Channel ID and Request ID are required"
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeRequired,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	status := model.RequestStatusCancelled
	var requestTimeOff *model.RequestTimeOff

	// auth user ID also should be passed in where clause
	err = r.DB.Model(&requestTimeOff).Where("channel_id = ? AND request_id = ?", channelID, requestID).Update("status", status).Error

	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	// get request time off
	err = r.DB.Where("channel_id = ? AND request_id = ?", channelID, requestID).First(&requestTimeOff).Error
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeNotFound,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	// get the user from the user service and return it
	user, err := util.GetUser(requestTimeOff.UserID)

	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	if user.ID == nil || *user.ID == string("") {
		errorMessage = "Something went wrong while fetching the user."
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	requestResponse := model.RequestResponse{
		ID:           requestTimeOff.ID,
		ChannelID:    *requestTimeOff.ChannelID,
		RequestID:    *requestTimeOff.RequestID,
		RequestNote:  requestTimeOff.RequestNote,
		Status:       &requestTimeOff.Status,
		ResponseNote: requestTimeOff.ResponseNote,
		ResponseAt:   requestTimeOff.ResponseAt,
		CreatedAt:    &requestTimeOff.CreatedAt,
		User:         user,
	}

	return &model.TimeOffResponse{
		Errors:  nil,
		Request: &requestResponse,
	}, nil
}

// ApproveRequestTimeOff is the resolver for the approveRequestTimeOff field.
func (r *mutationResolver) ApproveRequestTimeOff(ctx context.Context, id string, responseNote *string, authUserID *string) (*model.TimeOffResponse, error) {
	var shiftError []*model.ShiftError
	fieldError := "Approve Request Time Off"
	errorMessage := "Something went wrong while approving the Request Time Off." // default error message

	if authUserID == nil || *authUserID == string("") {
		errorMessage = "Authenticated user id is required"
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeRequired,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	var err error
	permission := false

	// validate permission
	permission, err = util.CheckPermission("request_time_off", "MANAGE_ALL", *authUserID)
	if err != nil {

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	if !permission {
		permission, err = util.CheckPermission("request_time_off", "MANAGE", *authUserID)
		if err != nil {
			sentry.CaptureException(err)
			defer sentry.Flush(2 * time.Second)

			errorMessage = err.Error()
			shiftError = append(shiftError, &model.ShiftError{
				Code:    model.ShiftErrorCodeInvalid,
				Field:   &fieldError,
				Message: &errorMessage,
			})

			return &model.TimeOffResponse{
				Errors:  shiftError,
				Request: nil,
			}, nil
		}
	}

	if !permission {
		errorMessage = "Permission denied: request_time_off.MANAGE, request_time_off.MANAGE_ALL"
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	if id == "" {
		errorMessage = "Request Time Off ID is required"
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeRequired,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	status := model.RequestStatusApproved
	var requestTimeOff *model.RequestTimeOff

	approve := &model.RequestTimeOff{
		Status:       status,
		ResponseNote: responseNote,
	}
	// update request time off status to accepted
	err = r.DB.Model(&requestTimeOff).Where("id= ?", id).Updates(approve).Error

	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	// get request time off by channel id and request id
	err = r.DB.Where("id= ?", id).First(&requestTimeOff).Error

	if err != nil {

		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeNotFound,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	// get the user from the user service and return it
	user, err := util.GetUser(requestTimeOff.UserID)

	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	if user.ID == nil || *user.ID == string("") {
		errorMessage = "Something went wrong while fetching the user."
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	requestResponse := model.RequestResponse{
		ID:           requestTimeOff.ID,
		ChannelID:    *requestTimeOff.ChannelID,
		RequestID:    *requestTimeOff.RequestID,
		RequestNote:  requestTimeOff.RequestNote,
		Status:       &requestTimeOff.Status,
		ResponseNote: requestTimeOff.ResponseNote,
		ResponseAt:   requestTimeOff.ResponseAt,
		CreatedAt:    &requestTimeOff.CreatedAt,
		User:         user,
	}

	return &model.TimeOffResponse{
		Errors:  nil,
		Request: &requestResponse,
	}, nil
}

// DenyRequestTimeOff is the resolver for the denyRequestTimeOff field.
func (r *mutationResolver) DenyRequestTimeOff(ctx context.Context, id string, responseNote *string, authUserID *string) (*model.TimeOffResponse, error) {
	var shiftError []*model.ShiftError
	fieldError := "Deny Request Time Off"
	errorMessage := "Something went wrong while denying the Request Time Off." // default error message

	if authUserID == nil || *authUserID == string("") {
		errorMessage = "Authenticated user id is required"
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeRequired,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	var err error
	permission := false

	// validate permission
	permission, err = util.CheckPermission("request_time_off", "MANAGE_ALL", *authUserID)
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeNotFound,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	if !permission {
		permission, err = util.CheckPermission("request_time_off", "MANAGE", *authUserID)
		if err != nil {
			sentry.CaptureException(err)
			defer sentry.Flush(2 * time.Second)

			errorMessage = err.Error()
			shiftError = append(shiftError, &model.ShiftError{
				Code:    model.ShiftErrorCodeNotFound,
				Field:   &fieldError,
				Message: &errorMessage,
			})

			return &model.TimeOffResponse{
				Errors:  shiftError,
				Request: nil,
			}, nil
		}
	}

	if !permission {
		errorMessage = "Permission denied: request_time_off.MANAGE, request_time_off.MANAGE_ALL"
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeRequired,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	if id == "" {
		errorMessage = "Request Time Off id is required"
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeRequired,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	status := model.RequestStatusDenied
	var requestTimeOff *model.RequestTimeOff
	deny := &model.RequestTimeOff{
		Status:       status,
		ResponseNote: responseNote,
	}
	// update request time off status to accepted
	err = r.DB.Model(&requestTimeOff).Where("id= ?", id).Updates(deny).Error

	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	// get request time off by channel id and request id
	err = r.DB.Where("id= ?", id).First(&requestTimeOff).Error

	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeNotFound,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	// get the user from the user service and return it
	user, err := util.GetUser(requestTimeOff.UserID)

	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		errorMessage = err.Error()
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	if user.ID == nil || *user.ID == string("") {
		errorMessage = "Something went wrong while fetching the user."
		shiftError = append(shiftError, &model.ShiftError{
			Code:    model.ShiftErrorCodeInvalid,
			Field:   &fieldError,
			Message: &errorMessage,
		})

		return &model.TimeOffResponse{
			Errors:  shiftError,
			Request: nil,
		}, nil
	}

	requestResponse := model.RequestResponse{
		ID:           requestTimeOff.ID,
		ChannelID:    *requestTimeOff.ChannelID,
		RequestID:    *requestTimeOff.RequestID,
		RequestNote:  requestTimeOff.RequestNote,
		Status:       &requestTimeOff.Status,
		ResponseNote: requestTimeOff.ResponseNote,
		ResponseAt:   requestTimeOff.ResponseAt,
		CreatedAt:    &requestTimeOff.CreatedAt,
		User:         user,
	}

	return &model.TimeOffResponse{
		Errors:  nil,
		Request: &requestResponse,
	}, nil
}

// GetRequestTimeOff is the resolver for the getRequestTimeOff field.
func (r *queryResolver) GetRequestTimeOff(ctx context.Context, id string, authUserID *string) (*model.RequestTimeOff, error) {
	if authUserID == nil || *authUserID == string("") {
		return nil, fmt.Errorf("Authenticated user id is required")
	}

	var err error
	permission := false

	// validate permission
	permission, err = util.CheckPermission("request_time_off", "READ_ALL", *authUserID)
	if err != nil {

		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		return nil, err
	}

	if !permission {
		permission, err = util.CheckPermission("request_time_off", "READ", *authUserID)
		if err != nil {
			sentry.CaptureException(err)
			defer sentry.Flush(2 * time.Second)

			return nil, err
		}
	}

	if !permission {
		return nil, fmt.Errorf("Permission denied: request_time_off.READ, request_time_off.READ_ALL")
	}
	if id == "" {
		return nil, errors.New("id is required")
	}

	// get request time off by channel id and request id
	var requestTimeOff model.RequestTimeOff
	err = r.DB.First(&requestTimeOff, "id= ?", id).Error

	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		return nil, err
	}

	return &requestTimeOff, nil
}

// GetRequestTimeOffs is the resolver for the getRequestTimeOffs field.
func (r *queryResolver) GetRequestTimeOffs(ctx context.Context, authUserID *string) ([]*model.RequestTimeOff, error) {
	if authUserID == nil || *authUserID == string("") {
		return nil, fmt.Errorf("Authenticated user id is required")
	}

	var err error
	permission := false

	// validate permission
	permission, err = util.CheckPermission("request_time_off", "READ_ALL", *authUserID)
	if err != nil {

		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		return nil, err
	}

	// if !permission {
	// 	permission, err = util.CheckPermission("request_time_off", "READ", *authUserID)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	if !permission {
		return nil, fmt.Errorf("Permission denied: request_time_off.READ_ALL")
	}

	// get all request time offs
	var requestTimeOffs []*model.RequestTimeOff
	err = r.DB.Find(&requestTimeOffs).Error

	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		return nil, err
	}

	return requestTimeOffs, nil
}

// GetRequestTimeOffsByChannelIDRequestID is the resolver for the getRequestTimeOffsByChannelIdRequestId field.
func (r *queryResolver) GetRequestTimeOffsByChannelIDRequestID(ctx context.Context, channelID string, requestID string) (*model.RequestTimeOff, error) {
	if channelID == "" || requestID == "" {
		return nil, errors.New("channelID and requestID are required")
	}

	var err error

	// get all request time offs
	var requestTimeOffs *model.RequestTimeOff
	err = r.DB.First(&requestTimeOffs, "channel_id = ? AND request_id = ?", channelID, requestID).Error

	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		return nil, err
	}

	return requestTimeOffs, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
