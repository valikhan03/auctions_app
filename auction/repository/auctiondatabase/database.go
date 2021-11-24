package auctiondatabase

import (
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuctionRepository struct {
	postgresdb *sqlx.DB
	mongodb    *mongo.Database
}

func NewAuctionRepository(postgres *sqlx.DB, mongodb *mongo.Database) *AuctionRepository {
	return &AuctionRepository{
		//postgres-db
		postgresdb: postgres,
		//mongo-db
		mongodb: mongodb,
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

	filter := bson.D{{"auction_id", auction_id}}

	update := bson.D{
		{"$push", bson.D{
			{"participants", user_id},
		}},
	}

	updRes, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(updRes)
	return err
}

func (a *AuctionRepository) GetAuction() {

}
