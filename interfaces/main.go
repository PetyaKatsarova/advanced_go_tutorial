package main

import "fmt"

type SoundMaker interface{ MakeSound() string }

type Dog struct{}
func (d Dog) MakeSound() string {
	return "Woof!"
}

type Cat struct{}
func (c Cat) MakeSound() string { return "Weow!" }

type Cow struct{}
func (c Cow) MakeSound() string { return "Moo!" }

func playSound(sm SoundMaker) {
	fmt.Println(sm.MakeSound())
}

func main() {
	dog := Dog{}
	cat := Cat{}
	cow := Cow{}

	playSound(dog)
	playSound(cat)
	playSound(cow)
}