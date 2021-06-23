# clean the cache
clean:
	-go clean -cache 

# builds the binary
build:
	go build

# checks code quality 
check:
	go vet -asmdecl -bools -assign ./

# rebuild container
reload:
	sudo docker-compose down
	sudo docker-compose build
	sudo docker-compose up

# start user-movie-rating-api
start:
	sudo docker-compose up
