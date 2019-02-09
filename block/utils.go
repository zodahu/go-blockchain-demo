package block

import "encoding/binary"

// Int64ToHex .
func Int64ToHex(i int64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i))
	return b
}
