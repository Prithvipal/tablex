package tablex

import (
	"fmt"
	"strings"
)

const (
	PADDING = 3
)

type Tablex struct {
	columns []string
	width   []int
	rows    [][]any
}

func NewTablex(columns []string) *Tablex {
	t := &Tablex{
		columns: columns,
		width:   make([]int, len(columns)),
	}
	for idx, col := range columns {
		t.width[idx] = len(col) + PADDING
	}

	return t
}

func (t *Tablex) AddRow(row []any) {
	for idx, ele := range row {
		eleS := ele.(string)
		if len(string(eleS))+PADDING > t.width[idx] {
			t.width[idx] = len(string(eleS)) + 3
		}
	}
	t.rows = append(t.rows, row)
}

func (t *Tablex) Draw() {
	s := strings.Builder{}
	s.WriteString("+")
	for idx := range t.columns {
		s.WriteString(strings.Repeat("-", t.width[idx]))
		s.WriteString("+")
	}
	s.WriteString("\n")
	for idx, col := range t.columns {
		s.WriteString("|")
		s.WriteString(" ")
		s.WriteString(col)
		s.WriteString(strings.Repeat(" ", t.width[idx]-len(col)-1))
	}
	s.WriteString("|")
	s.WriteString("\n")
	s.WriteString("+")
	for idx := range t.columns {
		s.WriteString(strings.Repeat("=", t.width[idx]))
		s.WriteString("+")
	}
	s.WriteString("\n")

	//===== ROWSSS

	for _, row := range t.rows {

		for idx, item := range row {
			s.WriteString("|")
			s.WriteString(" ")
			colItem := fmt.Sprintf("%v", item)
			s.WriteString(colItem)
			s.WriteString(strings.Repeat(" ", t.width[idx]-len(colItem)-1))
		}
		s.WriteString("|")
		s.WriteString("\n")
		s.WriteString("+")
		for idx := range t.columns {
			s.WriteString(strings.Repeat("-", t.width[idx]))
			s.WriteString("+")
		}
	}
	s.WriteString("\n")

	fmt.Println(s.String())

}
