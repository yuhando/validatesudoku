package main

// ItemsSet ..
type ItemsSet struct {
	items map[int]struct{}
}

// Add new item
func (set *ItemsSet) Add(item int) *ItemsSet {
	if set.items == nil {
		set.items = make(map[int]struct{})
	}
	_, itemExist := set.items[item]
	if !itemExist {
		set.items[item] = struct{}{}
	}

	return set
}

// Exist item
func (set *ItemsSet) Exist(item int) bool {
	_, exist := set.items[item]
	return exist
}

// Clear item
func (set *ItemsSet) Clear() {
	set.items = make(map[int]struct{})
}
