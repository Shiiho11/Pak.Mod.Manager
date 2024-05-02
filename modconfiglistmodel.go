package main

import (
	"github.com/lxn/walk"
)

type ModConfigListModel struct {
	walk.ListModelBase
}

func NewModConfigListModel() *ModConfigListModel {
	m := &ModConfigListModel{}
	return m
}

func (m *ModConfigListModel) ItemCount() int {
	return len(config.ModConfigSlice)
}

func (m *ModConfigListModel) Value(index int) interface{} {
	return config.ModConfigSlice[index].Name
}
