// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type AssignedShift struct {
	ID              string                     `json:"id"`
	Break           string                     `json:"break"`
	Color           string                     `json:"color"`
	StartTime       time.Time                  `json:"startTime"`
	EndTime         time.Time                  `json:"endTime"`
	Is24Hours       bool                       `json:"is24Hours"`
	Label           *string                    `json:"label,omitempty"`
	Note            *string                    `json:"note,omitempty"`
	ShiftToOffer    *AssignedShift             `json:"shiftToOffer,omitempty"`
	ShiftToSwap     *AssignedShift             `json:"shiftToSwap,omitempty"`
	ToSwapWith      *AssignedShift             `json:"toSwapWith,omitempty"`
	UserID          *string                    `json:"userId,omitempty"`
	ChannelID       *string                    `json:"channelId,omitempty"`
	ShiftGroupID    *string                    `json:"shiftGroupId,omitempty"`
	Type            *string                    `json:"type,omitempty"`
	IsOpen          *bool                      `json:"isOpen,omitempty"`
	IsShared        *bool                      `json:"isShared,omitempty"`
	ShiftActivities []*AssignedShiftActivities `json:"ShiftActivities,omitempty"`
}

type AssignedShiftActivities struct {
	ID              string    `json:"id"`
	ChannelID       *string   `json:"channelId,omitempty"`
	ShiftGroupID    *string   `json:"shiftGroupId,omitempty"`
	AssignedShiftID string    `json:"assignedShiftId"`
	UserID          *string   `json:"userId,omitempty"`
	Name            *string   `json:"name,omitempty"`
	Code            *string   `json:"code,omitempty"`
	Color           *string   `json:"color,omitempty"`
	StartTime       time.Time `json:"startTime"`
	EndTime         time.Time `json:"endTime"`
	IsPaid          bool      `json:"isPaid"`
}

type RequestResponse struct {
	ChannelID      string         `json:"channelId"`
	CreatedAt      *time.Time     `json:"createdAt,omitempty"`
	EndTime        *time.Time     `json:"endTime,omitempty"`
	ID             string         `json:"id"`
	IsAllDay       *bool          `json:"isAllDay,omitempty"`
	Reason         *string        `json:"reason,omitempty"`
	RequestID      string         `json:"requestId"`
	RequestNote    *string        `json:"requestNote,omitempty"`
	ResponseAt     *time.Time     `json:"responseAt,omitempty"`
	ResponseBy     *User          `json:"responseBy,omitempty"`
	ResponseNote   *string        `json:"responseNote,omitempty"`
	ShiftOfferedTo *User          `json:"shiftOfferedTo,omitempty"`
	ShiftToOffer   *AssignedShift `json:"shiftToOffer,omitempty"`
	ShiftToSwap    *AssignedShift `json:"shiftToSwap,omitempty"`
	StartTime      *time.Time     `json:"startTime,omitempty"`
	Status         *RequestStatus `json:"status,omitempty"`
	ToSwapWith     *AssignedShift `json:"toSwapWith,omitempty"`
	Type           *RequestType   `json:"type,omitempty"`
	User           *User          `json:"user"`
}

type RequestSwap struct {
	ID                        string        `json:"id"`
	ChannelID                 string        `json:"channelId"`
	RequestID                 *string       `json:"requestId,omitempty"`
	UserID                    string        `json:"userId"`
	AssignedUserShiftID       *string       `json:"assignedUserShiftId,omitempty"`
	AssignedUserShiftIDToSwap *string       `json:"assignedUserShiftIdToSwap,omitempty"`
	RequestNote               *string       `json:"requestNote,omitempty"`
	Status                    RequestStatus `json:"status"`
	ResponseNote              *string       `json:"responseNote,omitempty"`
	ResponseByUserID          *string       `json:"responseByUserId,omitempty"`
	ResponseAt                *time.Time    `json:"responseAt,omitempty"`
	CreatedAt                 time.Time     `json:"createdAt"`
}

type RequestSwapInput struct {
	ChannelID                 string     `json:"channelId"`
	UserID                    string     `json:"userId"`
	AssignedUserShiftID       string     `json:"assignedUserShiftId"`
	AssignedUserShiftIDToSwap *string    `json:"assignedUserShiftIdToSwap,omitempty"`
	RequestNote               *string    `json:"requestNote,omitempty"`
	ResponseNote              *string    `json:"responseNote,omitempty"`
	ResponseByUserID          *string    `json:"responseByUserId,omitempty"`
	ResponseAt                *time.Time `json:"responseAt,omitempty"`
}

type RequestSwapResponse struct {
	Errors  []*ShiftError    `json:"errors"`
	Request *RequestResponse `json:"request,omitempty"`
}

type ShiftError struct {
	Code    ShiftErrorCode `json:"code"`
	Field   *string        `json:"field,omitempty"`
	Message *string        `json:"message,omitempty"`
}

type User struct {
	ID        *string `json:"id,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
	Email     *string `json:"email,omitempty"`
	Avatar    *string `json:"avatar,omitempty"`
	IsActive  *bool   `json:"isActive,omitempty"`
	IsStaff   *bool   `json:"isStaff,omitempty"`
}

type RequestStatus string

const (
	RequestStatusPending   RequestStatus = "PENDING"
	RequestStatusApproved  RequestStatus = "APPROVED"
	RequestStatusDenied    RequestStatus = "DENIED"
	RequestStatusCancelled RequestStatus = "CANCELLED"
)

var AllRequestStatus = []RequestStatus{
	RequestStatusPending,
	RequestStatusApproved,
	RequestStatusDenied,
	RequestStatusCancelled,
}

func (e RequestStatus) IsValid() bool {
	switch e {
	case RequestStatusPending, RequestStatusApproved, RequestStatusDenied, RequestStatusCancelled:
		return true
	}
	return false
}

func (e RequestStatus) String() string {
	return string(e)
}

func (e *RequestStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RequestStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RequestStatus", str)
	}
	return nil
}

func (e RequestStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type RequestType string

const (
	RequestTypeOffer   RequestType = "OFFER"
	RequestTypeSwap    RequestType = "SWAP"
	RequestTypeTimeoff RequestType = "TIMEOFF"
)

var AllRequestType = []RequestType{
	RequestTypeOffer,
	RequestTypeSwap,
	RequestTypeTimeoff,
}

func (e RequestType) IsValid() bool {
	switch e {
	case RequestTypeOffer, RequestTypeSwap, RequestTypeTimeoff:
		return true
	}
	return false
}

func (e RequestType) String() string {
	return string(e)
}

func (e *RequestType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RequestType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RequestType", str)
	}
	return nil
}

func (e RequestType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ShiftErrorCode string

const (
	ShiftErrorCodeGraphqlError ShiftErrorCode = "GRAPHQL_ERROR"
	ShiftErrorCodeInvalid      ShiftErrorCode = "INVALID"
	ShiftErrorCodeNotFound     ShiftErrorCode = "NOT_FOUND"
	ShiftErrorCodeRequired     ShiftErrorCode = "REQUIRED"
)

var AllShiftErrorCode = []ShiftErrorCode{
	ShiftErrorCodeGraphqlError,
	ShiftErrorCodeInvalid,
	ShiftErrorCodeNotFound,
	ShiftErrorCodeRequired,
}

func (e ShiftErrorCode) IsValid() bool {
	switch e {
	case ShiftErrorCodeGraphqlError, ShiftErrorCodeInvalid, ShiftErrorCodeNotFound, ShiftErrorCodeRequired:
		return true
	}
	return false
}

func (e ShiftErrorCode) String() string {
	return string(e)
}

func (e *ShiftErrorCode) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ShiftErrorCode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ShiftErrorCode", str)
	}
	return nil
}

func (e ShiftErrorCode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
