package model

import (
	"errors"
	"fmt"
	"time"

	"github.com/Prosp3r/company/conf"
)

type Staff struct {
	ID        int64  `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Entrytime int64  `json:"entrytime,omitempty"`
}

type AddStaffInput struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
}

type DelStaffInput struct {
	ID int64 `json:"id,omitempty"`
}

var (
	AllStaffList []Staff
	AllStaffMap  = make(map[string]Staff)
)

/*
CREATE TABLE public.staff (
	id bigint NOT NULL AUTO_INCREMENT,
	"name" varchar NULL,
	email varchar NULL,
	phone varchar NULL,
	entrytime bigint NULL
);
CREATE INDEX staff_id_idx ON public.staff (id);
CREATE UNIQUE INDEX staff_email_idx ON public.staff (email);

*/
//CreateUser -
func CreateStaff(si AddStaffInput) (*Staff, error) {
	logTag := "Creating New Staff"

	// FakeName, err := GetPetName(1, "")
	// _ = FailOnError(err, "generaing random name", logTag)

	var nU Staff

	nU.Name = si.Name
	nU.Email = si.Email
	nU.Phone = si.Phone
	nU.Entrytime = time.Now().Unix()

	db, err := conf.GetDB()
	_ = FailOnError(err, "connecting to DB", logTag)
	// defer db.Close()
	var lastInsertId int64

	PQ := `INSERT INTO "staff"("name", "email", "phone", "entrytime") VALUES($1, $2, $3, $4) RETURNING id`
	Q, err := db.Prepare(PQ)
	_ = FailOnError(err, "Preparing qery", logTag)

	err = Q.QueryRow(nU.Name, nU.Email, nU.Phone, nU.Entrytime).Scan(&lastInsertId)
	fe := FailOnError(err, "Executing statement", logTag)
	if fe == true {
		return nil, err
	}

	nU.ID = lastInsertId
	return &nU, nil
}

//GetAllUsers - Populate the AllUsers slicee for quick manipulative access
func LoadAllStaff() bool {
	logTag := "LoadAllUsers"
	var AllStaffListx []Staff

	var u Staff

	db, err := conf.GetDB()
	_ = FailOnError(err, "connecting to DB", logTag)
	// defer db.Close()

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
		AllStaffListx = append(AllStaffListx, u)
	}
	AllStaffList = AllStaffListx
	return true
}

//isEmailUnique - checks if email has not been used previously.
func IsEmailUnique(email string) bool {
	//LoadAllUsers()
	if len(AllStaffList) > 0 {
		for _, v := range AllStaffList {
			if v.Email == email {
				return false
			}
		}
	}
	return true
}

//isPhoneUnique - checks if a phone number has not been used peviously.
func IsPhoneUnique(phone string) bool {
	// LoadAllStaff()
	if len(AllStaffList) > 0 {
		for _, v := range AllStaffList {
			fmt.Printf("%v : %v - %v - %v - %v", v.ID, v.Name, v.Email, v.Phone, v.Entrytime)
			if v.Phone == phone {
				return false
			}
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

func DeleteStaff(userid int64) bool {
	logTag := "DeleteStaff-model"
	db, err := conf.GetDB()
	fce := FailOnError(err, "connecting to DB", logTag)
	if fce == true {
		return false
	}

	PQ := `DELETE FROM "staff" WHERE id=$1`
	del, err := db.Exec(PQ, userid)
	fe := FailOnError(err, "Delete qery", logTag)
	if fe == true {
		return false
	}

	ar, err := del.RowsAffected()
	_ = FailOnError(err, "Checking Rows affected", logTag)

	if ar > 0 {
		return true
	}

	return false
}
