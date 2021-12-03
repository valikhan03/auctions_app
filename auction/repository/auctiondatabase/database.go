package auctiondatabase

import (
	"auctionservice/models"
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/go-redis/redis/v8"
)

type AuctionRepository struct {
	postgresdb *sqlx.DB
	mongodb    *mongo.Database
	redisdb    *redis.Client
}

func NewAuctionRepository(postgres *sqlx.DB, mongo *mongo.Database, redis *redis.Client) *AuctionRepository {
	return &AuctionRepository{
		//postgres-db
		postgresdb: postgres,
		//mongo-db
		mongodb: mongo,
		//redis
		redisdb: redis,
	}
}

func (a *AuctionRepository) NewAuction(user_id, auction_title string) (string, error) {
	auction_id := "gen-auction-id"
	_, err := a.postgresdb.Exec("insert into auctions (id, organizator_id, title) values ($1, $2, $3)", auction_id, user_id, auction_title)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return auction_id, err
}

func (a *AuctionRepository) AddParticipant(auction_id string, user_id string) error {
	collection := a.mongodb.Collection("")

	filter := bson.D{primitive.E{Key: "auction_id", Value: auction_id}}

	update := bson.D{
		{"$push", bson.D{
			primitive.E{Key: "participants", Value: user_id},
		}},
	}

	updRes, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(updRes)
	return err
}

func (a *AuctionRepository) GetAuctionData(auction_id string) (models.Auction, error) {
	auction_data := models.Auction{}
	err := a.postgresdb.Get(&auction_data, "select * from auction where auction_id=$1", auction_id)
	if err != nil {
		log.Println(err)
	}

	return auction_data, err
}

func (a *AuctionRepository) GetAuctionParticipants(auction_id string) ([]string, error) {
	collection := a.mongodb.Collection("")
	var participants []string

	filter := bson.D{{"auction_id", auction_id}}

	err := collection.FindOne(context.TODO(), filter).Decode(&participants)
	if err != nil {
		log.Println(err)
	}

	return participants, err
}
