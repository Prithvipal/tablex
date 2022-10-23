package tablex

import (
	"fmt"
	"strings"
)

const (
	PADDING = 3
)

type tablex struct {
	columns columns
	width   []int
	rows    []row
}

type columns []string

func (col columns) Draw(s *strings.Builder, width []int) {
	s.WriteString("+")
	for idx := range col {
		s.WriteString(strings.Repeat("-", width[idx]))
		s.WriteString("+")
	}
	s.WriteString("\n")
	for idx, col := range col {
		s.WriteString("|")
		s.WriteString(" ")
		s.WriteString(col)
		s.WriteString(strings.Repeat(" ", width[idx]-len(col)-1))
	}
	s.WriteString("|")
	s.WriteString("\n")
	s.WriteString("+")
	for idx := range col {
		s.WriteString(strings.Repeat("=", width[idx]))
		s.WriteString("+")
	}
	s.WriteString("\n")
}

type row []any

func (r row) Draw(s *strings.Builder, width []int) {
	for idx, item := range r {
		s.WriteString("|")
		s.WriteString(" ")
		colItem := fmt.Sprintf("%v", item)
		s.WriteString(colItem)
		s.WriteString(strings.Repeat(" ", width[idx]-len(colItem)-1))
	}
	s.WriteString("|")
	s.WriteString("\n")
	s.WriteString("+")
	for idx := range r {
		s.WriteString(strings.Repeat("-", width[idx]))
		s.WriteString("+")
	}
	s.WriteString("\n")
}

func NewTablex(columns []string) *tablex {
	t := &tablex{
		columns: columns,
		width:   make([]int, len(columns)),
	}
	for idx, col := range columns {
		t.width[idx] = len(col) + PADDING
	}

	return t
}

func (t *tablex) AddRow(row []any) {
	for idx, ele := range row {
		eleS := ele.(string)
		if len(string(eleS))+PADDING > t.width[idx] {
			t.width[idx] = len(string(eleS)) + 3
		}
	}
	t.rows = append(t.rows, row)
}

func (t *tablex) Draw() {
	s := strings.Builder{}
	t.columns.Draw(&s, t.width)

	for _, row := range t.rows {
		row.Draw(&s, t.width)
	}

	fmt.Println(s.String())
}
