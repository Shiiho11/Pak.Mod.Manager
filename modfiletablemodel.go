package main

import (
	"sort"

	"github.com/lxn/walk"
)

type ModFileTableItem struct {
	Name    string
	checked bool
}

type ModFileTableModel struct {
	walk.TableModelBase
	walk.SorterBase
	ConfigIndex int
	sortColumn  int
	sortOrder   walk.SortOrder
	items       []*ModFileTableItem
}

func NewModFileTableModel(configIndex int) *ModFileTableModel {
	m := new(ModFileTableModel)
	m.ConfigIndex = configIndex
	m.ResetRows()
	return m
}

// Called by the TableView from SetModel and every time the model publishes a RowsReset event.
func (m *ModFileTableModel) RowCount() int {
	return len(m.items)
}

// Called by the TableView when it needs the text to display for a given cell.
func (m *ModFileTableModel) Value(row, col int) interface{} {
	item := m.items[row]

	switch col {
	case 0:
		return item.Name
	}
	panic("unexpected col")
}

// Called by the TableView to retrieve if a given row is checked.
func (m *ModFileTableModel) Checked(row int) bool {
	return m.items[row].checked
}

// Called by the TableView when the user toggled the check box of a given row.
func (m *ModFileTableModel) SetChecked(row int, checked bool) error {
	m.items[row].checked = checked
	return nil
}

func (m *ModFileTableModel) CheckAll() {
	for i := 0; i < len(m.items); i++ {
		m.items[i].checked = true
	}
	// m.ResetRows()
}

func (m *ModFileTableModel) UncheckAll() {
	for i := 0; i < len(m.items); i++ {
		m.items[i].checked = false
	}
	// m.ResetRows()
}

// Called by the TableView to sort the model.
func (m *ModFileTableModel) Sort(col int, order walk.SortOrder) error {
	m.sortColumn, m.sortOrder = col, order

	sort.SliceStable(m.items, func(i, j int) bool {
		a, b := m.items[i], m.items[j]

		c := func(ls bool) bool {
			if m.sortOrder == walk.SortAscending {
				return ls
			}

			return !ls
		}

		switch m.sortColumn {
		case 0:
			return c(a.Name < b.Name)
		}

		panic("unreachable")
	})

	return m.SorterBase.Sort(col, order)
}

func (m *ModFileTableModel) ResetRows() {
	config.ModConfigSlice[m.ConfigIndex].RefreshModConfig()
	m.DataToModel()

	// Notify TableView and other interested parties about the reset.
	m.PublishRowsReset()

	m.Sort(m.sortColumn, m.sortOrder)
}

func (m *ModFileTableModel) ApplyChange() {
	config.ModConfigSlice[m.ConfigIndex].RefreshModConfig()
	m.ModelToData()
	config.ModConfigSlice[m.ConfigIndex].ApplyModFileInfoToactualFile()
}

func (m *ModFileTableModel) DataToModel() {
	m.items = make([]*ModFileTableItem, len(config.ModConfigSlice[m.ConfigIndex].ModFilesMap))
	var i = 0
	for name, disable := range config.ModConfigSlice[m.ConfigIndex].ModFilesMap {
		m.items[i] = &ModFileTableItem{
			Name:    name,
			checked: !disable,
		}
		i += 1
	}
}

func (m *ModFileTableModel) ModelToData() {
	FileInfo := make(map[string]bool)
	for _, item := range m.items {
		_, ok := config.ModConfigSlice[m.ConfigIndex].ModFilesMap[item.Name]
		if ok {
			FileInfo[item.Name] = !item.checked
		}
	}
	config.ModConfigSlice[m.ConfigIndex].ModFilesMap = FileInfo
}
