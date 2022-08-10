package main

import ("fmt"
)

type MyList[T any] struct {
  data []T
}

func (l *MyList[T]) head() T {
  return l.data[0]
}

func (l *MyList[T]) tail() MyList[T] {
  var ll = MyList[T]{
    data: l.data[1:],
  }
  return ll
}

func (l *MyList[T]) prepend(n T) MyList[T] {
  var nn = MyList[T]{}
  nn.data = append([]T {n} ,l.data...)
  return nn
}

func main(){
  var a = MyList[int]{}
  a.data = [] int {1}
  var b = a.prepend(2)
  fmt.Printf("%d", b.head())
}


