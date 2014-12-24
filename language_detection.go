package smallgear

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jbowles/nlpt_detect"
	"log"
	"net/http"
	"time"
)

type LangDetectResponse struct {
	Timestamp      time.Time
	ChoiceLanguage string
	ChoiceCode     string
	Detected2      string
	Code2          string
	Detected3      string
	Code3          string
	Detected4      string
	Code4          string
	Input          string
}

func LanguageDetect(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	text := params["text"]
	detected_lang := nlpt_detect.Detect(text, "name", len(text), 3, 3, 3)
	second_rank_lang := nlpt_detect.Detect(text, "name", len(text), 3, 0, 0)
	third_rank_lang := nlpt_detect.Detect(text, "name", len(text), 0, 3, 0)
	four_rank_lang := nlpt_detect.Detect(text, "name", len(text), 0, 0, 3)

	detected_code := nlpt_detect.Detect(text, "code", len(text), 3, 3, 3)
	second_rank_code := nlpt_detect.Detect(text, "code", len(text), 3, 0, 0)
	third_rank_code := nlpt_detect.Detect(text, "code", len(text), 0, 3, 0)
	four_rank_code := nlpt_detect.Detect(text, "code", len(text), 0, 0, 3)

	langres := &LangDetectResponse{
		Timestamp:      time.Now(),
		ChoiceLanguage: detected_lang,
		ChoiceCode:     detected_code,
		Detected2:      second_rank_lang,
		Code2:          second_rank_code,
		Detected3:      third_rank_lang,
		Code3:          third_rank_code,
		Detected4:      four_rank_lang,
		Code4:          four_rank_code,
		Input:          text,
	}

	response, err := json.Marshal(langres)
	if err != nil {
		log.Printf("Error", err)
	}
	w.Write(response)
}
