package auctiondatabase

import (
	"auction_api/models"
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

)

type AuctionRepository struct {
	postgresdb *sqlx.DB
	mongodb    *mongo.Database
}

func NewAuctionRepository(postgres *sqlx.DB, mongo *mongo.Database) *AuctionRepository {
	return &AuctionRepository{
		//postgres-db
		postgresdb: postgres,
		//mongo-db
		mongodb: mongo,
		//redis
	}
}

func (a *AuctionRepository) NewAuction(auctionTitle, auctionType, auctionStatus, auctionDate string) (string, error) {
	auction_id := uuid.New().String()
	_, err := a.postgresdb.Exec("insert into auctions (id, title, type, status, date) values ($1, $2, $3, $4, $5)", 
	auction_id, auctionTitle, auctionType, auctionStatus, auctionDate)
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
	err := a.postgresdb.Get(&auction_data, "select * from auctions where id=$1", auction_id)
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

func (a *AuctionRepository) GetAllPublicAuctions() (*[]models.Auction, error) {
	auctions := []models.Auction{}
	auction := models.Auction{}
	rows, err := a.postgresdb.Query("select * from auctions where type='public'")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&auction)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		auctions = append(auctions, auction)
	}

	return &auctions, nil
}
