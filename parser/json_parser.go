package parser

import (
	"encoding/json"
	"errors"
	"strings"
)

// ParseJsonData fills in a structure based on a json formatted string
// jsonData is the JSON formatted string
// jsonStructure is the interface/struct you want the data to be filled into
// err is nil if there wasn't an error parsing it into the struct, else it returns the error.
func ParseJsonData(jsonData string, jsonStructure interface{}) error {
	// check and make sure the json is valid
	if !json.Valid([]byte(jsonData)) {
		return errors.New("Invalid JSON format.")
	}

	// create decoder for parsing over the json string
	decoder := json.NewDecoder(strings.NewReader(jsonData))

	// decode the the jsonData into the json structure
	err := decoder.Decode(&jsonStructure)

	// if we got an error during the decoding process, handle it.
	if err != nil {
		return err
	}

	// no errors parsing the json
	return nil
}

// ParsePolicyData fills in a structure based on a json formatted string
// jsonData is the JSON formatted string
// jsonStructure is the interface/struct you want the data to be filled into
// err is nil if there wasn't an error parsing it into the struct, else it returns the error.
func ParsePolicyData(jsonData string, jsonStructure interface{}) error {
	return ParseJsonData(jsonData, jsonStructure)
}

// ParseXattrData fills in a structure based on a json formatted string
// jsonData is the JSON formatted string
// jsonStructure is the interface/struct you want the data to be filled into
// err is nil if there wasn't an error parsing it into the struct, else it returns the error.
func ParseXattrData(jsonData string, jsonStructure interface{}) (err error) {
	return ParseJsonData(jsonData, jsonStructure)
}

// DataToString takes some interface/struct and turns it into a json formatted string
// structure is the struct/interface to get the key: value from
// str is the json formatted string
// err is the error if something with wrong, if nothing went wrong it's nil
func DataToString(structure interface{}) (str string, err error) {
	// check that the structure is not a nil
	if structure == nil {
		return "", errors.New("Structure input should not be nil")
	}

	// try to parse the structure into a json formatted string
	data, err := json.Marshal(structure)

	// convert []byte -> string
	str = string(data)

	// if the str is "null" it means that it failed to convert it into a json formatted string
	if str == "null" {
		return "", errors.New("Invalid structure input")
	}

	// check if it's a valid json string
	if !json.Valid(data) {
		return "", errors.New("Output data didn't become a valid JSON formatted string")
	}

	// we were able to convert it into a proper json formatted string
	return str, err
}

func main() {

}
