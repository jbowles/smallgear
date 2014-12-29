package smallgear

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jbowles/wordlab"
	"log"
	"net/http"
	"strconv"
)

type KNNClassifierResponse struct {
	Classification  string
	ClassifierLabel int
	InputText       string
}

func KnnWildClassifyHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("knn wild text classifier"))
}

func KnnHotelErrorClassifyHandler(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	input := params["input"]
	bucket := wordlab.NewPredictionSentenceBucket(input, "bukt")
	attrs := bucket.BytePosSeqToFloat32()
	lookupID, err := strconv.Atoi(string(wordlab.AmitClassify(attrs)))
	if err != nil {
		log.Printf("Erro %v", err)
	}

	clsresp := &KNNClassifierResponse{
		Classification:  wordlab.HotelErrorIDTable[lookupID][0],
		ClassifierLabel: lookupID,
		InputText:       bucket.Sentence,
	}

	response, jsonerr := json.Marshal(clsresp)
	if jsonerr != nil {
		log.Printf("Error", jsonerr)
	}

	w.Write(response)
	//result := fmt.Sprintf("%v\n %v", wordlab.HotelErrorIDTable[lookupID][0], bucket.Sentence)
	//w.Write([]byte(result))
	//w.Write([]byte("knn wild text classifier"))
}
