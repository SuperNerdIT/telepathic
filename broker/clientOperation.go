package broker

type Operation byte

const (
	Remove Operation = iota
	Add
)

type ClientOperation struct {
	Operation Operation
	Client    *Client
}
