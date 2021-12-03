package lotsdatabase

import (
	"auctionservice/models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LotsRepository struct {
	mongodb         *mongo.Database
	lots_collection string
}

func NewLotsRepository(db *mongo.Database, lots_coll string) *LotsRepository {
	return &LotsRepository{
		mongodb:         db,
		lots_collection: lots_coll,
	}
}

func (l *LotsRepository) AddLot(lot models.Lot) (string, error) {

	collection := l.mongodb.Collection(l.lots_collection)

	//filter := bson.D{{"auction_id", lot.Auction}}

	data := bson.D{
		primitive.E{Key: "lot_id", Value: lot.LotID},
		primitive.E{Key: "auction_id", Value: lot.Auction},
		primitive.E{Key: "owner_id", Value: lot.Owner},
		primitive.E{Key: "products", Value: bson.A{lot.Products}},
		primitive.E{Key: "start_price", Value: lot.Start_price},
	}

	_, err := collection.InsertOne(context.TODO(), data)
	if err != nil{
		log.Println(err)
	}

	return lot.LotID, err
}


func (l *LotsRepository) GetLot(lot_id string)(error){
	collection := l.mongodb.Collection(l.lots_collection)

	filer := bson.D{{}}

	cursor, err := collection.Find(context.TODO(), filer)
	if err != nil{
		log.Println(err)
		return err
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil{
		log.Println(err)
		return err
	}

	return nil
}