# UniMorph

```go
package main

import (
	"fmt"
	"github.com/jonasknobloch/unimorph"
	"os"
	"path/filepath"
)

func main() {
	u := &unimorph.Unimorph{}

	_ := u.Init(filepath.Join(os.Getenv("HOME"), ".unimorph/ces/ces"))

	s1, _ := u.Split("afghánský", unimorph.SplitSiblings(1)) // [afghán sk ý]
	s2, _ := u.Split("afghánský", unimorph.SplitSiblings(2)) // [afghánsk ý]
	s3, _ := u.Split("afghánský", unimorph.SplitSiblings(3)) // [afghánsk ý]
	s4, _ := u.Split("afghánský", unimorph.SplitSiblings(4)) // [afghánský]
	
	fmt.Println(s1, s2, s3, s4)
}
```

![radix_afghansky.svg](assets/radix_afghansky.svg)
