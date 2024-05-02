package main

import (
	"os/exec"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

var config = NewConfig()

var mainWindow MainWindow

var configNameLE *walk.LineEdit
var modFolderLE *walk.LineEdit

var modConfigLB *walk.ListBox
var modConfigListModel *ModConfigListModel

var ModFilesTV *walk.TableView
var modFileTableModel *ModFileTableModel

var delectConfigPB *walk.PushButton

func main() {
	config.Load()
	config.ModConfigSlice[0].RefreshModConfig()
	modConfigListModel = NewModConfigListModel()
	modFileTableModel = NewModFileTableModel(0)

	mainWindow = MainWindow{
		Title:  "Pak Mod Manager",
		Size:   Size{Width: 700, Height: 500},
		Font:   Font{Family: "Arial", PointSize: 10},
		Layout: Grid{Rows: 1},
		Children: []Widget{
			Composite{
				Layout:  VBox{},
				MaxSize: Size{Width: 150},
				Children: []Widget{
					Label{
						Text: "Mod Config List",
						Font: Font{Family: "Arial", PointSize: 13, Bold: true},
					},
					PushButton{
						Text:      "New Config",
						OnClicked: newConfig,
					},
					ListBox{
						AssignTo:              &modConfigLB,
						Model:                 modConfigListModel,
						OnCurrentIndexChanged: modConfigIndexChanged,
						CurrentIndex:          0,
					},
					PushButton{
						AssignTo:  &delectConfigPB,
						Text:      "Delect Config",
						Enabled:   false,
						OnClicked: delectConfig,
					},
				},
			},
			Composite{
				Layout:     Grid{Columns: 1},
				ColumnSpan: 2,
				Children: []Widget{
					Label{
						Text: "Mod File Config",
						Font: Font{Family: "Arial", PointSize: 13, Bold: true},
					},
					Composite{
						Layout: Grid{Columns: 3},
						Children: []Widget{
							Label{
								Row:  0,
								Text: "Config Name:",
							},
							LineEdit{
								AssignTo: &configNameLE,
								Row:      1,
								Column:   0,
								Text:     config.ModConfigSlice[0].Name,
							},
							PushButton{
								Row:    1,
								Column: 1,
								Text:   "Rename",
								OnClicked: func() {
									rename(modConfigLB.CurrentIndex())
								},
							},
							PushButton{
								Row:       1,
								Column:    2,
								Text:      "Save",
								OnClicked: save,
							},
						},
					},
					Composite{
						Layout: Grid{Columns: 3},
						Children: []Widget{
							Label{
								Row:  0,
								Text: "Mod Folder:",
							},
							LineEdit{
								AssignTo: &modFolderLE,
								Row:      1,
								Column:   0,
								Text:     config.ModConfigSlice[0].ModFolder,
							},
							PushButton{
								Row:    1,
								Column: 1,
								Text:   "Set Folder",
								OnClicked: func() {
									setFolder(modConfigLB.CurrentIndex())
								},
							},
							PushButton{
								Row:    1,
								Column: 2,
								Text:   "Open Folder",
								OnClicked: func() {
									openFolder(modConfigLB.CurrentIndex())
								},
							},
						},
					},
					Composite{
						Name:    "Mod config top composite",
						RowSpan: 2,
						Layout:  Grid{Columns: 5},
						Children: []Widget{
							TableView{
								Name:             "Mode file table",
								AssignTo:         &ModFilesTV,
								AlternatingRowBG: true,
								CheckBoxes:       true,
								ColumnsOrderable: true,
								MultiSelection:   true,
								Columns: []TableViewColumn{
									{Title: "Name", Width: 400},
								},
								Model: modFileTableModel,
								// OnSelectedIndexesChanged: func() {
								// 	fmt.Printf("SelectedIndexes: %v\n", ModFilesTV.SelectedIndexes())
								// },
							},
							Composite{
								Layout: Grid{Columns: 1},
								Children: []Widget{
									PushButton{
										Text:      "Check All",
										OnClicked: checkAll,
									},
									PushButton{
										Text:      "Uncheck All",
										OnClicked: uncheckAll,
									},
									PushButton{
										Text: "Refresh",
										OnClicked: func() {
											refresh(modConfigLB.CurrentIndex())
										},
									},
									PushButton{
										Text:      "Apply",
										OnClicked: apply,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	mainWindow.Run()
	config.Save()

}

func newConfig() {
	config.ModConfigSlice = append(config.ModConfigSlice, ModConfig{
		Name:        "New Config",
		ModFolder:   config.ModConfigSlice[0].ModFolder,
		ModFilesMap: config.ModConfigSlice[0].ModFilesMap,
	})
	modConfigLB.SetModel(modConfigListModel)
	modConfigLB.SetCurrentIndex(len(config.ModConfigSlice) - 1)
}

func delectConfig() {
	index := modConfigLB.CurrentIndex()
	config.ModConfigSlice = append(config.ModConfigSlice[:index], config.ModConfigSlice[index+1:]...)
	modConfigLB.SetModel(modConfigListModel)
	modConfigLB.SetCurrentIndex(max(index-1, 0))
}

func modConfigIndexChanged() {
	index := modConfigLB.CurrentIndex()
	if index == 0 {
		delectConfigPB.SetEnabled(false)
	} else {
		delectConfigPB.SetEnabled(true)
	}
	if index >= 0 {
		refresh(index)
	}
}

func refreshModConfigList() {
	index := modConfigLB.CurrentIndex()
	modConfigLB.SetModel(modConfigListModel)
	modConfigLB.SetCurrentIndex(index)
}

func rename(index int) {
	config.ModConfigSlice[index].Name = configNameLE.Text()
	refreshModConfigList()
}

func save() {
	config.Save()
}

func setFolder(index int) {
	modFolder := modFolderLE.Text()
	config.ModConfigSlice[index].ModFolder = modFolder
	refresh(index)
}

func openFolder(index int) {
	exec.Command("cmd", "/c", "explorer", config.ModConfigSlice[index].ModFolder).Start()
}

func refresh(index int) {
	if index >= 0 {
		config.ModConfigSlice[index].RefreshModConfig()
		configNameLE.SetText(config.ModConfigSlice[index].Name)
		modFolderLE.SetText(config.ModConfigSlice[index].ModFolder)
		modFileTableModel.ConfigIndex = index
		modFileTableModel.ResetRows()
	}
}

func checkAll() {
	modFileTableModel.CheckAll()
	ModFilesTV.SetModel(modFileTableModel)
}

func uncheckAll() {
	modFileTableModel.UncheckAll()
	ModFilesTV.SetModel(modFileTableModel)
}

func apply() {
	config.ModConfigSlice[modFileTableModel.ConfigIndex].RefreshModConfig()
	modFileTableModel.ModelToData()
	config.ModConfigSlice[modFileTableModel.ConfigIndex].ApplyModFileInfoToactualFile()
	refresh(modFileTableModel.ConfigIndex)
}
