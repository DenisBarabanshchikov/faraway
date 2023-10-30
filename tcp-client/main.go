package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	// Read the PoW challenge from the server
	challenge, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading challenge:", err)
		return
	}

	// Solve the PoW challenge (implement your PoW solver logic here)
	response := solvePoW(challenge)

	// Send the PoW response to the server
	_, err = conn.Write([]byte(response + "\n"))
	if err != nil {
		fmt.Println("Error sending PoW response:", err)
		return
	}

	// Read and display the quote from the server
	quote, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading quote:", err)
		return
	}
	fmt.Println(quote)
}

func solvePoW(challenge string) string {
	// Extract the salt and difficulty from the challenge
	difficulty := len(challenge) - 12 // Assuming the last 12 characters are the base64-encoded salt
	saltBase64 := challenge[difficulty:]
	salt, err := base64.StdEncoding.DecodeString(saltBase64)
	if err != nil {
		fmt.Println("Error decoding base64 salt:", err)
		return ""
	}

	// Solve the PoW challenge (implement your PoW solver logic here)
	var response string
	for {
		// Generate a random string
		rand.Seed(time.Now().UnixNano())
		randomString := generateRandomString(8)

		// Generate a SHA1 hash of the random string and salt
		hash := sha1.New()
		hash.Write([]byte(randomString))
		hash.Write(salt)
		hashBytes := hash.Sum(nil)

		// Encode the hash in base64
		hashBase64 := base64.StdEncoding.EncodeToString(hashBytes)

		// Check if the hash meets the difficulty requirement
		if strings.HasPrefix(hashBase64, challenge[:difficulty]) {
			response = randomString
			break
		}
	}

	return response
}

func generateRandomString(length int) string {
	var result string
	for i := 0; i < length; i++ {
		result += strconv.Itoa(rand.Intn(10))
	}
	return result
}
