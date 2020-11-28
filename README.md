# Pokemon Shakespeare

A Rest API service to describe a [pokemon's](https://en.wikipedia.org/wiki/Pok%C3%A9mon) characteristics in [William Shakeperar's](https://en.wikipedia.org/wiki/William_Shakespeare) words. 

the service is written in `golang` and expose the following endpoint(s).

1) GET `/pokemon/<pokemon name>` :
    * Path Varibles:
        - pokemon name : `string` : name of a pokemon
    * Response:
        
        200 OK, content-type: application/json, response body as follow
        ```
        {
            name: <pokmon name>
            description: <dscription in Sir Shakespeares words>
        }
        ```
