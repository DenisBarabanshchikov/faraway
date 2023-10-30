#TCP server

## Description
This is a simple TCP server that can be used to test TCP clients. It will listen on port 8888.

Used PoW (Proof of Work) challenge with a specific difficulty level. The algorithm used in this code is a simple custom PoW algorithm. It combines the following steps:

1. Generate a random 8-byte salt value.
2. Encode the salt value in base64.
3. Construct the challenge string by concatenating a string of '0's with a base64-encoded salt.

The difficulty level is determined by the number of leading zeros required in the challenge string. The more leading zeros required, the more computationally intensive it is to find a valid solution.

This custom PoW algorithm is straightforward and serves as a basic example for illustrative purposes. In practice, more robust PoW algorithms like SHA-256 or Scrypt are commonly used, and their choice often depends on the specific security requirements of the system. Custom PoW algorithms can be prone to issues, and it's generally recommended to use well-established and tested PoW algorithms for real-world applications.

## Usage
```docker build -t word-of-wisdom-server .```

```docker run -p 8888:8888 word-of-wisdom-server```
