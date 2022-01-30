package httproutespokemon

import (
	"github.com/gorilla/mux"
	"github.com/iitheo/theopokemon/pkg/api/controllers/v1/pokemoncontroller"
	middlewarespokemon "github.com/iitheo/theopokemon/pkg/api/middlewares"
	"github.com/urfave/negroni"
)

func Router() *negroni.Negroni {
	route := mux.NewRouter()

	n := negroni.Classic()
	n.Use(middlewarespokemon.Cors())
	n.UseHandler(route)

	//BASE ROUTE
	//route.HandleFunc("/v1", homeHandler)

	//*****************
	// FILMS ROUTES
	//*****************
	pokemonRoute := route.PathPrefix("/v1/pokemon").Subrouter()
	pokemonRoute.HandleFunc("/get/{idOrName}", pokemoncontroller.GetPokemonSpecies).Methods("GET")

	return n
}
