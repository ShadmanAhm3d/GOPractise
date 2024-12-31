package main

import (
  "fmt"
  "st/morestructs"
  
)


func (e *Employee) updateName (naya_Name string) {
  e.Name =  naya_Name;
}



func updateKaro ( p *Employee, naya_Age int){
  p.Age =naya_Age;
}



func createaNewEmployee(Name string,Age int , IsRemote bool) *Employee{ //referencing employee struct

  var nayaLaunda Employee; //creating new Employee
  nayaLaunda.Name  = Name;
  nayaLaunda.Age =Age;
  nayaLaunda.IsRemote = IsRemote;

  return &nayaLaunda; // returning the address of the new employee , hence this is why the return type of this function is *Employee
}


func main () {

/*
  
  var employee1 Employee = Employee{
    Name: "Shadman",
    Age: 12,
    IsRemote: true,
  }

  var newEmp *Employee = createaNewEmployee("Shdaman", 12, false);

  updateKaro(&employee1, 12 )

  fmt.Print(employee1);
  fmt.Print(newEmp); //what da hell this both works 
  fmt.Print(*newEmp); // this works aswell




  fmt.Println("-----------------------------------------------------------------")


  var post1 morestructs.Post = morestructs.Post{
    Typeof: "descriptionPost",
    MemSize: 12.33,
    Year: 1221,
  }

   post2 := morestructs.Post{
    Typeof: "descriptionPost",
    MemSize: 12.33,
    Year: 1221,
  }
  
  fmt.Println(post1);
  fmt.Println(post2);

*/

  var result *morestructs.NameandAge = morestructs.SomeFunctionName(12,"Shadman");

  fmt.Printf("This is my adresss : %p" , &(*result))
  fmt.Println();
  fmt.Printf("This is my value   : %v", *result);



}
