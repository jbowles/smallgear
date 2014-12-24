package smallgear

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

const (
	BMSG = "Base message dude!"
)

func WebServerBase() {
	r := mux.NewRouter()
	// define root
	r.HandleFunc("/", RootHandler).Methods("GET")
	// Define subroutes
	languageDetection := r.PathPrefix("/language/detect").Subrouter()
	knnClassifier := r.PathPrefix("/knn/classify").Subrouter()
	bayesClassifier := r.PathPrefix("/bayes/classify").Subrouter()

	// define subroute function handlers
	languageDetection.HandleFunc("/{text}", LanguageDetectHandler)
	knnClassifier.HandleFunc("/wild/{input}", KnnWildClassifyHandler)
	bayesClassifier.HandleFunc("/wild/{input}", BayesWildClassifyHandler)
	//knnClassifier.HandleFunc("/error", KnnErrorClassifyHandler)
	//knnClassifier.HandleFunc("/policy/cancellation", KnnCancellationClassifyHandler)
	//bayesClassifier.HandleFunc("/error", BayesErrorClassifyHandler)
	//bayesClassifier.HandleFunc("/policy/cancellation", BayesCancellationClassifyHandler)
	//r.HandleFunc("/lang/kmeans/{chunks}", LanguageKmeans)
	http.Handle("/", r)
	log.Println("Listening... ")
	http.ListenAndServe(":3020", nil)
}

func RootHandler(w http.ResponseWriter, req *http.Request) {
	m := fmt.Sprintf("%v %v", time.Now(), BMSG)
	w.Write([]byte(m))
}
