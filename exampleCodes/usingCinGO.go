package main
//#include<stdio.h>
//int calc(){printf("hello, world!"); return 1;}
import "C"
import "fmt"
func main(){
C.calc();
}
