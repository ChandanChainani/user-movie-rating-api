package db

import (
	"fmt"
	"context"
	// "encoding/json"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"usermovieratingapi/request"
	e "usermovieratingapi/errors"
)

func (db *DB) GetUserInfoByID(user *request.SearchUser) (map[string]interface{}, error) {
	collection, err := db.GetCollection("users")
	if err != nil {
		return nil, err
	}

	fmt.Println(user)
	uID, _ := primitive.ObjectIDFromHex(user.ID)
	fmt.Println(uID)

	userMatchStage := bson.D{{ "$match", bson.M{ "_id": uID } }}
	// userRatedMoviesCommentsStage := bson.D{{ "$lookup", bson.M{ "from": "ratings", "let": bson.M{ "uID": "$_id" }, "pipeline": bson.A{ bson.M{ "$match": bson.M{ "$expr": bson.M{ "$eq": bson.A{ "$user_id", "$$uID" } } } }, bson.M{ "$lookup": bson.M{ "from": "movies", "let": bson.M{ "mID": "$movie_id" }, "pipeline": bson.A{ bson.M{ "$match": bson.M{ "$expr": bson.M{ "$eq": bson.A{ "$_id", "$$mID" } } } }, bson.M{ "$lookup": bson.M{ "from": "comments", "let": bson.M{ "mID": "$_id" }, "pipeline": bson.A{ bson.M{ "$match": bson.M{ "$expr": bson.M{ "$and": bson.A{ bson.M{ "$eq": bson.A{ "$movie_id", "$$mID" } }, bson.M{ "$eq": bson.A{ "$user_id", "$$uID" } } } } } } }, "as": "comments" } } }, "as": "movies" } } }, "as": "ratings" } }}
	// addFieldsStage := bson.D{{ "$addFields", bson.M{ "movies": bson.M{ "$map": bson.M{ "input": "$ratings.movies", "as": "m", "in": bson.M{ "name": bson.M{ "$arrayElemAt": bson.A{ "$$m.name", 0 } }, "comments": bson.M{ "$arrayElemAt": bson.A{ "$$m.comments", 0 } }, } } } } }}
	// projectStage := bson.D{{ "$project", bson.M{ "ratings": 0, "movies.comments._id": 0, "movies.comments.user_id": 0, "movies.comments.movie_id": 0 } }}
	userRatedMoviesCommentsStage := bson.D{{ "$lookup", bson.M{ "from": "ratings", "let": bson.M{ "uID": "$_id" }, "pipeline": bson.A{ bson.M{ "$match": bson.M{ "$expr": bson.M{ "$eq": bson.A{ "$user_id", "$$uID" } } } }, bson.M{ "$lookup": bson.M{ "from": "movies", "let": bson.M{ "mID": "$movie_id" }, "pipeline": bson.A{ bson.M{ "$match": bson.M{ "$expr": bson.M{ "$eq": bson.A{ "$_id", "$$mID" } } } }, bson.M{ "$lookup": bson.M{ "from": "comments", "let": bson.M{ "mID": "$_id" }, "pipeline": bson.A{ bson.M{ "$match": bson.M{ "$expr": bson.M{ "$and": bson.A{ bson.M{ "$eq": bson.A{ "$movie_id", "$$mID" } }, bson.M{ "$eq": bson.A{ "$user_id", "$$uID" } } } } } } }, "as": "comments" } } }, "as": "movies" } }, }, "as": "ratings" } }}
	unwindRatingsState := bson.D{{ "$unwind", "$ratings" }}
	groupStage := bson.D{{ "$group", bson.M{ "_id": "null", "movies": bson.M{ "$addToSet": bson.M{ "$arrayElemAt": bson.A{ "$ratings.movies", 0 } } } } }}
	projectStage := bson.D{{ "$project", bson.M{ "_id": 0, "movies._id": 0, "movies.comments._id": 0, "movies.comments.user_id": 0, "movies.comments.movie_id": 0 } }}

  // userPipeline := mongo.Pipeline{userMatchStage, userRatedMoviesCommentsStage, addFieldsStage, projectStage}
  userPipeline := mongo.Pipeline{userMatchStage, userRatedMoviesCommentsStage, unwindRatingsState, groupStage, projectStage}

	cur, err := collection.Aggregate(context.TODO(), userPipeline)
	if err != nil {
		return nil, err
	}
	// once exhausted, close the cursor
	defer cur.Close(context.TODO())

	if !cur.Next(context.TODO()) {
		return nil, e.NoDataFoundError
	}

	// var result bson.M
	var result map[string]interface{}
	if err = cur.Decode(&result); err != nil {
		return nil, err
	}

  // jsonBytes, err := json.Marshal(result)
	// if err != nil {
		// return "", err
	// }

	return result, nil
	// return string(jsonBytes), nil
}

func (db *DB) GetUserByEmail(u *request.User) (map[string]interface{}, error) {
	collection, err := db.GetCollection("users")
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = collection.FindOne(context.TODO(), bson.M{ "email": u.Email }).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
