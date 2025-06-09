package main

import "fmt"

func main() {

	print("Nader")

}

func print(Nader string) string {

	massage := fmt.Sprint("Haloooooooooooooooooooo-%v" + Nader)
	return massage

}
