package feeder

type Enclosure struct {
	Url    string
	Length int64 `json:",string"`
	Type   string
}
