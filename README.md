## Pokemon Shakespeare

Its a Rest API service to describe a given pokemon's characteristics in William Shakespeare's words.
The service is written in Go and provides the following endpoint(s).

1. GET /pokemon/<pokemon name> :
    * Path Variables:
        - pokemon name: string: Name of a pokemon

    * Response: 200 OK, content-type: application/json, response body as follow:
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
        "description": "Charizard flies 'round the sky in search of powerful opponents. 't breathes fire of such most wondrous heat yond 't melts aught. However, 't nev'r turns its fiery breath on any opponent weaker than itself.",
        "name": "charizard"
    }
    ```

Getting Started
Prerequisites
Install the Go
GNU Make
Docker (optional, but must be installed to build the container image)
Building
Build OSX Binary
make bin/pokemon-darwin-amd64
Build Linux Binary
make bin/pokemon-linux-amd64
Running Locally
once you build the binary, you can run the program using the following command:
For OSX
./bin/pokemon-darwin-amd64
For Linux
./bin/pokemon-linux-amd64
if you want to run it as a container,
docker run -d -p 5000:5000 quay.io/hriships/pokemon:v1
Execute HTTP GET to http://localhost:5000/pokemon/pikachu
curl --request GET http://localhost:5000/pokemon/pikachu   

{"name": "pikachu", "description": "Whenever pikachu cometh across something new, 't blasts 't with a jolt of electricity. If 't be true thee cometh across a blackened berry, 't's evidence yond this pokémon did misprision the intensity of its charge."}
Testing
Run Unit Tests
make unit-tests
Implementation Details
To handler HTTP request it uses gorilla/mux HTTP router for /pokemon/<name> endpoint. It simplifies the API testing and extracts path variable names.
The program implements the wrapper on PokeAPI to fetch the normal pokemon's description. Though Pokeapi does mention the go client library pokeapi-go, however, it has some issue returning the correct response in some case.
Upon fetching the pokemon description, another wrapper written on funtranslations gets the pokemon's description in the Shakespeares words style.
The http-mock package is to mock the HTTP response in unit tests.
Improvements
Switch to Pokeapi client-go: https://github.com/mtslzr/pokeapi-go once its resolved https://github.com/mtslzr/pokeapi-go/issues/29
Caching:
Given the rate limit and limited API quota(paid subscription) of the funtranslations API, it would be better to implement the server side cache to store previously obtained translations. Hence there is room to improve the overall resiliency of the API to a better extent and overall API response time could be reduced as well.
Switch to Pokeapi go client. Given the time, it's better to resolve an issue in the upstream or maintain the fork. This go-client fairly implements the cache to store the previous endpoint requests. This in turn, could help to reduce the overall API response time.
On the trade, cache invalidation may bring other problems. Despite this, if service is expected to sustain SLO's/SLI's then it's a better problem to solve.
Consistent endpoint test behaviour:
Right now test has the potential to fail due to funtranslations API rate limit. It's challenging to produce consistent test behaviour without subscription key support. Otherwise, the endpoint test needs to mock action given ample time.
