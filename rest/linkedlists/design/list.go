package design

type List interface {
	Get(index int) int
	AddAtHead(val int)
	AddAtTail(val int)
	AddAtIndex(index int, val int)
	DeleteAtIndex(index int)
}
