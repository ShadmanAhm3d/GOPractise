package main

import "fmt"


func someFunction (arr *[4]int){
  arr[2] = 2;
  fmt.Printf("Likh diya");
}

func main(){

  var arr[4] int ;

  for i :=0; i <4 ; i++ {
    fmt.Printf("Enter element number %d : ", i);
    //input element dynamically ? 
    fmt.Scan(&arr[i]);
  }

  for i:=0 ; i<len(arr);i++{
    fmt.Printf("%d",arr[i]);
    fmt.Print("\n");
  }

  someFunction(&arr) //why am i getting error here 

  fmt.Println(arr[2])

}
