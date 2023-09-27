package ransomwatch

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const groupsPath string = "groups.json"

// Getting the list of groups from ransomwatch repository.
func GetGroups() (groups []Group, err error) {

	req, err := http.NewRequest(http.MethodGet, mainURL+groupsPath, nil)
	if err != nil {
		err = fmt.Errorf("error creating request for '%s': %s", groupsPath, err)
		return
	}

	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("error getting '%s': %s", groupsPath, err)
		return
	}
	defer resp.Body.Close()

	if status := resp.StatusCode; status != http.StatusOK {
		data, _ := io.ReadAll(resp.Body)
		err = fmt.Errorf("'%d' bad status received code from '%s': %s", status, groupsPath, data)
		return
	}

	if err = json.NewDecoder(resp.Body).Decode(&groups); err != nil {
		err = fmt.Errorf("error decoding response from '%s': %s", groupsPath, err)
	}

	return
}
