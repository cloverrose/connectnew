package a

import (
	"connectrpc.com/connect"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Message struct {
	text string
}

func (m Message) ProtoReflect() protoreflect.Message {
	panic("implement me")
}

func useNewRequest() *connect.Request[Message] {
	return connect.NewRequest(&Message{ // OK
		text: "hello world",
	})
}

func useNewResponse() *connect.Response[Message] {
	return connect.NewResponse(&Message{ // OK
		text: "hello world",
	})
}

func useRequestDirectly() *connect.Request[Message] {
	return &connect.Request[Message]{ // want `use of &connect.Request\[T\]{} detected. Use connect.NewRequest\(\) instead.`
		Msg: &Message{
			text: "hello world",
		},
	}
}

func useResponseDirectly() *connect.Response[Message] {
	return &connect.Response[Message]{Msg: &Message{ // want `use of &connect.Response\[T\]{} detected. Use connect.NewResponse\(\) instead.`
		text: "hello world",
	}}
}
