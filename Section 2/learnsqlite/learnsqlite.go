package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type animal struct {
	id         int
	animalType string
	nickname   string
	zone       int
	age        int
}

func main() {
	log.Println("Connect to database...")
	//connect to the database
	db, err := sql.Open("sqlite3", "dino.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = createdatabase(db)
	if err != nil {
		log.Fatal(err)
	}
	err = insertsampledata(db)
	if err != nil {
		log.Fatal(err)
	}

	//general query with arguments
	rows, err := db.Query("select * from animals where age > $1", 5) //both $ and ? are supported
	handlerows(rows, err)

	//query a single row
	row := db.QueryRow("select * from animals where age > $1", 5)
	a := animal{}
	err = row.Scan(&a.id, &a.animalType, &a.nickname, &a.zone, &a.age)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a)
	/*
			//insert a row
			result, err := db.Exec("Insert into animals (animal_type,nickname,zone,age) values ('Carnotaurus', 'Carno', $1, $2)", 3, 22)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(result.LastInsertId())
			fmt.Println(result.RowsAffected())


		//update a row

		res, err := db.Exec("Update animals set age = $1 where id = $2", 16, 2)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res.LastInsertId())
		fmt.Println(res.RowsAffected())
	*/
	/*
		var id int
		db.QueryRow("Update animals set age = $1 where id = $2 returning id", 16, 2).Scan(&id)
		fmt.Println("id returned:", id)
	*/
	/*
		//prepare queries to use them multiple times, this also improves performance because
		fmt.Println("Prepared statements... ")
		stmt, err := db.Prepare(" select * from animals where age > $1 ")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()
		//let's try with age>5
		rows, err = stmt.Query(5)
		handlerows(rows, err)

		//let's try with age>10
		rows, err = stmt.Query(10)
		handlerows(rows, err)

		testTransaction(db)
	*/
}

func handlerows(rows *sql.Rows, err error) {
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	animals := []animal{}
	for rows.Next() {
		a := animal{}
		err := rows.Scan(&a.id, &a.animalType, &a.nickname, &a.zone, &a.age)
		if err != nil {
			log.Println(err)
			continue
		}
		animals = append(animals, a)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(animals)
}

func testTransaction(db *sql.DB) {
	fmt.Println("Transactions...")
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("select * from animals where age > $1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(15)
	handlerows(rows, err)
	rows, err = stmt.Query(16)
	handlerows(rows, err)
	results, err := tx.Exec("Update animals set age = $1 where id = $2", 17, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(results.RowsAffected())
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func createdatabase(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS
						animals(id INTEGER PRIMARY KEY AUTOINCREMENT, 
                        animal_type TEXT,
                        nickname TEXT,
                        zone INTEGER, 
                        age INTEGER) `)
	return err
}

func insertsampledata(db *sql.DB) error {
	_, err := db.Exec(`insert INTO animals (animal_type,nickname,zone,age) 
                        VALUES ('Tyrannosaurus rex','rex', 1, 10),
                        ('Velociraptor', 'rapto', 2, 15)`)
	return err
}
