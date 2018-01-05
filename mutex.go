package main

import (
	"sync"
	"fmt"
)

type Table struct {
	lock   sync.RWMutex
	onLine map[string]string
}

var Manage *Table

func init(){
	Manage = &Table{
		onLine:map[string]string{},
	}
}

func (t *Table) insert(key, value string) {
	t.lock.Lock()

	t.onLine[key] = value

	t.lock.Unlock()

	Manage.getTable()
}

func (t *Table) del(key string) {
	t.lock.Lock()

	if k, ok := t.onLine[key]; ok {
		fmt.Printf("delete: %s\n", k)
		delete(t.onLine, key)
	}
	// defer t.lock.Unlock()
	t.lock.Unlock()

	Manage.getTable()
}

func (t *Table) getTable() {
	t.lock.RLock()
	defer t.lock.RUnlock()

	for k, v := range t.onLine {
		fmt.Println(k,":" , v)

	}
	fmt.Println("\n")
}

func main() {
	Manage.insert("a", "z")
	Manage.insert("b", "y")
	Manage.insert("c", "w")
	Manage.del("a")
}