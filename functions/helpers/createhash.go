package helpers

import (
	"fmt"
	"hash"
	"hash/fnv"
)

var _hash hash.Hash32 = fnv.New32a()

// creates a hash, to check if any changes occured in the appsettings
func CreateHash(s string) string {
	_hash.Write([]byte(s))
	defer _hash.Reset()
	return fmt.Sprintf("%x", _hash.Sum(nil))
}
