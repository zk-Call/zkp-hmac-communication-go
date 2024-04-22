<div align="center">
  <img src="assets/zk-Call - Preview [Go].png">
</div>
<h1 align="center">zk-Call & Labs</h1>

<div align="center">
  <strong>"Zero-Knowledge" Proof Implementation with HMAC Communication in Go</strong>
</div>
<br />
<div align="center">
  <img src="http://badges.github.io/stability-badges/dist/experimental.svg" alt="Experimental" />
</div>
<div align="center">
  <sub>
    Built by <a href="https://github.com/zk-Call"> zk-Call</a> :)
  </sub>
</div>
<br />

# Table of Contents
- [Credits](#credits)
- [Purpose](#purpose)
- [How it Works](#how-it-works)
- [API](#api)
- [Example Usage](#example-usage)
<br>

# Credits

This repository hosts a refined implementation of [**Schnorr's Protocol**](https://en.wikipedia.org/wiki/Schnorr_signature), innovatively incorporating a state seed for enhanced security measures. While the underlying proofs may appear intricate, I aim to elucidate their functionality to the best of my ability. However, for a deeper understanding, I encourage referencing the seminal research papers underpinning this implementation, as they offer comprehensive insights.

---

![Detailed Schematic Overview of Schnorr's Protocol (Example)](assets/Schnorr's%20Protocol.png)
<br>
<br>

**For further exploration:**

[**Elliptic Curve Based "Zero-Knowledge" Proofs and Their Applicability on Resource Constrained Devices by Ioannis Chatzigiannakis, Apostolos Pyrgelis, Paul G. Spirakis, and Yannis C. Stamatiou**](https://arxiv.org/pdf/1107.1626.pdf)

---

![Detailed Schematic Overview of Elliptic Curves (Example)](assets/Elliptic%20Curve.png)
<br>

Additionally, this repository delves into the concepts of **"Zero-Knowledge" Proofs (ZKPs)** and **Hash-based Message Authentication Codes (HMACs)**. **ZKPs** are cryptographic protocols that allow one party **(the prover)** to prove to another party **(the verifier)** that a given statement is true, without revealing any additional information beyond the validity of the statement itself. This property is particularly valuable for preserving privacy while establishing trust.
<br>

On the other hand, **HMACs** are a type of cryptographic hash function used for message authentication. They involve a cryptographic hash function **(such as SHA-256)** and a secret cryptographic key. **HMACs** provide a way to verify both the data integrity and the authenticity of a message, ensuring that it has not been altered or tampered with during transmission and that it indeed originates from the purported sender.
<br>
<br>

# Purpose

In today's rapidly evolving IT and application development landscape, **"Zero-Knowledge" Proofs (ZKPs)** emerge as a pivotal paradigm for authentication security. Their capacity to affirm the validity of a claim, such as proving possession of a secret password — without revealing any sensitive information about the claim itself, such as passwords or hashes, revolutionizes the assurance of secure **AAA operations** (**authentication**, **authorization**, and **accounting**).

---

![The Purpose of our Repositories and The Overall Technology](assets/Purpose-1.png)
<br>

**zk-Call & Labs** represents an implementation of a [**Non-Interactive "Zero-Knowledge" Proof**](https://en.wikipedia.org/wiki/Non-interactive_zero-knowledge_proof) **(NIZKP)** protocol tailored specifically for validating text-based secrets. This framework proves invaluable for safeguarding passwords and other authentication mechanisms, ensuring robust security measures without compromising privacy. Additionally, the integration of **HMAC (Hash-Based Message Authentication Code)** further fortifies the authentication process, enhancing data integrity and thwarting potential security breaches.
<br>
<br>

# How It Works

The authentication protocol employed in this system operates based on two fundamental concepts:
**"Zero-Knowledge" Proofs (ZKPs)** and **Hash-Based Message Authentication Code (HMAC)**. Let's delve into each of these components and understand how they synergize to ensure secure authentication in messaging applications.
<br>

"Zero-Knowledge" Proofs (ZKPs)
---


#### **"Zero-Knowledge" Proofs (ZKPs):** 
**ZKPs** form the bedrock of privacy-preserving authentication mechanisms. These proofs allow one party **(the prover)** to demonstrate the validity of a claim to another party **(the verifier)** without revealing any additional information beyond the claim's validity. In essence, **ZKPs** enable authentication without the need for the prover to disclose sensitive data, such as passwords or cryptographic keys.

---

![Detailed Schematic Overview of "Zero-Knowledge" Technology (1)](assets/ZKP-HMAC-1.png)
![Detailed Schematic Overview of "Zero-Knowledge" Technology (2)](assets/ZKP-HMAC-2.png)
![Detailed Schematic Overview of "Zero-Knowledge" Technology (3)](assets/ZKP-HMAC-3.png)
![Detailed Schematic Overview of "Zero-Knowledge" Technology (4)](assets/ZKP-HMAC-4.png)
<br>


#### **Application in Authentication:** 
In the context of messaging applications, **ZKPs** play a pivotal role in verifying a user's identity without the need to transmit explicit credentials over the network. Instead, users can generate cryptographic proofs attesting to their identity or possession of certain credentials without exposing those credentials themselves. This ensures that sensitive information remains confidential during the authentication process, bolstering security and privacy.
<br>
<br>


Hash-Based Message Authentication Code (HMAC)
---

#### **Hash-Based Message Authentication Code (HMAC):**
**HMAC** provides a robust mechanism for verifying the integrity and authenticity of messages exchanged between parties. It involves the use of a cryptographic hash function in conjunction with a secret key to generate a unique code **(the HMAC)** for each message. This code serves as a digital signature, allowing the recipient to verify that the message has not been tampered with or altered during transmission.

---

![Detailed Schematic Overview of HMAC Encryption](assets/HMAC.png)


#### **Application in Authentication:**
In messaging applications, **HMAC** can be employed to authenticate message senders and ensure the integrity of communication channels. By appending an **HMAC** to each message using a shared secret key, both the sender and recipient can validate the message's authenticity upon receipt. Any unauthorized modifications to the message would result in a mismatch between the **computed HMAC** and the **received HMAC**, thereby alerting the recipient to potential tampering.
<br>
<br>

Synergistic Operation
---
When combined, **"Zero-Knowledge" Proofs** and **HMAC** create a formidable framework for secure authentication in messaging applications. **ZKPs** facilitate identity verification without divulging sensitive information, while **HMAC** ensures the integrity and authenticity of messages exchanged between parties. Together, these mechanisms uphold the confidentiality, integrity, and authenticity of communication channels, safeguarding users' privacy and security in the digital realm.

---

![The Advantages of Synergy between "Zero-Knowledge" Proof and HMAC](assets/Synergistic%20Operation.png)
<br>
<br>

# API

The **`"Zero-Knowledge"`** Go API is meant to be simple and intuitive:<br>

## Core Components

The **`Core Components`** are key for establishing a secure and efficient framework for cryptographic protocols;
streamlining the creation and validation of **"Zero-Knowledge" Proofs (ZKPs)**. They enhance anonymous, data-safe proof
validations.

![Detailed Schematic Overview of Core Components](assets/Core%20Components%20(Go).png)

---

#### ZeroKnowledge.models.ZeroKnowledgeParams

The parameters **used to initialize the "Zero-Knowledge"** crypto system.

    struct ZeroKnowledgeParams():
        """
        Parameters used to construct a Zero-Knowledge Proof state, utilizing an elliptic curve and a random salt
        """
        algorithm: string                    # Hashing algorithm name
        curve: string                  # Standard Elliptic Curve name to use
        s: []byte                      # Random salt for the state

#### ZeroKnowledge.models.ZeroKnowledgeSignature

A **cryptographic "Zero-Knowledge"** signature that can be used to verify future messages.

    struct ZeroKnowledgeSignature():
        """
        Cryptographic public signature designed to verify future messages
        """
        params: ZeroKnowledgeParams       # Reference ZeroKnowledge Parameters
        signature: []byte              # The public key derived from your original secret

#### ZeroKnowledge.models.ZeroKnowledgeProof

A **cryptographic proof** that can be verified against a signature.

    struct ZeroKnowledgeProof(NamedTuple):
        """
        Non-deterministic cryptographic Zero-Knowledge Proof designed to confirm that the
        private key creating the proof matches the key used to generate the signature
        """
        params: ZeroKnowledgeParams        # Reference ZeroKnowledge Parameters
        c: []byte                      # The hash of the signed data and random point, R
        m: []byte                      # The offset from the secret `r` (`R=r*g`) from c * Hash(secret)

#### ZeroKnowledge.models.ZeroKnowledgeData

**Wrapper** that contains **a proof and the necessary data** to validate the proof against a signature.

    class ZeroKnowledgeData(NamedTuple):
        """
        Wrapper designed to hold data along with its corresponding signed proof
        """
        data: string
        proof: ZeroKnowledgeProof

---

## ZeroKnowledge

The **`ZeroKnowledge`** class is the central component of **`ZeroKnowledge`** and its state (defined by **`ZeroKnowledgeParams`**) should be inherently known to both the **Client (Prover)** and **Server (Verifier)**.

![Comprehensive Visual Guide to ZeroKnowledge Framework](assets/ZeroKnowledge%20(Go).png)

---

#### Instance Methods

<table>
  <tr>
    <th width="9%">Method</th>
    <th width="46%">Params</th>
    <th width="10%">Role</th>
    <th width="35%">Purpose</th>
  </tr>
  <tr>
    <td><code>CreateSignature</code></td>
    <td><code>secret: []byte</code></td>
    <td>Prover</td>
    <td>Create a cryptographic <code>signature</code> derived from the value <code>secret</code> to be generated during initial registration and stored for subsequent <code>challenge</code> proofs.</td>
  </tr>
  <tr>
    <td><code>sign</code></td>
    <td><code>secret: []byte</code> <br /> <code>data: interface{}</code></td>
    <td>Prover</td>
    <td>Create a <code>ZeroKnowledgeData</code> object using the <code>secret</code> and any additional data.
  </tr>
  <tr>
    <td><code>verify</code></td>
    <td><code>challenge: interface{}</code> <br /> <code>signature: ZeroKnowledgeSignature</code> <br /> <code>data: Optional[Union[str, bytes, int]]</code></td>
    <td>Verifier</td>
    <td>Verify the user-provided <code>challenge</code> against the stored <code>signature</code> and randomly generated token to verify the validity of the <code>challenge</code>.</td>
  </tr>
</table>

---

# Example Usage

TODO: Include **`Example Usage`**

## Example 1

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

---

## Example 2

    package main // Declares that this file is part of the main package
    
    import (
        "fmt" // Import the "fmt" package for formatted I/O
        "sync" // Import the "sync" package for synchronization primitives
        zkx "tmp/src/ZeroKnowledge/core" // Import the ZeroKnowledge core package and alias it as "zkx"
        zkx_models "tmp/src/ZeroKnowledge/models" // Import the ZeroKnowledge models package and alias it as "zkx_models"
    )
    
    var DEBUG = true // Define a global variable DEBUG and set it to true
    
    func printMsg(who string, message string) { // Define a function named printMsg that takes two string parameters
        if DEBUG { // If DEBUG is true, execute the following block
            fmt.Printf("[%s] %s\n", who, message) // Print a formatted message to standard output
        }
    }
    
    func client(clientSocket chan string, serverSocket chan string, wg *sync.WaitGroup) { // Define a function named client with three parameters
        defer wg.Done() // Decrement the WaitGroup counter when this function exits
    
        // Create a ZeroKnowledge object for the client with specified curve and hash algorithm
        clientObject := zkx.New("Ed25519", "blake2b", nil, "HB2B", 16)
    
        // Generate a signature for the client identity
        identity := "John"
        signature := clientObject.CreateSignature(identity)
    
        // Send the signature to the server through the serverSocket channel
        serverSocket <- zkx_models.ZeroKnowledgeSignature.ToJSON(signature)
        printMsg("client", fmt.Sprintf("Sent signature: %s", zkx_models.ZeroKnowledgeSignature.ToJSON()))
    
        // Receive token from the server through the clientSocket channel
        token := <-clientSocket
        printMsg("client", fmt.Sprintf("Received token: %s", token))
    
        // Generate a proof using client identity and token
        proof := zkx_models.ZeroKnowledgeSignature.ToJSON(clientObject.Sign(identity, token))
        printMsg("client", fmt.Sprintf("Proof: %s", proof))
    
        // Send proof to the server through the serverSocket channel
        serverSocket <- proof
    
        // Receive result from the server through the clientSocket channel
        result := <-clientSocket
        printMsg("client", fmt.Sprintf("Result: %s", result))
    }
    
    func server(serverSocket chan string, clientSocket chan string, wg *sync.WaitGroup) { // Define a function named server with three parameters
        defer wg.Done() // Decrement the WaitGroup counter when this function exits
    
        // Set the server password
        serverPassword := "SecretServerPassword"
    
        // Create a ZeroKnowledge object for the server with specified curve and hash algorithm
        serverZK := zkx.New("Ed25519", "blake2b", nil, "HB2B", 16)
    
        // Receive client signature from the client through the serverSocket channel
        clientSig := <-serverSocket
        clientSignature := zkx_models.ZeroKnowledgeSignature.FromJSON(clientSig)
        printMsg("server", fmt.Sprintf("Received client signature: %s", clientSignature)
    
        // Generate a token signed by the server for the client
        token := serverZK.Sign(serverPassword, zkx.Token(serverZK))
        printMsg("server", fmt.Sprintf("Generated token: %s", token))
    
        // Send the token to the client through the clientSocket channel
        clientSocket <- zkx_models.ZeroKnowledgeData.ToJSON(token)
    
        // Receive proof from the client through the serverSocket channel
        proof := <-serverSocket
        clientProof := zkx_models.ZeroKnowledgeData.FromJSON(proof)
        printMsg("server", fmt.Sprintf("Received proof: %s", proof))
    
        // Verify the received proof
        tokenData := zkx_models.ZeroKnowledgeData.FromJSON(clientProof.Data)
        serverVerification := serverZK.Verify(tokenData, clientSignature)
        printMsg("server", fmt.Sprintf("Server verification result: %t", serverVerification))
    
        // If server verification fails, notify the client through the clientSocket channel
        if !serverVerification {
            clientSocket <- "Server verification failed"
        } else {
            // Otherwise, verify the proof using client signature
            clientVerification := clientSignature.Verify(clientProof, tokenData)
            printMsg("server", fmt.Sprintf("Client verification result: %t", clientVerification))
            if clientVerification {
                clientSocket <- "Verification successful"
            } else {
                clientSocket <- "Verification failed"
            }
        }
    }
    
    func main() { // Entry point of the program
        clientSocket := make(chan string) // Create a unbuffered channel of string type named clientSocket
        serverSocket := make(chan string) // Create a unbuffered channel of string type named serverSocket
        var wg sync.WaitGroup // Declare a WaitGroup variable named wg
        wg.Add(2) // Increment the WaitGroup counter by 2
    
        go func() { // Start a new goroutine
            defer close(clientSocket) // Close the clientSocket channel when this goroutine finishes
            defer close(serverSocket) // Close the serverSocket channel when this goroutine finishes
            wg.Wait() // Wait until the WaitGroup counter becomes zero
        }()
    
        go client(clientSocket, serverSocket, &wg) // Start a new goroutine for the client function
        go server(serverSocket, clientSocket, &wg) // Start a new goroutine for the server function
    
        wg.Wait() // Wait until all goroutines are finished
    }

---

## Example 3

    package main // Declares that this file is part of the main package
    
    import (
        "fmt"  // Import the "fmt" package for formatted I/O
        "sync" // Import the "sync" package for synchronization primitives
        HMAC_env "tmp/src/HMAC/core"
        seed_env "tmp/src/SeedGeneration/core"
        zkx "tmp/src/ZeroKnowledge/core"          // Import the ZeroKnowledge core package and alias it as "zkx"
        zkx_models "tmp/src/ZeroKnowledge/models" // Import the ZeroKnowledge models package and alias it as "zkx_models"
    )
    
    var DEBUG = true // Define a global variable DEBUG and set it to true
    
    func printMsg(who string, message string) { // Define a function named printMsg that takes two string parameters
        if DEBUG {                              // If DEBUG is true, execute the following block
          fmt.Printf("[%s] %s\n", who, message) // Print a formatted message to standard output
        }
    }
    
    func client(clientSocket chan string, serverSocket chan string, wg *sync.WaitGroup) { // Define a function named client with three parameters
        defer wg.Done() // Decrement the WaitGroup counter when this function exits
    
        // Create a ZeroKnowledge object for the client with specified curve and hash algorithm
        clientObject := zkx.New("Ed25519", "blake2b", nil, "HB2B", 16)
    
        // Generate a signature for the client identity
        identity := "John"
        signature := clientObject.CreateSignature(identity)
    
        // Send the signature to the server through the serverSocket channel
        serverSocket <- zkx_models.ZeroKnowledgeSignature.ToJSON(signature)
        printMsg("client", fmt.Sprintf("Sent signature: %s", zkx_models.ZeroKnowledgeSignature.ToJSON()))
    
        // Receive token from the server through the clientSocket channel
        token := <-clientSocket
        printMsg("client", fmt.Sprintf("Received token: %s", token))
    
        // Generate a proof using client identity and token
        proof := zkx_models.ZeroKnowledgeSignature.ToJSON(clientObject.Sign(identity, token))
        printMsg("client", fmt.Sprintf("Proof: %s", proof))
    
        // Send proof to the server through the serverSocket channel
        serverSocket <- proof
    
        // Receive result from the server through the clientSocket channel
        result := <-clientSocket
        printMsg("client", fmt.Sprintf("Result: %s", result))
        if result {
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
    }
    
    func server(serverSocket chan string, clientSocket chan string, wg *sync.WaitGroup) { // Define a function named server with three parameters
        defer wg.Done() // Decrement the WaitGroup counter when this function exits
    
        // Set the server password
        serverPassword := "SecretServerPassword"
    
        // Create a ZeroKnowledge object for the server with specified curve and hash algorithm
        serverZK := zkx.New("Ed25519", "blake2b", nil, "HB2B", 16)
    
        // Receive client signature from the client through the serverSocket channel
        clientSig := <-serverSocket
        clientSignature := zkx_models.ZeroKnowledgeSignature.FromJSON(clientSig)
        printMsg("server", fmt.Sprintf("Received client signature: %s", clientSignature)
    
        // Generate a token signed by the server for the client
        token := serverZK.Sign(serverPassword, zkx.Token(serverZK))
        printMsg("server", fmt.Sprintf("Generated token: %s", token))
    
        // Send the token to the client through the clientSocket channel
        clientSocket <- zkx_models.ZeroKnowledgeData.ToJSON(token)
    
        // Receive proof from the client through the serverSocket channel
        proof := <-serverSocket
        clientProof := zkx_models.ZeroKnowledgeData.FromJSON(proof)
        printMsg("server", fmt.Sprintf("Received proof: %s", proof))
    
        // Verify the received proof
        tokenData := zkx_models.ZeroKnowledgeData.FromJSON(clientProof.Data)
        serverVerification := serverZK.Verify(tokenData, clientSignature)
        printMsg("server", fmt.Sprintf("Server verification result: %t", serverVerification))
    
        // If server verification fails, notify the client through the clientSocket channel
        if !serverVerification {
            clientSocket <- "Server verification failed"
        } else {
            // Otherwise, verify the proof using client signature
            clientVerification := clientSignature.Verify(clientProof, tokenData)
            printMsg("server", fmt.Sprintf("Client verification result: %t", clientVerification))
            if clientVerification {
                clientSocket <- "Verification successful"
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
    
            } else {
                clientSocket <- "Verification failed"
            }
        }
    }
    
    func main() { // Entry point of the program
        clientSocket := make(chan string) // Create a unbuffered channel of string type named clientSocket
        serverSocket := make(chan string) // Create a unbuffered channel of string type named serverSocket
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
