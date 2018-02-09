package main
import (
. "fmt"
)
type Class struct{
strClassname string
nClassNum int
strAdress string
next *Class
}
func main(){
sc := Class{"one",200,"beijing",nil}
Println(sc)
}
