package main

import ("fmt"
)

type myIntList struct {
  data []int
}

func (l *myIntList) head() int {
  return l.data[0]
}
func (l *myIntList) tail() myIntList {
  var ll = myIntList{
    data: l.data[1:],
  }
  return ll
}
func (l *myIntList) prepend(n int) myIntList {
  var nn = myIntList{}
  nn.data = append([]int {n} ,l.data...)
  return nn
}
func main(){
  var a = myIntList{}
  a.data = [] int {1}
  var b = a.prepend(2)
  fmt.Printf("%d", b.head())
}


