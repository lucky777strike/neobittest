package shodan

import (
	"encoding/json"
	"fmt"
)

type QueryResp struct {
	Matches []QueryMatch `json:"matches"`
}

type QueryMatch struct {
	Http struct {
		Host string `json:"host"`
	} `json:"http"`
	Location struct {
		Lon float32 `json:"longitude"`
		Lan float32 `json:"latitude"`
	} `json:"location"`
	Port int    `json:"port"`
	Addr string `json:"ip_str"`
}

func (q QueryMatch) String() string {
	return fmt.Sprintf("Ip:%s \n Url: http://%s:%d \n longitude:%f \n latitude:%f",
		q.Addr, q.Http.Host, q.Port, q.Location.Lon, q.Location.Lan)
}

func (s *Shodan) QuerySearch(query string) (*QueryResp, error) {
	requestURL := "https://api.shodan.io/shodan/host/search"
	body, err := s.makeGet(requestURL, map[string]string{"query": query})
	if err != nil {
		return nil, err
	}
	//fmt.Println(string(*body))
	var p QueryResp
	err = json.Unmarshal(*body, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}
