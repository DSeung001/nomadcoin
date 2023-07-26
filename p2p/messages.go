package p2p

type MessageKind int

const (
	// iota 를 사용하는 순간 1,2,3 로 차례대로 매핑
	MessageNewestBlock MessageKind = iota
	MessageAllBlocksRequest
	MessageAllBlocksResponse
)

type Message struct {
	Kind    MessageKind
	Payload []byte
}
