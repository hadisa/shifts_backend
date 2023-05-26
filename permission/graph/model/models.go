package model

type PermissionGroupPermission struct {
	PermissionGroupID string `json:"permissionGroupId"`
	PermissionID      string `json:"permissionId"`
}

type Object struct {
	ID           string          `json:"id" gorm:"primaryKey; type:uuid; default:uuid_generate_v4()"`
	NameSpace    string          `json:"nameSpace"`
	ObjectDetail []*ObjectDetail `json:"objectDetail" gorm:"foreignKey:ObjectID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type ObjectDetail struct {
	ID       string `json:"id" gorm:"primaryKey; type:uuid; default:uuid_generate_v4()"`
	ObjectID string `json:"objectId" gorm:"type:uuid; not null"`
	Name     string `json:"name" gorm:"type:varchar(64); not null"`
}

type UserResponse struct {
	Data struct {
		User struct {
			ID        *string `json:"id"`
			Email     *string `json:"email"`
			FirstName *string `json:"firstName"`
			LastName  *string `json:"lastName"`
		}
	} `json:"data"`
}
