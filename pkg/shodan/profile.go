package shodan

import "encoding/json"

type Profile struct {
	Member  bool   `json:"member"`
	Credits int    `json:"credits"`
	Name    string `json:"display_name"`
}

func (s *Shodan) GetProfile() (*Profile, error) {
	requestURL := "https://api.shodan.io/account/profile"
	body, err := s.makeGet(requestURL, map[string]string{})
	if err != nil {
		return nil, err
	}
	var p Profile
	err = json.Unmarshal(*body, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}
