package main

import (
    "database/sql"
    "fmt"
    "github.com/go-sql-driver/mysql"
    "crypto/md5"
    "encoding/hex"
    //"golang.org/x/crypto/bcrypt"
    "crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
    "io"
)




func GetMD5Hash(text string) string {
    hasher := md5.New()
    
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}


// encrypt string to base64 crypto using AES
func encrypt(key []byte, text string) string {
	// key := []byte(keyText)
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// convert to base64
	return base64.URLEncoding.EncodeToString(ciphertext)
}

// decrypt from base64 to decrypted string
func decrypt(key []byte, cryptoText string) string {
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext)
}
func main() {
    
    originalText := "root"
	fmt.Println(originalText)

	key := []byte("DBSecretKey12121")

	// encrypt value to base64
	cryptoText := encrypt(key, originalText)
	fmt.Println(cryptoText)

	// encrypt base64 crypto to original value
	//text := decrypt(key, "_e9KzArAnb_xgB2UqfsKYv2p5LE=")
	//fmt.Printf(text)
    
    var config = mysql.Config {
    User: "root",
    Passwd: decrypt(key, "_e9KzArAnb_xgB2UqfsKYv2p5LE="),
    DBName: "rkd",
    Net: "tcp",
    Addr: "localhost:3306",
}
    //fmt.Println("Hash: ", GetMD5Hash("Vikash"))
    
    /*password := []byte("MyDarkSecret")

    // Hashing the password with the default cost of 10
    hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(hashedPassword))

    // Comparing the password with the hash
    err = bcrypt.CompareHashAndPassword(hashedPassword, password)
    
    fmt.Println(err) // nil means it is a match
   
    */
    
    db, err := sql.Open("mysql", config.FormatDSN())// "root:root@/rkd")
    if err != nil {
        panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
    }
    defer db.Close()

    // Prepare statement for inserting data
    stmtIns, err := db.Prepare("INSERT INTO numSq VALUES( ?, ? )") // ? = placeholder
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

    // Prepare statement for reading data
    stmtOut, err := db.Prepare("SELECT squareNumber FROM numSq WHERE number = ?")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    defer stmtOut.Close()

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
    fmt.Printf("\nThe square number of 13 is: %d\n", squareNum)

    // Query another number.. 1 maybe?
    err = stmtOut.QueryRow(1).Scan(&squareNum) // WHERE number = 1
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    fmt.Printf("The square number of 1 is: %d\n", squareNum)
    
    // Prepare statement for reading full table
    rows, err := db.Query("SELECT * FROM numSq")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    // Get column names
    columns, err := rows.Columns()
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    // Make a slice for the values
    values := make([]sql.RawBytes, len(columns))

    // rows.Scan wants '[]interface{}' as an argument, so we must copy the
    // references into such a slice
    // See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
    scanArgs := make([]interface{}, len(values))
    for i := range values {
        scanArgs[i] = &values[i]
    }

    // Fetch rows
    for rows.Next() {
        // get RawBytes from data
        err = rows.Scan(scanArgs...)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }

        // Now do something with the data.
        // Here we just print each column as a string.
        var value string
        for i, col := range values {
            // Here we can check if the value is nil (NULL value)
            if col == nil {
                value = "NULL"
            } else {
                value = string(col)
            }
            fmt.Println(columns[i], ": ", value)
        }
        fmt.Println("-----------------------------------")
    }
    if err = rows.Err(); err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    
    
}