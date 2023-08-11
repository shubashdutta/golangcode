package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/shubash/saibaba/moddel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

const connectingstring = "mongodb+srv://shubashduttasaibaba:LE19u3bje3SmmDl0@cluster0.ovbu7ak.mongodb.net/"

const dbname = "SaiBABA"
const colname = "user"

var Collection *mongo.Collection

func init() {
	ClientOptions := options.Client().ApplyURI(connectingstring)

	client, err := mongo.Connect(context.TODO(), ClientOptions)
	if err != nil {
		panic(err.Error())

	}
	fmt.Println("coonecting to mongo db")
	Collection = client.Database(dbname).Collection(colname)
}
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to sai baba turst web site")
}
func Singup(w http.ResponseWriter, r *http.Request) {
	 w.Header().Set("Access-Control-Allow-Origin", "https://shuas.onrender.com")
    w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    w.Header().Set("Access-Control-Allow-Credentials", "true")
	var user moddel.User

	_ = json.NewDecoder(r.Body).Decode(&user)
	count, err := Collection.CountDocuments(context.Background(), bson.M{"email": user.Email})
	// count, err := Collection.CountDocuments(context.Background(), bson.M{"email": user.Email})
	if err != nil {
		log.Println(err)
		return
	}
	if count != 0 {
		fmt.Println("this email is used allready ")
		return
	}
	num, err := Collection.CountDocuments(context.Background(), bson.M{"phone": user.Phone})
	if err != nil {
		log.Println(err)
	}
	if num != 0 {
		fmt.Println("this number is use to make a new id ")
	}

	password := Hashpassword(user.Password)
	user.Password = password

	insertoneuser(user)
	json.NewEncoder(w).Encode(user)
}
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json-x-www-from-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var user moddel.User

	var founduser moddel.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
	}
	err1 := Collection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&founduser)
	if err1 != nil {
		// log.Println(err1)
		log.Println("email is incorect")
		// json.NewEncoder(w).Encode("email incorrect")
		return

	}
	err2 := bcrypt.CompareHashAndPassword([]byte(founduser.Password), []byte(user.Password))
	if err2 != nil {
		log.Println("password does not match")
		return
	}

	json.NewEncoder(w).Encode(founduser)
	fmt.Println("user is on")
}

func Hashpassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
func insertoneuser(user moddel.User) {
	insert, err := Collection.InsertOne(context.Background(), user)
	if err != nil {
		panic(err)
	}
	fmt.Println("inserted one user in our database", insert.InsertedID)
}
