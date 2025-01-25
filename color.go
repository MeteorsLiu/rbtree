package rbtree

type Color uint8

const (
	RED Color = iota
	BLACK
)

func (c Color) String() string {
	if c == RED {
		return "RED"
	}
	return "BLACK"
}
