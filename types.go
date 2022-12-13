package camparser

type Cam struct {
	Id   int     `json:"id" db:"id"`
	Name string  `json:"name" db:"camname"`
	Lon  float32 `json:"lng" db:"lon"`
	Lan  float32 `json:"lat" db:"lat"`
	Addr string  `json:"url" db:"addr"`
}

type Task struct {
	Id     int    `json:"id" db:"id"`
	Query  string `json:"query" db:"query"`
	Status string `json:"status" db:"status"`
	Count  int    `json:"count" db:"count"`
}
