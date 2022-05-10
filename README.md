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

Here you will use the intra to extract the epitech calendar information

get your autologin link: 
https://intra.epitech.eu/admin/autolog

now in exercise3 handler, extract your autologin link from a POST request to localhost:8000/calendar
{
	"key": "{Your Key}" 
}

Then, using http library, make a get request to ("{autologin link}/module/board/?format=json&start=2022-05-09&end=2022-05-16")

and print the result