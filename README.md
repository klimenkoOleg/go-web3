# go-web3 for Ethereum and ZK SSO blockchain

`go-web3` is a Go library designed for interacting with the Ethereum blockchain. It provides utilities for working with Ethereum addresses, computing cryptographic hashes, making it easier to build applications that communicate with Ethereum nodes, send transactions, and interact with smart contracts.

## Features

- Compute HMAC digests using SHA-256
- Convert Ethereum addresses to checksummed format (EIP-55)
- Validate checksummed Ethereum addresses
- Compute Keccak-256 hashes
- Pad hexadecimal strings and byte slices to 32 bytes
- Concatenate multiple byte slices

## Requirements

- Go 1.13 or later

## Installation

To install the library, clone the repository and build the project:

```bash
git clone https://github.com/outofboxer/go-web3.git
cd go-web3
go build -o go-web3
```

## Getting Started

To get started with the `go-web3` library, follow these steps:

1. Import the `go-web3` package in your Go code:

    ```go
    import "github.com/outofboxer/go-web3"
    ```

2. Update your module dependencies by running the shell command:
    ```bash
    go mod tidy
    ```

## Usage

### Compute HMAC Digest

```go
message := []byte("your message")
secret := []byte("your secret key")
digest := web3.ComputeHMACDigest(message, secret)
```

### Convert to Checksum Address

```go
address := []byte{0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcd, 0xef, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcd, 0xef, 0x12, 0x34, 0x56, 0x78}
checksumAddress, err := web3.ToChecksumAddress(address)
if err != nil {
    // handle error
}
```

### Validate Checksum Address

```go
isValid := web3.IsChecksumAddress("0x1234567890abcdef1234567890abcdef12345678")
```

### Compute Keccak-256 Hash

```go
data := []byte("your data")
hash := web3.Keccak(data)
```

### Pad Hex String to 32 Bytes

```go
paddedData, err := web3.PadHexStringTo32Bytes("0x1234")
if err != nil {
    // handle error
}
```

### Pad Byte Slice to 32 Bytes

```go
paddedBytes := web3.PadTo32Bytes([]byte{0x12, 0x34})
```

### Concatenate Byte Slices

```go
concatenated := web3.ConcatBytes([]byte{0x12, 0x34}, []byte{0x56, 0x78})
```

## License

This project is licensed under the terms of the MIT license. See the [LICENSE](LICENSE) file for details.
```

This `README.md` provides an overview of the library, installation instructions, and examples of how to use its features. Adjust the content as needed to better fit your project's specifics.
