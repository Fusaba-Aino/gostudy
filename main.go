package main

import (
	"log"

	mydb "gostudy/db_study/modernc.org-sqlite"
)

func main() {
	db, err := mydb.Connect()
	if err != nil {
		log.Fatalf("connect: %v", err)
	}
	defer mydb.CloseDB(db)

	err = mydb.CreateTable(db)
	if err != nil {
		return
	}

	err = mydb.InsertUser(db, "test user", 18)
	if err != nil {
		return
	}
}
