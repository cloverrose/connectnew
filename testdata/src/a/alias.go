package a

import (
	cnc "connectrpc.com/connect"
)

func useRequestDirectlyAlias() *cnc.Request[Message] {
	return &cnc.Request[Message]{ // want `use of &connect.Request\[T\]{} detected \(imported as cnc\). Use cnc.NewRequest\(\) instead`
		Msg: &Message{
			text: "hello world",
		},
	}
}

func useResponseDirectlyAlias() *cnc.Response[Message] {
	return &cnc.Response[Message]{Msg: &Message{ // want `use of &connect.Response\[T\]{} detected \(imported as cnc\). Use cnc.NewResponse\(\) instead`
		text: "hello world",
	}}
}
