# Request Swaps

## Vault

> Shell Command to Create AppRole Role and Secret ID for `Request Swap` Service

```bash
vault write auth/approle/role/request_swap token_policies="shifts" token_ttl=10h token_max_ttl=18h
vault read auth/approle/role/request_swap/role-id
vault write -f auth/approle/role/request_swap/secret-id

```

## Role Creation statements

```text
CREATE ROLE "{{name}}" WITH LOGIN PASSWORD '{{password}}' VALID UNTIL '{{expiration}}' INHERIT;
GRANT USAGE ON SCHEMA public TO "{{name}}";
GRANT SELECT, UPDATE ,INSERT, DELETE ON ALL TABLES IN SCHEMA public TO "{{name}}";
ALTER USER "{{name}}" WITH SUPERUSER;
```

## Graphql

### Query

#### getRequestsSwaps

Returns a list of request swaps for the given channelId.

Arguments

- channelId (required): ID of the channel to fetch request swaps for.
- authUserId (optional): ID of the user making the request.

Returns

An array of RequestSwap objects.

```graphql
query GetRequestsSwapsQuery($channelId: ID!, $authUserId: ID) {
  getRequestsSwaps(channelId: $channelId, authUserId: $authUserId) {
    id
    userId
    channelId
    responseNote
    requestId
    requestNote
    status
    responseAt
    assignedUserShiftId
    assignedUserShiftIdToSwap
    createdAt
  }
}
```

Variables:

```json
{
  "authUserId": "58500165-593c-471d-b92b-ac1ebd7b1ea3",
  "channelId": "4a2cfef2-a6f9-4c2c-8ed5-062f4e2c9712"
}
```

#### Response

```json
{
  "data": {
    "getRequestsSwaps": [
      {
        "id": "b8bd010a-9b9f-452e-9267-518b75d67dc0",
        "userId": "58500165-593c-471d-b92b-ac1ebd7b1ea3",
        "channelId": "4a2cfef2-a6f9-4c2c-8ed5-062f4e2c9712",
        "responseNote": "a note here as well",
        "requestId": "138aa7e7-0a4d-41f9-b498-3bb085b0a679",
        "requestNote": "a note here 33",
        "status": "PENDING",
        "responseAt": "2022-10-15T04:30:00+04:30",
        "assignedUserShiftId": "b72a6635-f3e3-4ad1-8b85-d61addc94c41",
        "assignedUserShiftIdToSwap": "7261620f-657c-48c6-a0ed-e9ff447defc0",
        "createdAt": "2023-03-07T00:33:54.699689+04:30"
      }
    ]
  }
}
```

#### getRequestsSwapsByChannelIdRequestId

Returns a single request swap for the given channelId and requestId.

Arguments

- channelId (required): ID of the channel to fetch request swaps for.
- requestId (required): ID of the request swap to fetch.

Returns

A RequestSwap object.

#### getRequestSwap

Returns a single request swap for the given id.

Arguments

- id (required): ID of the request swap to fetch.
- authUserId (optional): ID of the user making the request.

Returns

A RequestSwap object.

```graphql
query GetRequestsSwapQuery($id: ID!, $authUserId: ID) {
  getRequestsSwap(id: $id, authUserId: $authUserId) {
    id
    userId
    channelId
    responseNote
    requestId
    requestNote
    status
    responseAt
    assignedUserShiftId
    assignedUserShiftIdToSwap
    createdAt
  }
}
```

Variables:

```json
{
  "id": "b8bd010a-9b9f-452e-9267-518b75d67dc0",
  "authUserId": "58500165-593c-471d-b92b-ac1ebd7b1ea3"
}
```

### Mutation

#### createRequestSwap

Creates a new request swap.

Arguments

- input (required): Input object containing the details of the request swap.
- authUserId (optional): ID of the user making the request.

Returns

A RequestSwapResponse object.

```graphql
mutation CreateRequestSwapMutation($input: RequestSwapInput!, $authUserId: ID) {
  createRequestSwap(input: $input, authUserId: $authUserId) {
    errors {
      code
      field
      message
    }
    request {
      channelId
      id
      requestId
      type
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
    "channelId": "4a2cfef2-a6f9-4c2c-8ed5-062f4e2c9712",
    "userId": "58500165-593c-471d-b92b-ac1ebd7b1ea3",
    "assignedUserShiftId": "b72a6635-f3e3-4ad1-8b85-d61addc94c41",
    "assignedUserShiftIdToSwap": "7261620f-657c-48c6-a0ed-e9ff447defc0",
    "requestNote": "a note here 33",
    "responseNote": "a note here as well",
    "responseByUserId": "58500165-593c-471d-b92b-ac1ebd7b1ea3",
    "responseAt": "2022-10-15T00:00:00.000Z"
  },
  "authUserId": "58500165-593c-471d-b92b-ac1ebd7b1ea3"
}
```

#### updateRequestSwap

Updates an existing request swap.

Arguments

- id (required): ID of the request swap to update.
- input (required): Input object containing the updated details of the request swap.
- authUserId (optional): ID of the user making the request.

Returns

A RequestSwapResponse object.

#### deleteRequestSwap

Deletes a request swap.

Arguments

- id (required): ID of the request swap to delete.
- authUserId (optional): ID of the user making the request.

Returns

A RequestSwapResponse object.

```graphql
mutation DeleteRequestSwapMutation($id: ID!, $authUserId: ID) {
  deleteRequestSwap(id: $id, authUserId: $authUserId) {
    errors {
      code
      field
      message
    }
    request {
      id
      requestNote
      reason
      responseNote
      status
    }
  }
}
```

Variables:

```json
{
  "id": "b72a6635-f3e3-4ad1-8b85-d61addc94c41",
  "authUserId": "58500165-593c-471d-b92b-ac1ebd7b1ea3"
}
```

#### cancelRequestSwap

Cancels a request swap.

Arguments

- channelId (required): ID of the channel the request swap belongs to.
- requestId (required): ID of the request swap to cancel.
- authUserId (optional): ID of the user making the request.

Returns

A RequestSwapResponse object.

```graphql
mutation CancelRequestSwapMutation(
  $channelId: ID!
  $requestId: ID!
  $authUserId: ID
) {
  cancelRequestSwap(
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
      requestNote
      reason
      responseNote
      status
    }
  }
}
```

Variables:

```json
{
  "channelId": "4a2cfef2-a6f9-4c2c-8ed5-062f4e2c9712",
  "requestId": "b72a6635-f3e3-4ad1-8b85-d61addc94c41",
  "authUserId": "58500165-593c-471d-b92b-ac1ebd7b1ea3"
}
```

#### approveRequestSwap

Approves a request swap.

Arguments

- id (required): ID of the request swap to approve.
- responseNote (optional): Note of Response
- authUserId (optional): ID of the user making the request.

Returns

A RequestSwapResponse object.

```graphql
mutation ApproveRequestSwapMutation(
  $id: ID!
  $responseNote: String
  $authUserId: ID
) {
  approveRequestSwap(
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
      requestNote
      reason
      responseNote
      status
    }
  }
}
```

Variables:

```json
{
  "id": "b8bd010a-9b9f-452e-9267-518b75d67dc0",
  "responseNote": "leave your note here",
  "authUserId": "58500165-593c-471d-b92b-ac1ebd7b1ea3"
}
```

#### denyRequestSwap

Denies a request swap.

Arguments

- id (required): ID of the request swap to deny.
- responseNote (optional): Note of Response
- authUserId (optional): ID of the user making the request.

Returns

A RequestSwapResponse object.

```graphql
mutation DenyRequestSwapMutation(
  $id: ID!
  $responseNote: String
  $authUserId: ID
) {
  denyRequestSwap(
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
      requestNote
      reason
      responseNote
      status
    }
  }
}
```

Variables:

```json
{
  "id": "b8bd010a-9b9f-452e-9267-518b75d67dc0",
  "responseNote": "leave your note here",
  "authUserId": "58500165-593c-471d-b92b-ac1ebd7b1ea3"
}
```

Note: All mutation operations return a RequestSwapResponse object which contains either an array of ShiftError objects in case of errors, or a RequestResponse object containing the details of the updated request swap.

**Note:** Replace `Variables` data with your actual data.
