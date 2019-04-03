package main

type OP struct {
	enumidx int
	cur     EnumVar
	origin  *OPS
}

func (op *OP) Identity() string {

	var ident string = ""

	ident += fmt.Sprintf("op%d", op.opx)

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

///////////////////////////

type OPS struct {
	cups []CUP
	ops  []OP
}

func (m *OPS) Clone() *OPS {

	var ops OPS

	ops.cups = make(CUP, 0, len(m.cups))
	copy(ops.cups, m.cups)

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

func (m *OPS) CalcBranches(prev OP) *OPS {

	d := m.Clone()
	d.Do(prev)

	for i, opx := range m_EnumSetting.forms {

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
