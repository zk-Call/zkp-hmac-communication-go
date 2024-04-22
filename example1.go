package main // Declares that this file is part of the main package

import (
	"fmt"                                  // Import the "fmt" package for formatted I/O
	"sync"                                 // Import the "sync" package for synchronization primitives
	HMAC_env "tmp/src/HMAC/core"           // Import the HMAC core package and alias it as "HMAC_env"
	seed_env "tmp/src/SeedGeneration/core" // Import the SeedGeneration core package and alias it as "seed_env"
)

var DEBUG = true // Define a global variable DEBUG and set it to true

func printMsg(who string, message string) { // Define a function named printMsg that takes two string parameters
	if DEBUG { // If DEBUG is true, execute the following block
		fmt.Printf("[%s] %s\n", who, message) // Print a formatted message to standard output
	}
}

func client(clientSocket chan string, serverSocket chan string, wg *sync.WaitGroup) { // Define a function named client with three parameters
	defer wg.Done() // Decrement the WaitGroup counter when this function exits

	// Generating the main seed
	mainSeed := seed_env.NewSeedGenerator("jack").Generate() // Generate a main seed using the SeedGenerator with a specified username
	obj := HMAC_env.NewHMACClient("sha256", mainSeed, 1)     // Create a new HMACClient object with SHA-256 hash algorithm, main seed, and iteration count
	obj.InitDecryptDict()                                    // Initialize the decryption dictionary for the HMACClient

	// Sending the main seed to the server through the serverSocket channel
	serverSocket <- string(mainSeed) // Convert the main seed to a string and send it to the serverSocket channel

	// Checking if the server has successfully received the seed
	if <-clientSocket == obj.EncryptMessage("") { // Wait for a response from the server via the clientSocket channel and compare it with an encrypted empty message
		// If successful, send a message to the server
		message := "hello"                                                 // Define a message to send to the server
		serverSocket <- obj.EncryptMessageByChunks(message)                // Encrypt and send the message to the server through the serverSocket channel
		printMsg("client", fmt.Sprintf("client sent message %s", message)) // Print a message indicating that the client has sent a message

		// Checking if the server has successfully decrypted the message
		if <-clientSocket == obj.EncryptMessage(message) { // Wait for a response from the server via the clientSocket channel and compare it with the encrypted message
			printMsg("client", "server has decrypt message") // If the message matches, print a message indicating that the server has decrypted the message
		}
	}
}

func server(serverSocket chan string, clientSocket chan string, wg *sync.WaitGroup) { // Define a function named server with three parameters
	defer wg.Done() // Decrement the WaitGroup counter when this function exits

	// Receiving the main seed from the client through the serverSocket channel
	mainSeed := <-serverSocket // Receive the main seed from the client via the serverSocket channel

	// Create a new HMACClient object with SHA-256 hash algorithm, received main seed, and iteration count
	obj := HMAC_env.NewHMACClient("sha256", []byte(mainSeed), 1)
	obj.InitDecryptDict() // Initialize the decryption dictionary for the HMACClient

	// Sending an empty message to the client as acknowledgment
	clientSocket <- obj.EncryptMessage("") // Encrypt and send an empty message to the client through the clientSocket channel

	// Receiving the message from the client through the serverSocket channel
	client_mes := <-serverSocket                                    // Receive the encrypted message from the client via the serverSocket channel
	printMsg("server", fmt.Sprintf("msg encrypted %s", client_mes)) // Print a message indicating the encrypted message received

	// Decrypt the received message by chunks
	msg := obj.DecryptMessageByChunks(client_mes)            // Decrypt the received message by chunks
	printMsg("server", fmt.Sprintf("msg decrypted %s", msg)) // Print a message indicating the decrypted message

	// Sending the decrypted message back to the client through the clientSocket channel
	clientSocket <- obj.EncryptMessageByChunks(msg) // Encrypt and send the decrypted message back to the client
}

func main() { // The entry point of the program
	clientSocket := make(chan string) // Create an unbuffered channel of string type named clientSocket
	serverSocket := make(chan string) // Create an unbuffered channel of string type named serverSocket
	var wg sync.WaitGroup             // Declare a WaitGroup variable named wg
	wg.Add(2)                         // Increment the WaitGroup counter by 2

	go func() { // Start a new goroutine
		defer close(clientSocket) // Close the clientSocket channel when this goroutine finishes
		defer close(serverSocket) // Close the serverSocket channel when this goroutine finishes
		wg.Wait()                 // Wait until the WaitGroup counter becomes zero
	}()

	go client(clientSocket, serverSocket, &wg) // Start a new goroutine for the client function
	go server(serverSocket, clientSocket, &wg) // Start a new goroutine for the server function

	wg.Wait() // Wait until all goroutines are finished
}
