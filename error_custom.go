package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"ID tidak boleh kosong"}
	}

	if id != "regi" {
		return &notFoundError{"Data tidak ditemukan"}
	}

	return nil
}

func main() {
	err := SaveData("budi", nil)
	if err != nil {
		// if validationErr, ok := err.(*validationError); ok {
		// 	fmt.Println("Validation Error:", validationErr.Error())
		// } else if notFoundErr, ok := err.(*notFoundError); ok {
		// 	fmt.Println("Not Found Error:", notFoundErr.Error())
		// } else {
		// 	fmt.Println("Error", err.Error())
		// }

		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("Validation Error:", finalError.Error())
		case *notFoundError:
			fmt.Println("Not Found Error:", finalError.Error())
		default:
			fmt.Println("Error", finalError.Error())
		}
	} else {
		fmt.Println("Data berhasil disimpan")
	}
}