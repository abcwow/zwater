package main

type OP struct {
	enumidx int
	cur     EnumVar
}

///////////////////////////

type OPS struct {
	cups []CUP
	ops  []OP
}

func (m *OPS) NewOps(prev OP) *OPS {

	var ops OPS

	ops.cups = make(CUP, 0, len(m.cups))
	copy(ops.cups, m.cups)

	// update to current state of all cups
	for i, v := range ops.cups {

		if v.id == prev.to.id {
			ops.cups[i] = prev.to
		}

		if v.id == prev.from.id {
			ops.cups[i] = prev.from
		}
	}

	return &ops
}

func (m *OPS) CalcBranches(prev OP) *OPS {

	d := m.NewOps(prev)

	for i, opx := range m_EnumSetting.forms {

		for _, cup1 := range d.cups {
			for _, cup2 := range d.cups {
				before := EnumVar{cup1, cup2}

				after, err := opx.enum(after)
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
