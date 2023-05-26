package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"request_swaps/graph/model"
	"time"

	"github.com/getsentry/sentry-go"
)

func httpRequest(method string, url string, daprAppId string, query map[string]string) ([]byte, error) {
	jsonResult, err := json.Marshal(query)
	newRequest, err := http.NewRequest(method, url, bytes.NewBuffer(jsonResult))
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		fmt.Printf("There was an error creating the new request %v", err)
		return nil, err
	}

	newRequest.Header.Set("Content-Type", "application/json")
	// DapR header
	newRequest.Header.Add("dapr-app-id", daprAppId)

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

func CreateRequest(channelId *string, userId *string) (string, error) {
	jsonMapInstance := map[string]string{
		"query": `
		mutation {
			createRequest(input: {
			  channelId: "` + *channelId + `"
			  userId: "` + *userId + `"
			  recipientId: "-1"
			  type: "requestOffer"
			}) {
			  id
			}
		  }
		`,
	}

	responseData, err := httpRequest("POST", os.Getenv("REQUEST_API"), os.Getenv("DAPR_REQUEST_APP_ID"), jsonMapInstance)
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		fmt.Printf("There was an error executing the request%v", err)
		return "", err
	}

	// convert responseData to struct
	var responseObject model.CreateRequestResponse
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		fmt.Printf("JSON Unmarshal error: %v", err)
	}

	return responseObject.Data.CreateRequest.ID, err
}

func DeleteRequest(id *string) (bool, error) {
	jsonMapInstance := map[string]string{
		"query": `
		mutation {
			deleteRequest(id: "` + *id + `") 
		  }
		`,
	}

	responseData, err := httpRequest("POST", os.Getenv("REQUEST_API"), os.Getenv("DAPR_REQUEST_APP_ID"), jsonMapInstance)
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		fmt.Printf("There was an error executing the request%v \n", err)
		return false, err
	}

	if string(responseData) != "" {
		return true, err
	} else {
		return false, err
	}

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

	responseData, err := httpRequest("POST", os.Getenv("PERMISSION_API"), os.Getenv("DAPR_PERMISSION_APP_ID"), jsonMapInstance)
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		fmt.Printf("There was an error executing the request%v \n", err)
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
			  avatar
			}
		  }
		`,
	}

	responseData, err := httpRequest("POST", os.Getenv("USER_ACCOUNT_API"), os.Getenv("DAPR_USER_APP_ID"), jsonMapInstance)
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(2 * time.Second)

		fmt.Printf("There was an error executing the request%v", err)
		fmt.Printf("There was an error executing the request%v \n", err)
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
		ID:        &userData.Data.User.ID,
		Email:     &userData.Data.User.Email,
		FirstName: &userData.Data.User.FirstName,
		LastName:  &userData.Data.User.LastName,
		IsStaff:   &userData.Data.User.IsStaff,
		// Avatar:       userData.Data.User.Avatar,
	}, nil
}
