In the function bodies we’ll generally use the following workflow:

Protbuf Message (Request) → Regular Go Struct → Convert to BSON + Mongo Action → Protobuf Message (Response)

Now let’s implement the CreateLaptop method. Extract the fields from the request message and provide them to the CreateLaptop struct.