package ransomwatch

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const postsPath string = "posts.json"

// Getting the list of posts from ransomwatch repository.
func GetPosts() (posts []Post, err error) {

	req, err := http.NewRequest(http.MethodGet, mainURL+postsPath, nil)
	if err != nil {
		err = fmt.Errorf("error creating request for '%s': %s", postsPath, err)
		return
	}

	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("error getting '%s': %s", postsPath, err)
		return
	}
	defer resp.Body.Close()

	if status := resp.StatusCode; status != http.StatusOK {
		data, _ := io.ReadAll(resp.Body)
		err = fmt.Errorf("'%d' bad status received code from '%s': %s", status, postsPath, data)
		return
	}

	if err = json.NewDecoder(resp.Body).Decode(&posts); err != nil {
		err = fmt.Errorf("error decoding response from '%s': %s", postsPath, err)
	}

	return
}
