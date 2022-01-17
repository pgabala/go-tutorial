package main

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	number     int `required number`
	actorName  string
	companions []string
}

func main() {
	// positional syntax not great for maintenance
	aDoctor := Doctor{
		3,
		"Jon Pertwee",
		[]string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}

	fmt.Println(aDoctor)
	tag := reflect.TypeOf(Doctor{})
	field, _ := tag.FieldByName("number")
	fmt.Println(field)

	// with field names much safer
	bDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}

	fmt.Println(bDoctor)
	fmt.Println(bDoctor.actorName)

	// anonoymous struct
	cDoctor := struct{ name string }{name: "John Pertwee"}
	fmt.Println(cDoctor)
}
