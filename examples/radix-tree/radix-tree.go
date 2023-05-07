package main

import (
	"fmt"
	"strings"
)

type RadixNode struct {
	label    string
	children map[string]*RadixNode
}

func NewRadixNode(label string) *RadixNode {
	return &RadixNode{
		label:    label,
		children: make(map[string]*RadixNode),
	}
}

type RadixTree struct {
	root *RadixNode
}

func NewRadixTree() *RadixTree {
	return &RadixTree{
		root: NewRadixNode(""),
	}
}

func longestPrefix(s1, s2 string) int {
	i := 0
	for i < len(s1) && i < len(s2) && s1[i] == s2[i] {
		i++
	}
	return i
}

func (t *RadixTree) Insert(key string) {
	node := t.root
	for len(key) > 0 {
		prefix := longestPrefix(node.label, key)

		if prefix == len(node.label) {
			key = key[prefix:]

			if len(key) == 0 {
				return
			}

			if child, ok := node.children[string(key[0])]; ok {
				node = child
			} else {
				newNode := NewRadixNode(key)
				node.children[string(key[0])] = newNode
				node = newNode
			}

		} else {
			child1 := NewRadixNode(node.label[prefix:])
			child2 := NewRadixNode(key[prefix:])

			for k, v := range node.children {
				child1.children[k] = v
				delete(node.children, k)
			}
			node.children[string(node.label[prefix])] = child1
			node.children[string(key[prefix])] = child2
			node.label = node.label[:prefix]

			return
		}
	}
}

func printNode(node *RadixNode, indent int) {
	var label string
	if node.label != "" {
		label = node.label
	} else {
		label = "root"
	}
	fmt.Printf("%s%s\n", strings.Repeat(" ", indent), label)
	for _, child := range node.children {
		printNode(child, indent+2)
	}
}

func printTree(tree *RadixTree) {
	printNode(tree.root, 0)
}

func main() {
	tree := NewRadixTree()
	tree.Insert("foo")
	tree.Insert("bar")
	tree.Insert("baz")
	tree.Insert("romane")
	tree.Insert("romanus")
	tree.Insert("romulus")
	tree.Insert("rubens")
	tree.Insert("ruber")
	printTree(tree)
}
