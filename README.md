ljatom reads new public posts from LiveJournal via public atom stream interface.

# Install

```
go get github.com/belfinor/ljatom
```

# Usage

All you need is:

```go
package main

import (
	"fmt"

	"github.com/belfinor/ljatom"
)

func main() {

	for msg := range ljatom.Read() {

		fmt.Println("Time: " + msg.Created.String())
		fmt.Println("Journal: " + msg.Journal)
		fmt.Println("Post URL: " + msg.Url)
		fmt.Println("Post title: " + msg.Title)
		fmt.Print("Post body: " + msg.Content + "\n\n\n")

	}
}
```

and wait some times. The atom stream returns data by chunks.

*ljatom.Read* return chan of *Entry* refs.  Each *Entry* object has following structure:

```go
type Entry struct {
	Journal      string
	JournalTitle string
	Url          string
	Created      time.Time
	Title        string
	Content      string
}
```

Moreover, *ljatom.Read* makes reconnect if connection is broken. All you need is read chan of Entry objects.
