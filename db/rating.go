package db

import (
	"fmt"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"usermovieratingapi/types/request"
)

func (db *DB) InsertUserMovieRating(r *request.UserMovieRating) (string, error) {
	collection, err := db.GetCollection("ratings")
	if err != nil {
		return "", err
	}
	uID, _ := primitive.ObjectIDFromHex(r.UserID)
	mID, _ := primitive.ObjectIDFromHex(r.MovieID)
	fmt.Println(r.Rating)

	res, err := collection.InsertOne(context.TODO(), bson.M{ "user_id": uID, "movie_id": mID, "rating": r.Rating})
	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}
