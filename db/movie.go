package db

import (
	"context"
	// "encoding/json"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"usermovieratingapi/request"
	e "usermovieratingapi/errors"
)

func (db *DB) GetMovieInfoByName(name string) (*map[string]interface{}, error) {
	collection, err := db.GetCollection("movies")
	if err != nil {
		return nil, err
	}

	movieMatchStage := bson.D{{ "$match", bson.D{{ "name", name }} }}
	movieRatingsLookupStage := bson.D{{ "$lookup", bson.M{ "from" : "ratings", "localField" : "_id", "foreignField" : "movie_id", "as" : "ratings" } }}
	// movieCommentsLookupStage := bson.D{{ "$lookup", bson.M{ "from": "comments", "localField": "_id", "foreignField": "movie_id", "as": "comments" } }}
	// movieCommentUserLookupStage := bson.D{{ "$lookup", bson.M{ "from": "users", "localField": "comments.user_id", "foreignField": "_id", "as": "users" } }}
	movieCommentsUserLookupStage := bson.D{{ "$lookup", bson.M{ "from": "comments", "let": bson.M{ "mID": "$_id" }, "pipeline": bson.A{ bson.M{ "$match": bson.M{ "$expr": bson.M{ "$eq": bson.A{ "$movie_id", "$$mID" } } } }, bson.M{ "$lookup": bson.M{ "let": bson.M{ "uID": "$user_id" }, "from": "users", "pipeline": bson.A{ bson.M{ "$match": bson.M{ "$expr": bson.M{ "$eq": bson.A{ "$_id", "$$uID" } }}} }, "as": "users" } }, bson.M{ "$unwind": "$users" } }, "as": "comments" }}}
	// addFieldsStage := bson.D{{ "$addFields", bson.M{ "comments": bson.M{ "$map": bson.M{ "input": bson.M{ "$range": bson.A{ 0, bson.M{ "$size": "$comments" } } }, "in": bson.M{ "author": bson.M{ "$arrayElemAt": bson.A{ "$users.name", "$$this" } }, "comment": bson.M{ "$arrayElemAt": bson.A{ "$comments.comment", "$$this" } }, } } } } }}
	addFieldsStage := bson.D{{ "$addFields", bson.M{ "comments": bson.M{ "$map": bson.M{ "input": bson.M{ "$range": bson.A{ 0, bson.M{ "$size": "$comments.users"} } }, "in": bson.M{ "author": bson.M{ "$arrayElemAt": bson.A{ "$comments.users.name", "$$this" } }, "comment": bson.M{ "$arrayElemAt": bson.A{ "$comments.comment", "$$this" } } } } } } }}
	projectStage := bson.D{{ "$project", bson.M{ "rating": bson.M{ "$trunc": bson.A{ bson.M{ "$avg": "$ratings.rating" }, 2} }, "count": bson.M{ "$size": "$ratings" }, "comments": 1 } }}

  moviePipeline := mongo.Pipeline{movieMatchStage, movieRatingsLookupStage, movieCommentsUserLookupStage, addFieldsStage, projectStage}

	cur, err := collection.Aggregate(context.TODO(), moviePipeline)
	if err != nil {
		return nil, err
	}
	// once exhausted, close the cursor
	defer cur.Close(context.Background())

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
		// return nil, err
	// }

	return &result, nil
	// return string(jsonBytes), nil
	// return jsonBytes, nil
}

func (db *DB) InsertMovie(m *request.Movie) (string, error) {
	collection, err := db.GetCollection("movies")
	if err != nil {
		return "", err
	}

	res, err := collection.InsertOne(context.TODO(), bson.M{ "name": m.Name, "description": m.Description })
	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}
