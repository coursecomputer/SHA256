package SHA256

import (
	"encoding/binary"
	"fmt"
)

// 64 bytes == 512 bits
// Adds a padding for the final part
func padding(message []byte, totalSize int) []byte {
	var size = len(message)
	var sizePart = make([]byte, 8) // 64 bits

	// get the number of padding
	var padSize = 64 - (size % 64)

	// Check if there is enough space for storage the size (64 bits == 8 bytes) + the byte (10000000 == 1 byte)
	// else adds 512 bits
	if padSize < (8 + 1) {
		padSize += 64
	}

	// Remove the size + the byte of end
	padSize -= 8 + 1

	// Adds end byte
	message = append(message, 0x80) // add 10000000
	// Adds zero padding
	message = append(message, make([]byte, padSize)...) // add 00000000

	// Convert the total size (uint64) to binary ([]byte)
	binary.BigEndian.PutUint64(sizePart, uint64(size + totalSize) * 8)

	// Adds size to data
	return append(message, sizePart...)
}

// Finalize the hashing and return the hash
func (s *SHA256) Digest() string {
	if s.remain != nil {
		s.remain = padding(s.remain, s.size)
		s.Update(nil)
		s.remain = nil
	}

	return fmt.Sprintf(
		"%08x%08x%08x%08x%08x%08x%08x%08x",
		s.states[0],
		s.states[1],
		s.states[2],
		s.states[3],
		s.states[4],
		s.states[5],
		s.states[6],
		s.states[7],
	)
}