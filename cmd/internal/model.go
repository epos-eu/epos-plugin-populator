package internal

import "time"

// PLUGIN FILE

type Plugin struct {
	Version      string     `json:"version"`
	Name         string     `json:"name"`
	Description  string     `json:"description"`
	VersionType  string     `json:"version_type"`
	Repository   string     `json:"repository"`
	Runtime      string     `json:"runtime"`
	Executable   string     `json:"executable"`
	Arguments    string     `json:"arguments"`
	Enabled      bool       `json:"enabled"`
	InputFormat  string     `json:"inputFormat"`
	OutputFormat string     `json:"outputFormat"`
	Relations    []Relation `json:"relations"`
}
type Relation struct {
	RelationID string `json:"relationId"`
}

// SEARCH RESPONSE

type SearchResponse struct {
	Results Results `json:"results"`
}
type AvailableFormats struct {
	Format         string `json:"format"`
	Href           string `json:"href"`
	Label          string `json:"label"`
	OriginalFormat string `json:"originalFormat"`
	Type           string `json:"type"`
	InputFormat    string `json:"inputFormat,omitempty"`
	PluginID       string `json:"pluginId,omitempty"`
}
type Distributions struct {
	AvailableFormats []AvailableFormats `json:"availableFormats"`
	Description      string             `json:"description"`
	Href             string             `json:"href"`
	HrefExtended     string             `json:"hrefExtended"`
	ID               string             `json:"id"`
	Status           int                `json:"status"`
	StatusTimestamp  *time.Time         `json:"statusTimestamp,omitempty"`
	Title            string             `json:"title"`
	UID              string             `json:"uid"`
}
type Results struct {
	Distributions []Distributions `json:"distributions"`
}

// DISTRIBUTION DETAILS RESPONSE

type DetailsResponse struct {
	AvailableContactPoints []struct {
		Href string `json:"href"`
		Type string `json:"type"`
	} `json:"availableContactPoints"`
	AvailableFormats []struct {
		Format         string `json:"format"`
		Href           string `json:"href"`
		Label          string `json:"label"`
		OriginalFormat string `json:"originalFormat"`
		Type           string `json:"type"`
		InputFormat    string `json:"inputFormat,omitempty"`
		PluginID       string `json:"pluginId,omitempty"`
	} `json:"availableFormats"`
	Categories struct {
		Children []struct {
			Children []struct {
				Children []struct {
					Ddss string `json:"ddss"`
					Name string `json:"name"`
				} `json:"children"`
				Ddss string `json:"ddss"`
				Name string `json:"name"`
			} `json:"children"`
			Code    string `json:"code"`
			Color   string `json:"color"`
			ID      string `json:"id"`
			ImgURL  string `json:"imgUrl"`
			LinkURL string `json:"linkUrl"`
			Name    string `json:"name"`
		} `json:"children"`
		Name string `json:"name"`
	} `json:"categories"`
	DataProvider []struct {
		Country               string `json:"country"`
		DataProviderLegalName string `json:"dataProviderLegalName"`
		DataProviderURL       string `json:"dataProviderUrl"`
		Instanceid            string `json:"instanceid"`
		Metaid                string `json:"metaid"`
		UID                   string `json:"uid"`
	} `json:"dataProvider"`
	Description     string   `json:"description"`
	EditorID        string   `json:"editorId"`
	Endpoint        string   `json:"endpoint"`
	FrequencyUpdate string   `json:"frequencyUpdate"`
	Href            string   `json:"href"`
	HrefExtended    string   `json:"hrefExtended"`
	ID              string   `json:"id"`
	InternalID      []string `json:"internalID"`
	Keywords        []string `json:"keywords"`
	License         string   `json:"license"`
	MetaID          string   `json:"metaId"`
	// this is the uid of the operation
	OperationID          string   `json:"operationid"`
	ScienceDomain        []string `json:"scienceDomain"`
	ServiceDescription   string   `json:"serviceDescription"`
	ServiceDocumentation string   `json:"serviceDocumentation"`
	ServiceEndpoint      string   `json:"serviceEndpoint"`
	ServiceName          string   `json:"serviceName"`
	ServiceParameters    []struct {
		Enum          []string `json:"Enum,omitempty"`
		Label         string   `json:"label"`
		Name          string   `json:"name"`
		Required      bool     `json:"required"`
		Type          string   `json:"type"`
		MaxValue      string   `json:"maxValue,omitempty"`
		MinValue      string   `json:"minValue,omitempty"`
		Property      string   `json:"property,omitempty"`
		ValuePattern  string   `json:"valuePattern,omitempty"`
		DefaultValue  string   `json:"defaultValue,omitempty"`
		ReadOnlyValue string   `json:"readOnlyValue,omitempty"`
	} `json:"serviceParameters"`
	ServiceProvider struct {
		Country               string `json:"country"`
		DataProviderLegalName string `json:"dataProviderLegalName"`
		DataProviderURL       string `json:"dataProviderUrl"`
		Instanceid            string `json:"instanceid"`
		Metaid                string `json:"metaid"`
		UID                   string `json:"uid"`
	} `json:"serviceProvider"`
	ServiceSpatial struct {
		Paths [][][]float64 `json:"paths"`
		Wkid  int           `json:"wkid"`
	} `json:"serviceSpatial"`
	ServiceTemporalCoverage struct {
		StartDate time.Time `json:"startDate"`
	} `json:"serviceTemporalCoverage"`
	Spatial struct {
		Paths [][][]float64 `json:"paths"`
		Wkid  int           `json:"wkid"`
	} `json:"spatial"`
	TemporalCoverage struct {
		StartDate time.Time `json:"startDate"`
	} `json:"temporalCoverage"`
	Title            string `json:"title"`
	Type             string `json:"type"`
	UID              string `json:"uid"`
	VersioningStatus string `json:"versioningStatus"`
}
