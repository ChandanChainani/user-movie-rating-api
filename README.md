# user-movie-rating-api

- Simple Authentication with email
- Search movie by name and get related information
- Logged in admin user can create movie
- Logged in user can rate or comment on movie
- Get user rated movies and related comments

# Requirements
## Minimum
- golang
- mongodb

## Optional
- docker
- docker-compose

# If docker and docker-compose is installed
Then you can simply run from root of the project `sudo docker-compose up`

# Mongo Setup
1) Start MongoDB without access control.
- `mongod --dbpath /data/db`

2) Connect to mongo instance
- `mongo`

3) switch to admin database and then create user
```
use admin
db.createUser(
  {
    user: "mongoadmin",
    pwd: "secret",
    roles: [ { role: "dbOwner", db: "rating_app" } ]
  }
)
```

4) Start mongo again with access control
- `mongod --auth --dbpath /data/db`

5) Connect and test authentication as the user
```
use rating_app
db.auth("mongoadmin", "secret")
```

# Steps to build and Run Test App
Before starting the service make sure
mongo is running and setup is complete if not follow [Mongo Setup](#mongo-setup)

- change `url` from `mongodb://database:27017` to `mongodb://localhost:27017` in [config/config.json](https://github.com/ChandanChainani/user-movie-rating-api/blob/main/config/config.json)
- open the project folder in terminal and run `go build -o main .`
- then you can run the `main` binary which has been generated

# Note
- I have added migration functionality that will generate sample data on app starts
- Postman request response collection is stored inside [postman](https://github.com/ChandanChainani/user-movie-rating-api/blob/main/postman/user-movie-rating-api.postman_collection.json) folder
- mongo queries are stored in [dump/queries.txt]( https://github.com/ChandanChainani/user-movie-rating-api/blob/main/dump/queries.txt )
