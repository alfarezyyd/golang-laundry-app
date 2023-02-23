package helper

import (
	"encoding/json"
	"golang-laundry-app/model/web"
	"net/http"
)

func ReadFromRequestBody(r *http.Request, resultData interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(resultData)
	PanicIfError(err)
}

func WriteToResponseBody(w http.ResponseWriter, responseData web.WebResponse) {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(responseData)
	PanicIfError(err)
}
