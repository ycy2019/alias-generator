package alias

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//go:embed nouns.json
var nouns []byte

//go:embed adjectives.json
var adjectives []byte

var nounsJson, adjectivesJson []string

func init() {
	nounsJson, _ = readJsonFile(nouns)
	adjectivesJson, _ = readJsonFile(adjectives)
}

func readJsonFile(b []byte) ([]string, error) {
	var jsonData []string

	err := json.Unmarshal(b, &jsonData)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return nil, err
	}

	return jsonData, nil
}

func GenerateUsername(separator string) string {
	rand.Seed(time.Now().UnixNano())

	nounIndex := rand.Intn(len(nounsJson))
	adjectiveIndex := rand.Intn(len(adjectivesJson))

	noun := nounsJson[nounIndex]
	adjective := adjectivesJson[adjectiveIndex]

	result := fmt.Sprintf("%s%s%s", strings.Title(adjective), separator, strings.Title(noun))

	return result
}
