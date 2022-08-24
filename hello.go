package main

import "fmt"

type Person struct {
	name string
}

func learn() {
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "genap")
		} else {
			fmt.Println(i, "ganjil")
		}
	}

	//slice & for loop
	var classFriends = []string{"Aulia Nurhady", "San Sayidul Akdam Agusta", "Wicaksana Pratama", "Kadek Anjasmara", "Muhammad Irfan Maulana", "Adhitya Giva Muhammad", "Iqbal Hamdani", "Fahmi Tajuddin", "Fajar Agus Maulana", "Ziyadatun Afisa"}

	for i, friend := range classFriends {
		fmt.Println("class friend no", i+1, friend)
	}

	//closure & pointer
	listFriend := func (l []*string) {
		for i, friend := range l {
			fmt.Println("class friend no", i+1, friend)
		}
	}
	
	var pointerClassFriends = make([]*string, 10)

	// copy(&classFriends, pointerClassFriends)
	// append(pointerClassFriends, classFriends&...)

	for i, friend := range classFriends {
		pFriend := friend
		pointerClassFriends[i] = &pFriend
	}


	listFriend(pointerClassFriends)


	//closure + pointer + struck
	var person = []*Person {
			{name: "Aulia Nurhady"},
			{name: "San Sayidul Akdam Agusta"},
			{name: "Kadek Anjasmara"},
	}

	listStructPerson := func (l []*Person) {
		for i, friend := range l {
			fmt.Println("class friend no", i+1, friend.name)
		}
	}

	listStructPerson(person)
}

// func main() {
// 	// learn()
// 	Bio()
// }
