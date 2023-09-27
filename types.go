package ransomwatch

import (
	"encoding/json"
	"fmt"
	"time"
)

const mainURL string = "https://raw.githubusercontent.com/joshhighet/ransomwatch/main/"

type Match struct {
	Post  Post  `json:"post"`
	Group Group `json:"group"`
}

type Group struct {
	Name    string     `json:"name"`
	Captcha bool       `json:"captcha"`
	Parser  bool       `json:"parser"`
	JS      bool       `json:"javascript_render"`
	Meta    string     `json:"meta"`
	Loc     []Location `json:"locations"`
	Profile []string   `json:"profile"`
}

type Location struct {
	FQDN       string     `json:"fqdn"`
	Title      string     `json:"title"`
	Version    int        `json:"version"`
	Slug       string     `json:"slug"`
	Available  bool       `json:"available"`
	Updated    RansomTime `json:"updated"`
	LastScrape RansomTime `json:"lastscrape"`
	Enabled    bool       `json:"enabled"`
}

type Post struct {
	Title      string     `json:"post_title"`
	GroupName  string     `json:"group_name"`
	Discovered RansomTime `json:"discovered"`
}

type RansomTime time.Time

func (t *RansomTime) UnmarshalJSON(b []byte) (err error) {
	var s string
	if err = json.Unmarshal(b, &s); err != nil {
		return
	}

	if s == "0000-00-00 00:00:00" || s == "" {
		*t = RansomTime(time.Time{})
		return
	}

	var tt time.Time
	if tt, err = time.Parse("2006-01-02 15:04:05.999999", s); err != nil {
		return
	}

	*t = RansomTime(tt)

	return
}

func (t RansomTime) MarshalJSON() ([]byte, error) {
	if tm := time.Time(t); tm.IsZero() {
		return []byte(`"0000-00-00 00:00:00"`), nil
	} else {
		return []byte(fmt.Sprintf(`"%s"`, tm.Format("2006-01-02 15:04:05.999999"))), nil
	}
}

func (t RansomTime) Format(format string) string {
	return time.Time(t).Format(format)
}
