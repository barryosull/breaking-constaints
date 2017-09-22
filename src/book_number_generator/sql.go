package book_number_generator

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"github.com/barryosull/breaking-constaints/src"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

type Sql struct {
	db *sql.DB
}

func Make() *Sql {

	username := os.Getenv("DB_USERNAME");
	password := os.Getenv("DB_PASSWORD");
	database := os.Getenv("DB_DATABASE");
	host := os.Getenv("DB_HOST");

	log.Fatal(username+":"+password+"@tcp("+host+":3306)/"+database)

	db, err := sql.Open("mysql", username+":"+password+"@tcp("+host+":3306)/"+database)
	if err != nil {
		log.Fatal(err)
	}

	return &Sql{db}
}

func (s *Sql) GenerateNumber(categoryId int) int {

	var number int;

	err := s.db.QueryRow("SELECT number FROM booking_numbers WHERE category_id = ? ORDER BY number DESC", categoryId).Scan(&number)

	switch {
		case err == sql.ErrNoRows:
			return 1
		case err != nil:
			log.Fatal(err)
	}

	return number+1;
}

func (s *Sql) NumberUsed(categoryId int, number int) error {

	stmt, err := s.db.Prepare(`
		INSERT INTO booking_numbers(category_id, number) VALUES(?, ?)
	`)

	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(categoryId, number)


	if err != nil {
		return &src.BookNumberNotUnique{}
	}

	return nil
}

func (s *Sql) Reset() {

	_, err := s.db.Exec("DELETE FROM booking_numbers")

	if err != nil {
		log.Fatal(err)
	}
}