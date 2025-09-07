package indexer

import (
	"fmt"
)

type Visitor interface {
	visitFolder(Node)
	visitFile(Node)
	// todo links and so on
}

type Node interface {
	Accept(Visitor)
	Name() string
	Type() NodeType
	SetParent(Node)
	GetParent() Node
	GetPath() string
}

// helper function to recursively print any node
func printNode(node Node, indent string) {
	switch n := node.(type) {
	case *Folder:
		fmt.Printf("%s+ Folder: %s\n", indent, n.Name())
		for _, child := range n.GetChildren() { // iterate folder children
			printNode(child, indent+"  ")
		}
	case *File:
		fmt.Printf("%s- File: %s (%s)\n", indent, n.Name(), n.Type())
	default:
		fmt.Printf("%s? Unknown node type\n", indent)
	}
}
