package football

import (
	"context"
	"github.com/ashlamp08/go-graphql-football/infrastructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func GetClubById(ctx context.Context, id int) (result interface{}) {
	var club Club
	data := infrastructure.Mongodb.Collection("clublist").FindOne(ctx, bson.M{"_id": id})
	data.Decode(&club)
	return club
}

func GetClubList(ctx context.Context, limit int) (result interface{}) {
	var club Club
	var clubs []Club

	option := options.Find().SetLimit(int64(limit))

	cur, err := infrastructure.Mongodb.Collection("clublist").Find(ctx, bson.M{}, option)
	defer cur.Close(ctx)
	if err != nil {
		log.Println(err)
		return nil
	}
	for cur.Next(ctx) {
		cur.Decode(&club)
		clubs = append(clubs, club)
	}
	return clubs
}

func CreateClub(ctx context.Context, club Club) error {
	club.Id, _ = getNextSequenceValue(context.Background(), "clubid")
	_, err := infrastructure.Mongodb.Collection("clublist").InsertOne(ctx, club)
	return err
}

func UpdateClub(ctx context.Context, club Club) error {
	filter := bson.M{"_id": club.Id}
	update := bson.M{"$set": club}
	upsertBool := true
	updateOption := options.UpdateOptions{
		Upsert: &upsertBool,
	}
	_, err := infrastructure.Mongodb.Collection("clublist").UpdateOne(ctx, filter, update, &updateOption)
	return err
}

func DeleteClubById(ctx context.Context, id int) error {
	_, err := infrastructure.Mongodb.Collection("clublist").DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func getNextSequenceValue(ctx context.Context, sequenceName string) (int32, error) {
	filter := bson.M{"_id": sequenceName}
	update := bson.M{"$inc": bson.M{"sequence_value": 1}}
	upsertBool := true
	after := options.After
	updateOption := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsertBool,
	}
	result := infrastructure.Mongodb.Collection("counters").FindOneAndUpdate(ctx, filter, update, &updateOption)
	if result.Err() != nil {
		return -1, result.Err()
	}

	// 9) Decode the result
	doc := bson.M{}
	decodeErr := result.Decode(&doc)
	return doc["sequence_value"].(int32), decodeErr
}
