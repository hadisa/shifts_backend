# Request Time Off

## Vault

> Shell Command to Create AppRole Role and Secret ID for `Request Time Off` Service

```bash
vault write auth/approle/role/request_time_off token_policies="shifts" token_ttl=10h token_max_ttl=18h
vault read auth/approle/role/request_time_off/role-id
vault write -f auth/approle/role/request_time_off/secret-id

```

### Role Creation statements

```text
CREATE ROLE "{{name}}" WITH LOGIN PASSWORD '{{password}}' VALID UNTIL '{{expiration}}' INHERIT;
GRANT USAGE ON SCHEMA public TO "{{name}}";
GRANT SELECT, UPDATE ,INSERT, DELETE ON ALL TABLES IN SCHEMA public TO "{{name}}";
ALTER USER "{{name}}" WITH SUPERUSER;
```

## Graphql

### Query

#### getRequestTimeOff

getRequestTimeOff(id: ID!, authUserId: ID): RequestTimeOff!

This query returns a single RequestTimeOff object based on the provided id. The authUserId parameter is optional and can be used to authenticate the user.

```graphql
query GetRequestTimeOffQuery($id: ID!, $authUserId: ID) {
  getRequestTimeOff(id: $id, authUserId: $authUserId) {
    id
    userId
    status
    reason
    requestNote
    responseNote
    startTime
  }
}
```

Variables:

```json
{
  "id": "8785714d-e7cb-4638-9891-dd073790b1d9",
  "authUserId": "58500165-593c-471d-b92b-ac1ebd7b1ea3"
}
```

#### getRequestTimeOffs

getRequestTimeOffs(authUserId: ID): [RequestTimeOff]!

This query returns a list of RequestTimeOff objects. If the authUserId parameter is provided, it will filter the results to only show requests made by the user.

```graphql
query GetRequestTimeOffsQuery($authUserId: ID) {
  getRequestTimeOffs(authUserId: $authUserId) {
    id
    userId
    status
    reason
    requestNote
    responseNote
    startTime
  }
}
```

Variables:

```json
{
  "authUserId": "58500165-593c-471d-b92b-ac1ebd7b1ea3"
}
```

#### Response

```json
{
  "data": {
    "getRequestTimeOffs": [
      {
        "id": "8785714d-e7cb-4638-9891-dd073790b1d9",
        "userId": "58500165-593c-471d-b92b-ac1ebd7b1ea3",
        "status": "DENIED",
        "reason": "note reason",
        "requestNote": "a note here",
        "responseNote": "for test",
        "startTime": "2022-10-15T04:30:00+04:30"
      }
    ]
  }
}
```

#### getRequestTimeOffsByChannelIdRequestId

> For other services

getRequestTimeOffsByChannelIdRequestId(channelId: ID!, requestId: ID!): RequestTimeOff!

This query returns a single RequestTimeOff object based on the provided channelId and requestId.

### Mutation

#### createRequestTimeOff

createRequestTimeOff(input: RequestTimeOffInput!, authUserId: ID): TimeOffResponse

This mutation creates a new RequestTimeOff object using the provided input parameters. The authUserId parameter is optional and can be used to authenticate the user. The response returns a TimeOffResponse object which contains any errors that may have occurred and the newly created request.

```graphql
mutation CreateRequestTimeOffMutation(
  $input: RequestTimeOffInput!
  $authUserId: ID
) {
  createRequestTimeOff(input: $input, authUserId: $authUserId) {
    request {
      id
      channelId
      requestNote
      reason
      responseNote
      isAllDay
      user {
        id
        firstName
        lastName
      }
    }
  }
}
```

Variables:

```json
{
  "input": {
    "userId": "58500165-593c-471d-b92b-ac1ebd7b1ea3",
    "channelId": "4a2cfef2-a6f9-4c2c-8ed5-062f4e2c9712",
    "startTime": "2022-10-15T00:00:00.000Z",
    "endTime": "2023-11-15T00:00:00.000Z",
    "is24Hours": false,
    "reason": "note reason",
    "requestNote": "a note here",
    "responseNote": "for test",
    "responseByUserId": "58500165-593c-471d-b92b-ac1ebd7b1ea3",
    "responseAt": "2023-10-15T00:00:00.000Z"
  },
  "authUserId": "58500165-593c-471d-b92b-ac1ebd7b1ea3"
}
```

#### updateRequestTimeOff

updateRequestTimeOff(id: ID!, input: RequestTimeOffInput!, authUserId: ID): TimeOffResponse

This mutation updates an existing RequestTimeOff object based on the provided id and input parameters. The authUserId parameter is optional and can be used to authenticate the user. The response returns a TimeOffResponse object which contains any errors that may have occurred and the updated request.

```graphql
mutation UpdateRequestTimeOffMutation(
  $id: ID!
  $input: RequestTimeOffInput!
  $authUserId: ID
) {
  updateRequestTimeOff(input: $input, authUserId: $authUserId) {
    request {
      id
      channelId
      requestNote
      reason
      responseNote
      isAllDay
      user {
        id
        firstName
        lastName
      }
    }
  }
}
```

Variables:

```json
{
  "input": {
    "userId": "58500165-593c-471d-b92b-ac1ebd7b1ea3",
    "channelId": "4a2cfef2-a6f9-4c2c-8ed5-062f4e2c9712",
    "startTime": "2022-10-15T00:00:00.000Z",
    "endTime": "2023-11-15T00:00:00.000Z",
    "is24Hours": false,
    "reason": "note reason",
    "requestNote": "a note here",
    "responseNote": "for test",
    "responseByUserId": "58500165-593c-471d-b92b-ac1ebd7b1ea3",
    "responseAt": "2023-10-15T00:00:00.000Z"
  },
  "authUserId": "58500165-593c-471d-b92b-ac1ebd7b1ea3",
  "id": "8785714d-e7cb-4638-9891-dd073790b1d9"
}
```

#### deleteRequestTimeOff

deleteRequestTimeOff(id: ID!, authUserId: ID): TimeOffResponse

This mutation deletes an existing RequestTimeOff object based on the provided id. The authUserId parameter is optional and can be used to authenticate the user. The response returns a TimeOffResponse object which contains any errors that may have occurred and the deleted request.

```graphql
mutation DeleteRequestTimeOffMutation($id: ID!, $authUserId: ID) {
  deleteRequestTimeOff(id: $id, authUserId: $authUserId) {
    errors {
      code
      field
      message
    }
    request {
      id
      status
    }
  }
}
```

Variables:

```json
{
  "id": "8785714d-e7cb-4638-9891-dd073790b1d9",
  "authUserId": "58500165-593c-471d-b92b-ac1ebd7b1ea3"
}
```

#### cancelRequestTimeOff

cancelRequestTimeOff(channelId: ID!, requestId: ID!, authUserId: ID): TimeOffResponse

This mutation cancels an existing RequestTimeOff object based on the provided channelId and requestId. The authUserId parameter is optional and can be used to authenticate the user. The response returns a TimeOffResponse object which contains any errors that may have occurred and the cancelled request.

```graphql
mutation CancelRequestTimeOffMutation(
  $channelId: ID!
  $requestId: ID!
  $authUserId: ID
) {
  cancelRequestTimeOff(
    channelId: $channelId
    requestId: $requestId
    authUserId: $authUserId
  ) {
    errors {
      code
      field
      message
    }
    request {
      id
      status
    }
  }
}
```

Variables:

```json
{
  "channelId": "4a2cfef2-a6f9-4c2c-8ed5-062f4e2c9712",
  "requestId": "8785714d-e7cb-4638-9891-dd073790b1d9",
  "authUserId": "58500165-593c-471d-b92b-ac1ebd7b1ea3"
}
```

#### approveRequestTimeOff

approveRequestTimeOff(id: ID!, responseNote: String, authUserId: ID): TimeOffResponse

This mutation approves an existing RequestTimeOff object based on the provided id. The authUserId parameter is optional and can be used to authenticate the user. The response returns a TimeOffResponse object which contains any errors that may have occurred and the approved request.

```graphql
mutation ApproveRequestTimeOffMutation(
  $id: ID!
  $responseNote: String
  $authUserId: ID
) {
  approveRequestTimeOff(
    id: $id
    responseNote: $responseNote
    authUserId: $authUserId
  ) {
    errors {
      code
      field
      message
    }
    request {
      id
      status
      responseNote
    }
  }
}
```

Variables:

```json
{
  "id": "8785714d-e7cb-4638-9891-dd073790b1d9",
  "responseNote": "leave your note here",
  "authUserId": "58500165-593c-471d-b92b-ac1ebd7b1ea3"
}
```

#### denyRequestTimeOff

denyRequestTimeOff(id: ID!, responseNote: String, authUserId: ID): TimeOffResponse

This mutation denies an existing RequestTimeOff object based on the provided id. The authUserId parameter is optional and can be used to authenticate the user. The response returns a TimeOffResponse object which contains any errors that may have occurred and the denied request.

```graphql
mutation DenyRequestTimeOffMutation(
  $id: ID!
  $responseNote: String
  $authUserId: ID
) {
  denyRequestTimeOff(
    id: $id
    responseNote: $responseNote
    authUserId: $authUserId
  ) {
    errors {
      code
      field
      message
    }
    request {
      id
      status
      responseNote
    }
  }
}
```

Variables:

```json
{
  "id": "8785714d-e7cb-4638-9891-dd073790b1d9",
  "responseNote": "leave your note here",
  "authUserId": "58500165-593c-471d-b92b-ac1ebd7b1ea3"
}
```

**Note:** Replace `Variables` data with your actual data.
