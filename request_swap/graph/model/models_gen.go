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
	Label           *string                    `json:"label"`
	Note            *string                    `json:"note"`
	ShiftToOffer    *AssignedShift             `json:"shiftToOffer"`
	ShiftToSwap     *AssignedShift             `json:"shiftToSwap"`
	ToSwapWith      *AssignedShift             `json:"toSwapWith"`
	UserID          *string                    `json:"userId"`
	ChannelID       *string                    `json:"channelId"`
	ShiftGroupID    *string                    `json:"shiftGroupId"`
	Type            *string                    `json:"type"`
	IsOpen          *bool                      `json:"isOpen"`
	IsShared        *bool                      `json:"isShared"`
	ShiftActivities []*AssignedShiftActivities `json:"ShiftActivities"`
}

type AssignedShiftActivities struct {
	ID              string    `json:"id"`
	ChannelID       *string   `json:"channelId"`
	ShiftGroupID    *string   `json:"shiftGroupId"`
	AssignedShiftID string    `json:"assignedShiftId"`
	UserID          *string   `json:"userId"`
	Name            *string   `json:"name"`
	Code            *string   `json:"code"`
	Color           *string   `json:"color"`
	StartTime       time.Time `json:"startTime"`
	EndTime         time.Time `json:"endTime"`
	IsPaid          bool      `json:"isPaid"`
}

type RequestResponse struct {
	ChannelID      string         `json:"channelId"`
	CreatedAt      *time.Time     `json:"createdAt"`
	EndTime        *time.Time     `json:"endTime"`
	ID             string         `json:"id"`
	IsAllDay       *bool          `json:"isAllDay"`
	Reason         *string        `json:"reason"`
	RequestID      string         `json:"requestId"`
	RequestNote    *string        `json:"requestNote"`
	ResponseAt     *time.Time     `json:"responseAt"`
	ResponseBy     *User          `json:"responseBy"`
	ResponseNote   *string        `json:"responseNote"`
	ShiftOfferedTo *User          `json:"shiftOfferedTo"`
	ShiftToOffer   *AssignedShift `json:"shiftToOffer"`
	ShiftToSwap    *AssignedShift `json:"shiftToSwap"`
	StartTime      *time.Time     `json:"startTime"`
	Status         *RequestStatus `json:"status"`
	ToSwapWith     *AssignedShift `json:"toSwapWith"`
	Type           *RequestType   `json:"type"`
	User           *User          `json:"user"`
}

type RequestSwap struct {
	ID                        string        `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	ChannelID                 string        `json:"channelId" gorm:"type:varchar(64)"`
	RequestID                 *string       `json:"requestId" gorm:"type:uuid; not null"`
	UserID                    string        `json:"userId" gorm:"type:varchar(64); not null"`
	AssignedUserShiftID       *string       `json:"assignedUserShiftId" gorm:"type:uuid; not null"`
	AssignedUserShiftIDToSwap *string       `json:"assignedUserShiftIdToSwap" gorm:"type:uuid; not null"`
	RequestNote               *string       `json:"requestNote"`
	Status                    RequestStatus `json:"status" gorm:"type:varchar(16); not null"`
	ResponseNote              *string       `json:"responseNote"`
	ResponseByUserID          *string       `json:"responseByUserId"`
	ResponseAt                *time.Time    `json:"responseAt"`
	CreatedAt                 time.Time     `json:"createdAt" gorm:"default:now()"`
}

type RequestSwapInput struct {
	ChannelID                 string     `json:"channelId"`
	UserID                    string     `json:"userId"`
	AssignedUserShiftID       string     `json:"assignedUserShiftId"`
	AssignedUserShiftIDToSwap *string    `json:"assignedUserShiftIdToSwap"`
	RequestNote               *string    `json:"requestNote"`
	ResponseNote              *string    `json:"responseNote"`
	ResponseByUserID          *string    `json:"responseByUserId"`
	ResponseAt                *time.Time `json:"responseAt"`
}

type RequestSwapResponse struct {
	Errors  []*ShiftError    `json:"errors"`
	Request *RequestResponse `json:"request"`
}

type ShiftError struct {
	Code    ShiftErrorCode `json:"code"`
	Field   *string        `json:"field"`
	Message *string        `json:"message"`
}

type User struct {
	ID        *string `json:"id"`
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
	Email     *string `json:"email"`
	Avatar    *string `json:"avatar"`
	IsActive  *bool   `json:"isActive"`
	IsStaff   *bool   `json:"isStaff"`
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
