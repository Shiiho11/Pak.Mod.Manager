[EN](README.md) | 中文

# Pak.Mod.Manager
一个基于GO [walk](https://github.com/lxn/walk)的小工具，用于配置是否启用Pak Mod文件。

## 下载
https://github.com/Shiiho11/Pak.Mod.Manager/releases

## 使用
请将可执行程序放在独立的文件夹中，程序会生成`config.json`用于保存配置。

左边的`Mod Config List`，可以创建、选择、删除配置。右边的`Mod File Config`是每个配置的具体内容。

`Config Name`可以修改后重命名，点击`Save`会保存所有配置到文件`config.json`中。在正常关闭软件时会自动保存，该按钮用于手动保存，防止软件不正常关闭的情况下丢失配置。

`Mod Folder`是当前配置管理的Mod文件夹，请输入文件夹路径后点击`Set Folder`保存，然后点击`Open Folder`测试是否正确打开指定的文件夹。

下面的表格用于配置是否启用对应的Pak文件。勾选是启用，不勾选是不启用。点击`Apply`会应用当前配置（不启用的Pak文件会被加上.disable后缀名）。`Refresh`按钮用于刷新文件夹中的文件信息，请在添加或删除Pak文件后刷新。

## 构建应用程序
请参考[walk](https://github.com/lxn/walk)的指南。