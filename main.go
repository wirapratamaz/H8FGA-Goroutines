package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	// //* error
	var number int
	var err error

	//case error 1
	number, err = strconv.Atoi("123456H") //false

	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	number, err = strconv.Atoi("123") //true

	//case error 2
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	//* custom error
	var password string

	fmt.Print("Masukkan password :")
	fmt.Scanln(&password)

	// cek password dengan fungsi 'validPassword()'
	if valid, err := validPassword(password); err != nil {
		//error message ke console
		fmt.Println(err.Error())
	} else {
		//password valid
		fmt.Println(valid)
	}

}

// * custom error
// validasi password
func validPassword(password string) (string, error) {
	// mendapatkan panjang password
	pl := len(password)

	// Jika panjang password kurang dari 5 karakter
	if pl < 5 {
		return "", errors.New("password has to have more than 4 characters")
	}
	// Jika password valid
	return "Valid password", nil
}
