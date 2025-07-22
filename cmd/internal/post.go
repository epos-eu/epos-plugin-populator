package internal

import (
	"bytes"
	"encoding/json"
	"epos-plugin-populator/cmd/internal/converter"
	"epos-plugin-populator/display"
	"fmt"
	"io"
	"net/url"
	"path"
)

func postPlugins(baseURL url.URL, plugins []Plugin, versionFlag string, operationUIDMap map[string]string) ([]*converter.Plugin, error) {
	pluginsURL := baseURL
	pluginsURL.Path = path.Join(pluginsURL.Path, "plugins")
	pluginPostURL := pluginsURL.String()

	relationsURL := baseURL
	relationsURL.Path = path.Join(relationsURL.Path, "plugin-relations")
	relationPostURL := relationsURL.String()

	pluginsPosted := 0
	pluginErrors := 0
	relationPosted := 0
	relationErrors := 0
	relationsTotal := 0
	postedPlugins := []*converter.Plugin{}

	display.Step("Starting plugin population process...")
	display.Info("Found %d plugins to process", len(plugins))

	if versionFlag != "" {
		display.Info("Using custom version: %s", versionFlag)
	}

	for i, p := range plugins {
		relationsTotal += len(p.Relations)

		display.Step("Processing plugin %d/%d: '%s'", i+1, len(plugins), p.Name)

		pluginToPost := converter.Plugin{
			Arguments:   p.Arguments,
			Description: p.Description,
			Enabled:     p.Enabled,
			Executable:  p.Executable,
			Name:        p.Name,
			Repository:  p.Repository,
			Runtime:     p.Runtime,
			Version:     p.Version,
			VersionType: p.VersionType,
		}

		if versionFlag != "" {
			pluginToPost.Version = versionFlag
		}

		postedPlugin, err := postAndUnmarshal(pluginPostURL, pluginToPost)
		if err != nil {
			display.Error("Failed to post plugin '%s': %v", p.Name, err)
			pluginErrors++
			continue
		}

		pluginsPosted++
		postedPlugins = append(postedPlugins, postedPlugin)
		display.Done("Plugin '%s' posted successfully (ID: %s)", postedPlugin.Name, postedPlugin.ID)

		if len(p.Relations) > 0 {
			display.Step("Processing %d relations for plugin '%s'", len(p.Relations), p.Name)

			for j, r := range p.Relations {
				if _, ok := operationUIDMap[r.RelationID]; !ok {
					display.Warn("  └─ Relation with operation UID '%s' has no mapping in the current environment", r.RelationID)
					relationErrors++
					continue
				}

				display.Info("  └─ Posting relation %d/%d (ID: %s)", j+1, len(p.Relations), r.RelationID)

				_, err := postAndUnmarshal(relationPostURL, converter.PluginRelation{
					InputFormat:  p.InputFormat,
					OutputFormat: p.OutputFormat,
					PluginID:     postedPlugin.ID,
					RelationID:   operationUIDMap[r.RelationID],
				})
				if err != nil {
					display.Warn("Failed to post relation '%s' for plugin '%s': %v", r.RelationID, p.Name, err)
					relationErrors++
					continue
				}

				relationPosted++
				display.Done("  └─ Relation '%s' posted successfully", r.RelationID)
			}
		} else {
			display.Info("No relations to process for plugin '%s'", p.Name)
		}
	}

	display.Info("Population complete - Summary:")
	display.Info("Plugins: %d successful, %d failed (out of %d total)", pluginsPosted, pluginErrors, len(plugins))
	display.Info("Relations: %d successful, %d failed (out of %d total)", relationPosted, relationErrors, relationsTotal)

	if pluginsPosted == 0 {
		display.Error("No plugins were successfully posted")
		return nil, fmt.Errorf("no plugin post has been successful")
	}

	if pluginErrors > 0 || relationErrors > 0 {
		display.Warn("Process completed with some errors - check logs above for details")
		return postedPlugins, fmt.Errorf("process completed with some errors - check logs for details")
	}

	display.Done("All plugins and relations posted successfully!")
	return postedPlugins, nil
}

func postAndUnmarshal[T any](url string, object T) (*T, error) {
	postBody, err := json.Marshal(object)
	if err != nil {
		return nil, fmt.Errorf("error converting post object '%+v' into json: %w", object, err)
	}

	resp, err := postHTTPClient.Post(url, "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		return nil, fmt.Errorf("error posting object '%+v': %w", object, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body for object '%+v' (status %s): %w", object, resp.Status, err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("object post for '%+v' returned non-success status %s. Body: %s", object, resp.Status, string(body))
	}

	var unmarshalObject T
	if err := json.Unmarshal(body, &unmarshalObject); err != nil {
		return nil, fmt.Errorf("failed to unmarshal successful response JSON for object '%+v': %w. Body: %s", object, err, string(body))
	}

	return &unmarshalObject, nil
}
