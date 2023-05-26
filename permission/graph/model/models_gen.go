package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type GetGrantedPermissionsResponse struct {
	FirstName   string               `json:"firstName"`
	LastName    *string              `json:"lastName"`
	Email       *string              `json:"email"`
	Permissions []*GrantedPermission `json:"permissions"`
}

// the number of users that have this permission or the number of users that have this permission group also can be counted by the number of granted permissions
type GrantedPermission struct {
	ID        string `json:"id"`
	NameSpace string `json:"nameSpace"`
	// Auth User (ID)
	UserID string `json:"userId"`
	// Relation between a Subject and a Object
	Permission string `json:"permission"`
	// eg. 'User', 'Post', 'open shift'
	Object    string    `json:"object"`
	GrantedAt time.Time `json:"grantedAt"`
}

type GrantedPermissionInput struct {
	NameSpace  NameSpaceEnum  `json:"nameSpace"`
	UserID     string         `json:"userId"`
	Permission PermissionEnum `json:"permission"`
	Object     string         `json:"object"`
}

type GrantedPermissionResponse struct {
	PermissionID *string    `json:"permissionId"`
	NameSpace    *string    `json:"nameSpace"`
	Permission   *string    `json:"permission"`
	Object       *string    `json:"object"`
	GrantedAt    *time.Time `json:"grantedAt"`
	User         *User      `json:"user"`
}

type User struct {
	ID        *string `json:"id"`
	Email     *string `json:"email"`
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
}

type NameSpaceEnum string

const (
	NameSpaceEnumShifts  NameSpaceEnum = "shifts"
	NameSpaceEnumBooking NameSpaceEnum = "booking"
)

var AllNameSpaceEnum = []NameSpaceEnum{
	NameSpaceEnumShifts,
	NameSpaceEnumBooking,
}

func (e NameSpaceEnum) IsValid() bool {
	switch e {
	case NameSpaceEnumShifts, NameSpaceEnumBooking:
		return true
	}
	return false
}

func (e NameSpaceEnum) String() string {
	return string(e)
}

func (e *NameSpaceEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = NameSpaceEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid NameSpaceEnum", str)
	}
	return nil
}

func (e NameSpaceEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PermissionEnum string

const (
	PermissionEnumRead      PermissionEnum = "READ"
	PermissionEnumReadAll   PermissionEnum = "READ_ALL"
	PermissionEnumWrite     PermissionEnum = "WRITE"
	PermissionEnumWriteAll  PermissionEnum = "WRITE_ALL"
	PermissionEnumManage    PermissionEnum = "MANAGE"
	PermissionEnumManageAll PermissionEnum = "MANAGE_ALL"
)

var AllPermissionEnum = []PermissionEnum{
	PermissionEnumRead,
	PermissionEnumReadAll,
	PermissionEnumWrite,
	PermissionEnumWriteAll,
	PermissionEnumManage,
	PermissionEnumManageAll,
}

func (e PermissionEnum) IsValid() bool {
	switch e {
	case PermissionEnumRead, PermissionEnumReadAll, PermissionEnumWrite, PermissionEnumWriteAll, PermissionEnumManage, PermissionEnumManageAll:
		return true
	}
	return false
}

func (e PermissionEnum) String() string {
	return string(e)
}

func (e *PermissionEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PermissionEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PermissionEnum", str)
	}
	return nil
}

func (e PermissionEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
