// Package file contains an example of a prototype interface.
// It has the methods Print and Clone.
// Clone is what makes this example interesting.
package file

import "fmt"

// Node is an interface used for File and Folder.
type Node interface {
	Print(string)
	Clone() Node
}

// File contains a fileName and implements Node.
type File struct {
	Name string
}

// Folder contains a folder structure and implements Node.
type Folder struct {
	Children []Node
	Name     string
}

// Local helper function.
func cloneName(name string) string {
	return name + "_clone"
}

// Local helper function.
func increaseIndentation(indentation string) string {
	return "   " + indentation
}

// Print outputs file to Stdout and is part of the Node interface.
func (f *File) Print(indentation string) {
	fmt.Println(indentation + f.Name)
}

// Clone copies the content of a file and is part of the Node interface.
func (f *File) Clone() Node {
	return &File{Name: cloneName(f.Name)}
}

// Print outputs folder to Stdout and is part of the Node interface.
func (f *Folder) Print(indentation string) {
	fmt.Println(indentation + f.Name)
	for _, child := range f.Children {
		child.Print(increaseIndentation(indentation))
	}
}

// Clone copies the content of a folder and is part of the Node interface.
func (f *Folder) Clone() Node {
	f2 := &Folder{Name: cloneName(f.Name)}
	c2 := []Node{}
	for _, child := range f.Children {
		c2 = append(c2, child.Clone())
	}
	f2.Children = c2
	return f2
}
