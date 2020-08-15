package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/oklog/ulid"
	"github.com/spf13/pflag"
)

var (
	isULIDFromUUID bool = false
	isUUIDv1       bool = false
	isUUIDv4       bool = false
	isUppercase    bool = false
)

func init() {
	pflag.BoolVarP(&isULIDFromUUID, "from-ulid", "", false, "generate UUID from ULID")
	pflag.BoolVarP(&isUUIDv1, "uuidv1", "", false, "generate UUID v1")
	pflag.BoolVarP(&isUUIDv4, "uuidv4", "", false, "generate UUID v4")
	pflag.BoolVarP(&isUppercase, "uppercase", "u", false, "output uppercase")
}

func main() {
	pflag.Parse()

	if isULIDFromUUID {
		fmt.Printf("%s\n", newUUIDFromULID(isUppercase))
		return
	} else if isUUIDv1 {
		fmt.Printf("%s\n", uuidv1(isUppercase))
	} else {
		fmt.Printf("%s\n", uuidv4(isUppercase))
	}

	return
}

func uuidv1(isUppercase bool) (s string) {
	id, _ := uuid.NewV1()

	s = id.String()
	if isUppercase {
		s = strings.ToUpper(id.String())
	}

	return
}

func uuidv4(isUppercase bool) (s string) {
	id, _ := uuid.NewV4()

	s = id.String()
	if isUppercase {
		s = strings.ToUpper(id.String())
	}

	return
}

func newUUIDFromULID(isUppercase bool) (s string) {
	// Returns 128-bit ULID in 36-character UUID format.
	ulidEntropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	id := uuid.UUID(ulid.MustNew(ulid.Timestamp(time.Now()), ulidEntropy))

	s = id.String()
	if isUppercase {
		s = strings.ToUpper(id.String())
	}

	return
}
