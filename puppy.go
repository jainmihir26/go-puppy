package puppy

import (
	dog "github.com/jainmihir26/go-dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Bark! Bark! Bark!"
}

func BarkFromDog() string {
	return dog.BarkInDog()
}
