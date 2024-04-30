package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type InputJSON struct {
	Number1 json.RawMessage `json:"number_1"`
	String1 json.RawMessage `json:"string_1"`
	String2 json.RawMessage `json:"string_2"`
	Map1    json.RawMessage `json:"map_1"`
}

type OutputJSON struct {
	Map1 struct {
		List1 []interface{} `json:"list_1"`
		Null1 interface{}   `json:"null_1"`
	} `json:"map_1"`
	Number1 float64 `json:"number_1"`
	String1 string  `json:"string_1"`
	String2 int64   `json:"string_2"`
}

func transformString(input json.RawMessage) (string, error) {
	var raw map[string]string
	if err := json.Unmarshal(input, &raw); err != nil {
		return "", err
	}
	if value, ok := raw["S"]; ok {
		trimmedValue := strings.TrimSpace(value)
		if trimmedValue == "" {
			return "", fmt.Errorf("empty string")
		}
		// Check if RFC3339 and convert to Unix timestamp if true
		if t, err := time.Parse(time.RFC3339, trimmedValue); err == nil {
			return strconv.FormatInt(t.Unix(), 10), nil
		}
		return trimmedValue, nil
	}
	return "", fmt.Errorf("invalid string field")
}

func transformNumber(input json.RawMessage) (float64, error) {
	var raw map[string]string
	if err := json.Unmarshal(input, &raw); err != nil {
		return 0, err
	}
	if value, ok := raw["N"]; ok {
		if trimmedValue := strings.TrimSpace(value); trimmedValue != "" {
			return strconv.ParseFloat(trimmedValue, 64)
		}
	}
	return 0, fmt.Errorf("invalid number field")
}

func main() {
	// This is a placeholder for where you would actually read your input JSON.
	inputJSON := []byte(`
    {
        "number_1": {"N": "1.50"},
        "string_1": {"S": "784498 "},
        "string_2": {"S": "2014-07-16T20:55:46Z"},
        "map_1": {"M": {
            "bool_1": {"BOOL": "true"},
            "null_1": {"NULL ": "true"},
            "list_1": {"L": [
                {"S": "  "},
                {"N": "011"},
                {"N": "5215"},
                {"BOOL": "false"},
                {"NULL": "0"}
            ]}
        }}
    }
    `)

	var input InputJSON
	if err := json.Unmarshal(inputJSON, &input); err != nil {
		log.Fatal(err)
	}

	var output OutputJSON

	// Transformation logic here
	// Example transformation for string
	if s, err := transformString(input.String1); err == nil {
		output.String1 = s
	} else {
		log.Println("Error transforming string_1:", err)
	}

	// Convert and output the final JSON
	outputJSON, err := json.Marshal(output)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(outputJSON))
}

