package main

import "fmt"

type JudgeItem struct {
	enumidx int //current op idx
	op      OP  //prev op info
}

func (d *JudgeItem) ToString() string {

	opident := d.op.Identity()

	str := fmt.Sprintf("op%d", d.enumidx+1) + "_" + opident

	return str

}

/////////////////////////////////////

const (
	REVERSE int = 0
	NEXT    int = 1
)

type JudgeTable struct {
	ops map[string]bool
}

func NewJudgeTable() *JudgeTable {
	var table JudgeTable
	table.ops = make(map[string]bool)
	return &table
}

func (m *JudgeTable) Add(item *JudgeItem) {

	ident := item.ToString()
	m.ops[ident] = true
}

func (m *JudgeTable) Find(item *JudgeItem) bool {

	ident := item.ToString()
	if _, ok := m.ops[ident]; ok == false {
		return false
	}

	return true

}

func (m *JudgeTable) Judge(cur, prev OP) int {
	item := JudgeItem{cur.enumidx, prev}

	if m.Find(&item) == true {
		return REVERSE
	}

	m.Add(&item)

	return NEXT

}
