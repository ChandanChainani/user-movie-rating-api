package db

import (
	"fmt"
	"context"
	"math/rand"

	"go.mongodb.org/mongo-driver/bson"

	// "usermovieratingapi/request"
)

func (db *DB) RunMigrations() {
	var users = bson.A{
		bson.M{"name": "John", "email": "john@mongodb.com", "admin": 1 },
		bson.M{"name": "Johnny", "email": "johnny@mongodb.com", "admin": 0 },
		bson.M{"name": "Sunny", "email": "sunny@mongodb.com", "admin": 0 },
	}

	var movies = bson.A{
		bson.M{"name": "The One", "description": "" },
		bson.M{"name": "Ghost Rider", "description": "" },
		bson.M{"name": "Ghost Buster", "description": "" },
	}

	u := db.connection.Collection("users")
	m := db.connection.Collection("movies")
	r := db.connection.Collection("ratings")
	c := db.connection.Collection("comments")

	u.Drop(context.TODO())
	m.Drop(context.TODO())
	r.Drop(context.TODO())
	c.Drop(context.TODO())

	userIds, _ := u.InsertMany(context.TODO(), users)
	movieIds, _ := m.InsertMany(context.TODO(), movies)

	r1 := rand.New(rand.NewSource(99))

	var comments bson.A
	var ratings bson.A
	for i := 1; i <= 10; i++ { 
		comments_count := r1.Intn(5)
		urand := r1.Intn(3)
		user := users[urand].(bson.M)
		uID, mID := userIds.InsertedIDs[urand], movieIds.InsertedIDs[r1.Intn(3)]
		for j := 1; j <= comments_count; j++ {
			comment := fmt.Sprintf("%v : %v comment", user["name"], j)
			comments = append(comments, bson.M{ "user_id": uID, "movie_id": mID, "comment": comment})
			ratings = append(ratings, bson.M{ "user_id": uID, "movie_id": mID, "rating":(r1.Float32() * 5) })
		}
	}

	c.InsertMany(context.TODO(), comments)
	r.InsertMany(context.TODO(), ratings)
}
