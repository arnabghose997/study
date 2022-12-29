# Protocol Buffers

- They are a data format that are used to serialise structured data
- It can be used for programs to communicate with each other, or for storing data
- Data definitions are defined in `.proto` files. These cannot be executed, but using a tool called `protoc`, we can generate code (in supported programming languages) based on the definitions in the `proto` files


## Alternates

- XML
- JSON
- Avro


## Messages

- Content of the payload and types of the fields
- Fields uses default values, if the payload doesn't have the value for a particular fields.
- Example
```proto
message User {
    string uuid = 1;
    uint32 salary = 2;
    // repeated indicates that a field accepts list of a particular type as values
    repeated Address address = 3; 
}

message Address {
    string street = 1;
    string city = 2;
}

```
- Using optional values for fields are not recommended.


## Enumerations

- Predefined the list of values
- Note: Always define the first field as zero
- Example

```proto
message User {
    string uuid = 1;
    uint32 salary = 2;
    // repeated indicates that a field accepts list of a particular type as values
    repeated Address address = 3;
    MaritalStatus marital_status = 4; 
}

enum MaritalStatus {
    MARITAL_STATUS_UNSPECIFIED = 0;
    MARITAL_STATUS_SINGLE = 1;
    MARITAL_STATUS_MARRIED = 2;
}

```
