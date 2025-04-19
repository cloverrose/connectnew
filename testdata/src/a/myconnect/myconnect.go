package myconnect

type Request[T any] struct {
	Msg *T
}

type Response[T any] struct {
	Msg *T
}
