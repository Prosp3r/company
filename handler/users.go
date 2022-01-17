package handler

import (
	"io/ioutil"
	"net/http"

	"github.com/Prosp3r/company/models"
)

func CreateStaff(w http.ResponseWriter, r *http.Request) {
	logTag := "Create Staff - handler"

	sentInfo, err := ioutil.ReadAll(r.Body)
	fe := models.FailOnError(err, "Reading sent datta", logTag)
	

}
func GetStaff(w http.ResponseWriter, r *http.Request)    {}
func UpdateStaff(w http.ResponseWriter, r *http.Request) {}
