// Package hash provides content hashing primitives used by future analyzers.
package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/zeebo/blake3"
)

// Digest contains primary and verification hashes for a byte stream.
type Digest struct {
	BLAKE3 string
	SHA256 string
}

// Reader calculates BLAKE3 and SHA-256 in one deterministic pass.
func Reader(r io.Reader) (Digest, error) {
	primary := blake3.New()
	verification := sha256.New()
	if _, err := io.Copy(io.MultiWriter(primary, verification), r); err != nil {
		return Digest{}, fmt.Errorf("hash content: %w", err)
	}
	return Digest{BLAKE3: hex.EncodeToString(primary.Sum(nil)), SHA256: hex.EncodeToString(verification.Sum(nil))}, nil
}
