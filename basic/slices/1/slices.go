package main

import "fmt"

func main() {

  
  var n int ;
  fmt.Scan(&n)

  slice := []int{10, 20, 30, 40, 50}

  for i:=0; i < n;i++{
    var val int ;
    fmt.Scan(&val);
    slice = append(slice, val)
  }


  for _,elem := range slice{

    fmt.Println(elem);

  }
  fmt.Println(" Lenghy : ", len(slice))
}
