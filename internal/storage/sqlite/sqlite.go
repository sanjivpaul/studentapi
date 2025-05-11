package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sanjivpaul/studentapi/internal/config"
)

type Sqlite struct{
	Db *sql.DB
}

// jo bhi conventions bnate hai usko New name dete
// receive config here
// * => receive only pointer (just take refrence)
func New(cfg *config.Config)(*Sqlite, error){
	// it return db instance and error
	db, err := sql.Open("sqlite3", cfg.StoragePath)

	if err != nil{
		return nil, err
	
	}

	// create table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT,
		age INTEGER
	)`)

	if err != nil{
		return nil, err
	}

	return &Sqlite{
		Db:db,
	}, nil

}