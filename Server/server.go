package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/google/uuid"
)

type Data struct {
	USDBRL struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

type AmericanDolarBrl struct {
	id       string
	dolarBrl string
}

func NewDolarBrl(id string, dolarBrl string) *AmericanDolarBrl {
	return &AmericanDolarBrl{id: uuid.New().String(), dolarBrl: dolarBrl}
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/cotacao", dolarBrl{})
	http.ListenAndServe(":8080", mux)
}

type dolarBrl struct {
}

func (d dolarBrl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./data/dolar_brl.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if r.URL.Path != "/cotacao" {
		http.Error(w, "404 Not Found.", http.StatusNotFound)
		return
	}

	ctx := r.Context()

	dolar, err := GetDolar(ctx)
	if err != nil {
		http.Error(w, "404 Not Found.", http.StatusNotFound)
		return
	}

	err = InsertDolarBrl(ctx, db, dolar.USDBRL.Bid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(dolar)
}

func GetDolar(ctx context.Context) (*Data, error) {

	select {
	case <-time.After(200 * time.Millisecond):
		return nil, ctx.Err()

	default:
		request, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/last/USD-BRL", nil)
		if err != nil {
			return nil, err
		}
		request.Header.Add("Accept", "application/json")

		response, err := http.DefaultClient.Do(request)
		if err != nil {
			return nil, err
		}
		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		var data Data
		err = json.Unmarshal(body, &data)
		if err != nil {
			return nil, err
		}

		return &data, nil
	}
}

func InsertDolarBrl(ctx context.Context,db *sql.DB, dolar string) error {
	select {
	case <-time.After(10 * time.Millisecond):
		return ctx.Err()

	default:
		stmt, err := db.Prepare("INSERT INTO dolar_brl (id, price, create_at) VALUES (?, ?, ?)")
		if err != nil {
			return err
		}
		defer stmt.Close()

		// Get current time in Brazil timezone
		loc, _ := time.LoadLocation("America/Sao_Paulo")
		now := time.Now().In(loc)

		_, err = stmt.Exec(uuid.New().String(), dolar, now.Format("2006-01-02 15:04:05"))
		if err != nil {
			return err
		}

		return nil
	}
}
