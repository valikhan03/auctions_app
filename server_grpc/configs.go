package grpc_server

import (
	"io/ioutil"
	"log"
	"os"
	"fmt"

	"gopkg.in/yaml.v2"
	"github.com/joho/godotenv"
)

type MongoConfs struct {
	DB_URI string `yaml:"uri"`
}

func ReadMongoConfigs() string{
	var configs MongoConfs

	var confsFile, err = os.Open("configs/mongo.yaml")
	if err != nil{
		log.Fatal(err)
	}

	defer func(){
		err = confsFile.Close()
		if err != nil{
			log.Fatal(err)
		}
	}()

	configData, err := ioutil.ReadAll(confsFile)
	if err != nil{
		log.Fatal(err)
	}

	err = yaml.Unmarshal(configData, &configs)
	if err != nil{
		log.Fatal(err)
	}

	return configs.DB_URI
}

type PostgresConfigs struct{
	Host    string `yaml:"host"`
	Port    string `yaml:"port"`
	User    string `yaml:"user"`
	DBName  string `yaml:"dbname"`
	SSLMode string `yaml:"sslmode"`
}


func ReadPostgresConfigs() string {
	var confs PostgresConfigs

	file, err := os.Open("configs/postgres.yaml")
	if err != nil {
		log.Fatal(err)
	}

	byteConfigData, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(byteConfigData, &confs)
	if err != nil {
		log.Fatal(err)
	}
	
	err = godotenv.Load()
	if err != nil{
		log.Fatal(err)
	}

	key := os.Getenv("POSTGRES_SECRET_KEY")

	conn_str := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s", confs.Host, confs.Port, confs.User,  confs.DBName, confs.SSLMode, key)

	return conn_str
}
