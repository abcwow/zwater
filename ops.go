package main

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

	ident += fmt.Sprintf("op%d", op.enumidx)

	ops := op.origin
	var enums string
	for i, v := range ops.cups {
		if v.id == op.to.id {
			enums += fmt.Sprintf("_%d", op.to.current)
		} else if v.id == op.from.id {
			enums += fmt.Sprintf("_%d", op.from.current)
		} else {
			enums += fmt.Sprintf("_%d", v.cups[i].current)
		}

	}

	ident += enums

	return ident

}

func (op *OP) Description() string {

	str := fmt.Sprintf("op%d ", op.enumidx)

	cup1 := &op.cur.from
	cup2 := &op.cur.to

	if cup1.id == cup2.id {
		str += fmt.Sprintf("cup%d to %d", cup1.id, cup1.current)
	} else {
		str += fmt.Sprintf("cup%d to %d from cup%d", cup1.id, cup1.current, cup2.id, cup2.current)
	}

	str += " (identity: " + op.Identity() + " )"
}

///////////////////////////

const (
	FOUND    int = 0
	NOTFOUND int = 1
)

type OPS struct {
	env *OpsEnv

	z    int
	cups []CUP
	ops  []OP
}

func NewOps() *OPS {

	var ops OPS

	ops.cups = make(CUP, ncups)

	return &ops

}

func (m *OPS) Init(z int, cups []CUP) {

	m.z = z

	m.cups = cups

	var env OpsEnv
	m.env = &env

	m.enums = m_EnumSetting
	m.judge = NewJudgeTable()
	m.path = NewSearchPath()
}

func (m *OPS) Clone() *OPS {

	var ops OPS

	ops.cups = make(CUP, 0, len(m.cups))
	copy(ops.cups, m.cups)

	ops.z = m.z
	ops.env = m.env

	return &ops
}

func (m *OPS) Do(op *OP) {
	// update to current state of all cups
	for i, v := range m.cups {

		if v.id == op.to.id {
			m.cups[i] = op.to
		}

		if op.to.id == op.to.from { //do on cup itself
			continue
		}

		if v.id == prev.from.id {
			m.cups[i] = op.from
		}
	}

}

func (m *OPS) CalcBranches(prev *OP) *OPS {

	d := m.Clone()
	d.Do(prev)

	for i, opx := range m.env.forms {

		for _, cup1 := range d.cups {
			for _, cup2 := range d.cups {
				before := EnumVar{cup1, cup2, &d}

				after, err := opx.enum(before)
				if err != nil {
					continue
				}
				op := OP{i, after}
				d.ops = append(d.ops, op)

			}

		}

	}

	return d
}

func (m *OPS) CheckEnd() {

	total := 0
	for i, op := range m.cups {
		total += op.cups[i].current
	}

	if total == m.z {
		return FOUND
	}

	return NOTFOUND

}

//////////////////

type OpsEnv struct {
	enums *EnumSetting
	judge *JudgeTable
	path  *SearchPath
}

//////////////////////////////////////

func (m *OPS) NextStep() {
	m_SearchPath
	m_JudgeTable

}
