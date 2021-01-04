package main

import "github.com/andreasatle/go-design-patterns/01-creational-design-patterns/05-prototype/file"

func main() {
	file1 := &file.File{Name: "File1"}
	file2 := &file.File{Name: "File2"}
	file3 := &file.File{Name: "File3"}

	folder1 := &file.Folder{
		Children: []file.Node{file1},
		Name:     "Folder1",
	}

	folder2 := &file.Folder{
		Children: []file.Node{folder1, file2, file3},
		Name:     "Folder1",
	}

	folder2.Print("")
	folder3 := folder2.Clone()
	folder3.Print("")

	folder4 := &file.Folder{
		Children: []file.Node{folder1, folder2, folder3, file1, file3},
		Name:     "Deeper structure",
	}
	folder5 := folder4.Clone()
	folder5.Print("")
}
