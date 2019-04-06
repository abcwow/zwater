package main

import "fmt"

type OP struct {
	enumidx int
	cur     EnumVar
	origin  *OPS
}

func NewOp() *OP {
	var op OP

	return &op
}

func (op *OP) Identity() string {

	var ident string = ""

	ident += fmt.Sprintf("op%d", op.enumidx+1)

	ops := op.origin
	var enums string
	for _, v := range ops.cups {
		if v.id == op.cur.to.id {
			enums += fmt.Sprintf("_%d", op.cur.to.current)
		} else if v.id == op.cur.from.id {
			enums += fmt.Sprintf("_%d", op.cur.from.current)
		} else {
			enums += fmt.Sprintf("_%d", v.current)
		}

	}

	ident += enums

	return ident

}

func (op *OP) Description() string {

	str := fmt.Sprintf("op%d ", op.enumidx+1)

	cup2 := &op.cur.from
	cup1 := &op.cur.to

	if cup1.id == cup2.id {
		str += fmt.Sprintf("cup%d to %d", cup1.id+1, cup1.current)
	} else {
		str += fmt.Sprintf("cup%d to %d from cup%d left %d", cup1.id+1, cup1.current, cup2.id+1, cup2.current)
	}

	str += " (identity: " + op.Identity() + " )"

	return str
}

///////////////////////////

const (
	FOUND    int = 400
	NOTFOUND int = 404
)

type OPS struct {
	env *OpsEnv

	z    int
	cups []CUP
	ops  []OP
}

func NewOps() *OPS {

	var ops OPS

	return &ops

}

func (m *OPS) Init(z int, cups []CUP) {

	m.z = z

	m.cups = cups

	var env OpsEnv
	m.env = &env

	m.env.enum = &m_EnumSetting
	m.env.judge = NewJudgeTable()
	m.env.path = NewSearchPath()
}

func (m *OPS) Clone() *OPS {

	var ops OPS

	n := len(m.cups)

	ops.cups = make([]CUP, n, n)
	copy(ops.cups, m.cups)

	ops.z = m.z
	ops.env = m.env

	return &ops
}

func (m *OPS) Do(op *OP) {
	// update to current state of all cups
	for i, v := range m.cups {

		if v.id == op.cur.to.id {
			m.cups[i] = op.cur.to
		}

		if op.cur.to.id == op.cur.from.id { //do on cup itself
			continue
		}

		if v.id == op.cur.from.id {
			m.cups[i] = op.cur.from
		}
	}

}

func (m *OPS) CalcBranches(prev *OP) *OPS {

	//fmt.Println("calcbranch -->")
	//defer fmt.Println("calcbranch <--")

	//fmt.Printf("ops::(len=%d) %v\n", len(m.cups), m.cups)

	d := m.Clone()
	if prev != OpInitial {
		//fmt.Println("ops::do update ", prev)
		d.Do(prev)
	}

	//fmt.Printf("ops::cloned (len=%d) %v\n", len(d.cups), d.cups)

	//fmt.Println("enum::len of enums ", len(m.env.enum.forms))

	fmt.Printf(".")
	for i, opx := range m.env.enum.forms {

		//var count int = 0
		for _, cup1 := range d.cups {
			for _, cup2 := range d.cups {
				before := EnumVar{cup1, cup2}

				//count++
				//fmt.Printf("enum::op%d round %d start %v\n", i+1, count, before)
				after, err := opx.enum(before)
				if err != nil {
					//fmt.Println("enum::err:", err.Error())
					continue
				}
				//fmt.Printf("enum::op%d round %d found %v\n", i+1, count, after)
				//fmt.Println("calcbranch::found one op!")
				op := OP{i, after, d}
				d.ops = append(d.ops, op)

			}

		}

	}

	return d
}

func (m *OPS) CheckEnd(prev *OP) int {

	d := m.Clone()
	if prev != OpInitial {
		d.Do(prev)
	}

	total := 0
	for _, op := range d.cups {
		total += op.current
	}

	if total == d.z {
		return FOUND
	}

	return NOTFOUND

}

//////////////////

type OpsEnv struct {
	enum  *EnumSetting
	judge *JudgeTable
	path  *SearchPath
}

//////////////////////////////////////

var OpInitial *OP = nil

func (m *OPS) NextStep(prev *OP) {

	//next round search
	if prev != OpInitial {
		m.env.path.Push(*prev)
	}

	if m.CheckEnd(prev) == FOUND {
		panic(FOUND)
	}

	d := m.CalcBranches(prev)

	//fmt.Printf("branches::(len=%d)%v\n", len(d.ops), d.ops)

	for _, opx := range d.ops {
		if prev == OpInitial || m.env.judge.Judge(opx, *prev) != REVERSE {
			d.NextStep(&opx)
		}
	}

	if prev != OpInitial {
		m.env.path.Pop() //NOTE: return here means: not found under this opx, so throw it
	}
}

func (m *OPS) RecurseEntry() {

	fmt.Println("begin to search --> ")
	defer func() {

		defer fmt.Println("stoped to search <-- ")
		if r := recover(); r != nil {
			if ret, ok := r.(int); ok && ret == FOUND {
				m.env.path.ShowPath()
				return
			}
		}

		fmt.Println("\n!path not found")
	}()

	m.NextStep(OpInitial)

}
