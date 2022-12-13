package shodan

import (
	"io"
	"net/http"
	"time"
)

type Shodan struct {
	token  string
	client *http.Client //Clients are safe for concurrent use by multiple goroutines.
}

func New(token string) *Shodan {
	client := &http.Client{Timeout: time.Second * 180} //somtimes shodan has freezes
	return &Shodan{token: token, client: client}
}

func (s *Shodan) makeGet(url string, params map[string]string) (*[]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	q := req.URL.Query()
	q.Add("key", s.token)
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	//fmt.Println(string(body))
	//fmt.Println(resp.StatusCode)
	if resp.StatusCode == 401 {
		return nil, &ReqError{
			Status: 401,
			Msg:    "Bad token",
		}
	}
	if resp.StatusCode == 429 {
		return nil, &ReqError{
			Status: 429,
			Msg:    "Rate limit",
		}
	}
	if resp.StatusCode == 200 {
		if err != nil {
			return nil, err
		}
		return &body, nil
	}
	return nil, &ReqError{
		Status: resp.StatusCode,
		Msg:    "Unknown error",
	}
}
