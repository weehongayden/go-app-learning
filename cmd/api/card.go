package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/shopspring/decimal"
	"github.com/weehongayden/bank-api/internal/query"
)

type CardRequest struct {
	Name          string          `json:"name"`
	StatementDate int             `json:"statement_date"`
	Amount        decimal.Decimal `json:"initial_amount"`
}

type CardResponse struct {
	ID            int             `json:"id"`
	Name          string          `json:"name"`
	StatementDate int             `json:"statement_date"`
	Amount        decimal.Decimal `json:"amount"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
}

func RegisterCardRoute(r chi.Router, args Route) {
	r.Route("/cards", func(card chi.Router) {
		card.Get("/", args.getCards)
		card.Get("/{id}", args.getCard)
		card.Post("/", args.createCard)
		card.Put("/{id}", args.updateCard)
		card.Delete("/{id}", args.deleteCard)
	})
}

func (app Route) getCards(w http.ResponseWriter, r *http.Request) {
	app.logger.Println("getCards() called.")

	row, err := app.db.Query(query.CardGetAll)
	if err != nil {
		app.logger.Printf("Failed to decode request body: %v\n", err)
		http.Error(w, "Unable to decode request body", http.StatusBadRequest)
		return
	}

	var cards []CardResponse
	for row.Next() {
		var card CardResponse
		if err := row.Scan(&card.ID, &card.Name, &card.Amount, &card.StatementDate, &card.CreatedAt, &card.UpdatedAt); err != nil {
			app.logger.Printf("Failed to retrieve record: %v\n", err)
		} else {
			cards = append(cards, card)
		}
	}

	resp := &Response{
		Status: http.StatusOK,
		Data:   cards,
	}

	jsonResponse, err := ResponseHandler(resp)
	if err != nil {
		app.logger.Printf("Unable to decode request body: %v\n", err)
		http.Error(w, "Unable to decode request body", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func (app Route) getCard(w http.ResponseWriter, r *http.Request) {
	app.logger.Println("getCard() called.")
	json.NewEncoder(w).Encode("Get Single Card")
}

func (app Route) createCard(w http.ResponseWriter, r *http.Request) {
	app.logger.Println("createCard() called.")

	var request CardRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		app.logger.Printf("Unable to decode request body: %v\n", err.Error())
		http.Error(w, "Unable to decode request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	row, err := app.db.Exec(query.CreateCard, request.Name, request.Amount, request.StatementDate, time.Now(), time.Now())
	if err != nil {
		app.logger.Printf("Failed to create record: %v\n", err.Error())
		http.Error(w, fmt.Sprintf("Failed to create record: %v\n", err), http.StatusBadRequest)
		return
	}

	_, err = row.LastInsertId()
	if err != nil {
		app.logger.Printf("Failed to create record: %v\n", err.Error())
		http.Error(w, fmt.Sprintf("Failed to create record: %v\n", err), http.StatusBadRequest)
		return
	}

	resp := &Response{
		Status: http.StatusCreated,
		Data:   "Record has been created.",
	}

	jsonResponse, err := ResponseHandler(resp)
	if err != nil {
		app.logger.Printf("Unable to decode request body: %v\n", err)
		http.Error(w, "Unable to decode request body", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}

func (app Route) updateCard(w http.ResponseWriter, r *http.Request) {
	app.logger.Println("updateCard() called.")
	json.NewEncoder(w).Encode("Update Card")
}

func (app Route) deleteCard(w http.ResponseWriter, r *http.Request) {
	app.logger.Println("deleteCard() called.")
	json.NewEncoder(w).Encode("Delete Card")
}
