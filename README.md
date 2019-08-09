# commnet
### Communications network library that links together NATS and Redis

DOCS NOT COMPLETE. WORKING ON IT.

General usage is as follows:

```go
redisConnection := "localhost"
natsConnection := "localhost"
c := comms.NewComm(redisConnection, natsConnection)
```

From here, the variable 'c' is used to access both the Redis database and the NATS messager in a simplified form.
It represents a communications key, which hold the string id of the key and allows you to read and write various objects to redis and publish and subscribe to the NATS channel that matches that key.

The full API is shown below:

```go
// Gets the 'users' key.
users := c.Key("users")

// Gets the 'users.4' key. Notice the addition of the '.' in between.
user4 := users.Key("4")

// Get the string of the communications key. Should just assign 'users.4' to the variable.
idString := user4.GetID()

// Get the value of a key.
name := user4.Key("name")
nameString := name.Get()

// Get the value of a key as an integer.
age := user4.Key("age")
ageValue := age.Integer()

// Set the value of a key.
name.Set("joe")

// Increment the value of the key by 1 and return the number.
newAge := age.Incr()

// TODO: Continue examples.
```

There is no error checking at the moment. I will refactor that in soon. (Read below)
Honestly the naming could be better, but this is currently only being used in a game library, so... Yeah.
