package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CreateJsonResponse(w http.ResponseWriter, data interface{}) []byte {
	jsonString, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	return jsonString
}

func WriteJsonResponse(w http.ResponseWriter, data interface{}) {
	jsonString := CreateJsonResponse(w, data)

	w.Header().Set("content-type", "application/json")

	_, err := w.Write(jsonString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ParseRequestBody(w http.ResponseWriter, r *http.Request, data interface{}) {
	b, err := ioutil.ReadAll(r.Body)
	defer func() {
		err = r.Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(b, &data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
