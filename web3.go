package web3

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"golang.org/x/crypto/sha3"
	"strings"
)

// ComputeHMACDigest calculates the HMAC (Hash-based Message Authentication Code) digest
// of a given message using SHA-256 as the underlying hash function.
//
// Parameters:
//   - message: A byte slice containing the message to be authenticated.
//   - secret: A byte slice containing the secret key used for HMAC computation.
//
// Returns:
//   - A byte slice containing the computed HMAC digest.
func ComputeHMACDigest(message, secret []byte) []byte {
	// Create a new HMAC hasher with SHA-256
	h := hmac.New(sha256.New, secret)

	// Write the message to the hasher
	h.Write(message)

	// Compute the final HMAC digest
	hashed := h.Sum(nil)

	// Convert the hash to a hexadecimal string
	return hashed
}

// ToChecksumAddress converts a given Ethereum address to a checksummed address
// according to EIP-55 standard.
//
// This function takes a byte slice representing an Ethereum address and returns
// the checksummed version of that address as a string. The checksum is calculated
// using the Keccak-256 hash of the lowercase hexadecimal encoding of the address.
//
// Parameters:
//   - a: A byte slice containing the 20-byte Ethereum address to be checksummed.
//
// Returns:
//   - string: The checksummed Ethereum address as a string, including the "0x" prefix.
//   - error: An error if the conversion process fails, otherwise nil.
func ToChecksumAddress(a []byte) (string, error) {
	address := hex.EncodeToString(a)
	// Remove the '0x' prefix if it exists
	// address = strings.ToLower(strings.TrimPrefix(address, "0x"))

	// Compute the Keccak-256 hash of the address
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write([]byte(address))
	hash := hasher.Sum(nil)

	// Encode hash to hexadecimal
	hashHex := hex.EncodeToString(hash)

	// Apply EIP-55 checksum rules
	var checksumAddress strings.Builder
	checksumAddress.WriteString("0x") // Prepend "0x"

	for i, c := range address {
		if hashHex[i] >= '8' { // Uppercase if the corresponding hash character is >= '8'
			checksumAddress.WriteRune(rune(strings.ToUpper(string(c))[0]))
		} else {
			checksumAddress.WriteRune(c)
		}
	}

	return checksumAddress.String(), nil
}

// IsChecksumAddress validates whether a given Ethereum address string is correctly checksummed.
//
// This function checks if the provided address adheres to the EIP-55 checksum format.
// It verifies that the address starts with "0x", has the correct length, and matches
// its own checksum when recalculated.
//
// Parameters:
//   - address: A string representing the Ethereum address to be validated.
//
// Returns:
//   - bool: true if the address is a valid checksummed address, false otherwise.
func IsChecksumAddress(address string) bool {
	if !strings.HasPrefix(address, "0x") || len(address) != 42 {
		return false
	}

	addressHex, err := hex.DecodeString(address[2:])
	if err != nil {
		return false
	}

	expectedChecksum, err := ToChecksumAddress(addressHex)
	if err != nil {
		return false
	}

	return address == expectedChecksum
}

// Keccak computes the Keccak-256 hash of the input data.
//
// This function uses the Keccak-256 algorithm, which is the original version of
// SHA-3 before it was standardized. It's commonly used in Ethereum and other
// blockchain applications.
//
// Parameters:
//   - input: A byte slice containing the data to be hashed.
//
// Returns:
//
//	A byte slice containing the 32-byte (256-bit) Keccak-256 hash of the input data.
func Keccak(input []byte) []byte {
	// Create a new Keccak-256 hasher
	hasher := sha3.NewLegacyKeccak256()

	// Write the input to the hasher
	hasher.Write(input)

	// Compute the hash digest
	hash := hasher.Sum(nil)

	return hash
}

// PadHexStringTo32Bytes converts a hexadecimal string to a byte slice and pads it to 32 bytes.
//
// This function takes a hexadecimal string (with or without '0x' prefix), converts it to bytes,
// and then pads the result to ensure it's exactly 32 bytes long. The original data is right-justified
// in the resulting byte slice, with any necessary zero padding added to the left.
//
// Parameters:
//   - hexString: A string containing a hexadecimal representation of bytes, optionally prefixed with '0x'.
//
// Returns:
//   - []byte: A byte slice that is exactly 32 bytes long, containing the original data right-justified
//     with zero padding on the left if necessary.
//   - error: An error if the input string is not a valid hexadecimal representation.
func PadHexStringTo32Bytes(hexString string) ([]byte, error) {
	if strings.HasPrefix(hexString, "0x") {
		hexString = hexString[2:]
	}

	// Decode hex string to byte slice
	data, err := hex.DecodeString(hexString)
	if err != nil {
		return nil, err
	}

	// Ensure the slice is exactly 32 bytes long by left-padding with 0x00
	paddedData := make([]byte, 32)
	copy(paddedData[32-len(data):], data) // Right-justify by copying to the right

	return paddedData, nil
}

// PadTo32Bytes pads the input byte slice to 32 bytes, right-justifying the original data.
//
// This function ensures that the returned byte slice is exactly 32 bytes long.
// If the input is already 32 bytes or longer, it is returned unchanged.
// For shorter inputs, the function pads with zero bytes on the left side.
//
// Parameters:
//   - sender: The input byte slice to be padded.
//
// Returns:
//   A new byte slice that is exactly 32 bytes long, containing the original data
//   right-justified (at the end of the slice), with any necessary zero padding at the beginning.

func PadTo32Bytes(sender []byte) []byte {
	if len(sender) >= 32 {
		return sender // Already 32 bytes long or longer
	}

	// Pad to 32 bytes (right-justify)
	paddedSender := make([]byte, 32)
	copy(paddedSender[32-len(sender):], sender)

	return paddedSender
}

// ConcatBytes concatenates multiple byte slices into a single byte slice.
//
// Parameters:
//   - bytes: A variadic parameter of type []byte, representing the byte slices to be concatenated.
//
// Returns:
//   - []byte: A new byte slice containing all input byte slices concatenated in the order they were provided.
func ConcatBytes(bytes ...[]byte) []byte {
	size := 0
	for _, b := range bytes {
		size += len(b)
	}
	result := make([]byte, 0, size)
	for _, b := range bytes {
		result = append(result, b...)
	}

	return result
}
