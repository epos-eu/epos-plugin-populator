package internal

import (
	"encoding/json"
	"epos-plugin-populator/display"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"
)

var getHTTPClient = &http.Client{
	Timeout: 30 * time.Second,
}

var postHTTPClient = &http.Client{
	Timeout: 60 * time.Second,
}

func Populate(baseURL url.URL, plugins []Plugin, versionFlag string) error {
	display.Step("Searching available distributions in environment")

	distributionIDs, err := findDistributionIDs(baseURL)
	if err != nil {
		return fmt.Errorf("error finding disstribution IDs in environment '%s': %w", baseURL.String(), err)
	}

	display.Done("Found %d distributions in environment '%s'", len(distributionIDs), baseURL.String())
	display.Step("Mapping distributions to theirs operation UID")

	operationUIDMap, err := getDistOperationUIDs(baseURL, distributionIDs)
	if err != nil {
		return fmt.Errorf("error mapping distribution IDs to their Operation UIDs: %w", err)
	}

	display.Done("Finished fetching details. Mapped %d Operation UIDs.", len(operationUIDMap))
	display.Step("Populating the converter with the plugins")

	postedPlugins, err := postPlugins(baseURL, plugins, versionFlag, operationUIDMap)
	if err != nil {
		return fmt.Errorf("error posting plugins: %w", err)
	}

	display.Done("Finished posting plugins. Posted %d plugins successfully.", len(postedPlugins))
	display.Step("Populating the converter with the plugin relations")

	return nil
}

func findDistributionIDs(baseURL url.URL) (map[string]struct{}, error) {
	baseURL.Path = path.Join(baseURL.Path, "resources/search")
	query := url.Values{}
	query.Set("facets", "false")
	query.Set("q", "")
	baseURL.RawQuery = query.Encode()
	searchURL := baseURL.String()

	display.Step("Fetching all distributions from: %s", searchURL)

	searchResp, err := getAndUnmarshal[SearchResponse](searchURL)
	if err != nil {
		return nil, fmt.Errorf("failed to get distributions from search endpoint '%s': %w", searchURL, err)
	}

	distributionIDs := map[string]struct{}{}
	if searchResp.Results.Distributions == nil {
		display.Warn("Search response from %s contained no distribution items. Nothing to do.", searchURL)
		return distributionIDs, nil
	}

	for _, distribution := range searchResp.Results.Distributions {
		distributionIDs[distribution.ID] = struct{}{}
	}

	return distributionIDs, nil
}

// getDistOperationUIDs returns a map[uid]instanceid for the distributionIDs
func getDistOperationUIDs(baseURL url.URL, distributionIDs map[string]struct{}) (map[string]string, error) {
	count := 0
	UIDFetchErrors := 0
	total := len(distributionIDs)
	operationUIDMap := map[string]string{}
	for distID := range distributionIDs {
		UID, err := getOperationIDForDistribution(baseURL, distID)
		if err != nil {
			display.Warn("Failed to get details for distribution ID '%s': %v", distID, err)
			UIDFetchErrors++
			continue
		}
		operationUIDMap[UID] = distID

		count++
		if count%100 == 0 || count == total {
			display.Info("Processed %d / %d distributions for Operation UIDs (%d errors so far)", count, total, UIDFetchErrors)
		}
	}
	if UIDFetchErrors > 0 {
		display.Warn("Encountered %d errors while fetching details for %d distributions.", UIDFetchErrors, total)
	}

	return operationUIDMap, nil
}

func getOperationIDForDistribution(baseURL url.URL, distID string) (string, error) {
	baseURL.Path = path.Join(baseURL.Path, "resources", "details", distID)
	detailsURL := baseURL.String()

	details, err := getAndUnmarshal[DetailsResponse](detailsURL)
	if err != nil {
		return "", fmt.Errorf("failed to get details for distribution at '%s': %w", detailsURL, err)
	}

	if details.OperationID == "" {
		return "", fmt.Errorf("operation ID empty for distribution with id '%s'", distID)
	}
	return strings.TrimPrefix(details.OperationID, "file:///"), nil
}

// getAndUnmarshal performs an HTTP GET request and unmarshals the JSON response.
func getAndUnmarshal[T any](url string) (T, error) {
	var result T
	resp, err := getHTTPClient.Get(url)
	if err != nil {
		return result, fmt.Errorf("http GET request to %s failed: %w", url, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, fmt.Errorf("failed to read response body from %s (status %s): %w", url, resp.Status, err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted {
		return result, fmt.Errorf("http GET request to %s returned status %s. Body: %s", url, resp.Status, string(body))
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return result, fmt.Errorf("failed to unmarshal JSON response from %s. Error: %w. Body: %s", url, err, string(body))
	}
	return result, nil
}
