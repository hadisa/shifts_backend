package model

type CreateRequestResponse struct {
	Data struct {
		CreateRequest struct {
			ID string `json:"id"`
		} `json:"createRequest"`
	} `json:"data"`
}

type PermissionResponse struct {
	Data struct {
		CheckPermission bool `json:"CheckPermission"`
	} `json:"data"`
}

type UserResponse struct {
	Data struct {
		User struct {
			ID        string  `json:"id"`
			Email     string  `json:"email"`
			FirstName string  `json:"firstName"`
			LastName  string  `json:"lastName"`
			IsStaff   bool    `json:"isStaff"`
			IsActive  bool    `json:"isActive"`
			Avatar    *string `json:"avatar"`
		}
	} `json:"data"`
}
