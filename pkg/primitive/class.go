package primitive

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"sync/atomic"
	"time"
)

var (
	classCounter = readRandomUint32()
)

func newClass(timestamp time.Time) string {
	var b [7]byte
	binary.BigEndian.PutUint32(b[0:4], uint32(timestamp.Unix()))
	putUint24(b[4:7], atomic.AddUint32(&classCounter, 1))
	return hex.EncodeToString(b[:])
}

func readRandomUint32() uint32 {
	var b [4]byte
	_, err := io.ReadFull(rand.Reader, b[:])
	if err != nil {
		panic(fmt.Errorf("cannot initialize objectid package with crypto.rand.Reader: %w", err))
	}

	return (uint32(b[0]) << 0) | (uint32(b[1]) << 8) | (uint32(b[2]) << 16) | (uint32(b[3]) << 24)
}

func putUint24(b []byte, v uint32) {
	b[0] = byte(v >> 16)
	b[1] = byte(v >> 8)
	b[2] = byte(v)
}
