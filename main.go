package main

import "fmt"

var decimalNumber = 2.62
var exist bool = true
var message string = "nama saya dobleh saya aslinya dua orang, salam kenal"
const usiaAwal int = 25
var totalUmur int = 0
var fruit = [4]string{"Apple", "Grape", " banana", "melon"}

type person struct{
	name string
	age int
}
type student struct{
	grade int
	person
	matakuliah string
}

func main() {
	fmt.Println("Hello world")
	fmt.Printf("Bilangan Desimal : %f\n", decimalNumber)
	fmt.Printf("Bilangan Desimal : %.3f\n", decimalNumber)
	fmt.Printf("exist? %t \n", exist)
	fmt.Println(message);
	totalUmur = usiaAwal+25
	fmt.Println(totalUmur)

	for i, fruit := range fruit {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}

	a, b := swap("hello", "world")
	fmt.Println(a,b)

	var  s1 = student{}
	s1.name = "dobleh"
	s1.age = 21
	s1.grade = 2 
	s1.matakuliah = "kalkulus"

	fmt.Println("nama : ", s1.name)
	fmt.Println("age : ", s1.person.age)
	fmt.Println("matakuliah : ",s1.matakuliah)

	var  allStudents = []person{
		{name: "dobleh", age: 17},
		{name: "jamal", age: 18},
		{name: "kabur", age: 19},
	}

	for _, student := range allStudents	{
		fmt.Println(student.name, "age is", student.age)
	}

}

func swap(x,y string) (string,string) {
 return y,x

}