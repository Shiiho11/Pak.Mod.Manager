package main

import (
	"os"
	"path/filepath"
	"strings"
)

const modFileSuffix = "pak"
const disableSuffix = "disable"

type ModConfig struct {
	Name        string
	ModFolder   string
	ModFilesMap map[string]bool
}

func (mc *ModConfig) RefreshModConfig() {
	actualFileInfo := getActualFileInfo(mc.ModFolder)
	mc.ModFilesMap = mergeFileInfo(mc.ModFilesMap, actualFileInfo)
}

func getActualFileInfo(modFolder string) map[string]bool {
	FileInfo := make(map[string]bool)
	fileNameSlice := getAllFiles(modFolder)
	for _, fileName := range fileNameSlice {
		name, disable, ok := checkModFileName(fileName)
		if ok {
			FileInfo[name] = disable
		}
	}
	return FileInfo
}

func getAllFiles(path string) []string {
	dir, err := os.ReadDir(path)
	if err != nil {
		return nil
	}

	var fileNameSlice []string
	for _, fi := range dir {
		if !fi.IsDir() {
			fileNameSlice = append(fileNameSlice, fi.Name())
		}
	}
	return fileNameSlice
}

func checkModFileName(fileName string) (name string, disable bool, ok bool) {
	fileName = filepath.Base(fileName)
	fileNameSplit := strings.Split(fileName, ".")
	if len(fileNameSplit) >= 3 && fileNameSplit[len(fileNameSplit)-2] == "pak" && fileNameSplit[len(fileNameSplit)-1] == "disable" {
		name := strings.Join(fileNameSplit[:len(fileNameSplit)-2], ".")
		return name, true, true
	} else if len(fileNameSplit) >= 2 && fileNameSplit[len(fileNameSplit)-1] == "pak" {
		name := strings.Join(fileNameSplit[:len(fileNameSplit)-1], ".")
		return name, false, true
	} else {
		return fileName, false, false
	}
}

func mergeFileInfo(currentFileInfo map[string]bool, actualFileInfo map[string]bool) map[string]bool {
	FileInfo := make(map[string]bool)
	for name, actual_disable := range actualFileInfo {
		current_disable, ok := currentFileInfo[name]
		if ok {
			FileInfo[name] = current_disable
		} else {
			FileInfo[name] = actual_disable
		}
	}
	return FileInfo
}

func (mc *ModConfig) ApplyModFileInfoToactualFile() {
	actualFileInfo := getActualFileInfo(mc.ModFolder)
	for name, actual_disable := range actualFileInfo {
		set_disable, ok := mc.ModFilesMap[name]
		if ok {
			if actual_disable == set_disable {
				continue
			}
			var fileName string
			var newFileName string
			if actual_disable {
				fileName = mc.ModFolder + "\\" + name + "." + modFileSuffix + "." + disableSuffix
			} else {
				fileName = mc.ModFolder + "\\" + name + "." + modFileSuffix
			}
			if set_disable {
				newFileName = mc.ModFolder + "\\" + name + "." + modFileSuffix + "." + disableSuffix
			} else {
				newFileName = mc.ModFolder + "\\" + name + "." + modFileSuffix
			}
			_ = os.Rename(fileName, newFileName)
		}
	}
}
