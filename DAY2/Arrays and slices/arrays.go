package main
import "fmt"

//arrays

// slicing is possible

// func main(){
// 	var arr1=[5]int{1,2,3,4,5}
// 	arr2:=[3]string{"abc","def","ghi"}
// 	fmt.Println(arr1)
// 	fmt.Println(arr2)

// 	a1:=[...]int{1,2,3}
// 	fmt.Println(a1)
// 	fmt.Println(arr2[0])
// 	a1[1]=4
// 	fmt.Println(a1)

// 	a2:=[5]int{}
// 	fmt.Println(a2)

// 	// initlaising only specific element
// 	a3:=[5]int{1:10,3:5}
// 	fmt.Println(a3)
// 	fmt.Println(len(a3))
// }

// slices
func main(){
	s1:=[]int{1,2,4}
	fmt.Println(s1)
	fmt.Println(len(s1),"and",cap(s1))

	// creating slice from array

	arr1:=[5]int{1,2,3,4,5}
	fmt.Printf("%v\n",arr1)
	s2:=arr1[2:4]
	fmt.Printf("%v\n",s2)
	fmt.Printf("%v\n",len(s2))
	// cap 3 as slice can grow tell the end of the array that is why
	fmt.Printf("%v\n",cap(s2))

	// make function
	s3:=make([]int,5,10)
	// or
	s4:=make([]int,5) // capacity==length
	fmt.Println(s3)
	fmt.Println(s4)

	// accessing changing element are same as array
	// append elements to slices
	s_1:=[]int{1,2,3,4,5}
	fmt.Printf("%v\n",s_1)
	fmt.Printf("%v\n",len(s_1))
	fmt.Printf("%v\n",cap(s_1))

	s_1=append(s_1,6,7,8)
	fmt.Printf("%v\n",s_1)
	fmt.Printf("%v\n",len(s_1))
	fmt.Printf("%v\n",cap(s_1))

	// appending one slice to another
	s3=append(s2,s1...)
	fmt.Printf("%v\n",s3)
	fmt.Printf("%v\n",len(s3))
	fmt.Printf("%v\n",cap(s3))

	// copy funtion
	n:=[]int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
	fmt.Printf("%v\n",n)
	fmt.Printf("%v\n",len(n))
	fmt.Printf("%v\n",cap(n))

	n1:=n[:len(n)-10]
	fmt.Printf("%v\n",n1)
	fmt.Printf("%v\n",len(n1))
	fmt.Printf("%v\n",cap(n1))
	n2:=make([]int,len(n1))
	copy(n2,n1)
	fmt.Printf("%v\n",n2)
	fmt.Printf("%v\n",len(n2))
	fmt.Printf("%v\n",cap(n2))
}