# User Account

## Vault

> Shell Command to Create AppRole Role and Secret ID for `User Account` Service

```
vault write auth/approle/role/user_account token_policies="general" token_ttl=10h token_max_ttl=18h

vault read auth/approle/role/user_account/role-id

vault write -f auth/approle/role/user_account/secret-id

```
<!-- role_id a0c78146-d146-5788-a66d-a99ee28709c2 -->
<!-- secret_id          e841cace-7d19-03ac-9a62-b1bb9c158dea
secret_id_accessor d7450fa7-8e2f-d7c4-a479-e033eb9e4e22
secret_id_num_uses 0                                   
secret_id_ttl      0 -->

### Role Creation statements:

```

CREATE ROLE "{{name}}" WITH LOGIN PASSWORD '{{password}}' VALID UNTIL '{{expiration}}' INHERIT;
GRANT USAGE ON SCHEMA public TO "{{name}}";
GRANT SELECT, UPDATE ,INSERT, DELETE ON ALL TABLES IN SCHEMA public TO "{{name}}";
ALTER USER "{{name}}" WITH SUPERUSER;
```

# Graphql

## Queries

Get user lists

```graphql
query GetUsersQuery {
  users {
    id
    firstName
    lastName
    email
    isStaff
    isActive
    phone
    avatar
    note
  }
}
```

### Response

```json
{
  "data": {
    "users": [
      {
        "id": "58500165-593c-471d-b92b-ac1ebd7b1ea3",
        "firstName": "John",
        "lastName": "Doe",
        "email": "john.doe@gmail.com",
        "isStaff": false,
        "isActive": true,
        "phone": "+48 123 456 789",
        "avatar": "no avatar",
        "note": "no note"
      },
      {
        "id": "58500165-593c-471d-b92b-ac1ebd7b1ea3",
        "firstName": "John",
        "lastName": "Doe",
        "email": "john.doe@yahoo.com",
        "isStaff": true,
        "isActive": true,
        "phone": "+48 123 456 789",
        "avatar": "no avatar",
        "note": "no note"
      }
    ]
  }
}
```

Get a specific `User` by id or email

```graphql
query GetUserQuery($id: ID, $email: String) {
  user(id: $id, email: $email) {
    id
    firstName
    lastName
    email
    isStaff
    isActive
    phone
    avatar
    note
  }
}
```

Variables:

```json
{
  "id": "58500165-593c-471d-b92b-ac1ebd7b1ea3",
  "email": ""
}
```

### Response

```json
{
  "data": {
    "user": {
      "id": "58500165-593c-471d-b92b-ac1ebd7b1ea3",
      "firstName": "Yaseen",
      "lastName": "Akbari",
      "email": "akbari01.dev@gmail.com",
      "isStaff": true,
      "isActive": true,
      "phone": "+48 123 456 789",
      "avatar": "https://www.gravatar.com/avatar/205e460b479e2e5b48aec07710c08d50?f=y",
      "note": "no note"
    }
  }
}
```

## Mutations

> AccountRegister

Register a new user.

```graphql
mutation AccountRegister($input: AccountRegisterInput!) {
  accountRegister(input: $input) {
    requiresConfirmation
    errors {
      field
      message
      code
      addressType
    }
    user {
      id
      email
      firstName
      lastName
      isStaff
      isActive
      phone
      whatsapp
      note
      avatar
      languageCode
      lastLogin
      dateJoined
      updatedAt
    }
  }
}
```

> Input Arguments:

input - required. Input object with user data:

- identityId - required. The unique identities.id from the Kratos user.
- firstName - optional. Given name.
- lastName - optional. Family name.
- phone - optional. Phone number.
- whatsapp - optional. WhatsApp number.
- languageCode - required. User language code.
- email - required. The email address of the user.
- password - required. Password.
- redirectUrl - optional. Base of frontend URL that will be needed to create confirmation URL.
- channel - optional. Slug of a channel which will be used to notify users. Optional when only one channel exists.

> Return Fields:

- requiresConfirmation - A boolean that informs whether users need to confirm their email address.registration process.
- errors - An array of AccountError objects that represent errors that occurred during the registration process.

user - An object that represents the registered user. It has the following fields:

- id - The ID of the user.
- email - The email address of the user.
- firstName - Given name.
- lastName - Family name.
- isStaff - A boolean that indicates whether the user is a staff member.
- isActive - A boolean that indicates whether the user is active.
- phone - Phone number.
- whatsapp - WhatsApp number.
- note - A note about the customer. Requires one of the following permissions: MANAGE_USERS, MANAGE_STAFF.
- avatar - Avatar.
- languageCode - User language code.
- lastLogin - Time of last login.
- dateJoined - Time of user registration.
- updatedAt - Time of the last update of the user.

> Variables Example

```json
{
  "input": {
    "identityId": "0a90b5dc-6f40-483b-b0ba-bd42b1fd2bb1",
    "firstName": "John",
    "lastName": "Doe",
    "phone": "+48 123 456 789",
    "whatsapp": "+48 123 456 789",
    "languageCode": "EN",
    "email": "john.doe@example.com",
    "redirectUrl": "https://example.com/confirm",
    "avatar": "no avatar"
  }
}
```

### Response

```json
{
  "data": {
    "accountRegister": {
      "requiresConfirmation": false,
      "errors": [],
      "user": {
        "id": "0a90b5dc-6f40-483b-b0ba-bd42b1fd2bb1",
        "email": "john.doe@example.com",
        "firstName": "John",
        "lastName": "Doe",
        "isStaff": false,
        "isActive": true,
        "phone": "+48 123 456 789",
        "whatsapp": "+48 123 456 789",
        "note": "no note",
        "avatar": "no avatar",
        "languageCode": "EN",
        "lastLogin": "2021-05-20T12:00:00.000Z",
        "dateJoined": "2021-05-20T12:00:00.000Z",
        "updatedAt": "2021-05-20T12:00:00.000Z"
      }
    }
  }
}
```
