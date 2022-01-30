package tests

import (
	"errors"
	helperpokemon "github.com/iitheo/theopokemon/pkg/app/helper"
	"testing"
)

func TestValidatePokemonIdOrName(t *testing.T) {
	type dataModel struct {
		name     string
		nameOrId string
		err      error
	}
	var tableTestCases = []dataModel{
		{name: "valid_name", nameOrId: "misdreavus", err: nil},
		{name: "valid_id", nameOrId: "200", err: nil},
		{name: "empty_name_or_id", nameOrId: "", err: errors.New("pokemon id or name is required")},
	}

	for _, testCase := range tableTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			myInput := helperpokemon.ValidatePokemonIdOrName(testCase.nameOrId)
			if (myInput != nil) && (myInput.Error() != testCase.err.Error()) {
				t.Errorf("Input: %v, Expected: %v, Outcome: %v", testCase.nameOrId, testCase.err, myInput)
			}
		})

	}
}
