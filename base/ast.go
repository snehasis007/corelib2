package base

// type KEY interface {
// 	Read()
// 	ToString() string
// 	Less() bool
// 	Equals() bool
// }
// type KeyImpl struct {
// 	value int
// 	KEY
// }

// func (key *KeyImpl) Read() {
// 	key.Read()
// }

// type Item struct {
// 	key  KEY
// 	info interface{}
// }

// type Node struct {
// 	item  Item
// 	left  interface{}
// 	right interface{}
// }

type DataItem interface {
	Compare(data *DataItemImpl) int
}
type DataItemImpl struct {
	Index int
	Value interface{}
}

func (item *DataItemImpl) Compare(item2 *DataItemImpl) int {
	switch vType := item.Value.(type) {
	case int:
		if vType > item2.Value.(int) {
			return 1
		} else if vType < item2.Value.(int) {
			return -1
		} else {
			return 0
		}
	default:
		return -1111

	}

}
