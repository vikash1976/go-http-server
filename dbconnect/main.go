package main

import (
	"database/sql"
	"encoding/json"

	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/vikash1976/dbconnect/encdec"
	"os"
)

//KeyConfig ... data type that holds the key value used in encryption/decryption
type KeyConfig struct {
	Key string `json:"key"`
}

var secretKey KeyConfig

var config mysql.Config

var db *sql.DB
var err error

//Function taking file name to read and return json's struct form in result
func readConfigFile(fileName string, result interface{}) {

	configFile, err := os.Open(fileName + ".json")
	if err != nil {
		fmt.Println("opening config file", err.Error())
	}

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(result); err != nil {
		fmt.Println("parsing config file", err.Error())
	}

}

//Init function to read required configuration files and open connection with database
func init() {
	readConfigFile("key", &secretKey)
	readConfigFile("config", &config)

	//key := []byte("DBSecretKey12121")

	//decrypt the passwd and store it back to config
	config.Passwd = encdec.Decrypt([]byte(secretKey.Key), config.Passwd)

	db, err = sql.Open("mysql", config.FormatDSN())
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	//fmt.Println("DB Stats: ", db.Stats().OpenConnections)
}

//function to read data from database
func fetchData() {
	
	/*stmtIns, err := db.Prepare("INSERT INTO numSq VALUES( ?, ? )") // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates
	*/

	// Prepare statement for reading data
	fmt.Println("DB Stats, Before Statement: ", db.Stats().OpenConnections)
	stmtOut, err := db.Prepare("SELECT squareNumber FROM numSq WHERE number = ?")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()
	fmt.Println("DB Stats, After Statement: ", db.Stats().OpenConnections)

	// Insert square numbers for 0-24 in the database
	/*for i := 0; i < 25; i++ {
	    _, err = stmtIns.Exec(i, (i * i)) // Insert tuples (i, i^2)
	    if err != nil {
	        panic(err.Error()) // proper error handling instead of panic in your app
	    }
	})*/

	var squareNum int // we "scan" the result in here

	// Query the square-number of 13
	
	err = stmtOut.QueryRow(17).Scan(&squareNum) // WHERE number = 13
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	
	fmt.Printf("\nThe square number of 17 is: %d\n", squareNum)

	// Query another number.. 1 maybe?
	err = stmtOut.QueryRow(1).Scan(&squareNum) // WHERE number = 1
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Printf("The square number of 1 is: %d\n", squareNum)
}


func main() {
	defer db.Close()
	fmt.Println("DB Stats: ", db.Stats().OpenConnections)

	if len(os.Args) == 3 {
		fmt.Println("Usage: encrypt <text to encrypt>")

		if os.Args[1] == "encrypt" {
			encryptedText := encdec.Encrypt([]byte(secretKey.Key), os.Args[2])
			fmt.Println("Encrypted Text: ", encryptedText)
			return
		}
	}
	fetchData()
	

}
