package model

type Plugin struct {
	PluginName    string `json:"plugin_name"`
	PluginCode    string `json:"plugin_code"`
	PluginVersion string `json:"plugin_version"`
	PluginMemo    string `json:"plugin_memo"`
}
