package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"
)

type DataUSDBRL struct {
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
	} `json:"DataUSDBRL"`
}

type DolarBrl struct {
	Brl string
}

func main() {
	http.HandleFunc("/cotacao", BuscaDolarHandler)
	http.ListenAndServe(":8090", nil)
}

func BuscaDolarHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/cotacao" {
		http.Error(w, "404 Not Found.", http.StatusNotFound)
		return
	}

	dolarParam := r.URL.Query().Get("code")
	if dolarParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dolar, err := GetDolar()
	if err != nil {
		http.Error(w, "Timeout exceeded", http.StatusRequestTimeout)
		return
	}

	file, err := os.OpenFile("cotacao.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "404 Not Found.", http.StatusNotFound)
		return
	}
	defer file.Close()

	currentTime := time.Now()
	_, err = file.WriteString("Dolar: " + dolar + " create at " + currentTime.Format("02-01-2006 15:04:05") + ";" + "\n")
	if err != nil {
		http.Error(w, "404 Not Found.", http.StatusNotFound)
		return
	}

	w.Write([]byte("Dolar: " + dolar))
}

func GetDolar() (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx,"GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		return "reponse error", err
	}
	request.Header.Add("Accept", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return "error create header", err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "erro body", err
	}

	var data DataUSDBRL
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "erro desserialização", err
	}

	return data.USDBRL.Bid, nil
}
