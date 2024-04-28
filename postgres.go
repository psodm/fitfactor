package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

var counts int64

func Connect(dsn string, maxOpenConnections, maxIdleConnections, maxIdleTime int) (*sql.DB, error) {
	for {
		connection, err := openDB(dsn, maxOpenConnections, maxIdleConnections, maxIdleTime)
		if err != nil {
			counts++
		} else {
			return connection, nil
		}
		if counts > 10 {
			return nil, err
		}
		time.Sleep(2 * time.Second)
	}
}

func openDB(dsn string, maxOpenConnections, maxIdleConnections, maxIdleTime int) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(maxOpenConnections)
	db.SetMaxIdleConns(maxIdleConnections)
	db.SetConnMaxIdleTime(time.Duration(maxIdleTime) * time.Minute)
	if err := db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("Pinged")
	return db, nil
}

type UserStore struct{ db *sql.DB }

func (u *UserStore) FindUserById(ctx context.Context, id string) (*User, error) {
	query := "SELECT * FROM user_account WHERE id=$1"
	row, err := u.db.QueryContext(ctx, query, id)
	if err != nil {
		fmt.Printf("error 1: %s\n", err)
		return nil, err
	}
	var user User
	for row.Next() {
		err = row.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.DateCreated, &user.LastUpdated)
		if err != nil {
			fmt.Printf("error 2: %s\n", err)
			return nil, err
		}
	}
	return &user, nil
}

func (u *UserStore) FindUserByUsername(ctx context.Context, username string) (*User, error) {
	query := "SELECT * FROM user_account WHERE username=$1"
	row, err := u.db.QueryContext(ctx, query, username)
	if err != nil {
		fmt.Printf("error 1: %s\n", err)
		return nil, err
	}
	var user User
	for row.Next() {
		err = row.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.DateCreated, &user.LastUpdated)
		if err != nil {
			fmt.Printf("error 2: %s\n", err)
			return nil, err
		}
	}
	return &user, nil
}

func (u *UserStore) CreateUser(ctx context.Context, user *User) error {
	query := "INSERT INTO user_account (username, email, password) VALUES ($1, $2, $3) RETURNING id, date_created, last_updated"
	hash, err := HashPassword(user.Password)
	if err != nil {
		return err
	}
	err = u.db.QueryRowContext(ctx, query, user.Username, user.Email, hash).Scan(&user.Id, &user.DateCreated, &user.LastUpdated)
	if err != nil {
		fmt.Printf("err1: %v\n", err)
		return err
	}
	return nil
}
