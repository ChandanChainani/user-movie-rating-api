# user-movie-rating-api

- Simple Authentication with using email
- Search movie by name and get related information
- Logged in admin user can create movie
- Logged in user can rate or comment on movie
- Get user rates movies and related comments

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
Login in to mongoshell through terminal and run these commands
- use rating_app
- db.auth("mongoadmin", "secret")

# Steps to build and Run Api
Before starting the service make sure
mongo is running is setup if not follow [Mongo Setup](#mongo-setup)

Then open the project folder in terminal and run this commands
- `go build -o main .`
- then you can run the main binary which has been generate `./main`

# Note
- I have added migration functionality that will generate sample data on app starts
- Postman request response collection is stored inside `postman` folder
- mongo queries are stored in dump/queries.txt
