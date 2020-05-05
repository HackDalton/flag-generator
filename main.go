package main

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"strings"
)

const prefix = "hackDalton"

var vowelMap = map[string]string{
	"a": "4",
	"e": "3",
	"i": "1",
	"o": "0",
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		panic(errors.New("Invalid arguments"))
	}

	description := args[0]

	fmt.Println(GenerateFlag(description))
}

// GenerateFlag generates a flag with the given description
func GenerateFlag(description string) string {
	description = strings.ToLower(description)
	description = strings.ReplaceAll(description, " ", "_")

	for oldChar, newChar := range vowelMap {
		description = strings.ReplaceAll(description, oldChar, newChar)
	}

	str, err := generateRandomString(10)
	if err != nil {
		panic(err)
	}

	return prefix + "{" + description + "_" + str + "}"
}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func generateRandomString(s int) (string, error) {
	b, err := generateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b)[:s], err
}
