package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Person struct {
	Age            int
	Name, LastName string
}

func TestParserWithStringInsteadOfInt(t *testing.T) {
	invalidJsonString := `{
		"Age": "50",
		"Name": "Anderson",
		"LastName": "Cooper"
	}`

	var person Person
	assert.Error(t, ParseJsonData(invalidJsonString, &person), "The invalid JSON string, should throw an error")
}

func TestParserWithInvalidJsonInput(t *testing.T) {
	invalidJsonString := `
		"Age": "50",
		"Name": "Anderson",
		"LastName": "Cooper"
	}`

	var person Person
	assert.Error(t, ParseJsonData(invalidJsonString, &person), "The invalid JSON string, should throw an error")
}

func TestParserWithInvalidJsonInput2(t *testing.T) {
	invalidJsonString := `
		"Age": "50",
		"Name": Anderson,
		"LastName": "Cooper"
	}`

	var person Person
	assert.Error(t, ParseJsonData(invalidJsonString, &person), "The invalid JSON string, should throw an error")
}

func TestVerifyOutputOfValidJsonString(t *testing.T) {
	invalidJsonString := `{
		"Age": 50,
		"Name": "Anderson",
		"LastName": "Cooper"
	}`

	var person Person
	assert.Nil(t, ParseJsonData(invalidJsonString, &person), "The valid json string should return a Nil.")
	assert.Equal(t, 50, person.Age, "The age is not the same as the expected age")
	assert.Equal(t, "Anderson", person.Name, "The name is not the same as the expected name")
	assert.Equal(t, "Cooper", person.LastName, "The last name is not the same as the expected last name")
}

func TestInvalidStructureToString(t *testing.T) {
	str, err := DataToString(nil)

	assert.NotEqual(t, "{}", str, "The output of Structure -> String on an invalid structure should not be {}")
	assert.NotNil(t, err, "Error should not be nil on a invalid structure")
	assert.Error(t, err, "A error should be returned on a invalid structure")
}

func TestValidStructureToString(t *testing.T) {
	person := Person{50, "Anderson", "Cooper"}

	str, err := DataToString(&person)

	assert.Equal(t, `{"Age":50,"Name":"Anderson","LastName":"Cooper"}`, str, "The output of the Structure -> string is not accurate")
	assert.Nil(t, err, "Error should not be nil on a valid structure")
}

type JsonWithArray struct {
	name string
	Age  int
	cars []string
}

func TestVerifyOutputOfValidJsonStringWithArray(t *testing.T) {
	invalidJsonString := `{
		"name":"John",
		"Age":30,
		"cars":[ "Ford", "BMW", "Fiat" ]
		}`

	var data JsonWithArray
	assert.Nil(t, ParseJsonData(invalidJsonString, &data), "The valid json string should return a Nil.")
	assert.Equal(t, 30, data.Age, "The age is not the same as the expected age")
}
