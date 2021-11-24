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
	"auctionservice/auth"
	"auctionservice/auth/repository/authdatabase"

	auctionUsecase "auctionservice/auction/usecase"
	deliveryhttp "auctionservice/auth/delivery/http"
	authUsecase "auctionservice/auth/usecase"


	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type App struct {
	httpServer *http.Server
	authUC     auth.UseCase
	auctionUC  auction.UseCase
}

func newApp() *App {

	postgresDB := initPostgreDB()
	mongoDB := initMongoDB()

	authRepos := authdatabase.NewUserRepository(postgresDB)
	auctionRepos := auctiondatabase.NewAuctionRepository(postgresDB, mongoDB)

	return &App{
		authUC:    authUsecase.NewAuthUseCase(authRepos, "Pstre12e_9fQz", []byte("pwr12qxk90"), 10),
		auctionUC: auctionUsecase.NewAuctionUseCase(auctionRepos),
	}
}

func initPostgreDB() *sqlx.DB {
	db, err := sqlx.Connect("", "")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func initMongoDB() *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("")

	return db
}

func (a *App) Run() error {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	//needed to set up http handlers as endpoints
	deliveryhttp.RegisterAuthHTTPEndpoints(router, a.authUC)

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
