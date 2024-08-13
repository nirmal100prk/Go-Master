package main

import (
	"fmt"
	"go-datastructures/Generics/mapper"
)

func main() {
	// S := Stack[int]{}

	// S.Push(1)
	// S.Push(2)
	// fmt.Println(S)
	// S.Pop()
	// fmt.Println(S)
	p := Profile{Name: "nirmal", Age: 24}
	mapper.Configure(&mapper.MapperConfig{
		MapUnexportedFields: false,
	})

	err := mapper.CreateMap[*Profile, *NewProfile]()
	if err != nil {
		fmt.Println(err)
	}
	np, err := mapper.Map[*NewProfile](p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(np)
}

type Profile struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

type NewProfile struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() {
	s.items = s.items[:len(s.items)-1]
}

//////////////////////////////////////////

// func main() {
// 	fmt.Println(Add(1, 2))
// 	fmt.Println(Add(1.2, 2.8))
// }

// type Adder interface {
// 	int | float64
// }

// func Add[T Adder](a, b T) T {
// 	return a + b
// }
