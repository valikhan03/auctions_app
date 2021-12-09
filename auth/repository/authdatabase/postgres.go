package authdatabase

import (
	"auctionservice/models"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	database *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		database: db,
	}
}

func (rep *UserRepository) CreateUser(user *models.User) error {
	_, err := rep.database.Exec("insert into users (id, email, firstname, lastname, password) values ($1, $2, $3, $4, $5)",
		user.Id, user.Email, user.Firstname, user.Lastname, user.Password)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (rep *UserRepository) GetUserID(email, password string) (string, error) {
	fmt.Println(email, password)
	var userid string
	err := rep.database.Get(&userid, "select id from users where email=$1 and password=$2 LIMIT 1", email, password)
	
	if err != nil {
		log.Println(err)
	}
	return userid, err
}
