package pokemoncontroller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/iitheo/theopokemon/pkg/app/config/dbconfigs"
	"github.com/iitheo/theopokemon/pkg/app/models/pokemonmodels"
	"github.com/iitheo/theopokemon/pkg/app/services/pokemonservices"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func FetchPokemonSpecies(IdOrName string) (pokemonData pokemonmodels.PokemonSpeciesResponseVM, myErr error) {
	var (
		resultPokemonData              pokemonmodels.PokemonSpeciesResponseVM
		fetchedPokemonData             pokemonmodels.PokemonSpeciesResponse
		translatedDescriptionFromRedis = pokemonmodels.TranslateResponseVM{}
		redisKeyFunTranslation         = ""
		translationType                string
	)

	baseUrlPokemon := "https://pokeapi.co/api/v2"

	getUrlPokemon := fmt.Sprintf("%s/%s/%s", baseUrlPokemon, "pokemon-species", IdOrName)

	redisConn, err := dbconfigs.RedisConn()
	defer redisConn.Close()
	if err != nil {
		return pokemonData, err
	}

	request, err := http.NewRequest("GET", getUrlPokemon, nil)
	if err != nil {
		return pokemonData, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	client := &http.Client{}

	// Do sends an HTTP request and
	response, err := client.Do(request)
	if err != nil {
		return pokemonData, err
	}
	if response.StatusCode != http.StatusOK {
		return pokemonData, errors.New(fmt.Sprintf("error fetching data, status - %d", response.StatusCode))
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return pokemonData, err
	}

	err = json.Unmarshal(body, &fetchedPokemonData)

	if err != nil {
		return pokemonData, err
	}

	//check if isLegendary or habitat is 'cave'
	if fetchedPokemonData.IsLegendary || strings.TrimSpace(strings.ToLower(fetchedPokemonData.Habitat.Name)) == "cave" {
		redisKeyFunTranslation = fmt.Sprintf("yoda-%s", IdOrName)
		translationType = "yoda"
	} else {
		redisKeyFunTranslation = fmt.Sprintf("sp-%s", IdOrName)
		translationType = "shakespeare"
	}

	//check if fun translation data exists in redis
	translatedString, err := redis.String(redisConn.Do("GET", redisKeyFunTranslation))
	if err == nil {
		b := []byte(translatedString)
		err = json.Unmarshal(b, &translatedDescriptionFromRedis)
		if err != nil {
			return pokemonData, err
		}
		if len(fetchedPokemonData.FlavorTextEntries) > 0 {
			resultPokemonData.IsLegendary = fetchedPokemonData.IsLegendary
			resultPokemonData.Habitat = fetchedPokemonData.Habitat.Name
			resultPokemonData.Name = fetchedPokemonData.Name
			resultPokemonData.Description = fetchedPokemonData.FlavorTextEntries[0].FlavorText
			resultPokemonData.TranslatedDescription = translatedDescriptionFromRedis.Contents.Translated
		}
		return resultPokemonData, nil
	}

	translatedInfo, err := pokemonservices.TranslateDescription(fetchedPokemonData.FlavorTextEntries[0].FlavorText, translationType)
	resultPokemonData.IsLegendary = fetchedPokemonData.IsLegendary
	resultPokemonData.Habitat = fetchedPokemonData.Habitat.Name
	resultPokemonData.Name = fetchedPokemonData.Name
	if err != nil {
		log.Printf("error occurred translating description: %s", err.Error())
		if len(fetchedPokemonData.FlavorTextEntries) > 0 {
			resultPokemonData.Description = fetchedPokemonData.FlavorTextEntries[0].FlavorText
			resultPokemonData.TranslatedDescription = fetchedPokemonData.FlavorTextEntries[0].FlavorText
		}
	} else {
		if len(fetchedPokemonData.FlavorTextEntries) > 0 {
			resultPokemonData.Description = fetchedPokemonData.FlavorTextEntries[0].FlavorText
			resultPokemonData.TranslatedDescription = translatedInfo.Contents.Translated
		}

		//save fun translation data in redis
		b, err := json.Marshal(&translatedInfo)
		if err == nil {
			_, err = redisConn.Do("SET", redisKeyFunTranslation, string(b))
			if err != nil {
				log.Printf("error saving data to redis: %s\n", err.Error())
			}
		}
	}

	return resultPokemonData, nil
}
