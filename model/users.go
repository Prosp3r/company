package logic

import (
	"errors"
	"fmt"
)

type Staff struct {
	ID        int64  `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Entrytime int64  `json:"entrytime,omitempty"`
}

type AddStaffInput struct {
	Phone string `json:"phone,omitempty"`
	Email string `json:"email"`
}

var (
	AllStaffList []Staff
	AllStaffMap  = make(map[string]Staff)
)

//CreateUser -
func CreateStaff(si AddStaffInput) (*Staff, error) {
	logTag := "Creating New Staff"

	FakeName, err := GetPetName(1, "")
	_ = FailOnError(err, "generaing random name", logTag)

	var nU Staff

	nU.Name = FakeName
	nU.Email = si.Email
	nU.Phone = si.Phone

	db := DbConn()
	defer db.Close()
	PQ, err := db.Prepare("INSERT INTO staff(name, email, phone, entrytime) VALUES(?, ?, ?, ?)")
	fe := FailOnError(err, "Preparing db insert", logTag)
	if fe == true {
		return nil, err
	}
	ins, err := PQ.Exec(nU.Name, nU.Email, nU.Phone, nU.Entrytime)
	fe = FailOnError(err, "Executing statement", logTag)
	if fe == true {
		return nil, err
	}

	lastInsertId, err := ins.LastInsertId()
	fe = FailOnError(err, "Executing LastInsertId", logTag)
	if fe == true {
		return nil, err
	}

	nU.ID = lastInsertId
	return &nU, nil
}

//GetAllUsers - Populate the AllUsers slicee for quick manipulative access
func LoadAllStaff() bool {
	logTag := "LoadAllUsers"

	var u Staff

	db := DbConn()
	defer db.Close()

	PQ, err := db.Query("SELECT id, name, email, phone, entrytime FROM staff ORDER BY id DESC")
	em := FailOnError(err, "Preparing verifications Query", logTag)
	if em == true {
		return false
	}

	for PQ.Next() {
		err = PQ.Scan(&u.ID, &u.Name, &u.Email, &u.Phone, &u.Entrytime)
		em = FailOnError(err, "Reading user list from select", logTag)
		if em == true {
			return false
		}

		AllStaffMap[u.Phone] = u
		AllStaffList = append(AllStaffList, u)
	}
	return true
}

//isEmailUnique - checks if email has not been used previously.
func IsEmailUnique(email string) bool {
	//LoadAllUsers()
	for _, v := range AllStaffList {
		if v.Email == email {
			return false
		}
	}
	return true
}

//isPhoneUnique - checks if a phone number has not been used peviously.
func IsPhoneUnique(phone string) bool {
	LoadAllStaff()
	for _, v := range AllStaffList {
		if v.Phone == phone {
			return false
		}
	}
	return true
}

func GetDetailsPhone(phone string) (*Staff, error) {
	for _, v := range AllStaffList {
		if v.Phone == phone {
			return &v, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Could not find user with phone %v", phone))
}

func GetDetailsEmail(email string) (*Staff, error) {
	for _, v := range AllStaffList {
		fmt.Println(v.Email)
		if v.Email == email {
			return &v, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Could not find user with email address %v", email))
}

func UpadateStaff(userid int, content interface{}) (*Staff, error) {

	return nil, nil
}
