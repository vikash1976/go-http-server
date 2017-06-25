package main

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

var logger *log.Entry

const logFilePath = "./log/logFile_"
const (
	// File Path to key files used for signing JWT token
	privKeyFile = "jwtKeys/priKey.rsa"
	pubKeyFile  = "jwtKeys/pubKey.rsa.pub"
)

var (
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
)

func handleFatal(err error) {
	if err != nil {
		log.Fatalf("%+v", err)
	}
}

func initializeKeys() {
	privateKeyBytes, err := ioutil.ReadFile(privKeyFile)
	handleFatal(err)
	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	handleFatal(err)
	publicKeyBytes, err := ioutil.ReadFile(pubKeyFile)
	handleFatal(err)
	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	handleFatal(err)
}

// UserCredentials - will be used to store decoded post login data
type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Response - will be used to create Http response
type Response struct {
	Data string `json:"data"`
}

// Token - will be used to create token as Http response for /login endpoint
type Token struct {
	Token string `json:"token"`
}

func ConfigureNLaunchServer() {
	// Un-Protected Endpoint(s), will be used by user to post credentials to receive token
	http.HandleFunc("/getMeToken", TokenCreationHandler)
	// Protected Endpoint(s), can be accessed only with JWT token as Authorization header
	http.Handle("/iAmProtected", negroni.New(
		negroni.HandlerFunc(ValidateTokenMiddleware),
		negroni.Wrap(http.HandlerFunc(ProtectedResourceHandler)),
	))
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	t := time.Now()
	logFileSuffix := fmt.Sprintf("%d%02d%02d%02d%02d%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	logFilename := logFilePath + logFileSuffix + ".txt"
	file, err := os.OpenFile(logFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	initializeKeys()
	ConfigureNLaunchServer()
}

func ProtectedResourceHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{"Hurrey!!!, I am in as i have JWT token with me."}
	logger.Info("Access provided to protected resource")
	WriteJsonResponse(response, w)
}

func TokenCreationHandler(w http.ResponseWriter, req *http.Request) {
	logger = log.WithFields(log.Fields{
		"URL":           req.URL.Path,
		"Method":        req.Method,
		"RemoteAddress": req.RemoteAddr,
		"RequestURI":    req.RequestURI,
	})
	var user UserCredentials
	err := json.NewDecoder(req.Body).Decode(&user)
	logger.Data["userId"] = user.Username
	if err != nil {
		logger.Errorf("%v", "Request has error")
		http.Error(w, "Request has error", http.StatusForbidden)
		return
	}

	if user.Password != user.Username+"!!" {
		logger.Errorf("%v", "Invalid credentials")
		http.Error(w, "Invalid credentials", http.StatusForbidden)
		return
	}

	// preparing token, create a new one with given Signing Method,
	// make claim and set claim to token
	token := jwt.New(jwt.SigningMethodRS256)
	claims := make(jwt.MapClaims)
	// set expiry of the token
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(10)).Unix()
	// set issued at timestamp
	claims["iat"] = time.Now().Unix()
	claims["username"] = user.Username
	// set claims to the token
	token.Claims = claims

	if err != nil {
		logger.Errorf("%v", "Error extracting the key")
		http.Error(w, "Error extracting the key", http.StatusInternalServerError)
		return
	}

	// now sign the token with private key, token is ready for distribution
	tokenString, err := token.SignedString(privateKey)

	if err != nil {
		logger.Errorf("%v", "Error signing the token")
		http.Error(w, "Error signing the token", http.StatusInternalServerError)
		return
	}
	response := Token{tokenString}
	logger.Info("Token generated successfully.")
	WriteJsonResponse(response, w)
}

func ValidateTokenMiddleware(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	logger = log.WithFields(log.Fields{
		"URL":           req.URL.Path,
		"Method":        req.Method,
		"RemoteAddress": req.RemoteAddr,
		"RequestURI":    req.RequestURI,
	})
	// parse the token....
	// keyFunc receives the parsed token and should return the key for validating, i.e. publicKey
	token, err := request.ParseFromRequest(req, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return publicKey, nil
		})
	//tokenAsMap := token.Claims.(jwt.MapClaims)["username"]
	fmt.Println(tokenAsMap)
	if err == nil {
		if token.Valid {
			next(w, req)
		} else {
			logger.Errorf("%v", "Invalid token")
			http.Error(w, "Your token is invalid", http.StatusUnauthorized)

		}
	} else {
		logger.Errorf("%v by %v", "Unauthorized access to this resource", token.Claims.(jwt.MapClaims)["username"])
		http.Error(w, "Got a token? You need one to access to this resource.", http.StatusUnauthorized)
	}
}

func WriteJsonResponse(response interface{}, w http.ResponseWriter) {
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error creating response", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
