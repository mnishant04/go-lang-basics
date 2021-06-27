package main

import "fmt"

func main(){
  ids :=[]int{1,2,3,4,5,6,7,8,9}
  for i,id:=range ids{   // id = ids[i] i<len(ids)-1  ,   id is int variable not array
	  fmt.Printf("%d - ID = %d \n",i,id) 
  }
// when index is not required
  for _,id:=range ids{  // error will be generated when we use i insted of _
	fmt.Printf("ID = %d \n",id) 
  }

  //sum and product of ids
sum:=0
pro:=1
  for _,id:=range ids {
	sum +=id
	pro *=id
}
  fmt.Printf("sum of numbers are : %d ,product of numbers of : %d\n",sum,pro)
//range with maps
student := make(map[string]string)

//assigning value to map
student["nishant"] = "900000000"
student["pradeep"] = "900000001"
student["sangam"] = "900000002"
for key,value:=range student{
	fmt.Printf("%s:%s\n",key,value)
} 
for key:=range student{
	fmt.Println("name : " +key)
} 




}