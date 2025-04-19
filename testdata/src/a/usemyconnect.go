package a

import (
	connect "a/myconnect"
)

func useRequestDirectlyMyConnect() *connect.Request[Message] {
	return &connect.Request[Message]{ // OK
		Msg: &Message{
			text: "hello world",
		},
	}
}

func useResponseDirectlyMyConnect() *connect.Response[Message] {
	return &connect.Response[Message]{Msg: &Message{ // OK
		text: "hello world",
	}}
}
