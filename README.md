### What is this repository for? ###

* Pokemon API
* [Learn Markdown](https://bitbucket.org/tutorials/markdowndemo)

You can quickly test the app with the url:
https://theomypokemon.herokuapp.com/v1/pokemon/get/1

# How do I get set up? #
## Requirements: ##
1. Ensure you have docker installed on your computer.

2. Start up docker on your computer.

## Steps to Start App ##
### Next build your docker image. ###

To build your docker image:
1. Open your Terminal. Clone the app from the github repo by running the command below.
2. git clone https://github.com/iitheo/mypokemon.git
3. Navigate to the root path of the app. This is the same path as the Dockerfile. Run ls in your terminal and verify you can see the Dockerfile
4. Build the docker image by running the command below:
5. docker build -t mypokemonapp:1 .
6. This will build the docker image.
7. Next run the docker container by running the command below:
8. docker run -dp 8081:8081 mypokemonapp:1
9. Now, you have the app running.
10. Open your browser or Postman.
11. type this localhost:8081/v1/pokemon/get/misdreavus
12. You will get a response like this or something similar:

{
"data": {
"name": "misdreavus",
"description": "It likes playing\nmischievous tricks\nsuch as screaming\fand wailing to\nstartle people at\nnight.",
"habitat": "cave",
"isLegendary": false,
"translatedDescription": "Playing mischievous tricks such as screaming and wailing to startle people at night,  it likes."
},
"message": "data successfully fetched",
"status": true
}


CONGRATULATIONS!!!

### The following documents are provided: ###

### Postman documentation ###
https://www.getpostman.com/collections/02ec0f5ec7b0e4833f7b

Postman Documentation
6 endpoints are listed together with sample request and response.

The first 5 endpoints provide context on the pokemon api and also
provides sample data to test the MyPokemon App.

The custom endpoint for getting fun description from MyPokemon App is:
GET Fun Pokemon Data
