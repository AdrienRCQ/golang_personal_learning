package main

import "fmt"

func main() {
	todos := Todos{}
	todos.add("Buy Milk")
	todos.add("Buy Avocado")
	fmt.Printf("%+v\n\n", todos)
	todos.delete(0)
	fmt.Printf("%+v\n\n", todos)
}