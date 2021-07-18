package main

import (
	"SAICCodingAssessment/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type algorithmInfo models.Info

func info(w http.ResponseWriter, r *http.Request) {
	//articles := Articles{
	//	Article{Title:"Test Title", Desc: "Test Description", Content: "Hello World"},
	//}

	algorithmInfo := models.Info{
		AlgorithmName: "Face Recognition API Implmentation",
		AlgorithmVersion: "1.0.1",
		AlgorithmType: "Face",
		CompanyName: "SAIC",
		TechnicalContactEmail: "areebkhalfay@gmail.com",
		RecommendedCPUs: 4,
		RecommendedMem: 2048,
	}
	fmt.Println("Algorithm Information Outputted")
	json.NewEncoder(w).Encode(algorithmInfo)
}

func testPostallArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test POST endpoint worked")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HomePage")
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/v1/info", info).Methods("GET")
	//myRouter.HandleFunc("/articles", testPostallArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	handleRequests()
}