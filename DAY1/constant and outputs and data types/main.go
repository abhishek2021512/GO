package main
import "fmt"

const A int=1

func main(){
	const(
		B int=2
		C string="abc"
	)
	const D=14
	fmt.Println(A," ",B," ",C," ",D)
	fmt.Print(A,B,"\n")
	fmt.Println(B)

	var a string="abc"
	var b int=123
	var c float32=213.1245
	var d bool=true;

	fmt.Printf("a value is %v and type is %T\n",a,a)

	// int format
	fmt.Printf("%b\n",b)
	fmt.Printf("%d\n",b)
	fmt.Printf("%+d\n",b)
	fmt.Printf("%o\n",b)
	fmt.Printf("%O\n",b)
	fmt.Printf("%x\n",b)
	fmt.Printf("%X\n",b)
	fmt.Printf("%#x\n",b)
	fmt.Printf("%4d\n",b)
	fmt.Printf("%-4d\n",b)
	fmt.Printf("%04d\n",b)

	// string format
	fmt.Println("--------------------------------------------------")
	fmt.Printf("%s\n",a)
	fmt.Printf("%q\n",a)
	fmt.Printf("%8s\n",a)
	fmt.Printf("%-8s\n",a)
	fmt.Printf("%x\n",a)
	fmt.Printf("% x\n",a)
// bool fomrat
	fmt.Println("--------------------------------------------------")
	fmt.Printf("%t\n",d)
// float 
	fmt.Println("--------------------------------------------------")
	fmt.Printf("%e\n",c)
	fmt.Printf("%f\n",c)
	fmt.Printf("%.2f\n",c)
	fmt.Printf("%6.2f\n",c)
	fmt.Printf("%g\n",c)
}