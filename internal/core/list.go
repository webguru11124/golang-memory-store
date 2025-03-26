package core

type List struct {
	values []interface{}
}

func NewList() *List {
	return &List{values: make([]interface{}, 0)}
}

func (l *List) Push(value interface{}) {
	l.values = append(l.values, value)
}

func (l *List) Pop() (interface{}, bool) {
	if len(l.values) == 0 {
		return nil, false
	}
	val := l.values[len(l.values)-1]
	l.values = l.values[:len(l.values)-1]
	return val, true
}

func (l *List) GetAll() []interface{} {
	return l.values
}
