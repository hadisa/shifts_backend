# Shift Group Member

## Vault

> Shell Command to Create AppRole Role and Secret ID for `Shift Group Member` Service

```bash
vault write auth/approle/role/shift_group_member token_policies="shifts" token_ttl=10h token_max_ttl=18h
vault read auth/approle/role/shift_group_member/role-id
vault write -f auth/approle/role/shift_group_member/secret-id

```

### Role Creation statements

```text
CREATE ROLE "{{name}}" WITH LOGIN PASSWORD '{{password}}' VALID UNTIL '{{expiration}}' INHERIT;
GRANT USAGE ON SCHEMA public TO "{{name}}";
GRANT SELECT, UPDATE ,INSERT, DELETE ON ALL TABLES IN SCHEMA public TO "{{name}}";
ALTER USER "{{name}}" WITH SUPERUSER;
```

## GraphQL

### Query

#### getNonShiftGroupMembers

This query returns all users in the given channel that are not members of the given shift group. It takes the following arguments:

- channelId: ID! - Required ID of the channel to search for users.
- shiftGroupId: ID! - Required ID of the shift group to exclude its members.
- authUserId: ID - Optional ID of the authenticated user.
  It returns a getNonShiftGroupMembersResponse object that contains:
  - message: String
    A message returned by the query.
  - result: [User]!
    A list of User objects that are not members of the shift group in the channel.
  - status: String
    A status message returned by the query.

```graphql
query GetNonShiftGroupMembersQuery(
  $shiftGroupId: ID!
  $channelId: ID!
  $authUserId: ID
) {
  getNonShiftGroupMembers(
    shiftGroupId: $shiftGroupId
    channelId: $channelId
    authUserId: $authUserId
  ) {
    message
    status
    result {
      id
      firstName
      lastName
      email
      languageCode
      avatar
      isStaff
      isActive
    }
  }
}
```

Variables:

```json
{
  "shiftGroupId": "c562f1a0-86a3-4e3e-a5bc-07168eb159bd",
  "channelId": "4a2cfef2-a6f9-4c2c-8ed5-062f4e2c9712",
  "authUserId": "58500165-593c-471d-b92b-ac1ebd7b1ea3"
}
```

#### getShiftGroupMembers

This query returns all members of the given shift group in the given channel. It takes the following arguments:

- shiftGroupId: ID! - Required ID of the shift group to get its members.
- channel: String! - Required name of the channel where the shift group is located.
- authUserId: ID - Optional ID of the authenticated user.
  It returns a list of User objects that are members of the shift group in the channel.

```graphql
query GetShiftGroupMembersQuery(
  $shiftGroupId: ID!
  $channel: String!
  $authUserId: ID
) {
  getShiftGroupMembers(
    shiftGroupId: $shiftGroupId
    channel: $channel
    authUserId: $authUserId
  ) {
    id
    firstName
    lastName
    email
    languageCode
    avatar
    isStaff
    isActive
  }
}
```

Variables:

```json
{
  "shiftGroupId": "c562f1a0-86a3-4e3e-a5bc-07168eb159bd",
  "channel": "default-channel",
  "authUserId": "58500165-593c-471d-b92b-ac1ebd7b1ea3"
}
```

#### getAllShiftMembers

This query returns all users in the system. It takes the following arguments:

- first: Int - Optional number of items to retrieve from the beginning of the list.
- last: Int - Optional number of items to retrieve from the end of the list.
- authUserId: ID - Optional ID of the authenticated user.
  It returns a list of all User objects in the system.

```graphql
query GetAllShiftMembersQuery($authUserId: ID) {
  getAllShiftMembers(authUserId: $authUserId) {
    id
    firstName
    lastName
    email
  }
}
```

Variables:

```json
{
  "authUserId": "58500165-593c-471d-b92b-ac1ebd7b1ea3"
}
```

#### getAllUniqueShifts

This query returns all unique assigned shifts and open shifts in the given channel and shift group. It takes the following arguments:

- channelId: ID! - Required ID of the channel to search for shifts.
- shiftGroupId: ID! - Required ID of the shift group to search for shifts.
- authUserId: ID - Optional ID of the authenticated user.
  It returns a getAllUniqueShiftsResponse object that contains:
  - message: String - A message returned by the query.
  - result: UniqueShifts - An object that contains two lists of shifts:
    - assignedShifts: [AssignedShift] - A list of unique AssignedShift objects.
    - openShifts: [OpenShift] - A list of unique OpenShift objects.
    - status: String - A status message returned by the query.

```graphql
query GetAllUniqueShiftsMutation(
  $shiftGroupId: ID!
  $channelId: ID!
  $authUserId: ID
) {
  getAllUniqueShifts(
    shiftGroupId: $shiftGroupId
    channelId: $channelId
    authUserId: $authUserId
  ) {
    message
    status
    result {
      assignedShifts {
        id
        color
        label
        startTime
        endTime
        ShiftActivities {
          id
          name
          color
          startTime
          endTime
        }
      }
      openShifts {
        id
        color
        label
        startTime
        endTime
        ShiftActivities {
          code
          id
          color
          startTime
          endTime
        }
      }
    }
  }
}
```

Variables:

```json
{
  "shiftGroupId": "c562f1a0-86a3-4e3e-a5bc-07168eb159bd",
  "channelId": "4a2cfef2-a6f9-4c2c-8ed5-062f4e2c9712",
  "authUserId": "58500165-593c-471d-b92b-ac1ebd7b1ea3"
}
```

#### getShiftsByPeople

The getShiftsByPeople query is used to retrieve shifts of a group of people within a specific time range. This query returns a GetShiftsResponse object that includes a Shifts object with information about assigned shifts and open shifts.
Arguments

- channelId (required): The ID of the channel in which the shifts are assigned.
- endDate (required): The end date of the time range for which shifts are being requested.
  filter (optional): An input object that allows you to filter which types of shifts to include in the response. The input object includes the following fields:
  - includeOpenShifts (optional): Boolean value indicating whether to include open shifts in the response.
  - includeRequests (optional): Boolean value indicating whether to include shift requests in the response.
  - includeShifts (optional): Boolean value indicating whether to include assigned shifts in the response.
  - shiftGroupIds (optional): An array of shift group IDs. Only shifts from the specified shift groups will be returned.
  - shiftGroupMemberIds (optional): An array of shift group member IDs. Only shifts assigned to the specified shift group members will be returned.
- shiftGroupId (required): The ID of the shift group containing the specified people.
- startDate (required): The start date of the time range for which shifts are being requested.
- authUserId (optional): The ID of the user making the request. This field is used for authorization purposes.

Return Values

- message: A message indicating the status of the request.
- result: A Shifts object with information about assigned shifts and open shifts.
  assignedShifts: An array of UserAssignedShifts objects, each representing a specific user's assigned shifts. Each UserAssignedShifts object includes the following fields:
  - image (optional): The URL of the user's avatar image.
  - name: The user's name.
  - numberOfHours: The total number of hours assigned to the user within the specified time range.
  - shifts: An array of AssignedShift objects, each representing a specific assigned shift for the user.
    Each AssignedShift object includes the following fields:
  - id: The unique ID of the assigned shift.
  - break: The break time for the - assigned shift.
  - color: The color associated with the assigned shift.
  - startTime: The start time of the assigned shift.
  - endTime: The end time of the assigned shift.
  - is24Hours: A boolean value indicating whether the assigned shift lasts 24 hours.
  - label (optional): A label associated with the assigned shift.
  - note (optional): A note associated with the assigned shift. - shiftToOffer (optional): An AssignedShift object representing a shift that the user can offer to someone else.
  - shiftToSwap (optional): An AssignedShift object representing a shift that the user can swap with someone else.
  - toSwapWith (optional): An AssignedShift object representing a shift that the user has requested to swap with.
  - userId: The ID of the user to whom the assigned shift is assigned.
  - channelId: The ID of the channel in which the assigned shift is assigned.
  - shiftGroupId: The ID of the shift group in which the assigned shift is assigned.
  - type: A string value indicating the type of shift

```graphql
query GetShiftsByPeopleQuery(
  $channelId: ID!
  $startDate: Time!
  $endDate: Time!
  $shiftGroupId: ID!
  $includeOpenShifts: Boolean!
  $includeShifts: Boolean!
  $authUserId: ID
) {
  getShiftsByPeople(
    channelId: $channelId
    startDate: $startDate
    endDate: $endDate
    shiftGroupId: $shiftGroupId
    filter: {
      includeOpenShifts: $includeOpenShifts
      includeShifts: $includeShifts
    }
    authUserId: $authUserId
  ) {
    message
    status
    result {
      assignedShifts {
        userId
        name
        image
        numberOfHours
        shifts {
          id
          label
          note
          color
          ShiftActivities {
            id
            startTime
            endTime
          }
        }
      }
      openShifts {
        numberOfShifts
        title
        shifts {
          id
          channelId
          color
          label
          note
          ShiftActivities {
            id
            name
            code
          }
        }
      }
    }
  }
}
```

Variables:

```json
{
  "channelId": "4a2cfef2-a6f9-4c2c-8ed5-062f4e2c9712",
  "startDate": "2022-02-15T00:00:00.000Z",
  "endDate": "2023-12-15T00:00:00.000Z",
  "shiftGroupId": "c562f1a0-86a3-4e3e-a5bc-07168eb159bd",
  "includeOpenShifts": true,
  "includeShifts": true,
  "authUserId": "58500165-593c-471d-b92b-ac1ebd7b1ea3"
}
```

#### getShiftsByTask

The getShiftsByTask query returns the shifts of users in a given channel based on their assigned tasks during a specified time period.

getShiftsByTask(
channelId: ID!
startDate: Time!
endDate: Time!
filter: getShiftsFilter
authUserId: ID
): GetShiftsByTaskResponse

Arguments

The getShiftsByTask query takes the following arguments:

- channelId: ID! - The ID of the channel.
- startDate: Time! - The start date and time of the time period for which shifts are to be retrieved.
- endDate: Time! - The end date and time of the time period for which shifts are to be retrieved.
- filter: getShiftsFilter - An optional filter to limit the query results. The default value is null.
- authUserId: ID - The ID of the authenticated user.

Return Value

The getShiftsByTask query returns a GetShiftsByTaskResponse object, which contains the following fields:

- message: String
  - A message that indicates the status of the operation.
- result: [ShiftGroups]
  An array of ShiftGroups objects. Each ShiftGroup object contains the following fields:
  - groupId: ID!
    The ID of the shift group.
  - groupName: String!
    The name of the shift group.
  - shifts: Shifts
    A Shifts object that contains the following fields:
  - assignedShifts: [UserAssignedShifts]!
    An array of UserAssignedShifts objects. Each UserAssignedShifts object contains the following fields:
  - image: String
    The URL of the user's avatar.
  - name: String!
    The name of the user.
  - numberOfHours: Int!
    The total number of hours the user is scheduled to work during the specified time period.
  - shifts: [AssignedShift]
    An array of AssignedShift objects that represent the shifts assigned to the user during the specified time period.
  - userId: ID! - The ID of the user.
  - openShifts: OpenShiftInfo
    An OpenShiftInfo object that contains the following fields:
    - numberOfShifts: Int
      The number of open shifts in the specified time period.
    - shifts: [OpenShift]
      An array of OpenShift objects that represent the open shifts in the specified time period.
    - title: String
      The title of the open shifts section

```graphql
query GetShiftsByTaskQuery(
  $channelId: ID!
  $startDate: Time!
  $endDate: Time!
  $filter: GetShiftsFilter!
  $authUserId: ID
) {
  getShiftsByTask(
    channelId: $channelId
    startDate: $startDate
    endDate: $endDate
    filter: $filter
    authUserId: $authUserId
  ) {
    message
    status
    result {
      groupId
      groupName
      shifts {
        assignedShifts {
          userId
          name
          image
          numberOfHours
          shifts {
            id
            type
            label
            note
            color
            startTime
            endTime
            break
            is24Hours
            ShiftActivities {
              id
              name
              code
              color
              startTime
              endTime
              isPaid
            }
          }
        }
        openShifts {
          numberOfShifts
          title
          shifts {
            id
            label
            note
            color
            startTime
            endTime
            break
            is24Hours
            slots
            ShiftActivities {
              id
              name
              code
              color
              startTime
              endTime
              isPaid
            }
          }
        }
      }
    }
  }
}
```

Variables:

```json
{
  "channelId": "4a2cfef2-a6f9-4c2c-8ed5-062f4e2c9712",
  "startDate": "2020-10-15T00:00:00.000Z",
  "endDate": "2023-12-15T00:00:00.000Z",
  "filter": {
    "includeOpenShifts": true,
    "includeShifts": true,
    "includeRequests": true,
    "shiftGroupIds": [
      "your_shift_group_1",
      "your_shift_group_2",
      "your_shift_group_n"
    ]
  },
  "authUserId": "58500165-593c-471d-b92b-ac1ebd7b1ea3"
}
```

#### getShiftGroupMembersList

> Only For test

The getShiftGroupMembersList query retrieves a list of all ShiftGroupMember objects for a specified channelId and shiftGroupId.
Arguments

- channelId (required): ID of the channel to retrieve ShiftGroupMember objects from.
- shiftGroupId (required): ID of the shift group to retrieve ShiftGroupMember objects from.
- authUserId (optional): ID of the authenticated user. If provided, the query will only retrieve ShiftGroupMember objects that the authenticated user has permission to view.

Return Value

The query returns an array of ShiftGroupMember objects that belong to the specified channelId and shiftGroupId.

```graphql
query GetShiftGroupMembersListQuery(
  $channelId: ID!
  $shiftGroupId: ID!
  $authUserId: ID
) {
  getShiftGroupMembersList(
    channelId: $channelId
    shiftGroupId: $shiftGroupId
    authUserId: $authUserId
  ) {
    id
    channelId
    shiftGroupId
    userId
    position
  }
}
```

```json
{
  "channelId": "4a2cfef2-a6f9-4c2c-8ed5-062f4e2c9712",
  "shiftGroupId": "c562f1a0-86a3-4e3e-a5bc-07168eb159bd",
  "authUserId": "58500165-593c-471d-b92b-ac1ebd7b1ea3"
}
```

### Mutation

#### shiftGroupMemberAdd

shiftGroupMemberAdd mutation creates a new member of a shift group.
Input

input: Required. A ShiftGroupMemberInput object that contains:

- channelId: Required. The ID of the channel that the shift group is in.
- shiftGroupId: Required. The ID of the shift group to add the member to.
- userId: Required. The ID of the user to add as a member of the shift group.

Response

ShiftGroupMemberAddResponse: An object that contains:

- errors: An array of ShiftError objects that indicate any errors that occurred during the mutation.
- user: The User object that was added as a member of the shift group.

```graphql
mutation AddShiftGroupMemberMutation(
  $channelId: ID!
  $shiftGroupId: ID!
  $userId: ID!
  $authUserId: ID
) {
  shiftGroupMemberAdd(
    input: {
      channelId: $channelId
      shiftGroupId: $shiftGroupId
      userId: $userId
    }
    authUserId: $authUserId
  ) {
    errors {
      code 
      field
      message
    }
    user {
      id
      firstName
      lastName
    }
  }
}
```

Variables:

```json
{
  "channelId": "4a2cfef2-a6f9-4c2c-8ed5-062f4e2c9712",
  "shiftGroupId": "6850aac4-4dcf-4a39-bbb3-1c245770ddd4",
  "userId": "91da8f38-31e0-4d27-9dc1-53cc0e43c5bb",
  "authUserId": "58500165-593c-471d-b92b-ac1ebd7b1ea3"
}
```

#### shiftGroupMembersReorder

shiftGroupMembersReorder mutation reorders the members of a shift group.
Input

- channelId: Required. The ID of the channel that the shift group is in.
- shiftGroupId: Required. The ID of the shift group to reorder the members of.
- userIds: Required. An array of the user IDs in the order that the members should be sorted.
- authUserId: Optional. The ID of the user making the request. If specified, the user must be a member of the channel that the shift group is in.

Response

ResponseStatus: An object that contains:

- message: A message indicating the status of the request.
- status: A string indicating the status of the request.

```graphql
mutation ShiftGroupMembersReorderMutation(
  $channelId: ID!
  $shiftGroupId: ID!
  $userIds: [ID!]!
  $authUserId: ID
) {
  shiftGroupMembersReorder(
    channelId: $channelId
    shiftGroupId: $shiftGroupId
    userIds: $userIds
    authUserId: $authUserId
  ) {
    errors {
      field
      message
      code
    }
    user {
      id
      firstName
      lastName
    }
  }
}
```

Variables:

```json
{
  "channelId": "4a2cfef2-a6f9-4c2c-8ed5-062f4e2c9712",
  "shiftGroupId": "6850aac4-4dcf-4a39-bbb3-1c245770ddd4",
  "userId": ["id_1", "id_2", "id_3"],
  "authUserId": "58500165-593c-471d-b92b-ac1ebd7b1ea3"
}
```

#### shiftGroupMemberRemove

shiftGroupMemberRemove mutation removes a member from a shift group.
Input

- channelId: Required. The ID of the channel that the shift group is in.
- shiftGroupId: Required. The ID of the shift group to remove the member from.
- userId: Required. The ID of the user to remove from the shift group.
- authUserId: Optional. The ID of the user making the request. If specified, the user must be a member of the channel that the shift group is in.

Response

ShiftGroupMemberRemoveResponse: An object that contains:

- errors: An array of ShiftError objects that indicate any errors that occurred during the mutation.
- user: The User object that was removed from the shift group.

```graphql
mutation ShiftGroupMemberRemoveMutation(
  $channelId: ID!
  $shiftGroupId: ID!
  $userId: ID!
  $authUserId: ID
) {
  shiftGroupMemberRemove(
    channelId: $channelId
    shiftGroupId: $shiftGroupId
    userId: $userId
    authUserId: $authUserId
  ) {
    errors {
      field
      message
      code
    }
    user {
      id
      firstName
      lastName
    }
  }
}
```

Variables:

```json
{
  "channelId": "4a2cfef2-a6f9-4c2c-8ed5-062f4e2c9712",
  "shiftGroupId": "6850aac4-4dcf-4a39-bbb3-1c245770ddd4",
  "userId": "91da8f38-31e0-4d27-9dc1-53cc0e43c5bb",
  "authUserId": "58500165-593c-471d-b92b-ac1ebd7b1ea3"
}
```

**Note:** Replace `Variables` data with your actual data.
