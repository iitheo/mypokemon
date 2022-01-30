package helperpokemon

import "errors"

func ValidatePokemonIdOrName(IdOrName string) error {
	if IdOrName == "" {
		return errors.New("pokemon id or name is required")
	}

	return nil
}
