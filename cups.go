package main

/////////////////////////////////

type CUP struct {
	id       int
	capacity int
	current  int
}

func NewCup(capacity int) *CUP {

	cup := CUP{0, capacity, 0}

	return &cup

}
