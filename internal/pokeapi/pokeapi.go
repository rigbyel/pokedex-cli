package pokeapi

import (
	"net/http"
	"time"

	"github.com/rigbyel/pokedex-cli/internal/pokecache"
)

type HttpClient struct {
	httpClient http.Client
	cache pokecache.Cache
}

func NewHttpClient(interval time.Duration) HttpClient {
	return HttpClient{
		httpClient: http.Client { 
			Timeout: 5*time.Second * time.Duration(1),
		},
		cache: *pokecache.NewCache(interval),
	}
}

var baseUrl = "https://pokeapi.co/api/v2"

