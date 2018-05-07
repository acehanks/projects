# nextAI

This project consists of two main modules; randimage.go and imageapi.go

randimage.go makes a request to the unsplash api and reterives information including image urls. The image is downloaded to the local machine and added to the sqlite database. Each time randimage.go is run, it retrieves a random image.

imageapi.go is an api server built on negroni, gorilla mux. It runs on port 8080 and has endpoints "/", "/all" and "/img/{id}".
