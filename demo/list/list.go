package list

import (
	"reflect"
)

type List struct {
  contents []interface{}
  contentsType reflect.Type
}

func (l *List) Add(item interface{}) {
  itemType := reflect.TypeOf(item)

  if itemType != l.contentsType {
    panic("OH GOD NOOOOO")
  }

  l.contents = append(l.contents, item)
}

func (l *List) GetContents() []interface{} {
    return l.contents
}

func ListFactory(contentsType reflect.Type) *List {
	return &List{ contentsType: contentsType, contents: []interface{}{} }
}
