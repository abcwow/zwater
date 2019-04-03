package main

type SearchPath struct {
	ops []OP
}

func NewSearchPath() *SearchPath {
	var s SearchPath

	return &s
}

func (s *SearchPath) Init() {
}

func (s *SearchPath) Push(cur OP) {
	s.ops = append(s.ops, op)
}

func (s *SearchPath) Pop() {

	n := len(s.ops)
	if n <= 0 {
		return
	}

	s.ops = s.ops[:n-1]
}

func (s *SearchPath) ShowPath() {

	fmt.Println("total steps: ", len(s.ops))
	fmt.Println("!path found --> ")
	for i, op := range s.ops {
		str := fmt.Sprintf("step%04d: \n", op.Description())
	}
	fmt.Println("!path found <-- ")
}
