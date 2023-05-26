package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"permissions_aws/graph/model"
	"time"

	"github.com/getsentry/sentry-go"
)

func httpRequest(method string, url string, query map[string]string) ([]byte, error) {
	jsonResult, err := json.Marshal(query)
	newRequest, err := http.NewRequest(method, url, bytes.NewBuffer(jsonResult))
	if err != nil {

		SentryLogError(err)

		fmt.Printf("There was an error creating the new request %v", err)
		return nil, err
	}
	newRequest.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: time.Second * 5}
	response, err := client.Do(newRequest)

	if err != nil {
		SentryLogError(err)
		fmt.Printf("There was an error executing the request%v", err)
	}

	// check if the response is nil
	if response == nil {
		SentryLogError(err)

		fmt.Println("check the " + url + " url ")
		return nil, fmt.Errorf("Something went wrong with the request: " + url)
	}

	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		SentryLogError(err)

		fmt.Printf("Data Read Error%v", err)
	}

	return responseData, nil
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
			}
		  }
		`,
	}

	responseData, err := httpRequest("POST", os.Getenv("USER_ACCOUNT_API"), jsonMapInstance)
	if err != nil {
		SentryLogError(err)

		fmt.Printf("There was an error executing the request%v", err)
		return nil, err
	}

	// convert into User struct
	var userData model.UserResponse
	err = json.Unmarshal(responseData, &userData)

	if err != nil {
		SentryLogError(err)

		fmt.Printf("There was an error unmarshaling the JSON instance %v", err)
	}

	return &model.User{
		ID:        userData.Data.User.ID,
		Email:     userData.Data.User.Email,
		FirstName: userData.Data.User.FirstName,
		LastName:  userData.Data.User.LastName,
	}, nil
}

func SentryLogError(err error) {
	sentry.CaptureException(err)
	defer sentry.Flush(2 * time.Second)
}
