package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Prosp3r/company/model"
)

func CreateStaff(w http.ResponseWriter, r *http.Request) {
	logTag := "Create Staff - handler"

	sentInfo, err := ioutil.ReadAll(r.Body)
	fe := model.FailOnError(err, "Reading sent datta", logTag)
	if fe == true {
		ResponseJSON(w, http.StatusBadRequest, http.StatusBadRequest,
			fmt.Sprintf("%v Could not read sent information. Error: %v", logTag, err))
		return
	}

	// loadStaff := model.LoadAllStaff()
	_ = model.LoadAllStaff()
	// if loadStaff == false {
	// 	ResponseJSON(w, http.StatusInternalServerError, http.StatusInternalServerError,
	// 		fmt.Sprintf("%v Could not load user information. Error: %v", logTag, err))
	// 	return
	// }

	var addStaffInput model.AddStaffInput
	if err := json.Unmarshal(sentInfo, &addStaffInput); err != nil {
		ResponseJSON(w, http.StatusBadRequest, http.StatusBadRequest,
			fmt.Sprintf("%s Failed to convert sent user data. Error: %v", logTag, err))
	}

	//validate unique email
	if model.IsEmailUnique(addStaffInput.Email) == false {
		ResponseJSON(w, http.StatusBadRequest, http.StatusBadRequest,
			fmt.Sprintf("The Email address %v has been used before. Error: duplicate email", addStaffInput.Email))
		return
	}

	//validate unique Phone
	if model.IsPhoneUnique(addStaffInput.Phone) == false {
		ResponseJSON(w, http.StatusBadRequest, http.StatusBadRequest,
			fmt.Sprintf("The Phone number %v has been used before. Error: duplicate Phone number", addStaffInput.Phone))
		return
	}

	createStaff, err := model.CreateStaff(addStaffInput)
	fe = model.FailOnError(err, "Could not create staff", logTag)
	if fe == true {
		ResponseJSON(w, http.StatusInternalServerError, http.StatusInternalServerError, fmt.Sprintf("Could not add staff %v", err))
		return
	}

	var msg string
	if createStaff.ID > 0 {
		msg = "Adding new staff was successful."
	}
	ResponseJSON(w, http.StatusOK, http.StatusOK, msg)
	return
}

func GetStaff(w http.ResponseWriter, r *http.Request) {
	logTag := "Get Staff"
	loadStaf := model.LoadAllStaff()
	if loadStaf == false {
		ResponseJSON(w, http.StatusInternalServerError, http.StatusInternalServerError, fmt.Sprintf("Please try again later"))
		return
	}

	staffJson, err := json.Marshal(model.AllStaffList)
	_ = model.FailOnError(err, "Converting to json", logTag)

	ResponseJSON(w, http.StatusOK, http.StatusOK, staffJson)
	return 
}

func UpdateStaff(w http.ResponseWriter, r *http.Request) {}
