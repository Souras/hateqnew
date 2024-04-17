package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Souras/hateqnew/service/hateqService/internal/db"
	"github.com/Souras/hateqnew/service/hateqService/internal/models"
	"github.com/gorilla/mux"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	u := r.URL
	query := u.Query()
	id := query.Get("date")

	products, err := db.GetProducts(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	response := models.ResponseData{Status: true, Data: products}
	jsonData, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error marshalling JSON response: %v", err)
		return
	}

	// w.WriteHeader(response.Status)
	// w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(jsonData)
	// json.NewEncoder(w).Encode(jsonData)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64) // Assuming ID is an int64
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid product ID"})
		return
	}
	product, err := db.GetProduct(id)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "Product not found"})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	json.NewEncoder(w).Encode(product)
}

func parseDatetime(dateTimeStr string) (time.Time, error) {
	// inputLayout := "MM/DD/YYYY HH:MM:SS" // Adjust based on your chosen input format
	layout := "01/02/2006 15:04:05"
	return time.Parse(layout, dateTimeStr)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	// inputLayout := "MM/DD/YYYY HH:MM:SS" // Adjust layout as needed for your specific input format

	// Parse the input time string
	var p models.QueueData

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	insertTime, err := parseDatetime(p.InsertTime)
	if err != nil {
		// Handle error (return bad request status code with error message)
		fmt.Println("Error:", err)
		return
	}

	startTime, err := parseDatetime(p.StartTime)
	if err != nil {
		// Handle error (return bad request status code with error message)
		fmt.Println("Error:", err)
		return
	}

	endTime, err := parseDatetime(p.EndTime)
	if err != nil {
		// Handle error (return bad request status code with error message)
		fmt.Println("Error:", err)
		return
	}

	createdProduct, err := db.CreateProduct(p, insertTime, startTime, endTime)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	json.NewEncoder(w).Encode(createdProduct)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid product ID"})
		return
	}

	var p models.QueueData
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid product data"})
		return
	}

	updatedProduct, err := db.UpdateProduct(id, p)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "Product not found"})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	json.NewEncoder(w).Encode(updatedProduct)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid product ID"})
		return
	}

	var p models.QueueData
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid product data"})
		return
	}

	err = db.DeleteProduct(id)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "Product not found"})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"msg": "Deleted Successfully"})
}

func GetLatestTokenIDByDate(w http.ResponseWriter, r *http.Request) {
	var currentTokenReq models.CurrentTokenReq
	err := json.NewDecoder(r.Body).Decode(&currentTokenReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	currentToken, err := db.GetCurrentTokenByAdminID(currentTokenReq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	// json.NewEncoder(w).Encode(currentToken)

	response := models.ResponseData{Status: true, Data: currentToken}
	jsonData, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error marshalling JSON response: %v", err)
		return
	}

	// w.WriteHeader(response.Status)
	// w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(jsonData)
}

func GetDocStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	adminID := vars["id"]

	currentToken, err := db.GetDocStatus(adminID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	// json.NewEncoder(w).Encode(currentToken)

	response := models.ResponseData{Status: true, Data: currentToken}
	jsonData, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error marshalling JSON response: %v", err)
		return
	}
	w.Write(jsonData)
}
