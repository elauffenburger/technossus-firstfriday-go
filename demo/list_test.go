package main

import (
	"github.com/elauffenburger/technossus-firstfriday-go/demo/list"
	"reflect"
	"testing"
)

func modifyArr(arr []int) {
	for i := range arr {
		val := arr[i]

		arr[i] = val + 1
	}
}

func printListSize(l *list.List, t *testing.T) {
	t.Logf("contents has %d elements", len(l.GetContents()))
}

func makeIntList() *list.List {
	// yes, this is really how you do it
	return list.ListFactory(reflect.TypeOf(int(0)))
}

func TestCreateAndAddToIntList(t *testing.T) {
	l := makeIntList()

	// print size
	printListSize(l, t)

	// add something and print size
	l.Add(1)
	l.Add(2)
	l.Add(3)

	printListSize(l, t)

	contents := l.GetContents()
	for i := range contents {
		t.Logf("%d\n", contents[i])
	}

	length := len(contents)
	if length != 3 {
		t.Errorf("Expected 3 items in contents, had %d", length)
	}
}

func TestAddBoolToIntList(t *testing.T) {
	l := makeIntList()

	defer func() {
		if r := recover(); r != nil {
			// recovered here, so the test passed
			t.Log("List successfully panicked when a bad data type was added to an int")
		}
	}()

	l.Add(false)

	t.Error("List should have failed to add a boolean to an integer list")
}

func TestArray(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	modifyArr(arr)

	t.Logf("arr: %v\n", arr)
}
