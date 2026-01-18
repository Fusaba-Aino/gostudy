package main

import (
	"log"

	github_com_golang_jwt_jwt "gostudy/jwt_study/github.com-golang-jwt-jwt"
	"gostudy/web_study/github.com-gin-gonic-gin"

	mydb "gostudy/db_study/modernc.org-sqlite"
)

func mydbtest() {
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

func mygintest() {
	github_com_gin_gonic_gin.Controller()
}
func main() {
	github_com_golang_jwt_jwt.TokenTest()
}
