package main

import (
	"SAICCodingAssessment/models"
	"encoding/json"
	"fmt"
	"github.com/DataDog/go-python3"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)
type Score struct{
	Score float64
	NormalizedScore float64
}

func info(w http.ResponseWriter, r *http.Request) {
	var algorithmInfo models.Info
	algorithmInfo = models.Info{
		AlgorithmName:         "Face Recognition API Implmentation",
		AlgorithmVersion:      "1.0.1",
		AlgorithmType:         "Face",
		CompanyName:           "SAIC",
		TechnicalContactEmail: "areebkhalfay@gmail.com",
		RecommendedCPUs:       4,
		RecommendedMem:        2048,
	}
	err := json.NewEncoder(w).Encode(algorithmInfo)
	if err != nil {
		http.Error(w, "The internal license has expired.", http.StatusInternalServerError)
		fmt.Println("500 Internal Server Error")
	}
	fmt.Println("Algorithm Information Outputted")
}

func compareList(w http.ResponseWriter, r *http.Request) {
	jsonDecoder := json.NewDecoder(r.Body)
	var compareListRequest models.CompareListRequest
	err := jsonDecoder.Decode(&compareListRequest)
	if err != nil {
		http.Error(w, "Unable to decode image data as a PNG.", http.StatusBadRequest)
		fmt.Println("400 Bad Request")
	}
	singleTemplate := compareListRequest.SingleTemplate.Imagedata
	var templateList []models.Image = compareListRequest.TemplateList
	singleTemplateString := python3.PyUnicode_FromString(singleTemplate)

	//Code Refactored from https://github.com/christian-korneck/python-go/blob/master/py-bindings/outliers/main.go
	defer python3.Py_Finalize()
	python3.Py_Initialize()
	if !python3.Py_IsInitialized() {
		fmt.Println("Error initializing the python interpreter")
		os.Exit(1)
	}
	//Path setting does not work as intended. If functional, API would be in order.
	//1st method for setting Path
	//dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// we could also use PySys_GetObject("path") + PySys_SetPath,
	////but this is easier (at the cost of less flexible error handling)
	//ret := python3.PyRun_SimpleString("import sys\nsys.path.append(\"" + dir + "\")")
	//if ret != 0 {
	//	log.Fatalf("error appending '%s' to python sys.path", dir)
	//}

	//2nd Method for setting Path
	//python3.PySys_GetObject("path")
	//Would need to change to directory on home machine
	//err = python3.PySys_SetPath("/home/areebk/go/src/SAICCodingAssessment")
	//if err != nil {
	//	return
	//}

	//Ideally, both should work. However, when running either, separate issues occur. If this problem were to be fixed,
	//the api would be up and running and working properly.
	pythonFilesImport := python3.PyImport_ImportModule("python_files") //ret val: new ref
	if !(pythonFilesImport != nil && python3.PyErr_Occurred() == nil) {
		python3.PyErr_Print()
		log.Fatal("failed to import module 'python_files'")
	}

	defer pythonFilesImport.DecRef()

	pythonFilesModule := python3.PyImport_AddModule("python_files") //ret val: borrowed ref (from oImport)
	if !(pythonFilesModule != nil && python3.PyErr_Occurred() == nil) {
		python3.PyErr_Print()
		log.Fatal("failed to add module 'python_files'")
	}
	pythonFilesDict := python3.PyModule_GetDict(pythonFilesModule)
	match_images := python3.PyDict_GetItemString(pythonFilesDict, "match_images")
	var scoreList []float64
	for i := 0; i < len(templateList); i++ {
		var listTemplateString *python3.PyObject = python3.PyUnicode_FromString(templateList[i].Imagedata)
		args := python3.PyTuple_New(2) //retval: New reference
		if args == nil {
			listTemplateString.DecRef()
			return
		}
		ret := python3.PyTuple_SetItem(args, 0, singleTemplateString)
		ret1 := python3.PyTuple_SetItem(args, 1, listTemplateString)
		if ret != 0 {
			if python3.PyErr_Occurred() != nil {
				python3.PyErr_Print()
			}
			singleTemplateString.DecRef()
		}
		if ret1 != 0 {
			if python3.PyErr_Occurred() != nil {
				python3.PyErr_Print()
			}
			listTemplateString.DecRef()
		}
		imagesDataPy := match_images.CallObject(args)
		defer imagesDataPy.DecRef()
		scoreList = append(scoreList, python3.PyFloat_AsDouble(imagesDataPy))
	}
	var scores []Score
	for i := 0; i < len(templateList); i++ {
		scores = append(scores, Score{Score: scoreList[i] * 10000, NormalizedScore: scoreList[i]})
	}
	err = json.NewEncoder(w).Encode(scores)
	if err != nil {
		http.Error(w, "The internal license has expired.", http.StatusInternalServerError)
		fmt.Println("500 Internal Server Error")
	}
	fmt.Println("Lists Compared")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Biometric Matching Service")
}

func handleRequests() {
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/v1/info", info).Methods("GET")
	myRouter.HandleFunc("/v1/compare-list", compareList).Methods("POST")
	http.ListenAndServe(":8080", myRouter)
}

func main() {
	handleRequests()
}