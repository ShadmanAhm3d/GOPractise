package main

import (
  "fmt"
)



func (e *Employee) updateName (naya_name string) {
  e.name =  naya_name;
}



func updateKaro ( p *Employee, naya_age int){
  p.age =naya_age;
}



func createaNewEmployee(name string,age int , isRemote bool) *Employee{ //referencing employee struct

  var nayaLaunda Employee; //creating new Employee
  nayaLaunda.name  = name;
  nayaLaunda.age =age;
  nayaLaunda.isRemote = isRemote;

  return &nayaLaunda; // returning the address of the new employee , hence this is why the return type of this function is *Employee
}


func main () {


  
  var employee1 Employee = Employee{
    name: "Shadman",
    age: 12,
    isRemote: true,
  }

  var newEmp *Employee = createaNewEmployee("Shdaman", 12, false);

  updateKaro(&employee1, 12 )

  fmt.Print(employee1);
  fmt.Print(newEmp); //what da hell this both works 
  fmt.Print(*newEmp); // this works aswell


}
