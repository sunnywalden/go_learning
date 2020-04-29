package kernel_plugin

import (
	"container/list"
	"errors"
	"fmt"
)

type Plugin struct {
	pluginName string
	pluginVersion string
}


type PluginManage struct {
	plugins []Plugin
}


func ListIndex(list []interface{},target interface{}) int {
	for i,s := range list {
		if s == target {
			result = append(list[:i],list[i+1:]...)
			return true,nil
		} else {
			fmt.Printf("no %s matched found", target)
			return -1
		}
	}
}

func(pm *PluginManage) PluginRegister(pl Plugin) (bool, error) {
	pm.plugins = append(pm.plugins, pl)
	return true, nil
}

func(pm *PluginManage) PluginUnregister(pl Plugin) (bool, error) {
	for i,plugin := range pm.plugins {
		if plugin == pl {
			pm.plugins = append(pm.plugins[:i],pm.plugins[i+1:]...)
			return true,nil
		} else {
			fmt.Printf("no plugin %s matched found", plugin.pluginName)
			return false, errors.New("no plugin matched found")
		}
	}
	return true, nil
}

func (pm *PluginManage) PluginInstall(pl Plugin) (bool, error) {
	fmt.Printf("Install version %s of %s to manage", pl.pluginVersion, pl.pluginName)
	_, err := pm.PluginRegister(pl)
	if err != nil {
		return false, err
	} else {
		fmt.Printf("Version %s of plugin %s installed!", pl.pluginVersion, pl.pluginName)
		return true,nil
	}
}

func (pm *PluginManage) PluginUninstall(pl Plugin) (bool, error) {
	fmt.Printf("Uninstall version %s of %s to manage", pl.pluginVersion, pl.pluginName)
	_, err := pm.PluginUnregister(pl)
	if err != nil {
		return false, err
	} else {
		fmt.Printf("Version %s of plugin %s uninstalled!", pl.pluginVersion, pl.pluginName)
		return true,nil
	}
}

func (pm *PluginManage) PluginUpdate(pl Plugin) (bool, error) {
	fmt.Printf("Update plugin %s from version %s to %s", pl.pluginName, pl.pluginVersion)
	_, err := pm.PluginRegister(pl)
	if err != nil {
		return false, err
	} else {
		fmt.Printf("Version %s of plugin %s installed!", pl.pluginVersion, pl.pluginName)
		return true,nil
	}
}

func (pm *PluginManage) PluginList() ([]Plugin, error) {
	if pm.plugins != nil {
		return pm.plugins,nil
		} else {
			return nil, errors.New("no plugin found")
	}
}


