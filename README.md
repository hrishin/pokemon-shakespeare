# Pokemon Shakespeare

Its a Rest API service to describe a given [pokemon's](https://en.wikipedia.org/wiki/Pok%C3%A9mon) characteristics in [William Shakeperar's](https://en.wikipedia.org/wiki/William_Shakespeare) words. 

The service is written in [Go](https://golang.org) and provides the following endpoint(s).

1) GET `/pokemon/<pokemon name>` :
* Path Varibles:
    - pokemon name : `string` : name of a pokemon
* Response:
    200 OK, content-type: application/json, response body as follow
    ```
    {
        name: <pokmon name>
        description: <dscription in Shakespeares words>
    }
    ```
* Example:
    ```
    ➜ http http://192.168.99.104:32300/pokemon/charizard       
    HTTP/1.1 200 OK
    Content-Length: 244
    Content-Type: application/json
    Date: Sun, 29 Nov 2020 17:36:23 GMT

    {
        "description": "Charizard flies 'round the sky in search of powerful opponents. 't breathes fire of such most wondrous heat yond 't melts aught. However,  't nev'r turns its fiery breath on any opponent weaker than itself.",
        "name": "charizard"
    }

    ```

# Getting Started

## Prerequisites
* Install the [Go](https://golang.org/doc/install)
* GNU [Make](https://www.gnu.org/software/make/)
* Docker (optional, but must be installed to build the container image)

## Building

### Build OSX Binary
```
make bin/pokemon-darwin-amd64
```

### Build Linux Binary
```
make bin/pokemon-linux-amd64
```

## Running Locally

* once you build the binary, you can run the program using the following command:

For OSX
```
./bin/pokemon-darwin-amd64
```

For Linux

```
./bin/pokemon-linux-amd64
```

* if you want to run it as a container,

```
docker run -d -p 5000:5000 quay.io/hriships/pokemon:v1
```

Execute HTTP GET to `http://localhost:5000/pokemon/pikachu`
```
curl --request GET http://localhost:5000/pokemon/pikachu   

{"name":"pikachu","description":"Whenever pikachu cometh across something new,  't blasts 't with a jolt of electricity. If 't be true thee cometh across a blackened berry,  't’s evidence yond this pokémon did misprision the intensity of its charge."}
```

## Testing

### Run Unit Tests
```
make unit-tests
```

### Run Integration Test
```
make integration-test
```

## Implementation Details
 * To handler HTTP request it uses [gorrila/mux](https://github.com/gorilla/mux) HTTP router for `/pokemon/<name>` endpoint. It simplifies the API testing and extract path variable names.
 * Program implements the wrapper on [PokeAPI](https://pokeapi.co/docs/v2) to fetch the normal pokemon’s description. Though [Pokeapi](https://pokeapi.co/docs/v2) does mention the go client library [pokeapi-go](https://github.com/mtslzr/pokeapi-go), however it has somes [issue](https://github.com/mtslzr/pokeapi-go/issues/29) returning the correct response in some case.
 * Upon fetching the pokemon description, another wrapper written on [funtranslations](https://funtranslations.com/api/shakespeare) gets the
 pokemon description in the Shakespeares words style.
 * Some utility packages are written to mock the HTTP response behaviour
 * Use [ginkgo](https://github.com/onsi/ginkgo) and [omega](https://github.com/onsi/gomega) for the integration test and assertions respectively.

## Improvements
- Switch to Pokeapi client-go:
 https://github.com/mtslzr/pokeapi-go once its resolved https://github.com/mtslzr/pokeapi-go/issues/29

- Caching: 
    * Given the rate limit and limited API quota(paid subscription) of the [funtranslations](https://funtranslations.com/api/shakespeare) API, it would be better to implement the server-side cache to store previously obtained translations. Hence the overall resiliency of the API could be improved to the better extent and overall API response time can be reduced as well.

    * Switch to [Pokeapi](https://github.com/mtslzr/pokeapi-go) go client. Given the time, it's better to resolve [issue](https://github.com/mtslzr/pokeapi-go/issues/29) in the upstream or maintain the fork. This go-client fairly implements the cache to store the previous endpoint requests. This could in turn, help to reduce the overall API response time.

    * On the trade cache invalidation may bring other problems. Despite this, if service is expected to sustain SLO/SLI then it's a better problem to solve.

- Consistent endpoint test behaviour: 
    * Right now test has a potential to fail due to [funtranslations](https://funtranslations.com/api/shakespeare) API rate limit. It’s challenging to produce consistent test behaviour without subscription key support. Otherwise, the endpoint test needs to mock behaviour given ample time.

- Metrics:
    * It would be nice to capture and expose essential metrics that could help in understand the behaviourß about API response time and error counts from
    the external API integrations of pokeapi and funtranslation


