package core

import (
	"testing"
)

func TestListPush(t *testing.T) {
	list := NewList()

	list.Push("item1")
	list.Push("item2")
	list.Push("item3")

	if len(list.GetAll()) != 3 {
		t.Errorf("Expected list size 3, got %d", len(list.GetAll()))
	}
}

func TestListPop(t *testing.T) {
	list := NewList()

	list.Push("item1")
	list.Push("item2")

	item, found := list.Pop()
	if !found || item != "item2" {
		t.Errorf("Expected 'item2', got %v", item)
	}

	item, found = list.Pop()
	if !found || item != "item1" {
		t.Errorf("Expected 'item1', got %v", item)
	}

	_, found = list.Pop()
	if found {
		t.Error("Expected no items left in the list")
	}
}

func TestListGetAll(t *testing.T) {
	list := NewList()
	list.Push("item1")
	list.Push("item2")

	items := list.GetAll()
	if len(items) != 2 {
		t.Errorf("Expected list size 2, got %d", len(items))
	}

	if items[0] != "item1" || items[1] != "item2" {
		t.Errorf("Expected [item1, item2], got %v", items)
	}
}

func TestListEmpty(t *testing.T) {
	list := NewList()

	_, found := list.Pop()
	if found {
		t.Error("Expected false, but got true for popping an empty list")
	}

	items := list.GetAll()
	if len(items) != 0 {
		t.Errorf("Expected empty list, got %v", items)
	}
}
