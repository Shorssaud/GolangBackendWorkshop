# GolangBackendWorkshop
A small workshop in Golang

# Exercise 1
# Step 1

Install Go and a http request software ( Insomnia or Postman )

# Step 2

Try using
```bash
go run exercise1.go
```
then using insomnia or postman to make a go request to http://localhost:8000

# Step 3

In Exercise1.go, change the "wdym" function so that when we call the url "http://localhost:8000/wdym?name=Steven" the program prints "Hello Steven!"

# Exercise 2

Congrats! You have officially created your first handler! Now we are going to use request bodies to pass information more discreetly to our program

In Exercise2.go use the encoding/json library to handle the JSON post request:
{
	"name": "Steven",
	"lastname": "Universe",
	"dateofbirth": "August 14"
}

and print "Hello Steven Universe born August 14th"

# Exercise 3

Here you will use the weather api to get the weather

Then, using http library, make a get request to ("https://api.open-meteo.com/v1/forecast?latitude=52.52&longitude=13.41&current_weather=true&hourly=temperature_2m,relativehumidity_2m,windspeed_10m")

and print the result

# Ending

You now have the ability to get information from any api using golang. Try using it for other apps like steam, twitter, or even yugioh! (ALL of this can and will be useful for those who haven't done the area yet.
