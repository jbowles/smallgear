//package smallgear
package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jbowles/nlpt_detect"
	"log"
	"net/http"
	"time"
)

const (
	BMSG = "Base message dude!"
)

type LangDetectResponse struct {
	Timestamp time.Time
	Detected1 string
	Code1     string
	Detected2 string
	Code2     string
	Detected3 string
	Code3     string
	Detected4 string
	Code4     string
	Input     string
}

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

func LanguageDetect(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	text := params["text"]
	detected_lang := nlpt_detect.Detect(text, "name", len(text), 3, 3, 3)
	second_rank_lang := nlpt_detect.Detect(text, "name", len(text), 2, 2, 2)
	third_rank_lang := nlpt_detect.Detect(text, "name", len(text), 1, 1, 1)
	four_rank_lang := nlpt_detect.Detect(text, "name", len(text), 0, 0, 0)

	detected_code := nlpt_detect.Detect(text, "code", len(text), 3, 3, 3)
	second_rank_code := nlpt_detect.Detect(text, "code", len(text), 2, 2, 2)
	third_rank_code := nlpt_detect.Detect(text, "code", len(text), 1, 1, 1)
	four_rank_code := nlpt_detect.Detect(text, "code", len(text), 0, 0, 0)

	//m := fmt.Sprintf("%v\nDetected1: %v \tCode: %v \nDetected2: %v\nDetected3: %v\n INPUT: \n%v", time.Now(), detected_lang, detected_code, second_rank_lang, third_rank_lang, text)
	//w.Write([]byte(m))

	langres := &LangDetectResponse{
		Timestamp: time.Now(),
		Detected1: detected_lang,
		Code1:     detected_code,
		Detected2: second_rank_lang,
		Code2:     second_rank_code,
		Detected3: third_rank_lang,
		Code3:     third_rank_code,
		Detected4: four_rank_lang,
		Code4:     four_rank_code,
		Input:     text,
	}

	response, err := json.Marshal(langres)
	if err != nil {
		log.Printf("Error", err)
	}
	w.Write(response)
}

func main() {
	WebServerBase()
}
