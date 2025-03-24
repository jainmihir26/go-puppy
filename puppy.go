package puppy

import (
	dog "github.com/jainmihir26/go-dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BarkFromDog() string {
	return dog.BarkInDog()
}
