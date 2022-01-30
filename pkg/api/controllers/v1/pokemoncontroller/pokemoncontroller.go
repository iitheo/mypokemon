package pokemoncontroller

import (
	"fmt"
	"github.com/iitheo/theopokemon/pkg/api/libs/httplibrs"
	"github.com/iitheo/theopokemon/pkg/app/config/httpresppokemon"
	helperpokemon "github.com/iitheo/theopokemon/pkg/app/helper"
	"github.com/iitheo/theopokemon/pkg/app/services/pokemonservices"
	"net/http"
)

func GetPokemonSpecies(res http.ResponseWriter, req *http.Request) {
	c := httplibrs.C{Req: req, Res: res}
	var (
		resp httpresppokemon.HttpResponse
	)
	pokemonIdOrName := c.Params("idOrName")
	err := helperpokemon.ValidatePokemonIdOrName(pokemonIdOrName)
	if err != nil {
		resp.Success = false
		resp.Message = fmt.Sprintf("error validating pokemon id or name: %s", err.Error())
		resp.Error = fmt.Sprintf("error validating pokemon id or name: %s", err.Error())
		httplibrs.Response400(res, resp)
		return
	}

	result, err := pokemonservices.FetchPokemonSpecies(pokemonIdOrName)
	if err == nil {
		resp.Success = true
		resp.Data = result
		resp.Message = fmt.Sprintf("data successfully fetched")
		httplibrs.Response(res, resp)
	} else {
		resp.Success = false
		resp.Error = err.Error()
		resp.Message = err.Error()
		httplibrs.Response500(res, resp)
	}
}
