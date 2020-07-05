package main

import (
  "fmt"
)

func main () {
  var h, w, k, pattern int
  var tmp string
  var c [6][6]rune

  fmt.Scanf("%d %d %d", &h, &w, &k)

  for i := 0; i < h; i++ {
    fmt.Scanf("%s", &tmp)
    for j, r := range tmp {
      c[i][j] = r
    }
  }
  var hlist, vlist [6]bool
  var maxv, maxh int
  for h1:=0; h1 < 2; h1++ {
    hlist[0] = h1== 1
    if h1== 1 {
      maxh = 1
    }
    for h2:=0; h2 < 2; h2++ {
      hlist[1] = h2== 1
      if h2== 1 {
        maxh = 2
      }
      for h3:=0; h3 < 2; h3++ {
        hlist[2] = h3== 1
        if h3== 1 {
          maxh = 3
        }    

        for h4:=0; h4 < 2; h4++ {
          hlist[3] = h4== 1
          if h4== 1 {
            maxh = 4
          }
      
          for h5:=0; h5 < 2; h5++ {
            hlist[4] = h5== 1
            if h5== 1 {
              maxh = 5
            }
        
            for h6:=0; h6 < 2; h6++ {
              hlist[5] = h6== 1
              if h6 == 1 {
                maxh = 6
              }
          
              for v1:=0; v1 < 2; v1++ {
                vlist[0] = v1== 1
                if v1== 1 {
                  maxv = 1
                }
            
                for v2:=0; v2 < 2; v2++ {
                  vlist[1] = v2== 1
                  if v2== 1 {
                    maxv = 2
                  }
  
                  for v3:=0; v3 < 2; v3++ {
                    vlist[2] = v3== 1
                    if v3== 1 {
                      maxv = 3
                    }
                    for v4:=0; v4 < 2; v4++ {
                      vlist[3] = v4== 1
                      if v4== 1 {
                        maxv = 4
                      }
    
                      for v5:=0; v5 < 2; v5++ {
                        vlist[4] = v5== 1
                        if v5== 1 {
                          maxv = 5
                        }
      
                        for v6:=0; v6 < 2; v6++ {
                          vlist[5] = v6== 1
                          if v6== 1 {
                            maxv = 6
                          }
        
                          if maxh > h || maxv > w {
                            continue
                          }
                          if k == resultfunc(c, hlist, vlist) {
                            pattern++
                          }
                        }            
                      }
                    }
                  }
                }
              }
            }            
          }
        }
      }
    }
  }
  fmt.Printf("%d\n", pattern)
}

func resultfunc(base [6][6]rune, hlist, vlist [6]bool) int {
  for i := 0; i < 6; i++ {
    for j := 0; j < 6; j++ {
      if hlist[i] || vlist[j] {
        base[i][j] = '='
      }
    }
  }
  return countBlack(base)
}

func countBlack(c [6][6]rune) int {
  const TARGET = '#'
  var b int
  for i := 0; i < 6; i++ {
    for j := 0; j < 6; j++ {
      if c[i][j] == TARGET {
        b++
      }
    }
  }
  return b
}
