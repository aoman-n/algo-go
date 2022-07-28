package hashtable

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"math/big"
)

type Pair struct {
	key   string
	value string
}

func (p *Pair) SetValue(val string) {
	p.value = val
}

func NewPair(key, val string) *Pair {
	return &Pair{
		key:   key,
		value: val,
	}
}

type HashTable struct {
	table [][]*Pair
}

func NewHashTable() *HashTable {
	table := make([][]*Pair, 10)
	return &HashTable{
		table: table,
	}
}

func (h *HashTable) toHash(key string) uint64 {
	bi := big.NewInt(0)
	m := md5.New()
	io.WriteString(m, key)
	hexStr := hex.EncodeToString(m.Sum(nil))
	bi.SetString(hexStr, 16)
	return bi.Uint64() % 10
}

func (h *HashTable) Add(key, val string) {
	index := h.toHash(key)

	exists := false
	for _, data := range h.table[index] {
		if data.key == key {
			exists = true
			data.SetValue(val)
		}
	}

	if !exists {
		h.table[index] = append(h.table[index], NewPair(key, val))
	}
}

func (h *HashTable) Get(key string) string {
	index := h.toHash(key)

	for _, data := range h.table[index] {
		if data.key == key {
			return data.value
		}
	}

	return ""
}

func (h *HashTable) Print() {
	for i, data := range h.table {
		fmt.Printf("%v ", i)
		for _, p := range data {
			fmt.Printf("--> [%v, %v] ", p.key, p.value)
		}
		fmt.Println()
	}
}
