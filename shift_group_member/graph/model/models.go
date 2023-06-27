package model

import "time"

type GetAssignedShiftsResponse struct {
	Data struct {
		GetAssignedShiftsByChannelIDShiftGroupIDUserID []struct {
			ID           string    `json:"id"`
			Break        string    `json:"break"`
			Label        *string   `json:"label"`
			Note         *string   `json:"note"`
			StartTime    time.Time `json:"startTime"`
			UserID       *string   `json:"userId"`
			ChannelID    *string   `json:"channelId"`
			ShiftGroupID *string   `json:"shiftGroupId"`
			Is24Hours    bool      `json:"is24Hours"`
			EndTime      time.Time `json:"endTime"`
			IsShared     *bool     `json:"isShared"`
			Type         *string   `json:"type"`
			IsOpen       *bool     `json:"isOpen"`
		} `json:"getAssignedShiftsByChannelIdShiftGroupIdUserId"`
	} `json:"data"`
}

type GetUniqueAssignedShiftsResponse struct {
	Data struct {
		GetAssignedShiftsByChannelIDShiftGroupID []struct {
			ID              string    `json:"id"`
			Break           string    `json:"break"`
			Label           *string   `json:"label"`
			Color           *string   `json:"color"`
			StartTime       time.Time `json:"startTime"`
			EndTime         time.Time `json:"endTime"`
			Is24Hours       bool      `json:"is24Hours"`
			UserID          *string   `json:"userId"`
			ChannelID       *string   `json:"channelId"`
			ShiftGroupID    *string   `json:"shiftGroupId"`
			Type            *string   `json:"type"`
			IsOpen          *bool     `json:"isOpen"`
			Note            *string   `json:"note"`
			IsShared        *bool     `json:"isShared"`
			ShiftActivities []struct {
				ID        string    `json:"id"`
				Name      *string   `json:"name"`
				Code      *string   `json:"code"`
				Color     *string   `json:"color"`
				StartTime time.Time `json:"startTime"`
				EndTime   time.Time `json:"endTime"`
				UserID    *string   `json:"userId"`
				IsPaid    bool      `json:"isPaid"`
			} `json:"ShiftActivities"`
		} `json:"getAssignedShiftsByChannelIdShiftGroupId"`
	} `json:"data"`
}

type GetUniqueOpenShiftsResponse struct {
	Data struct {
		GetOpenShifts []struct {
			ID              string     `json:"id"`
			ChannelID       *string    `json:"channelId"`
			Break           *string    `json:"break"`
			ShiftGroupID    *string    `json:"shiftGroupId"`
			Color           *string    `json:"color"`
			EndTime         *time.Time `json:"endTime"`
			StartTime       *time.Time `json:"startTime"`
			Slots           *int       `json:"slots"`
			CreatedAt       time.Time  `json:"createdAt"`
			Is24Hours       bool       `json:"is24Hours"`
			Label           *string    `json:"label"`
			Note            *string    `json:"note"`
			ShiftActivities []struct {
				ID           string    `json:"id"`
				ChannelID    *string   `json:"channelId"`
				ShiftGroupID *string   `json:"shiftGroupId"`
				OpenShiftID  string    `json:"openShiftId"`
				StartTime    time.Time `json:"startTime"`
				EndTime      time.Time `json:"endTime"`
				Name         string    `json:"name"`
				Code         *string   `json:"code"`
				Color        *string   `json:"color"`
				CreatedAt    time.Time `json:"createdAt"`
				IsPaid       bool      `json:"isPaid"`
			} `json:"ShiftActivities"`
		} `json:"getOpenShifts"`
	} `json:"data"`
}

type AssignedShiftDeleteResponse struct {
	Data struct {
		DeleteAssignedShift struct {
			AssignedShift struct {
				ID    string `json:"id"`
				Label string `json:"label"`
			} `json:"assignedShift"`
		} `json:"deleteAssignedShift"`
	} `json:"data"`
}

type UserResponse struct {
	Data struct {
		User struct {
			ID           string     `json:"id"`
			Email        string     `json:"email"`
			FirstName    string     `json:"firstName"`
			LastName     string     `json:"lastName"`
			IsStaff      bool       `json:"isStaff"`
			IsActive     bool       `json:"isActive"`
			Note         *string    `json:"note"`
			Avatar       *string    `json:"avatar"`
			LanguageCode string     `json:"languageCode"`
			LastLogin    *time.Time `json:"lastLogin"`
			DateJoined   *time.Time `json:"dateJoined"`
		}
	} `json:"data"`
}

type UserIsStaffResponse struct {
	Data struct {
		GetUserByIsStaff []struct {
			ID           string    `json:"id"`
			Email        string    `json:"email"`
			FirstName    string    `json:"firstName"`
			LastName     string    `json:"lastName"`
			Avatar       *string   `json:"avatar"`
			DateJoined   time.Time `json:"dateJoined"`
			IsActive     bool      `json:"isActive"`
			Note         *string   `json:"note"`
			LanguageCode string    `json:"languageCode"`
			IsStaff      bool      `json:"isStaff"`
		} `json:"getUserByIsStaff"`
	} `json:"data"`
}

type GetUsersResponse struct {
	Data struct {
		Users []struct {
			ID           string    `json:"id"`
			Email        string    `json:"email"`
			FirstName    string    `json:"firstName"`
			LastName     string    `json:"lastName"`
			Avatar       *string   `json:"avatar"`
			DateJoined   time.Time `json:"dateJoined"`
			IsActive     bool      `json:"isActive"`
			Note         *string   `json:"note"`
			LanguageCode string    `json:"languageCode"`
			IsStaff      bool      `json:"isStaff"`
		} `json:"users"`
	} `json:"data"`
}

type GetShiftGroupResponse struct {
	Data struct {
		ShiftGroupsByChannel []struct {
			ID        string `json:"id"`
			ChannelID string `json:"channelId"`
			Name      string `json:"name"`
		} `json:"shiftGroupsByChannel"`
	} `json:"data"`
}

type ShiftGroup struct {
	ID        string `json:"id"`
	ChannelID string `json:"channelId"`
	Name      string `json:"name"`
	Position  *int   `json:"position"`
}

type GetOpenShiftsResponse struct {
	Data struct {
		GetOpenShiftsByTime []struct {
			ID              string     `json:"id"`
			ChannelID       *string    `json:"channelId"`
			Break           *string    `json:"break"`
			ShiftGroupID    *string    `json:"shiftGroupId"`
			Color           *string    `json:"color"`
			EndTime         *time.Time `json:"endTime"`
			StartTime       *time.Time `json:"startTime"`
			Slots           *int       `json:"slots"`
			Is24Hours       bool       `json:"is24Hours"`
			Label           *string    `json:"label"`
			Note            *string    `json:"note"`
			ShiftActivities []struct {
				ID           string    `json:"id"`
				ChannelID    *string   `json:"channelId"`
				ShiftGroupID string    `json:"shiftGroupId"`
				OpenShiftID  string    `json:"openShiftId"`
				StartTime    time.Time `json:"startTime"`
				EndTime      time.Time `json:"endTime"`
				Name         string    `json:"name"`
				Code         *string   `json:"code"`
				Color        *string   `json:"color"`
				IsPaid       bool      `json:"isPaid"`
			} `json:"ShiftActivities"`
		} `json:"getOpenShiftsByTime"`
	} `json:"data"`
}

type GetTimeOffResponse struct {
	Data struct {
		GetTimeOffs []struct {
			ID           string     `json:"id"`
			UserID       *string    `json:"userId,omitempty"`
			ChannelID    *string    `json:"channelId,omitempty"`
			ShiftGroupID *string    `json:"shiftGroupId,omitempty"`
			StartTime    time.Time  `json:"startTime"`
			EndTime      time.Time  `json:"endTime"`
			Is24Hours    bool       `json:"is24Hours"`
			Label        string     `json:"label"`
			Color        string     `json:"color"`
			Note         string     `json:"note"`
			CreatedAt    *time.Time `json:"createdAt,omitempty"`
		} `json:"getTimeOffs"`
	} `json:"data"`
}
type GetAssignedShiftsByTimeResponse struct {
	Data struct {
		GetAssignedShiftsByTime []struct {
			ID              string    `json:"id"`
			Break           string    `json:"break"`
			Label           *string   `json:"label"`
			Color           string    `json:"color"`
			StartTime       time.Time `json:"startTime"`
			EndTime         time.Time `json:"endTime"`
			Is24Hours       bool      `json:"is24Hours"`
			UserID          *string   `json:"userId"`
			ChannelID       *string   `json:"channelId"`
			ShiftGroupID    *string   `json:"shiftGroupId"`
			Type            *string   `json:"type"`
			IsOpen          *bool     `json:"isOpen"`
			Note            *string   `json:"note"`
			IsShared        *bool     `json:"isShared"`
			ShiftActivities []struct {
				ID        string    `json:"id"`
				Name      *string   `json:"name"`
				Code      *string   `json:"code"`
				Color     *string   `json:"color"`
				StartTime time.Time `json:"startTime"`
				EndTime   time.Time `json:"endTime"`
				UserID    *string   `json:"userId"`
				IsPaid    bool      `json:"isPaid"`
			} `json:"ShiftActivities"`
		} `json:"getAssignedShiftsByTime"`
	} `json:"data"`
}

type ChannelResPonse struct {
	Data struct {
		Channel struct {
			ID   string `json:"id"`
			Slug string `json:"slug"`
			Name string `json:"name"`
		} `json:"channel"`
	} `json:"data"`
}

type PermissionResponse struct {
	Data struct {
		CheckPermission bool `json:"CheckPermission"`
	} `json:"data"`
}

type GetRequestsResponse struct {
	Data struct {
		GetRequestsByUser struct {
			Edges []struct {
				Node struct {
					ID             string    `json:"id"`
					RequestID      string    `json:"requestId"`
					Status         string    `json:"status"`
					StartTime      time.Time `json:"startTime"`
					EndTime        time.Time `json:"endTime"`
					RequestNote    *string   `json:"requestNote"`
					Reason         string    `json:"reason"`
					ResponseNote   string    `json:"responseNote"`
					ChannelID      string    `json:"channelId"`
					IsAllDay       bool      `json:"isAllDay"`
					Type           *string   `json:"type"`
					ShiftOfferedTo struct {
						ID        string `json:"id"`
						FirstName string `json:"firstName"`
						LastName  string `json:"lastName"`
						Email     string `json:"email"`
					} `json:"shiftOfferedTo"`
					ShiftToSwap AssignedShift `json:"shiftToSwap"`
					ToSwapWith  AssignedShift `json:"toSwapWith"`
				} `json:"node"`
			} `json:"edges"`
		} `json:"getRequestsByUser"`
	} `json:"data"`
}
