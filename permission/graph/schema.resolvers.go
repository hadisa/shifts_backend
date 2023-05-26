package graph

import (
	"context"
	"fmt"
	"permissions_aws/auth"
	"permissions_aws/graph/generated"
	"permissions_aws/graph/model"
	"permissions_aws/util"
	"time"

	"github.com/google/uuid"
)

// GrantPermission is the resolver for the grantPermission field.
func (r *mutationResolver) GrantPermission(ctx context.Context, input model.GrantedPermissionInput) (*model.GrantedPermissionResponse, error) {
	var err error

	// TODO: create keto tuple
	/*
		explanation:
		- create keto relation tuple

		data:
		- namespace (application name)
		- subject (user id)
		- permission (relation or action or role)
		- object (functionality or resource or object or entity name)
	*/

	// check if the permission exists
	var permission model.GrantedPermission
	err = r.DB.First(&permission, "user_id =? AND name_space =? AND permission =? AND object =?", input.UserID, input.NameSpace.String(), input.Permission.String(), input.Object).Error
	if err != nil && err.Error() != "record not found" {
		util.SentryLogError(err)

		return nil, err
	}

	if permission.ID != "" {
		return nil, fmt.Errorf("Permission already exists: %s", permission.ID)
	}

	grantedPermission := &model.GrantedPermission{
		ID:         uuid.New().String(),
		NameSpace:  input.NameSpace.String(),
		UserID:     input.UserID,
		Permission: input.Permission.String(),
		Object:     input.Object,
		GrantedAt:  time.Now().UTC(),
	}

	// create keto relation tuple
	result, err := auth.GrantPermission(input.NameSpace.String(), input.Object, input.Permission.String(), input.UserID)
	if err != nil {
		util.SentryLogError(err)
		return nil, err
	}

	if !result {
		return nil, fmt.Errorf("Failed to grant permission (Keto)")
	}

	err = r.DB.Create(grantedPermission).Error
	if err != nil {
		util.SentryLogError(err)
		return nil, err
	}

	user, err := util.GetUser(input.UserID)

	if err != nil {
		util.SentryLogError(err)

		return nil, err
	}

	if user.ID == nil {
		return nil, fmt.Errorf("User not found")
	}

	return &model.GrantedPermissionResponse{
		PermissionID: &grantedPermission.ID,
		NameSpace:    &grantedPermission.NameSpace,
		Permission:   &grantedPermission.Permission,
		Object:       &grantedPermission.Object,
		GrantedAt:    &grantedPermission.GrantedAt,
		User:         user,
	}, nil
}

// RevokePermission is the resolver for the revokePermission field.
func (r *mutationResolver) RevokePermission(ctx context.Context, id string) (*string, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	// get granted permission with its permission
	var grantedPermission model.GrantedPermission
	err := r.DB.First(&grantedPermission, "id =?", id).Error
	if err != nil {
		util.SentryLogError(err)
		return nil, err
	}

	// TODO: delete the tuple from keto
	// delete keto relation tuple
	result, err := auth.RevokePermission(grantedPermission.NameSpace, grantedPermission.Object, grantedPermission.Permission, grantedPermission.UserID)
	if err != nil {
		util.SentryLogError(err)
		return nil, err
	}

	if !result {
		return nil, fmt.Errorf("Failed to revoke permission (Keto)")
	}

	// delete granted permission
	err = r.DB.Delete(&grantedPermission).Error
	if err != nil {
		util.SentryLogError(err)
		return nil, err
	}

	message := "Permission revoked successfully"

	return &message, nil
}

// GetGrantedPermissions is the resolver for the getGrantedPermissions field.
func (r *queryResolver) GetGrantedPermissions(ctx context.Context, userID string) (*model.GetGrantedPermissionsResponse, error) {
	if userID == "" {
		return nil, fmt.Errorf("User ID is required")
	}

	var grantedPermissions []*model.GrantedPermission
	// get all granted permissions with their permission group and permission
	err := r.DB.Find(&grantedPermissions, "user_id =?", userID).Error
	if err != nil {
		util.SentryLogError(err)
		return nil, err
	}

	user, err := util.GetUser(userID)
	if err != nil {
		util.SentryLogError(err)

		return nil, err
	}

	if user.ID == nil {
		return nil, fmt.Errorf("User not found")
	}

	return &model.GetGrantedPermissionsResponse{
		FirstName:   *user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		Permissions: grantedPermissions,
	}, nil
}

// GetAllGrantedPermissions is the resolver for the getAllGrantedPermissions field.
func (r *queryResolver) GetAllGrantedPermissions(ctx context.Context) ([]*model.GrantedPermissionResponse, error) {
	// TODO: get all the users with its granted permissions
	var grantedPermissions []*model.GrantedPermission
	// get all granted permissions with their permission group and permission
	err := r.DB.Find(&grantedPermissions).Error
	if err != nil {
		util.SentryLogError(err)
		return nil, err
	}

	var grantedPermissionsResponse []*model.GrantedPermissionResponse

	for _, grantedPermission := range grantedPermissions {
		user, err := util.GetUser(grantedPermission.UserID)

		if err != nil {
			util.SentryLogError(err)

			return nil, err
		}

		if user.ID == nil {
			return nil, fmt.Errorf("User not found")
		}

		grantedPermissionsResponse = append(grantedPermissionsResponse, &model.GrantedPermissionResponse{
			PermissionID: &grantedPermission.ID,
			NameSpace:    &grantedPermission.NameSpace,
			Permission:   &grantedPermission.Permission,
			Object:       &grantedPermission.Object,
			GrantedAt:    &grantedPermission.GrantedAt,
			User:         user,
		})
	}

	return grantedPermissionsResponse, nil
}

// CheckPermission is the resolver for the CheckPermission field.
func (r *queryResolver) CheckPermission(ctx context.Context, nameSpace model.NameSpaceEnum, userID string, permission model.PermissionEnum, object string) (bool, error) {
	if userID == "" {
		return false, fmt.Errorf("User ID is required")
	}

	if object == "" {
		return false, fmt.Errorf("Object is required")
	}

	// check permission from keto
	result, err := auth.CheckPermission(nameSpace.String(), object, permission.String(), userID)
	if err != nil {
		util.SentryLogError(err)
		return false, err
	}

	return result, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
