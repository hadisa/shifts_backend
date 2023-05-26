package graph

import (
	"account_user/graph/generated"
	"account_user/graph/model"
	"account_user/util"
	"context"
	"fmt"
	"time"
)

// AccountRegister is the resolver for the accountRegister field.
func (r *mutationResolver) AccountRegister(ctx context.Context, input model.AccountRegisterInput) (*model.AccountRegister, error) {
	// channel := "default-channel"
	avatar := ""
	accountUser := model.User{
		ID:           input.IdentityID,
		Email:        input.Email,
		Phone:        input.Phone,
		Whatsapp:     input.Whatsapp,
		FirstName:    *input.FirstName,
		LastName:     *input.LastName,
		IsStaff:      true,
		IsActive:     true,
		Note:         input.Channel,
		LanguageCode: *input.LanguageCode,
		DateJoined:   time.Now(),
		UpdatedAt:    time.Now(),
		Avatar:       &avatar,
		
	}

	err := r.DB.Create(&accountUser).Error
	if err != nil {
		util.SentryLogError(err)

		field := "Register Account"
		errors := []*model.AccountError{}

		regErr := err.Error()

		errors = append(errors, &model.AccountError{
			Field:       &field,
			Message:     &regErr,
			Code:        model.AccountErrorCodeInvalid,
			AddressType: nil,
		})

		return &model.AccountRegister{
			User:   nil,
			Errors: errors,
		}, nil
	}

	return &model.AccountRegister{
		User:   &accountUser,
		Errors: nil,
	}, nil
}

// AccountUpdate is the resolver for the accountUpdate field.
func (r *mutationResolver) AccountUpdate(ctx context.Context, id string, input model.AccountInput) (*model.AccountUpdate, error) {
	field := "Update Account"

	var user model.User
	err := r.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		util.SentryLogError(err)

		regErr := err.Error()

		errors := []*model.AccountError{}
		errors = append(errors, &model.AccountError{
			Field:       &field,
			Message:     &regErr,
			Code:        model.AccountErrorCodeNotFound,
			AddressType: nil,
		})

		return &model.AccountUpdate{
			User:   nil,
			Errors: errors,
		}, nil
	}

	var phone string
	var whatsapp string
	if input.Phone != nil || *input.Phone != "" {
		phone = *input.Phone
	} else {
		phone = *user.Phone
	}

	if input.Whatsapp != nil || *input.Whatsapp != "" {
		whatsapp = *input.Whatsapp
	} else {
		whatsapp = *user.Whatsapp
	}

	user.FirstName = *input.FirstName
	user.LastName = *input.LastName
	user.Phone = &phone
	user.Whatsapp = &whatsapp
	user.LanguageCode = *input.LanguageCode
	user.UpdatedAt = time.Now()

	err = r.DB.Save(&user).Error
	if err != nil {
		util.SentryLogError(err)

		regErr := err.Error()

		errors := []*model.AccountError{}
		errors = append(errors, &model.AccountError{
			Field:       &field,
			Message:     &regErr,
			Code:        model.AccountErrorCodeInvalid,
			AddressType: nil,
		})

		return &model.AccountUpdate{
			User:   nil,
			Errors: errors,
		}, nil
	}

	return &model.AccountUpdate{
		User: &user,
	}, nil
}

// AccountRequestDeletion is the resolver for the accountRequestDeletion field.
func (r *mutationResolver) AccountRequestDeletion(ctx context.Context, channel *string, redirectURL string) (*model.AccountRequestDeletion, error) {
	// context := GetContext(ctx)

	// var user model.User
	return &model.AccountRequestDeletion{}, nil
}

// AccountDelete is the resolver for the accountDelete field.
func (r *mutationResolver) AccountDelete(ctx context.Context, id string) (*model.AccountDelete, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	// delete user account by id (also delete form Kratos)

	var user model.User
	var err error

	if id != "" {
		err = r.DB.Where("id = ?", id).Delete(&user).Error
	}

	if err != nil {
		util.SentryLogError(err)

		field := "Delete Account"
		errors := []*model.AccountError{}

		regErr := err.Error()

		errors = append(errors, &model.AccountError{
			Field:       &field,
			Message:     &regErr,
			Code:        model.AccountErrorCodeInvalid,
			AddressType: nil,
		})

		return &model.AccountDelete{
			User:   nil,
			Errors: errors,
		}, nil
	}

	return &model.AccountDelete{
		User: &user,
	}, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id *string, email *string) (*model.User, error) {
	var user model.User
	var err error
	// check if id or email in not nill
	if id != nil && email != nil {
		err = r.DB.Where("id = ? AND email = ?", id, email).First(&user).Error
	}
	if id != nil {
		err = r.DB.Where("id = ?", id).First(&user).Error
	}
	if email != nil {
		err = r.DB.Where("email = ?", email).First(&user).Error
	}

	if err != nil {
		util.SentryLogError(err)

		return nil, err
	}

	return &user, nil
}

// GetUserByIsStaff is the resolver for the getUserByIsStaff field.
func (r *queryResolver) GetUserByIsStaff(ctx context.Context, isStaff bool) ([]*model.User, error) {
	var users []*model.User
	err := r.DB.Where("is_staff = ?", isStaff).Find(&users).Error
	if err != nil {
		util.SentryLogError(err)

		return nil, err
	}
	return users, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, first *int, last *int) ([]*model.User, error) {
	var users []*model.User
	var err error

	if first != nil {
		err = r.DB.Limit(*first).Find(&users).Error
		if err != nil {

			util.SentryLogError(err)

			return nil, err
		}
	} else if last != nil {
		err = r.DB.Limit(*last).Order("id desc").Find(&users).Error
		if err != nil {
			util.SentryLogError(err)

			return nil, err
		}
	} else {
		err = r.DB.Find(&users).Error
		if err != nil {
			util.SentryLogError(err)

			return nil, err
		}
	}

	return users, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
