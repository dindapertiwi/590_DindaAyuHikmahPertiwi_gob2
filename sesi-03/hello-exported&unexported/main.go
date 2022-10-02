package main

import "hello-exported&unexported/helpers"

func main()  {

	helpers.Greet()

	var person = helpers.Person{}

	person.Invokegreet()
}