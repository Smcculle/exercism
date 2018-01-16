package robotname

import "time"
import (
	"fmt"
	"math/rand"
)

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	letterBytes   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

type Robot struct {
	name string
}

// For concurrent use, change issuedNames to sync.Map
var issuedNames = make(map[string]bool)
var src = rand.NewSource(time.Now().UnixNano())

func (r *Robot) Name() string {
	if r.name == "" {
		r.name = noDupRandomName()
		issuedNames[r.name] = true
	}
	return r.name
}

func (r *Robot) Reset() {
	r.name = ""
}

func noDupRandomName() string {
	name := randomName(2)
	for issuedNames[name] {
		name = randomName(2)
	}
	return name
}

// randomName generates a random string of characters from letterBytes of length n.
func randomName(n int) string {
	prefix := make([]byte, n)
	var suffix string
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		suffix = fmt.Sprintf("%03d", cache%1000)
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			prefix[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(prefix) + suffix
}
