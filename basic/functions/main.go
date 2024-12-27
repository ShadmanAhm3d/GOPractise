package main

import(
  "fmt"
)

func addSub (a int, b int )(sum int, diff int){

  sum = a+b;
  diff =a-b;

  return
}

//All about functions
func main(){

  fmt.Println(addSub(9, 3));
  fmt.Println(func(a ,b int )int {
    return a+b;
  }(4,5))
}
