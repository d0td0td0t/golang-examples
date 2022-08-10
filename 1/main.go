package main

import ("fmt";
)
type user struct {
  name string 
  age int
}

func main(){
  var a user
  var b []user
  a = addUser("jopa", 666)
  b = append(b,a)
  b[0].setAge(669)
  fmt.Printf("%s %d \n", b[0].name, b[0].getAge())
}

func (u user) getAge () int {
  return u.age
}
func (u *user) setAge (age int) {
  u.age = age
}
func addUser(name string, age int) user {
  return user {
    name: name,
    age: age,
  }
}
