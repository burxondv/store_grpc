package postgres

import (
	"database/sql"
	"fmt"
	"log"

	pb "store/proto"

	"github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", "localhost", 5432, "postgres", "bnnfav", "storedb")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	return db
}

func CreateStore(store *pb.Store) (*pb.Store, error) {
	db := ConnectDB()

	defer db.Close()

	newStore := pb.Store{}

	err := db.QueryRow("insert into stores(name, description, addresses, is_open) values($1, $2, $3, $4) returning id, name, description, addresses, is_open",
		store.Name, store.Description, pq.Array(store.Addresses), store.IsOpen).Scan(
		&newStore.Id,
		&newStore.Name,
		&newStore.Description,
		pq.Array(&newStore.Addresses),
		&newStore.IsOpen,
	)
	if err != nil {
		log.Fatalf("failed to insert store: %v", err)
	}

	return &newStore, nil
}

func GetStore(id int64) (*pb.Store, error) {
	db := ConnectDB()

	defer db.Close()

	newStore := pb.Store{}

	err := db.QueryRow("select id, name, description, addresses, is_open from stores where id = $1", id).Scan(
		&newStore.Id,
		&newStore.Name,
		&newStore.Description,
		pq.Array(&newStore.Addresses),
		&newStore.IsOpen,
	)
	if err != nil {
		log.Fatalf("failed to selecting store: %v", err)
	}

	return &newStore, nil
}

func UpdateStore(store *pb.Store) error {
	db := ConnectDB()

	defer db.Close()

	newStore := pb.Store{
		Id: 11,
		Name: "Ariana",
		Description: "22 years old",
		Addresses: []string{
			"Time square",
		},
		IsOpen: false,
	}

	res, err := db.Exec("update stores set name = $1, description = $2, addresses = $3, is_open = $4 where id = $5",
		newStore.Name, newStore.Description, pq.Array(newStore.Addresses), newStore.IsOpen, newStore.Id,
	)
	if err != nil {
		log.Fatalf("failed to update store in postgres: %v", err)
	}

	fmt.Println(res.RowsAffected())
	return nil
}

func DeleteStore(id int64) error {
	db := ConnectDB()
	
	defer db.Close()

	res, err := db.Exec("delete from stores where id = $1", id)
	if err != nil {
		log.Fatalf("failed to delete store in postgres: %v", err)
	}
	fmt.Println(res.RowsAffected())

	return nil
}
