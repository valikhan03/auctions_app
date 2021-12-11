package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"auctionservice/auction"
	"auctionservice/auction/repository/auctiondatabase"


	auctionhttp "auctionservice/auction/delivery/http"
	auctionUsecase "auctionservice/auction/usecase"


	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	_ "github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
)

type App struct {
	httpServer *http.Server
	auctionUC  auction.UseCase
}

func NewApp() *App {

	postgresDB := initPostgreDB()
	mongoDB := initMongoDB()
	redisDB := initRedisDB()

	auctionRepos := auctiondatabase.NewAuctionRepository(postgresDB, mongoDB, redisDB)

	return &App{
		auctionUC: auctionUsecase.NewAuctionUseCase(auctionRepos),
	}
}

func initPostgreDB() *sqlx.DB {
	db, err := sqlx.Connect("pgx", ReadPostgresConfigs())
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func initMongoDB() *mongo.Database {
	mongo, err := mongo.NewClient(options.Client().ApplyURI(ReadMongoConfigs()))
	if err != nil {
		log.Fatal(err)
	}
	db := mongo.Database("")

	return db
}

func initRedisDB() *redis.Client {
	db := redis.NewClient(&redis.Options{
		Addr:     "",
		Password: "",
		DB:       0,
	})

	return db
}

func (a *App) Run() error {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	//needed to set up http handlers as endpoints


	auctionhttp.RegisterAuctionHttpEndpoints(router, a.auctionUC)

	a.httpServer = &http.Server{
		Addr:           ":8090",
		Handler:        router,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Server failed : %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}
