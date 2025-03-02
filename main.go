package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const data_folder = "data"

type MyData struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Passion string `json:"passion"`
}

var dummy_data []MyData

type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (sr *statusRecorder) WriteHeader(statusCode int) {
	sr.statusCode = statusCode
	sr.ResponseWriter.WriteHeader(statusCode)
}

func populate_sample_data() {
	dummy_data = append(dummy_data, MyData{ID:1, Name:"jimson",Passion:"food"})
	dummy_data = append(dummy_data, MyData{ID:2, Name:"janey",Passion:"gym"})
	dummy_data = append(dummy_data, MyData{ID:3, Name:"joey",Passion:"running"})
}


func handleGet(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dummy_data)
	sr := &statusRecorder{ResponseWriter: w, statusCode: http.StatusOK}
	logLine := fmt.Sprintf("%s %s %d", req.Method, req.RequestURI, sr.statusCode)
	log.Print(logLine)
}
   
func getLastItemID() int{ //returns my data object ID
	var output MyData
	if len(dummy_data)-1 >= 0 {
		output = dummy_data[len(dummy_data)-1]
	}

	return output.ID
}

func ID_tests(w http.ResponseWriter, req *http.Request) {
	
	t := []int{1,2,3,4,8,9}
	for _, tt := range t {
		fmt.Println(checkIDs(tt))
	}
}

func checkIDs(checkID int) bool {
	for _, item := range dummy_data {
		if item.ID == checkID {
			return true
		}
	}
	return false
}

func handleDelete(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
  	params := mux.Vars(req)
	myid_str := params["id"]
	myid, err := strconv.Atoi(myid_str)
	if err != nil {
		fmt.Printf("error with id! %s", myid_str)
	}
	if checkIDs(myid){
		for index, item := range dummy_data {
			if item.ID == myid {
			  dummy_data = append(dummy_data[:index], dummy_data[index+1:]...)
			  break
			}
		  }
		  json.NewEncoder(w).Encode(dummy_data)
		  sr := &statusRecorder{ResponseWriter: w, statusCode: http.StatusOK}
		  logLine := fmt.Sprintf("%s %s %d", req.Method, req.RequestURI, sr.statusCode)
		  log.Print(logLine)
	} else {
		fmt.Printf("id: %s not found!", myid_str)
	}
}

func handlePost(w http.ResponseWriter, req *http.Request) {
	var data MyData 
	//fmt.Println(req.Header)
	//fmt.Println(req.URL)
	_ = json.NewDecoder(req.Body).Decode(&data)
	if data.Name == "" || data.Passion == "" {
		fmt.Fprintf(w, "invalid post values")
		return
	}

	data.ID = getLastItemID() + 1
	output := []string{data.Name, data.Passion}
	//fmt.Fprintf(w, "%v, %v", output[0], output[1])
	dummy_data = append(dummy_data, data)
	json.NewEncoder(w).Encode(&data)

	sr := &statusRecorder{ResponseWriter: w, statusCode: http.StatusCreated}
	logLine := fmt.Sprintf("%s %s %s %d", req.Method, output, req.RequestURI, sr.statusCode)
	log.Print(logLine)
   }

func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	i, _ := strconv.Atoi(params["id"])
	for _, item := range dummy_data {
	  if item.ID == i {
		json.NewEncoder(w).Encode(item)
		return
	  }
	}
	json.NewEncoder(w).Encode(&MyData{})
}   

func updatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	i, _ := strconv.Atoi(params["id"])
	for index, item := range dummy_data {
		if item.ID == i {
			dummy_data = append(dummy_data[:index], dummy_data[index+1:]...)      
			var data MyData
			_ = json.NewDecoder(r.Body).Decode(&data)
			data.ID = i
			dummy_data = append(dummy_data, data)
			json.NewEncoder(w).Encode(&data)      
			return
		}
	}
	json.NewEncoder(w).Encode(dummy_data)

}

func main() {
	populate_sample_data()
	fmt.Println(getLastItemID())

	r := mux.NewRouter()
	r.HandleFunc("/data", handlePost).Methods("POST")
	r.HandleFunc("/data", handleGet).Methods("GET")
	r.HandleFunc("/", handleGet).Methods("GET")
	r.HandleFunc("/test", ID_tests).Methods("GET")
	r.HandleFunc("/data/{id}", handleDelete).Methods("DELETE")
	r.HandleFunc("/data/{id}", getPost).Methods("GET")
	r.HandleFunc("/data/{id}", updatePost).Methods("PUT")
   
	srv := &http.Server{
	 Addr:    "localhost:8080",
	 Handler: r,
	}
	srv.ListenAndServe()
   }