package main

/////////////////////////////////

type CUP struct {
	id       int
	capacity int
	current  int
}

func NewCup(id, capacity int) *CUP {

	cup := CUP{id, capacity, 0}

	return &cup

}
