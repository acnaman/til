
package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
  var n int
  var ac, wa, tle, re int
	fmt.Scanf("%d", &n)

	sc.Split(bufio.ScanWords)

	for i := 0; i < n; i++ {
    s := nextString()
    switch s {
      case "AC":
        ac++
      case "WA":
        wa++
      case "TLE":
        tle++
      case "RE":
        re++
    }
	}

	fmt.Printf("AC x %d\nWA x %d\nTLE x %d\nRE x %d\n", ac, wa, tle, re)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

