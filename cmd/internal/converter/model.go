package converter

type Plugin struct {
	Arguments   string `json:"arguments"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
	Executable  string `json:"executable"`
	ID          string `json:"id"`
	Installed   bool   `json:"installed"`
	Name        string `json:"name"`
	Repository  string `json:"repository"`
	Runtime     string `json:"runtime"`
	Version     string `json:"version"`
	VersionType string `json:"version_type"`
}
type PluginRelation struct {
	ID           string `json:"id"`
	InputFormat  string `json:"input_format"`
	OutputFormat string `json:"output_format"`
	PluginID     string `json:"plugin_id"`
	RelationID   string `json:"relation_id"`
}
