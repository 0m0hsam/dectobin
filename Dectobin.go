# dectobin
package main

 import (
 "fmt"
 )

 func main() {  
  fmt.Println("Program to convert input into \nBinary, Octal and Decimal")
 fmt.Println("Enter your number conversion command \n Binary Input: B \n Octal Input: O \n Decimal Input: D")
 var cmd string 
  fmt.Scan(&cmd)
     if cmd =="B"{
 fmt.Print("Input value to convert to Binary:")
 var num int
  fmt.Scan(&num)
 bin := fmt.Sprintf("%b", num)
 fmt.Println("binary value of",num,":",bin)
            }else if cmd =="O"{
fmt.Print("Input value to convert to Octal:")
var num int
  fmt.Scan(&num)
            oct := fmt.Sprintf("%o",num)
fmt.Println("binary value of",num,":",oct)
            } else if cmd =="D"{
              fmt.Print("Error: Pls check later \n this Section of Program is under maintainace and development \n Try others!!!")
              //fmt.Print("Input value to convert to Decimal:")
//var num int
 // fmt.Scan(&num)
               //dec := fmt.Sprintf("%d",num)
//fmt.Print("Input value to convert to Decimal", num, ":", dec)    
            } else {
            fmt.Print("Pls enter the right command!!!")
            }
 }

