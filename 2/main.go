package main

import ("fmt"
)

type myIntList struct {
  head_ int;
  tail_ *myIntList
}

func (l *myIntList) head() int {
  return l.head_
}
func (l *myIntList) tail() *myIntList {
  return l.tail_
}
func (l *myIntList) prepend(n int) myIntList {
  return myIntList { 
    head_: n,
    tail_: l,
  }
}
func main(){
  var a = myIntList {
    head_: 1,
  }
  var b = a.prepend(2)
  fmt.Printf("%d", b.head())
}


