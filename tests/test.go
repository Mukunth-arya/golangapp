package test

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/Mukunth-arya/golangapp/models"
	"github.com/dgrijalva/jwt-go"
	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type LoggedDetails struct {
	Email string
	jwt.StandardClaims
}

func Validate(Data3 models.Data) {

	err := validation.ValidateStruct(&Data3,
		validation.Field(&Data3.CakeName, validation.Required,
			validation.Match(regexp.MustCompile(`[a-z]+`)).Error("Letters between a to z  are only allowed")),
		validation.Field(&Data3.Cakeflavour, validation.Required,
			validation.Match(regexp.MustCompile(`[a-z]+`)).Error("Letters between a to z  are only allowed")),
		validation.Field(&Data3.TypeofCream, validation.Required,
			validation.Match(regexp.MustCompile(`[a-z]+`)).Error("Letters between a to z  are only allowed")),
		validation.Field(&Data3.Toppings, validation.Required,
			validation.Match(regexp.MustCompile(`[a-z]+`)).Error("Letters between a to z  are only allowed")),
		validation.Field(&Data3.Shape, validation.Required,
			validation.Match(regexp.MustCompile(`[a-z]+`)).Error("Letters between a to z  are only allowed")),
		//validation.Field(&Data3.Satisfied, validation.Required,
		//validation.Match(regexp.MustCompile(`[a-z]+`)).Error("Letters between a to z  are only allowed")),
	)
	if err != nil {
		log.Fatal(err)
	}

}

func Validate1(Data2 models.Jwtmodel) {

	err := validation.ValidateStruct(&Data2,
		validation.Field(&Data2.Email, validation.Required,
			validation.Match(regexp.MustCompile(`[^0-9A-Za-z_]+`)).Error("The  value  you have enetered is  not in  proper value")),
		//validation.Field(&Data2.Username, validation.Required,
		//validation.Match(regexp.MustCompile(`[\W]`)).Error("The entered value is in not proper value")),
		validation.Field(&Data2.Password, validation.Required,
			validation.Match(regexp.MustCompile(`[^0-9A-Za-z_]+`)).Error("The  value  you have enetered is  not in  proper value")),
		//validation.Field(&Data2.Password, validation.Required, validation.Required,
		//validation.Match(regexp.MustCompile(`[\W]`)).Error("The entered value is in not proper value")),
	)
	if err != nil {
		log.Fatal(err)
	}

}

const db_urLS = ""
const db_names = "myseconddatabase"
const db_collection = "mysecondcollection"

var collections *mongo.Collection
var user models.Jwtmodel
var Founddata models.Jwtmodel
var secretkey = "sample"
var mysigningkey = []byte(secretkey)

func strcon() int {

	value2, err := strconv.Atoi(Founddata.Email)
	if err != nil {
		log.Fatal(err)
	}
	return value2
}

func init() {
	// Here  we make the url action
	data1 := options.Client().ApplyURI(db_urLS)
	// here we take the make the database connection
	data2, err := mongo.Connect(context.TODO(), data1)
	if err != nil {

		log.Fatal(err)
	}
	fmt.Println("Mongodb connection was successfully established")
	//here we establish the databases and collection
	collections = data2.Database(db_names).Collection(db_collection)
	//fmt.Println("here we established the connection successfully:", data3)
}
func Hashingpassword(password string) string {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {

		log.Panic(err)
	}
	return string(bytes)
}
func Verifypassword(userpassword, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(userpassword))

	return err == nil

}
func Signup(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/x-www-form-urlencoded")
	rw.Header().Set("Allow-Control-Allow-Methods", "PUT")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
	}
	Validate1(user)
	//json.NewEncoder(rw).Encode(&user)
	//value, err := primitive.ObjectIDFromHex(user.Email)
	//if err != nil {
	//log.Fatal(err)
	//}
	count, err := collections.CountDocuments(ctx, bson.M{"email": user.Email})
	if err != nil {
		log.Panic(err)
		rw.WriteHeader(http.StatusBadRequest)
	}
	user.Password = Hashingpassword(user.Password)
	if count > 0 {

		fmt.Println("Username  already exists ")
	}

	insert, err := collections.InsertOne(ctx, user)
	if err != nil {
		rw.WriteHeader(http.StatusForbidden)
	}
	fmt.Println(insert.InsertedID)
}
func Login(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	rw.Header().Set("Allow-Control-Allow-Methods", "PUT")

	var user models.Jwtmodel
	var Founddata models.Jwtmodel
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := json.NewDecoder(r.Body).Decode(&Founddata)
	if err != nil {
		log.Fatal(err)
	}
	Validate1(Founddata)

	//value, err := primitive.ObjectIDFromHex(user.Email)
	//if err != nil {
	//log.Fatal(err)
	//}

	err = collections.FindOne(ctx, bson.M{"email": Founddata.Email}).Decode(&Founddata)
	if err != nil {
		log.Fatal(err)
	}

	if Founddata.Email == "" {
		log.Println("Username or password is incorrect")
	}
	//if strcon(); !true {
	//fmt.Println("Please enter the proper value")
	//}

	check := Verifypassword(Founddata.Password, user.Password)
	if !check {

		fmt.Println("Username and password are incorrect")

	}

	GenerateToken(Founddata.Email)
	var tokens models.Tokens
	tokens.Email = Founddata.Email

	//collections.FindOne(ctx, bson.M{"email": Founddata.User_id})
	http.SetCookie(rw, &http.Cookie{
		Name:  "Token",
		Value: GenerateToken(Founddata.Email),
	})
	json.NewEncoder(rw).Encode(tokens)

}
func GenerateToken(email string) string {

	token := jwt.New(jwt.SigningMethodHS256)
	claim := token.Claims.(jwt.MapClaims)
	claim["authorized"] = true
	claim["email"] = email
	claim["exp"] = time.Now().Add(time.Minute * 30).Unix()
	tokenstring, err := token.SignedString(mysigningkey)
	if err != nil {
		fmt.Errorf("Something went wrong %s", err.Error())
	}
	return tokenstring

}

// Middileware for authetication
func Homepage(rw http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("Token")
	if err != nil {
		if err == http.ErrNoCookie {

			rw.WriteHeader(http.StatusUnauthorized)
			return
		}
		rw.WriteHeader(http.StatusBadRequest)
	}

	tokenStr := cookie.Value

	tkn, err := jwt.Parse(tokenStr, func(*jwt.Token) (interface{}, error) {

		return mysigningkey, nil

	})
	if !tkn.Valid {
		rw.WriteHeader(http.StatusUnauthorized)
	}
	fmt.Println("hello welcome")
}
