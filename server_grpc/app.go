package grpc_server

import (
	"auction_api/auction"
	"auction_api/auction/repository/auctiondatabase"
	"log"
	"net"
	"os"
	"fmt"
	"os/signal"
	"syscall"

	"auction_api/auction/delivery/grpc"
	"auction_api/auction/delivery/grpc/pb"

	auctionUsecase "auction_api/auction/usecase"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"

)

type GRPCApp struct {
	gRPCServer pb.AuctionServiceServer
	auctionUC  auction.UseCase
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



func Run()  {

	postgresDB := initPostgreDB()
	mongoDB := initMongoDB()

	auctionRepos := auctiondatabase.NewAuctionRepository(postgresDB, mongoDB)
	auction_uc := auctionUsecase.NewAuctionUseCase(auctionRepos)
	endpoints := auctiongrpc.MakeEndpoints(auction_uc)

	gRPCServer := auctiongrpc.NewGRPCServer(endpoints)
	

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	grpcListener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Println("during", "Listen", "err", err)
		os.Exit(1)
	}

	go func() {
		baseServer := grpc.NewServer()
		pb.RegisterAuctionServiceServer(baseServer, gRPCServer)
		log.Println("msg", "Server started successfully 🚀")
		baseServer.Serve(grpcListener)
	}()

	log.Println("exit", <-errs)

	
}
