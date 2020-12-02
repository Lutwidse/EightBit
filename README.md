### Example

```go
package main

import (
	"github.com/Lutwidse/EightBit"
	"fmt"
)

func main() {
	btc := EightBit.Get_btc_jpy()
	fmt.Println(btc.Sell)
	fmt.Println(btc.Buy)
	fmt.Println(btc.High)
	fmt.Println(btc.Low)
	fmt.Println(btc.Last)
	fmt.Println(btc.Vol)
	fmt.Println(btc.Timestamp)
}
```
