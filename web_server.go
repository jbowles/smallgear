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
	r.HandleFunc("/", RootHandler).Methods("GET")
	r.HandleFunc("/lang/detect/{text}", LanguageDetect).Methods("GET")
	//r.HandleFunc("/lang/classify", LanguageClassify)
	//r.HandleFunc("/lang/kmeans/{chunks}", LanguageKmeans)
	http.Handle("/", r)
	log.Println("Listening... ")
	http.ListenAndServe(":3020", nil)
}

func RootHandler(w http.ResponseWriter, req *http.Request) {
	m := fmt.Sprintf("%v %v", time.Now(), BMSG)
	w.Write([]byte(m))
}
