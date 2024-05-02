EN | [中文](README_ZH.md)

# Pak.Mod.Manager
A small tool based on GO [walk](https://github.com/lxn/walk) for configuring whether pak mod files are enabled.

## Download
https://github.com/Shiiho11/Pak.Mod.Manager/releases

## Using
Please place the executable program in a separate folder, the program will generate `config. json` to save the configuration.

The `Mod Config List` on the left can create, select, and delete configurations. The `Mod File Config` on the right is the specific content of each configuration.

`Config Name` can be modified and renamed. Clicking `Save` will save all configurations to the file `config.json`. When the software is closed normally, it will automatically save. This button is used for manual saving to prevent configuration loss in the event of abnormal software shutdown.

`Mod Folder` is the Mod folder managed by the current configuration. Please enter the folder path, click `Set Folder` to save, and then click `Open Folder` to test whether the specified folder is opened correctly.

The table below is used to configure whether to enable the corresponding Pak file. Check to enable, uncheck to disable. Clicking `Apply` will apply the current configuration (disable Pak files will have disable suffix added). The `Refresh` button is used to refresh the file information in the folder. Please refresh after adding or deleting Pak files.

## Build app
Please refer to [walk](https://github.com/lxn/walk)