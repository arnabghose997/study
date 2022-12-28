# REST v/s GraphQL v/s gRPC

## REST (Representation State Transfer)

- Interact with resources from a given URI using HTTP requests such as GET, POST, etc
- Resources are identified are using paths. For example: `/users/` and `/users/{id}`
- Request and Response are composed in a JSON payload

**Pros**

- Mature (its there since 2000)
- Programming language agnostic since HTTP is involved
- It's easy to debug

**Cons**

- Prone to underfetching. So, if there's a `/users` endpoint which returns a list of users, however we can only do so much and can't expect the endpoint to breif us about user's details. For that, we have to loop through the response, get the `user_id` and call the endpoint `/user/{id}`.

- Pront to overfetching. Sometimes we don't need all the data from the response 


## GrapQL

- It utilises a Single API endpoint, and we specify the schema that represents the necessary fields that are requried to be fetched


**Pros**

- Eliminates over and under fetching
- Queries are strongly typed


**Cons**

- Querying logic is moved to backend
- Caching is complex, each client requests field based on their need.

**Tools in GO**: `gqlgen`

## gRPC

- It is a specification that typically uses Protocol Buffers for communication
- Protocol Buffers can thought of as an agreement between client and servers, based on structure and types of data fields (both request and response).

**Pros**

- It generates code (related to request and response structure) in many programming languages
- It is highly performatant because of binary format the bandwidth of communication will be relatively smaller compared to JSOn

**Cons**
- Hard to implement
- No browser support

**Tools**: `grpc-go`, `protoc`, `buf`
