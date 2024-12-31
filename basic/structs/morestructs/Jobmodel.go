package morestructs

import "fmt"




type Post struct{
  Typeof string
  MemSize float64
  Year int 
}

type NameandAge struct {
  Name string
  Age int
}

func SomeFunctionName(age int,name string) *NameandAge {

  /* var returner string = fmt.Sprintf("Name: %s , Age : %d \n", name , age);
  return returner; */

  var somePersonNameandAge NameandAge = NameandAge{
    Name: "Shadman from NameandAge struct",
    Age: 12,
  }


  fmt.Printf("The adress of somePersonNameandAge  %p: ", &somePersonNameandAge)
  fmt.Println();
  return &somePersonNameandAge;

}
