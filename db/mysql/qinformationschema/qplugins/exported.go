// Code generated by Lingo for table information_schema.PLUGINS - DO NOT EDIT

package qplugins

import "github.com/weworksandbox/lingo/core/path"

var instance = New()

func Q() QPlugins {
	return instance
}

func PluginName() path.StringPath {
	return instance.pluginName
}

func PluginVersion() path.StringPath {
	return instance.pluginVersion
}

func PluginStatus() path.StringPath {
	return instance.pluginStatus
}

func PluginType() path.StringPath {
	return instance.pluginType
}

func PluginTypeVersion() path.StringPath {
	return instance.pluginTypeVersion
}

func PluginLibrary() path.StringPath {
	return instance.pluginLibrary
}

func PluginLibraryVersion() path.StringPath {
	return instance.pluginLibraryVersion
}

func PluginAuthor() path.StringPath {
	return instance.pluginAuthor
}

func PluginDescription() path.StringPath {
	return instance.pluginDescription
}

func PluginLicense() path.StringPath {
	return instance.pluginLicense
}

func LoadOption() path.StringPath {
	return instance.loadOption
}
