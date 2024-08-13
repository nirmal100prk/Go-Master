package main

import "fmt"

/*
The Strategy Pattern is a behavioral design pattern that enables selecting an
algorithm at runtime. This pattern defines a family of algorithms,
encapsulates each one, and makes them interchangeable. The pattern lets
 the algorithm vary independently from the clients that use it.
*/

func main() {
	s := Spiderman{}
	h := NewHero(&s)
	h.Reciter.Recite()
}

type DailogueReciter interface {
	Recite()
}

type Spiderman struct{}

func (S *Spiderman) Recite() {
	fmt.Println("With great power comes great responsibility")
}

type Superman struct{}

func (s *Superman) Recite() {
	fmt.Println("There is a superhero in all of us" + "We just need the courage to put on the cape")
}

type Batman struct{}

func (b *Batman) Recite() {
	fmt.Println("It is not who I am underneath but what I do that defines me")
}

type Hero struct {
	Reciter DailogueReciter
}

func NewHero(dr DailogueReciter) *Hero {
	return &Hero{Reciter: dr}
}
