package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"sample/sample"

	"github.com/gorilla/mux"
)



const data_folder = "data"

const data_file = "dataset3.json"


// import the type from the sample package
type MyData = sample.MyData

var dummy_data []MyData 

func set_dummy_dataset(dataSet string) {
	switch dataSet {
	case "1":
		dummy_data = sample.ExampleData1()
	case "2":
		dummy_data = sample.ExampleData2()
	case "3":
		dummy_data = sample.ExampleData3()
	default:
		panic(fmt.Sprintf(" dataset %s does not exist!", dataSet))
	}
}

type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (sr *statusRecorder) WriteHeader(statusCode int) {
	sr.statusCode = statusCode
	sr.ResponseWriter.WriteHeader(statusCode)
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


func local_data_loader(dataFile string) {
	
	jsonData, err := os.ReadFile(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("{\"error\": \"File %s not found\"}", dataFile)
		} else {
			fmt.Printf("{\"error\": \"Internal Server Error\"}")
		}
		return
	}

	var dat []MyData
	m_error := json.Unmarshal(jsonData, &dat)
	if m_error != nil {
		panic(m_error)
	} 
	dummy_data = dat
	fmt.Println(dummy_data)

}

func TEST_ZONE(w http.ResponseWriter, req *http.Request) {
	testFile := "data/dataset2.json"
	local_data_loader(testFile)
	//test_marshal()
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

func saveData(outPutFile string) {
	file, _ := os.OpenFile(outPutFile, os.O_CREATE, os.ModePerm)
	defer file.Close()
	encoder := json.NewEncoder(file) 
	encoder.Encode(dummy_data)
}

func t_save(w http.ResponseWriter, r *http.Request) {
	saveData("output.json")
}

func main() {
	
	//this loads data sets from sample.go
	//set_dummy_dataset("1")

	local_data_loader(data_folder +"/"+ data_file)

	fmt.Printf("Number of objects loaded: %d ", getLastItemID())

	r := mux.NewRouter()

	r.HandleFunc("/", handleGet).Methods("GET")

	// test new functions here
	r.HandleFunc("/test0", ID_tests).Methods("GET")
	r.HandleFunc("/test", TEST_ZONE).Methods("GET")
	
	//save dummy_data to json file
	r.HandleFunc("/save", t_save).Methods("GET")

	// handle POST GET DELETE PUT
	r.HandleFunc("/"+data_folder, handlePost).Methods("POST")
	r.HandleFunc("/"+data_folder, handleGet).Methods("GET")
	r.HandleFunc("/"+data_folder+"/{id}", handleDelete).Methods("DELETE")
	r.HandleFunc("/"+data_folder+"{id}", getPost).Methods("GET")
	r.HandleFunc("/"+data_folder+"{id}", updatePost).Methods("PUT")
   
	srv := &http.Server{
	 Addr:    "localhost:8080",
	 Handler: r,
	}
	srv.ListenAndServe()
   }