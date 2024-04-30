package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
)

func handleLogin(ctx context.Context, userStore *UserStore) http.Handler {
	type LoginPayload struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			payload, err := decode[LoginPayload](r)
			if err != nil {
				encode(w, r, http.StatusBadRequest, struct {
					Message string `json:"message"`
				}{"bad request"})
				return
			}
			user, err := userStore.FindUserByUsername(ctx, payload.Username)
			if err != nil {
				encode(w, r, http.StatusNotFound, struct {
					Message string `json:"message"`
				}{"user does not exist"}) // Improve error handling here
				return
			}
			if !CheckPasswordHash(payload.Password, user.Password) {
				encode(w, r, http.StatusNotFound, struct {
					Message string `json:"message"`
				}{"invalid username or password"}) // Improve error handling here
				return
			}
			tokenString, err := CreateToken(user.Id.String(), user.Username, getenv)
			if err != nil {
				encode(w, r, http.StatusInternalServerError, struct {
					Message string `json:"message"`
				}{"failed to create token"}) // Improve error handling here
				return
			}
			encode(w, r, http.StatusOK, struct {
				Token string `json:"token"`
			}{tokenString})
		},
	)
}

func handleFindUser(ctx context.Context, userStore *UserStore) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			id := r.PathValue("id")
			user, err := userStore.FindUserById(ctx, id)
			if err != nil {
				switch err {
				case sql.ErrNoRows:
					// log raw error to logs
					encode(w, r, http.StatusNotFound, struct {
						Message string `json:"message"`
					}{"user does not exist"}) // Improve error handling here
					return
				default:
					// log raw error to logs
					encode(w, r, http.StatusInternalServerError, struct {
						Message string `json:"message"`
					}{"internal server error"}) // Improve error handling here
					return

				}
			}
			if user.Username == "" {
				encode(w, r, http.StatusNotFound, struct {
					Message string `json:"message"`
				}{"user does not exist"}) // Improve error handling
				return
			}
			encode(w, r, http.StatusOK, user)
		},
	)
}

func handleCreateUserAccount(ctx context.Context, userStore *UserStore) http.Handler {
	type NewUser struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			usr, err := decode[NewUser](r)
			if err != nil {
				encode(w, r, http.StatusBadRequest, err) // fix this raw error to return sanirised message
				return
			}
			hash, err := HashPassword(usr.Password)
			if err != nil {
				encode(w, r, http.StatusInternalServerError, err) // fix this raw error to return sanitised message
				return
			}
			user := User{Username: usr.Username, Email: usr.Email, Password: hash}
			if err = userStore.CreateUser(ctx, &user); err != nil {
				// Fix this for duplicate key error
				encode(w, r, http.StatusInternalServerError, err) // fix this raw error to return sanitised message
				return
			}
			fmt.Printf("username: %s, created: %v\n", user.Username, user.DateCreated)
			encode(w, r, http.StatusOK, user)
		},
	)
}
