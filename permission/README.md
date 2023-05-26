# Permission (User Permissions)

> implies (User Permissions) with Ory Keto

# Vault

> Shell Command to Create AppRole Role and Secret ID for `User Permission` Service

```
vault write auth/approle/role/user_permission token_policies="general" token_ttl=10h token_max_ttl=18h
vault read auth/approle/role/user_permission/role-id
vault write -f auth/approle/role/user_permission/secret-id
```

## Role Creation statements:

```
CREATE ROLE "{{name}}" WITH LOGIN PASSWORD '{{password}}' VALID UNTIL '{{expiration}}' INHERIT;
GRANT USAGE ON SCHEMA public TO "{{name}}";
GRANT SELECT, UPDATE ,INSERT, DELETE ON ALL TABLES IN SCHEMA public TO "{{name}}";
ALTER USER "{{name}}" WITH SUPERUSER;
```

## Name Spaces

- shifts
- booking

## Shifts Objects

- setting
- day_note
- assigned_shift
- open_shift
- request
- request_offer
- request_swap
- request_time_off
- user_time_off
- shared_schedule
- shift_group_member
- shift_group

## Booking Objects

- booking_appointment
- business_booking
- booking_custom_question
- booking_service
- booking_staff_member

## Permissions

> READ
> | READ_ALL
> | WRITE
> | WRITE_ALL
> | MANAGE
> | MANAGE_ALL

# Graphql

## GrantedPermission

This type defines a granted permission object which has the following fields:

- id (ID!) - The unique ID of the granted permission object.
- nameSpace (String!) - The namespace of the granted permission.
- userId (String!) - The ID of the user who was granted the permission.
- permission (String!) - The type of permission granted.
- object (String!) - The object to which the permission was granted.
- grantedAt (Time!) - The timestamp at which the permission was granted.

## GrantedPermissionInput

This input type is used when creating a new granted permission object. It has the following fields:

- nameSpace (NameSpaceEnum!) - The namespace of the permission being granted.
- userId (ID!) - The ID of the user who will be granted the permission.
- permission (PermissionEnum!) - The type of permission to be granted.
- object (ID!) - The ID of the object to which the permission will be granted.

## GrantedPermissionResponse

This type is used to represent the response returned when a new granted permission is created. It has the following fields:

- permissionId (ID) - The ID of the newly created granted permission object.
- nameSpace (String) - The namespace of the granted permission.
- permission (String) - The type of permission granted.
- object (String) - The object to which the permission was granted.
- grantedAt (Time) - The timestamp at which the permission was granted.
- user (User!) - The user who was granted the permission.

## GetGrantedPermissionsResponse

This type is used to represent the response returned when retrieving a user's granted permissions. It has the following fields:

- firstName (String!) - The first name of the user.
- lastName (String) - The last name of the user.
- email (String) - The email address of the user.
- permissions ([GrantedPermission!]!) - An array of the user's granted permissions.

## User

This type defines a user object with the following fields:

- id (ID) - The unique ID of the user.
- email (String) - The email address of the user.
- firstName (String) - The first name of the user.
- lastName (String) - The last name of the user.

## NameSpaceEnum

This enum defines the available namespaces for the granted permission object. The available options are:

- shifts
- booking

## PermissionEnum

This enum defines the available permission types that can be granted. The available options are:

- READ
- READ_ALL
- WRITE
- WRITE_ALL
- MANAGE
- MANAGE_ALL

## Mutation

### grantPermission

Creates a new granted permission object and returns the newly created object.

```graphql
mutation GrantPermissionMutation(
  $nameSpace: NameSpaceEnum!
  $userId: ID!
  $permission: PermissionEnum!
  $object: String!
) {
  grantPermission(
    input: {
      nameSpace: $nameSpace
      userId: $userId
      permission: $permission
      object: $object
    }
  ) {
    permissionId
    nameSpace
    permission
    object
    grantedAt
    user {
      id
      email
      firstName
      lastName
    }
  }
}
```

Variables:

```json
{
  "nameSpace": "shifts",
  "userId": "58500165-593c-471d-b92b-ac1ebd7b1ea3",
  "permission": "WRITE_ALL",
  "object": "shared_schedule"
}
```

### Response

```json
{
  "data": {
    "grantPermission": {
      "permissionId": "b5f9b9a0-5b1f-4b1f-9f5a-1b1b1b1b1b1b",
      "nameSpace": "shifts",
      "permission": "WRITE_ALL",
      "object": "shared_schedule",
      "grantedAt": "2020-10-01T15:00:00.000Z",
      "user": {
        "id": "58500165-593c-471d-b92b-ac1ebd7b1ea3",
        "email": "john.doe@example.com",
        "firstName": "John",
        "lastName": "Doe"
      }
    }
  }
}
```

### revokePermission

Deletes a granted permission object and returns a success message.

```graphql
mutation RevokePermissionMutation($id: ID!) {
  revokePermission(id: $id)
}
```

Variables:

```json
{
  "id": "user goes here"
}
```

### Response

```json
{
  "data": {
    "revokePermission": "Permission successfully revoked."
  }
}
```

## Query

### getGrantedPermissions

Retrieves all the granted permissions for a given user and returns a GetGrantedPermissionsResponse object.

```graphql
query GetGrantedPermissionsQuery($userId: ID!) {
  getGrantedPermissions(userId: $userId) {
    firstName
    lastName
    email
    permissions {
      id
      nameSpace
      object
      grantedAt
    }
  }
}
```

Variables:

```json
{
  "userId": "58500165-593c-471d-b92b-ac1ebd7b1ea3"
}
```

### Response

```json
{
  "data": {
    "getGrantedPermissions": {
      "firstName": "John",
      "lastName": "Doe",
      "email": "john.doe@example.com",
      "permissions": [
        {
          "id": "b5f9b9a0-5b1f-4b1f-9f5a-1b1b1b1b1b1b",
          "nameSpace": "shifts",
          "object": "shared_schedule",
          "grantedAt": "2020-10-01T15:00:00.000Z"
        }
      ]
    }
  }
}
```

### getAllGrantedPermissions

Retrieves all the granted permissions for all users and returns an array of GrantedPermissionResponse objects.

```graphql
query GetAllGrantedPermissionsQuery {
  getAllGrantedPermissions {
    permissionId
    nameSpace
    permission
    object
    grantedAt
    user {
      id
      firstName
      lastName
      email
    }
  }
}
```

### Response

```json
{
  "data": {
    "getAllGrantedPermissions": [
      {
        "permissionId": "b5f9b9a0-5b1f-4b1f-9f5a-1b1b1b1b1b1b",
        "nameSpace": "shifts",
        "permission": "WRITE_ALL",
        "object": "shared_schedule",
        "grantedAt": "2020-10-01T15:00:00.000Z",
        "user": {
          "id": "58500165-593c-471d-b92b-ac1ebd7b1ea3",
          "firstName": "John",
          "lastName": "Doe",
          "email": "john.doe@example.com"
        }
      }
    ]
  }
}
```

### CheckPermission

Checks if a user has a specific permission for a given object and returns a boolean value. this query is using by other api to check if a user has a specified permission or not

```graphql
query (
  $userId: ID!
  $nameSpace: NameSpaceEnum!
  $permission: PermissionEnum!
  $object: String!
) {
  CheckPermission(
    userId: $userId
    nameSpace: $nameSpace
    permission: $permission
    object: $object
  )
}
```

Variables:

```json
{
  "userId": "58500165-593c-471d-b92b-ac1ebd7b1ea3",
  "nameSpace": "shifts",
  "permission": "WRITE_ALL",
  "object": "shared_schedule"
}
```

### Response

```json
{
  "data": {
    "CheckPermission": true
  }
}
```

**Note:** Replace `Variables` data with your actual data.
