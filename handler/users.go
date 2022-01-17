package handler

import (
	"net/http"
	"io/ioutil"

	"github.com/Prosp3r/company/logic"
)

func CreateStaff(w http.ResponseWriter, r *http.Request) {
	logTag := "Create Staff - handler"

	sentInfo, err := ioutil.ReadAll(r.Body)
	fe := logic.FailOnError()
	
}
func GetStaff(w http.ResponseWriter, r *http.Request)    {}
func UpdateStaff(w http.ResponseWriter, r *http.Request) {}
