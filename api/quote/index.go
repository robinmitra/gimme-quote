package handler

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// Uppercase fields since json package only accesses exported fields.
type quote struct {
	Id       string `json:"id"`
	Author   string `json:"author"`
	Quote    string `json:"quote"`
	Category string `json:"category"`
	Year     string `json:"year,omitempty"`
}

type errorResponse struct {
	Message string `json:"message"`
}

func getQuotesByCategory(category string) ([]quote, error) {
	var quotes []quote
	path := "quotes/" + category + ".json"
	if _, err := os.Stat(path); err != nil {
		return quotes, errors.New(fmt.Sprintln("Error while reading quote file:", err))
	}
	quotesJson, err := ioutil.ReadFile(path)
	if err != nil {
		return quotes, errors.New(fmt.Sprintln("Error while reading quote file:", err))
	}
	if err := json.Unmarshal(quotesJson, &quotes); err != nil {
		return quotes, errors.New(fmt.Sprintln("Error while un-marshalling JSON:", err))
	}
	return quotes, nil
}

func getQuotesByCategories(categories []string) ([]quote, error) {
	var quotes []quote
	for _, c := range categories {
		q, err := getQuotesByCategory(c)
		if err != nil {
			return quotes, err
		}
		quotes = append(quotes, q...)
	}
	return quotes, nil
}

func sendResponse(data interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	res, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func sendError(err error, w http.ResponseWriter) {
	e := errorResponse{Message: err.Error()}
	sendResponse(e, w)
	return
}

func Handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	category := query.Get("category")
	random, _ := strconv.ParseBool(query.Get("random"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	var quotes []quote
	if len(category) > 0 {
		c := strings.Split(category, ",")
		q, err := getQuotesByCategories(c)
		if err != nil {
			sendError(err, w)
			return
		}
		quotes = q
	} else {
		c := []string{"inspirational", "movie", "programming"}
		q, err := getQuotesByCategories(c)
		if err != nil {
			sendError(err, w)
			return
		}
		quotes = q
	}
	if random {
		rand.Shuffle(len(quotes), func(i, j int) { quotes[i], quotes[j] = quotes[j], quotes[i] })
	}
	if limit > 0 {
		quotes = quotes[0:limit]
	}
	sendResponse(quotes, w)
	return
}
