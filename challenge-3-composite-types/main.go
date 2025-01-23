package main

import (
	"fmt"
	"strings"
)

type author struct {
	name string
}

type book struct {
	title  string
	author author
}

type library map[string][]book

func (l library) addBook(b book) {
	loweredName := strings.ToLower(b.author.name)
	l[loweredName] = append(l[loweredName], b)
}

func (l library) lookupByAuthorName(authorName string) []book {
	return l[strings.ToLower(authorName)]
}

func main() {
	// prepare data
	l := library{}
	l.addBook(book{
		title: "Not A Fake Book",
		author: author{
			name: "Alex Tan",
		},
	})

	l.addBook(book{
		title: "Not Another Fake Book",
		author: author{
			name: "Alex Tan",
		},
	})

	l.addBook(book{
		title: "Not Another Fake Book 2025 Edition",
		author: author{
			name: "Alex Tan",
		},
	})

	l.addBook(book{
		title: "A Fake Book",
		author: author{
			name: "Not Alex Tan",
		},
	})

	// perform search
	for _, lookupName := range []string{
		"Alex Tan",
		"Not Alex Tan",
		"Someone Else",
	} {
		ret := l.lookupByAuthorName(lookupName)

		fmt.Printf("listing book(s) by %s:\n", lookupName)
		if len(ret) == 0 {
			fmt.Println(" - nothing found")
		}
		for i := range ret {
			fmt.Printf(" - author: %s, title: %s\n", ret[i].author.name, ret[i].title)
		}
	}
}
