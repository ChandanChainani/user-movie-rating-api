package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"usermovieratingapi/request"
)

func (db *DB) InsertUserMovieComment(c *request.UserMovieComment) (string, error) {
	collection, err := db.GetCollection("comments")
	if err != nil {
		return "", err
	}
	uID, _ := primitive.ObjectIDFromHex(c.UserID)
	mID, _ := primitive.ObjectIDFromHex(c.MovieID)

	res, err := collection.InsertOne(context.TODO(), bson.M{ "user_id": uID, "movie_id": mID, "comment": c.Comment })
	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}
