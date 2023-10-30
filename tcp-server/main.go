package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
)

const (
	listenAddress = "localhost:8888" // Change this to the address you want to listen on
	difficulty    = 4                // Adjust the difficulty level as needed
)

var quotes = []string{
	// Populate with your quotes
	"Quote 1",
	"Quote 2",
	"Quote 3",
	// ...
}

func main() {
	listener, err := net.Listen("tcp", listenAddress)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("Listening on", listenAddress)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	// Generate a random quote to send to the client
	rand.Seed(time.Now().UnixNano())
	quote := quotes[rand.Intn(len(quotes))]

	// Send the PoW challenge to the client
	challenge := generateChallenge(difficulty)
	_, err := conn.Write([]byte("PoW Challenge: " + challenge + "\n"))
	if err != nil {
		return
	}

	// Read the response from the client
	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Verify the PoW
	if !verifyPoW(challenge, response) {
		fmt.Println("PoW verification failed")
		return
	}

	// Send the quote to the client
	_, err = conn.Write([]byte("Quote of Wisdom: " + quote + "\n"))
	if err != nil {
		return
	}
}

func generateChallenge(difficulty int) string {
	// Generate a random salt value of 8 bytes
	salt := make([]byte, 8)
	rand.Read(salt)

	// Encode the salt in base64
	saltBase64 := base64.StdEncoding.EncodeToString(salt)

	// Construct the challenge using the salt and difficulty
	challenge := strings.Repeat("0", difficulty) + saltBase64

	return challenge
}

func verifyPoW(challenge, response string) bool {
	// Verify the PoW response from the client
	return strings.HasPrefix(response, challenge)
}
