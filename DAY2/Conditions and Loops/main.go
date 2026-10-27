package main
import "fmt"
func main(){
	a:=10
	b:=14
	c:=4
	if (a>b && a>c) {
		fmt.Println("a is greater")
	}else if(b>a && b>c){
		fmt.Println("b is greater")
	}else{
		fmt.Println("c is greater")
	}

	for i:=0;i<5;i++{
		fmt.Println(i)
	}
}