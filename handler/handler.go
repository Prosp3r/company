package handler 

import(
	"net/http"
	"encoding/json"
)

//ResponseData model for displaying the status
type ResponseData struct {
	ErrorCode int         `json:"errorCode"`
	Data      interface{} `json:"data"`
}

//RespondJSON return the http response in json format
func ResponseJSON(w http.ResponseWriter, AppStatus int, HTTPStatus int, payload interface{}) {
	var res ResponseData
	res.ErrorCode = AppStatus
	res.Data = payload
	response, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(HTTPStatus)
	w.Write([]byte(response))
}

//RespondError return the http method error
func ResponseError(w http.ResponseWriter, AppStatus int, HTTPStatus int, message string) {
	var res ResponseData
	res.ErrorCode = AppStatus
	res.Data = message
	response, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(HTTPStatus)
	w.Write([]byte(response))
}