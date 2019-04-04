package main

import "fmt"

func main() {

	z, cups := CheckInput()

	ops := NewOps()
	ops.Init(z, cups)
	ops.RecurseEntry()
}

func CheckInput() (z int, cups []CUP) {

	var (
		n     int = 0
		c     []int
		total int = 0
	)

	for {
		fmt.Println("Please input the z value: ")
		fmt.Scanln(&z)
		if z <= 0 {
			fmt.Println("z should > 0!input again!")
			continue
		}
		break
	}
	for {
		fmt.Println("Please input n of cups: ")
		fmt.Scanln(&n)
		if n <= 0 || n > 5 {
			fmt.Println("only 1~5 cups allowed! input again!")
			continue
		}
		break
	}

	c = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Please input capacity of cup %d:\n", i+1)
		fmt.Scanln(&c[i])
		if c[i] < 0 || c[i] >= 1000 {
			fmt.Println("capacity of each cup is 0 < 0 < 1000! input again!")
			i--
			continue
		}

		total += c[i]
		if i == n-1 {
			if total < 0 || total < z {
				fmt.Println("z if greater than sum of capacity of all cups! input again from cup1!")
				i = -1
				continue
			}
		}
	}

	cups = make([]CUP, 0, n)
	for i := 0; i < n; i++ {

		cup := NewCup(c[i])
		cups = append(cups, *cup)
	}

	return

}
