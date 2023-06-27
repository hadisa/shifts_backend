package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"shift_group_members/graph/model"
	"strconv"
	"time"

	"github.com/getsentry/sentry-go"
)

func GetAssignedShiftsByChannelIDShiftGroupIDUserID(channelId *string, shiftGroupId *string, userId *string) ([]*model.AssignedShift, error) {

	jsonMapInstance := map[string]string{
		"query": `
		{
			getAssignedShiftsByChannelIdShiftGroupIdUserId(
			  channelId: "` + *channelId + `"
			  shiftGroupId: "` + *shiftGroupId + `"
			  userId: "` + *userId + `"
			) {
			  id
			  break
			  label
			  note
			  startTime
			  userId
			  channelId
			  shiftGroupId
			  is24Hours
			  endTime
			  isShared
			  type
			  isOpen
			}
		  }
		`,
	}

	responseData, err := httpRequest("POST", os.Getenv("ASSIGNED_SHIFT_API"), jsonMapInstance)
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		fmt.Printf("There was an error executing the request%v", err)
		return nil, err
	}

	// convert the responseData to GetAssignedShiftsResponse
	var responseObject model.GetAssignedShiftsResponse
	json.Unmarshal(responseData, &responseObject)

	// map all responseObject GetAssignedShiftsResponse
	// to AssignedShifts
	var assignedShifts []*model.AssignedShift
	for _, assignedShift := range responseObject.Data.GetAssignedShiftsByChannelIDShiftGroupIDUserID {
		assignedShifts = append(assignedShifts, &model.AssignedShift{
			ID:           assignedShift.ID,
			Break:        assignedShift.Break,
			Label:        assignedShift.Label,
			Note:         assignedShift.Note,
			StartTime:    assignedShift.StartTime,
			UserID:       assignedShift.UserID,
			ChannelID:    assignedShift.ChannelID,
			ShiftGroupID: assignedShift.ShiftGroupID,
			Is24Hours:    assignedShift.Is24Hours,
			EndTime:      assignedShift.EndTime,
			IsShared:     assignedShift.IsShared,
			Type:         assignedShift.Type,
			IsOpen:       assignedShift.IsOpen,
		})
	}

	return assignedShifts, nil
}

func GetAssignedShiftsByChannelIDShiftGroupID(channelId *string, shiftGroupId *string) ([]*model.AssignedShift, error) {

	jsonMapInstance := map[string]string{
		"query": `
		{
			getAssignedShiftsByChannelIdShiftGroupId(
			  channelId: "` + *channelId + `"
			  shiftGroupId: "` + *shiftGroupId + `"
			) {
			  id
			  break
			  label
			  color
			  startTime
			  endTime
			  is24Hours
			  userId
			  channelId
			  shiftGroupId
			  type
			  isOpen
			  note
			  isShared
			  ShiftActivities {
				id
				name
				code
				color
				startTime
				endTime
				userId
				isPaid
			  }
			}
		  }
		`,
	}

	responseData, err := httpRequest("POST", os.Getenv("ASSIGNED_SHIFT_API"), jsonMapInstance)
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		fmt.Printf("There was an error executing the request%v", err)
		return nil, err
	}

	// convert the responseData to GetAssignedShiftsResponse
	var responseObject model.GetUniqueAssignedShiftsResponse
	json.Unmarshal(responseData, &responseObject)

	var assignedShifts []*model.AssignedShift
	for _, assignedShift := range responseObject.Data.GetAssignedShiftsByChannelIDShiftGroupID {
		assignedShifts = append(assignedShifts, &model.AssignedShift{
			ID:           assignedShift.ID,
			Break:        assignedShift.Break,
			Label:        assignedShift.Label,
			Note:         assignedShift.Note,
			Color:        *assignedShift.Color,
			StartTime:    assignedShift.StartTime,
			UserID:       assignedShift.UserID,
			ChannelID:    assignedShift.ChannelID,
			ShiftGroupID: assignedShift.ShiftGroupID,
			Is24Hours:    assignedShift.Is24Hours,
			EndTime:      assignedShift.EndTime,
			IsShared:     assignedShift.IsShared,
			Type:         assignedShift.Type,
			IsOpen:       assignedShift.IsOpen,
		})

		for _, shiftActivity := range assignedShift.ShiftActivities {
			assignedShifts[len(assignedShifts)-1].ShiftActivities = append(assignedShifts[len(assignedShifts)-1].ShiftActivities, &model.AssignedShiftActivities{
				ID:        shiftActivity.ID,
				Name:      shiftActivity.Name,
				Code:      shiftActivity.Code,
				Color:     shiftActivity.Color,
				StartTime: shiftActivity.StartTime,
				EndTime:   shiftActivity.EndTime,
				UserID:    shiftActivity.UserID,
				IsPaid:    shiftActivity.IsPaid,
			})
		}
	}

	return assignedShifts, nil
}

func GetOpenShiftsByChannelIDShiftGroupID(channelId *string, shiftGroupId *string, authUserId *string) ([]*model.OpenShift, error) {

	jsonMapInstance := map[string]string{
		"query": `
		{
			getOpenShifts(
			  channelId: "` + *channelId + `"
			  shiftGroupId: "` + *shiftGroupId + `"
			  authUserId: "` + *authUserId + `"
			) {
			  id
			  channelId
			  break
			  shiftGroupId
			  color
			  endTime
			  startTime
			  slots
			  createdAt
			  is24Hours
			  label
			  note
			  ShiftActivities {
				id
				channelId
				shiftGroupId
				openShiftId
				startTime
				endTime
				name
				code
				color
				createdAt
				isPaid
			  }
			}
		  }
		`,
	}

	responseData, err := httpRequest("POST", os.Getenv("OPEN_SHIFT_API"), jsonMapInstance)
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		fmt.Printf("There was an error executing the request%v", err)
		return nil, err
	}

	var responseObject model.GetUniqueOpenShiftsResponse
	json.Unmarshal(responseData, &responseObject)

	var openShifts []*model.OpenShift
	for _, openShift := range responseObject.Data.GetOpenShifts {
		openShifts = append(openShifts, &model.OpenShift{
			ID:           openShift.ID,
			ChannelID:    openShift.ChannelID,
			Break:        openShift.Break,
			ShiftGroupID: openShift.ShiftGroupID,
			Color:        openShift.Color,
			EndTime:      openShift.EndTime,
			StartTime:    openShift.StartTime,
			Slots:        openShift.Slots,
			CreatedAt:    openShift.CreatedAt,
			Is24Hours:    openShift.Is24Hours,
			Label:        openShift.Label,
			Note:         openShift.Note,
		})

		for _, shiftActivity := range openShift.ShiftActivities {
			openShifts[len(openShifts)-1].ShiftActivities = append(openShifts[len(openShifts)-1].ShiftActivities, &model.OpenShiftActivities{
				ID:           shiftActivity.ID,
				ChannelID:    shiftActivity.ChannelID,
				ShiftGroupID: shiftActivity.ShiftGroupID,
				OpenShiftID:  shiftActivity.OpenShiftID,
				StartTime:    shiftActivity.StartTime,
				EndTime:      shiftActivity.EndTime,
				Name:         shiftActivity.Name,
				Code:         shiftActivity.Code,
				Color:        shiftActivity.Color,
				IsPaid:       shiftActivity.IsPaid,
			})

		}

	}

	return openShifts, nil
}

func DeleteAssignedShift(assignedShiftId *string, authUserId *string) (string, error) {

	jsonMapInstance := map[string]string{
		"query": `
		mutation {
			deleteAssignedShift(id: "` + *assignedShiftId + `" authUserId: "` + *authUserId + `") {
				assignedShift{
					id
					label
				}
			}
		  }
		`,
	}

	responseData, err := httpRequest("POST", os.Getenv("ASSIGNED_SHIFT_API"), jsonMapInstance)
	if err != nil {

		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		fmt.Printf("There was an error executing the request%v", err)
		return "", err
	}

	// convert the responseData to AssignedShiftDeleteResponse
	var responseObject model.AssignedShiftDeleteResponse
	json.Unmarshal(responseData, &responseObject)

	if responseObject.Data.DeleteAssignedShift.AssignedShift.ID == "" {
		return "Failed to delete assigned shift", nil
	}
	return "success", nil

}

func DeleteTimeOff(channelId *string, shiftGroupId *string, userId *string, authUserId *string) (string, error) {

	jsonMapInstance := map[string]string{
		"query": `
		mutation {
			deleteTimeOffs(channelId: "` + *channelId + `" shiftGroupId: "` + *shiftGroupId + `" userId: "` + *userId + `" authUserId: "` + *authUserId + `")
		  }
		`,
	}

	responseData, err := httpRequest("POST", os.Getenv("TIME_OFF_API"), jsonMapInstance)
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		fmt.Printf("There was an error executing the request%v", err)
		return "", err
	}

	// convert the responseData to TimeOffDeleteResponse
	var responseObject model.TimeOffDeleteResponse
	json.Unmarshal(responseData, &responseObject)

	// if responseObject.Data.DeleteTimeOffs == "" {
	// 	return "Failed to delete time off", nil
	// }
	return "success", nil

}

func GetUser(id string) (*model.User, error) {
	jsonMapInstance := map[string]string{
		"query": `
		{
			user(id: "` + id + `") {
			  id
			  email
			  firstName
			  lastName
			  isStaff
			  isActive
			  note
			  avatar
			  languageCode
			  dateJoined
			}
		  }
		`,
	}

	responseData, err := httpRequest("POST", os.Getenv("USER_ACCOUNT_API"), jsonMapInstance)
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		fmt.Printf("There was an error executing the request%v", err)
		return nil, err
	}

	// convert into User struct
	var userData model.UserResponse
	err = json.Unmarshal(responseData, &userData)

	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		fmt.Printf("There was an error unmarshaling the JSON instance %v", err)
	}

	return &model.User{
		ID:           userData.Data.User.ID,
		Email:        userData.Data.User.Email,
		FirstName:    userData.Data.User.FirstName,
		LastName:     userData.Data.User.LastName,
		IsStaff:      userData.Data.User.IsStaff,
		IsActive:     userData.Data.User.IsActive,
		Note:         userData.Data.User.Note,
		Avatar:       userData.Data.User.Avatar,
		LanguageCode: userData.Data.User.LanguageCode,
		DateJoined:   *userData.Data.User.DateJoined,
	}, nil
}

func GetUsers(first *int, last *int) ([]*model.User, error) {

	var query string
	if first != nil {
		query = `
		{
		users(first: ` + strconv.Itoa(*first) + `) {
		  id
		  email
		  firstName
		  lastName
		  avatar
		  dateJoined
		  isActive
		  note
		  languageCode
		  isStaff
		}
	  }
	`
	} else if last != nil {
		query = `
		{
		users(last: ` + strconv.Itoa(*last) + `) {
		  id
		  email
		  firstName
		  lastName
		  avatar
		  dateJoined
		  isActive
		  note
		  languageCode
		  isStaff
		}
	  }
	`
	} else {
		query = `
			{
			users {
			  id
			  email
			  firstName
			  lastName
			  avatar
			  dateJoined
			  isActive
			  note
			  languageCode
			  isStaff
			}
		  }
		`
	}

	jsonMapInstance := map[string]string{
		"query": query,
	}

	responseData, err := httpRequest("POST", os.Getenv("USER_ACCOUNT_API"), jsonMapInstance)
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		fmt.Printf("There was an error executing the request%v", err)
		return nil, err
	}

	// convert into User struct
	var userData model.GetUsersResponse
	err = json.Unmarshal(responseData, &userData)

	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		fmt.Printf("There was an error unmarshaling the JSON instance %v", err)
	}

	var users []*model.User
	for _, user := range userData.Data.Users {
		users = append(users, &model.User{
			ID:           user.ID,
			Email:        user.Email,
			FirstName:    user.FirstName,
			LastName:     user.LastName,
			IsStaff:      user.IsStaff,
			IsActive:     user.IsActive,
			Note:         user.Note,
			Avatar:       user.Avatar,
			LanguageCode: user.LanguageCode,
			DateJoined:   user.DateJoined,
		})
	}

	return users, nil
}

func GetUsersByIsStaff(isStaff bool) ([]*model.User, error) {
	jsonMapInstance := map[string]string{
		"query": `
			{
			getUserByIsStaff(isStaff: ` + strconv.FormatBool(isStaff) + `) {
			  id
			  email
			  firstName
			  lastName
			  avatar
			  dateJoined
			  isActive
			  note
			  languageCode
			  isStaff
			}
		  }
		`,
	}

	responseData, err := httpRequest("POST", os.Getenv("USER_ACCOUNT_API"), jsonMapInstance)
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		fmt.Printf("There was an error executing the request%v", err)
		return nil, err
	}
	// convert into User struct
	var userData model.UserIsStaffResponse
	err = json.Unmarshal(responseData, &userData)

	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		fmt.Printf("There was an error unmarshaling the JSON instance %v", err)
	}

	var users []*model.User
	for _, user := range userData.Data.GetUserByIsStaff {
		users = append(users, &model.User{
			ID:           user.ID,
			Email:        user.Email,
			FirstName:    user.FirstName,
			LastName:     user.LastName,
			IsStaff:      user.IsStaff,
			IsActive:     user.IsActive,
			Note:         user.Note,
			Avatar:       user.Avatar,
			LanguageCode: user.LanguageCode,
			DateJoined:   user.DateJoined,
		})
	}

	return users, nil
}

/*
------------- function for get Shifts by task --------------------
*/
func GetShiftGroups(channelId *string, authUserId *string) ([]*model.ShiftGroup, error) {
	jsonMapInstance := map[string]string{
		"query": `
		{
			shiftGroupsByChannel(channelId: "` + *channelId + `" authUserId: "` + *authUserId + `") {
			  id
			  channelId
			  name
			}
		  }
		`,
	}

	responseData, err := httpRequest("POST", os.Getenv("SHIFT_GROUP_API"), jsonMapInstance)
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		fmt.Printf("There was an error executing the request%v", err)
		return nil, err
	}

	// convert into GetShiftGroupResponse struct
	var shiftGroupData model.GetShiftGroupResponse
	err = json.Unmarshal(responseData, &shiftGroupData)

	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		fmt.Printf("There was an error unmarshaling the JSON instance %v", err)
	}

	var shiftGroups []*model.ShiftGroup
	for _, shiftGroup := range shiftGroupData.Data.ShiftGroupsByChannel {
		shiftGroups = append(shiftGroups, &model.ShiftGroup{
			ID:        shiftGroup.ID,
			ChannelID: shiftGroup.ChannelID,
			Name:      shiftGroup.Name,
		})
	}

	return shiftGroups, nil

}

func GetOpenShiftsByTime(channelId *string, shiftGroupId *string, endTime *time.Time, startTime *time.Time) ([]*model.OpenShift, error) {

	jsonMapInstance := map[string]string{
		"query": `
		{
			getOpenShiftsByTime(
			  channelId: "` + *channelId + `"
			  shiftGroupId: "` + *shiftGroupId + `"
			  endTime: "` + fmt.Sprintf("%s", endTime.Format(time.RFC3339Nano)) + `"
			  startTime: "` + fmt.Sprintf("%s", startTime.Format(time.RFC3339Nano)) + `"
			) {
			  id
			  channelId
			  break
			  shiftGroupId
			  color
			  endTime
			  startTime
			  slots
			  is24Hours
			  label
			  note
			  ShiftActivities {
				id
				channelId
				shiftGroupId
				openShiftId
				startTime
				endTime
				name
				code
				color
				isPaid
			  }
			}
		  }
		`,
	}

	responseData, err := httpRequest("POST", os.Getenv("OPEN_SHIFT_API"), jsonMapInstance)
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		fmt.Printf("There was an error executing the request%v", err)
		return nil, err
	}

	var responseObject model.GetOpenShiftsResponse
	json.Unmarshal(responseData, &responseObject)

	// fmt.Println("getOpenShiftsByTime = ===>", responseData)

	var openShifts []*model.OpenShift

	// map responseObject into OpenShift struct

	if len(responseObject.Data.GetOpenShiftsByTime) == 0 {
		return nil, nil
	}

	for _, openShift := range responseObject.Data.GetOpenShiftsByTime {
		openShifts = append(openShifts, &model.OpenShift{
			ID:           openShift.ID,
			ChannelID:    openShift.ChannelID,
			Break:        openShift.Break,
			ShiftGroupID: openShift.ShiftGroupID,
			Color:        openShift.Color,
			EndTime:      openShift.EndTime,
			StartTime:    openShift.StartTime,
			Slots:        openShift.Slots,
			Is24Hours:    openShift.Is24Hours,
			Label:        openShift.Label,
			Note:         openShift.Note,
		})

		var shiftActivities []*model.OpenShiftActivities
		for _, shiftActivity := range openShift.ShiftActivities {

			shiftActivities = append(shiftActivities, &model.OpenShiftActivities{
				ID:        shiftActivity.ID,
				ChannelID: shiftActivity.ChannelID,
				StartTime: shiftActivity.StartTime,
				EndTime:   shiftActivity.EndTime,
				Name:      shiftActivity.Name,
				Code:      shiftActivity.Code,
				Color:     shiftActivity.Color,
				IsPaid:    shiftActivity.IsPaid,
			})
		}

		openShifts[len(openShifts)-1].ShiftActivities = shiftActivities

	}

	return openShifts, nil
}

// ------------------------------------------ i added

func GetTimeOffByTime(channelId *string, shiftGroupId *string, endTime *time.Time, startTime *time.Time, userId *string, authUserId *string) ([]*model.TimeOff, error) {

	jsonMapInstance := map[string]string{
		"query": `
		{
			getTimeOffs(
			  channelId: "` + *channelId + `"
			  shiftGroupId: "` + *shiftGroupId + `"
			  startTime: "` + fmt.Sprintf("%s", startTime.Format(time.RFC3339)) + `"
			  endTime: "` + fmt.Sprintf("%s", endTime.Format(time.RFC3339)) + `"
			  authUserId: "` + *authUserId + `"
			  userId:  "` + *userId + `"
			) {
				id
				userId
				channelId
				shiftGroupId
				is24Hours
				startTime
				endTime
				label
				color
				note
			  
			}
		  }
		`,
	}
	responseData, err := httpRequest("POST", os.Getenv("TIME_OFF_API"), jsonMapInstance)
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		fmt.Printf("There was an error executing the request%v", err)
		return nil, err
	}

	// convert the responseData to TimeOffDeleteResponse
	var responseObject model.GetTimeOffResponse
	json.Unmarshal(responseData, &responseObject)

	var assignedShifts []*model.TimeOff
	for _, item := range responseObject.Data.GetTimeOffs {
		var labelText = item.Label
		var noteText = item.Note
		fmt.Println("labelText", labelText)
		fmt.Println("noteText", noteText)
		// taskType := "USER_TIMEOFF"
		assignedShifts = append(assignedShifts, &model.TimeOff{
			ID:           item.ID,
			StartTime:    item.StartTime,
			EndTime:      item.EndTime,
			Label:        labelText,
			Note:         noteText,
			Color:        item.Color,
			Is24Hours:    item.Is24Hours,
			UserID:       item.UserID,
			ChannelID:    item.ChannelID,
			ShiftGroupID: item.ShiftGroupID,
		})

		// newShift = &model.AssignedShift{
		// 	ID:        item.ID,
		// 	Type:      item.Type,
		// 	Color:     item.Color,
		// 	Note:      item.Note,
		// 	Label:     item.Label,
		// 	Is24Hours: item.Is24Hours,
		// 	StartTime: item.StartTime,
		// 	EndTime:   item.EndTime,
		// }

	}

	return assignedShifts, nil
}

// ------------------------------------------end i add

func GetAssignedShiftsByTime(channelId *string, shiftGroupId *string, userId *string, endTime *time.Time, startTime *time.Time) ([]*model.AssignedShift, error) {

	jsonMapInstance := map[string]string{
		"query": `
			{
			getAssignedShiftsByTime(
			  channelId: "` + *channelId + `"
			  shiftGroupId: "` + *shiftGroupId + `"
			  userId: "` + *userId + `"
			  startTime: "` + fmt.Sprintf("%s", startTime.Format(time.RFC3339)) + `"
			  endTime: "` + fmt.Sprintf("%s", endTime.Format(time.RFC3339)) + `"
			) 
			{
			  id
			  break
			  label
			  color
			  startTime
			  endTime
			  is24Hours
			  userId
			  channelId
			  shiftGroupId
			  type
			  isOpen
			  note
			  isShared
			  ShiftActivities {
				id
				name
				code
				color
				startTime
				endTime
				userId
				isPaid
			  }
			}
		  }
		`,
	}

	responseData, err := httpRequest("POST", os.Getenv("ASSIGNED_SHIFT_API"), jsonMapInstance)

	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		fmt.Printf("There was an error executing the request%v", err)
		return nil, err
	}

	var responseObject model.GetAssignedShiftsByTimeResponse
	json.Unmarshal(responseData, &responseObject)

	var assignedShifts []*model.AssignedShift

	for _, assignedShift := range responseObject.Data.GetAssignedShiftsByTime {
		assignedShifts = append(assignedShifts, &model.AssignedShift{
			ID:           assignedShift.ID,
			Break:        assignedShift.Break,
			Label:        assignedShift.Label,
			Note:         assignedShift.Note,
			StartTime:    assignedShift.StartTime,
			Color:        assignedShift.Color,
			UserID:       assignedShift.UserID,
			ChannelID:    assignedShift.ChannelID,
			ShiftGroupID: assignedShift.ShiftGroupID,
			Is24Hours:    assignedShift.Is24Hours,
			EndTime:      assignedShift.EndTime,
			IsShared:     assignedShift.IsShared,
			Type:         assignedShift.Type,
			IsOpen:       assignedShift.IsOpen,
		})

		for _, shiftActivity := range assignedShift.ShiftActivities {
			assignedShifts[len(assignedShifts)-1].ShiftActivities = append(assignedShifts[len(assignedShifts)-1].ShiftActivities, &model.AssignedShiftActivities{
				ID:        shiftActivity.ID,
				Name:      shiftActivity.Name,
				Code:      shiftActivity.Code,
				Color:     shiftActivity.Color,
				StartTime: shiftActivity.StartTime,
				EndTime:   shiftActivity.EndTime,
				UserID:    shiftActivity.UserID,
				IsPaid:    shiftActivity.IsPaid,
			})
		}
	}

	return assignedShifts, nil
}

func DiffHours(date1, date2 time.Time) int {
	diff := (date1.Sub(date2)) / time.Second
	diff /= (60 * 60)
	return int(diff)
}

func httpRequest(method string, url string, query map[string]string) ([]byte, error) {
	jsonResult, err := json.Marshal(query)
	newRequest, err := http.NewRequest(method, url, bytes.NewBuffer(jsonResult))
	if err != nil {

		// sentry.CaptureException(err)
		// defer sentry.Flush(2 * time.Second)

		fmt.Printf("There was an error creating the new request %v", err)
		return nil, err
	}
	newRequest.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: time.Second * 5}
	response, err := client.Do(newRequest)

	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)
		fmt.Printf("There was an error executing the request%v", err)
	}

	// check if the response is nil
	if response == nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		fmt.Println("check the " + url + " url ")
		return nil, fmt.Errorf("Something went wrong with the request: " + url)
	}

	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		fmt.Printf("Data Read Error%v", err)
	}

	return responseData, nil
}

func GetChannelBySlug(slug string) (*model.ChannelResPonse, error) {
	jsonMapInstance := map[string]string{
		"query": `
		{
			channel(slug: "` + slug + `") {
			  id
			  slug
			  name
			}
		  }	  
		`,
	}

	responseData, err := httpRequest("POST", os.Getenv("CHANNEL_API"), jsonMapInstance)
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		fmt.Printf("There was an error executing the request%v", err)
		return nil, err
	}

	var channelResponse model.ChannelResPonse
	err = json.Unmarshal(responseData, &channelResponse)
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		fmt.Printf("There was an error unmarshalling the response data%v", err)
	}

	return &channelResponse, nil
}
func GetChannelById(id string) (*model.ChannelResPonse, error) {
	jsonMapInstance := map[string]string{
		"query": `
		{
			channel(id: "` + id + `") {
			  id
			  slug
			  name
			}
		  }	  
		`,
	}

	responseData, err := httpRequest("POST", os.Getenv("CHANNEL_API"), jsonMapInstance)
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		fmt.Printf("There was an error executing the request%v", err)
		return nil, err
	}

	var channelResponse model.ChannelResPonse
	err = json.Unmarshal(responseData, &channelResponse)
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		fmt.Printf("There was an error unmarshalling the response data%v", err)
	}

	return &channelResponse, nil
}

func GetRequests(channelId *string, userId *string, authUserId *string) (*model.GetRequestsResponse, error) {
	jsonMapInstance := map[string]string{
		"query": `
		{
			getRequestsByUser(
			  channelId: "` + *channelId + `"
			  userId: "` + *userId + `"
			  authUserId: "` + *authUserId + `"
			) {
			  edges {
			 {
				  id
				  requestId
				  status
				  startTime
				  endTime
				  requestNote
				  reason
				  responseNote
				  channelId
				  isAllDay
				  type
				  shiftOfferedTo {
					id
					firstName
					lastName
					email
				  }
				  shiftToSwap {
					id
					color
					label
					note
				  }
				  toSwapWith {
					id
					color
					label
					note
				  }
				}
			  }
			}
		  }
		`,
	}

	responseData, err := httpRequest("POST", os.Getenv("REQUEST_API"), jsonMapInstance)
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		fmt.Printf("There was an error executing the request%v", err)
		return nil, err
	}

	var getRequestsResponse model.GetRequestsResponse
	err = json.Unmarshal(responseData, &getRequestsResponse)

	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		fmt.Printf("There was an error unmarshalling the response data%v", err)
	}

	return &getRequestsResponse, nil

}

func CheckPermission(object string, permission string, userId string) (bool, error) {
	jsonMapInstance := map[string]string{
		"query": `
			query {
				CheckPermission(
					NameSpace: ` + os.Getenv("NAMESPACE") + `
					object: "` + object + `"
					permission: ` + permission + `
					userId: "` + userId + `"
				)
			}
		`,
	}

	responseData, err := httpRequest("POST", os.Getenv("PERMISSION_API"), jsonMapInstance)
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		fmt.Printf("There was an error executing the request%v", err)
		return false, err
	}

	var permissionResponse model.PermissionResponse
	err = json.Unmarshal(responseData, &permissionResponse)
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		fmt.Printf("Unmarshal Error%v", err)
	}

	return permissionResponse.Data.CheckPermission, nil
}

func GetShiftsByTaskHandleError(message *string, status *string) (*model.GetShiftsByTaskResponse, error) {
	sentry.CaptureException(fmt.Errorf(*message))
	defer sentry.Flush(2 * time.Second)

	return &model.GetShiftsByTaskResponse{
		Message: message,
		Status:  status,
		Result:  nil,
	}, nil
}

func ShiftGroupMemberAddRHandleError(errorMessage *string, code model.ShiftErrorCode) (*model.ShiftGroupMemberAddResponse, error) {

	var shiftError []*model.ShiftError
	fieldError := "Add Shift Group Member"

	shiftError = append(shiftError, &model.ShiftError{
		Code:    code,
		Field:   &fieldError,
		Message: errorMessage,
	})

	return &model.ShiftGroupMemberAddResponse{
		Errors: shiftError,
		User:   nil,
	}, nil
}

func ShiftGroupMemberRemoveRHandleError(errorMessage *string, code model.ShiftErrorCode) (*model.ShiftGroupMemberRemoveResponse, error) {

	var shiftError []*model.ShiftError
	fieldError := "Remove shift GroupMember"

	shiftError = append(shiftError, &model.ShiftError{
		Code:    code,
		Field:   &fieldError,
		Message: errorMessage,
	})

	return &model.ShiftGroupMemberRemoveResponse{
		Errors: shiftError,
		User:   nil,
	}, nil
}
