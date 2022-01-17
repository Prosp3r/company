package logic

import (
	"fmt"
	"math/rand"
	"time"

	petname "github.com/dustinkirkland/golang-petname"
	_ "github.com/go-sql-driver/mysql"
)

type Messages struct {
	Content     string `json:"content,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Email       string `json:"email,omitempty"`
	MessageType string `json:"message_type,omitempty"`
}

func init() {
	rand.Seed(time.Now().UTC().Unix())
}

// func DbConn() (db *sql.DB) {
// 	dbDriver := os.Getenv("SWIPEDBDRIVER")
// 	dbUser := os.Getenv("SWIPEDBUSER")
// 	dbPass := os.Getenv("SWIPEDBPASS")
// 	dbName := os.Getenv("SWIPEDBNAME")
// 	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return db
// }

func RandomCode(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	code := make([]byte, length)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}

	return string(code)
}

//GetPetName - returns a human readable name
func GetPetName(size int, separator string) (string, error) {
	names := petname.Generate(size, separator)
	return names, nil
}

func FailOnError(err error, msg, logTag string) bool {
	if err != nil {
		fmt.Printf("Could not carry out operation %s - %s, failed with error => %v\n", msg, logTag, err)
		return true
	}
	return false
}

//SendEmail -
func SendEmail() {}

//SendSMS -
func SendSMS() {}

func SendMessage(msg Messages) bool {
	if len(msg.Phone) > 5 {
		fmt.Printf("\nSending message through mobile. \n Recipient: %v \n Message: %v\n", msg.Phone, msg.Content)
	}

	if len(msg.Email) > 6 {
		fmt.Printf("\nSending message through email. \n Recipient: %v \n Message: %v\n", msg.Email, msg.Content)
	}

	return true
}
