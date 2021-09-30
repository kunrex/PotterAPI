package main

import "HarryPotterAPI/api"

func main() {
	err := api.Start()

	if err != nil { panic(err) }
}
