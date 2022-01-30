### What is this repository for? ###

* Pokemon API
* [Learn Markdown](https://bitbucket.org/tutorials/markdowndemo)

### How do I get set up? ###
Requirements:
Ensure you have docker installed in your computer.
Next build your docker image and run the app.
To build your docker image:
Open your Terminal
Clone the app from the github repo.
Navigate to the directory where the app has been cloned or downloaded.
Ensure you are in the same path as the Dockerfile
Next run the command:
docker build -t mypokemon:1 .
This will build the docker image.
Next run the docker container by running the command
docker run -dp 8081:8081 mypokemon:1

Now, you have the app running.

Open your browser or Postman.
localhost:8081/v1/pokemon/get/misdreavus

You will get a response like this or something similar:

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






The following documents are provided:

Postman documentation
https://www.getpostman.com/collections/02ec0f5ec7b0e4833f7b

Start Up Steps


Postman Documentation
6 endpoints are listed together with sample request and response.
The first 5 endpoints provide context on the pokemon api and also
provides sample data to test the MyPokemon App.

The custom endpoint for getting fun description from MyPokemon App is:
GET Fun Pokemon Data
