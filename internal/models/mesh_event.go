package models

// MeshEvent is one decoded FromRadio frame: any of Node, Message or Channel may
// be set (a packet can carry both a Node update and a Message).
type MeshEvent struct {
	Node    *Node
	Message *Message
	Channel *Channel
}
