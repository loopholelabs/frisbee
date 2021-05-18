// Code generated by frisbeegenerator. DO NOT EDIT.

package pubsub

import (
	"github.com/loophole-labs/frisbee"
)

type PubSubClientHandler interface {
	HandlePub(incomingMessage frisbee.Message, incomingContent []byte) (outgoingMessage *frisbee.Message, outgoingContent []byte, action frisbee.Action)
	HandleSub(incomingMessage frisbee.Message, incomingContent []byte) (outgoingMessage *frisbee.Message, outgoingContent []byte, action frisbee.Action)
}
type PubSubServerHandler interface {
	HandlePub(c *frisbee.Conn, incomingMessage frisbee.Message, incomingContent []byte) (outgoingMessage *frisbee.Message, outgoingContent []byte, action frisbee.Action)
	HandleSub(c *frisbee.Conn, incomingMessage frisbee.Message, incomingContent []byte) (outgoingMessage *frisbee.Message, outgoingContent []byte, action frisbee.Action)
}

var messageTypes = map[string]uint16{"Pub": 1, "Sub": 2}

func initPubSubClientRouter(h PubSubClientHandler) frisbee.ClientRouter {
	router := make(frisbee.ClientRouter)
	router[messageTypes["Pub"]] = h.HandlePub
	router[messageTypes["Sub"]] = h.HandleSub
	return router
}
func initPubSubServerRouter(h PubSubServerHandler) frisbee.ServerRouter {
	router := make(frisbee.ServerRouter)
	router[messageTypes["Pub"]] = h.HandlePub
	router[messageTypes["Sub"]] = h.HandleSub
	return router
}
func NewPubSubClient(addr string, h PubSubClientHandler, opts ...frisbee.Option) *frisbee.Client {
	return frisbee.NewClient(addr, initPubSubClientRouter(h), opts...)
}
func NewPubSubServer(addr string, h PubSubServerHandler, opts ...frisbee.Option) *frisbee.Server {
	return frisbee.NewServer(addr, initPubSubServerRouter(h), opts...)
}
