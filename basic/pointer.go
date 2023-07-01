package main

import (
	"log"
)

/**
 * ! (*)
 *      tanda asterisk (*) => digunakan untuk mendeklarasikan, mengakses nilai yang disimpan dilokasi memory
 * ! (&)
 *      tanda ampersand (&) => mengambil alamat memori dari suatu variable
 */

type brith struct {
	Place   string
	Country string
}

type User struct {
	Name  string
	Email string
	TTL   *brith
}

func Pointer() {
	// a := &brith{Place: "Tegal", Country: "Indoneisa"}
	// user := &User{Name: "Beta", Email: "andreas@gmail.com", TTL: a}

	// log.Println(user)
	// log.Printf("%+v\n", user)
	// log.Printf("%+v\n", a)

	X()

}

func X() {
	var ptr *int

	a := 4
	ptr = &a

	/**
	 * *ptr itu means mengambil dan mengembalikan value yang disimpan di alamat mmemori
	 * ptr itu akan mengembalikan nilai alokasi memori
	 */

	log.Println(ptr)

}
