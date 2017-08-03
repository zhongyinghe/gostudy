package main

import "fmt"

func main() {
	u := struct {
		name string
		age  byte
	}{
		name: "Tom",
		age:  12,
	}

	type file struct {
		name string
		attr struct {
			owner int
			perm  int
		}
	}

	f := file{
		name: "test.ini",
		/*attr: {	//missing type in composite literal
			owner: 1,
			perm:  0755,
		},*/
	}

	f.attr.owner = 1
	f.attr.perm = 0755

	fmt.Println(u, f)
}
