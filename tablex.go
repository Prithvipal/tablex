package tablex

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

const (
	PADDING = 3
)

type table struct {
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

func NewTablex(columns []string) *table {
	t := &table{
		columns: columns,
		width:   make([]int, len(columns)),
	}
	for idx, col := range columns {
		t.width[idx] = len(col) + PADDING
	}

	return t
}

func FromCSV(path string) (*table, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	t := NewTablex(records[0])
	for i := 1; i < len(records); i++ {
		row := make([]any, len(records[i]))
		for idx, item := range records[i] {
			row[idx] = item
		}
		t.AddRow(row)
	}

	return t, nil
}

func (t *table) AddRow(row []any) {
	for idx, ele := range row {
		eleS := ele.(string)
		if len(string(eleS))+PADDING > t.width[idx] {
			t.width[idx] = len(string(eleS)) + 3
		}
	}
	t.rows = append(t.rows, row)
}

func (t *table) DeleteLastRow() {
	t.rows = t.rows[:len(t.rows)-1]
}

func (t *table) DeleteFirstRow() {
	t.rows = t.rows[1:len(t.rows)]
}

func (t *table) DeleteRowNByIndex(i int) error {
	if i < 0 || i >= len(t.rows) {
		return fmt.Errorf("the given index is out of bounds")
	}

	newRow := append(t.rows[:i], t.rows[i+1:]...)
	t.rows = newRow
	return nil
}

func (t *table) Draw() {
	s := strings.Builder{}
	t.columns.Draw(&s, t.width)

	for _, row := range t.rows {
		row.Draw(&s, t.width)
	}

	fmt.Println(s.String())
}
