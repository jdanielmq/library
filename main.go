package main

import (
	"fmt"
	"library/animal"
	"library/book"
)

func main() {
	// creando una instancia de book.Book
	var myBook = book.Book{
		Title:  "Moby Dick",
		Author: "Herman Melville",
		Pages:  300,
	}
	myBook.PrintInfo()
	fmt.Println("###############################")

	//aca estamos creando una instancia, donde se ultiza el metodo set para modificar el texto, luego se imprimer la informacion
	mybook1 := book.NewBook("Moby Dick 2", "Herman Malville", 450)
	mybook1.PrintInfo()
	mybook1.SetTitle("Moby Dick 3")
	fmt.Println(mybook1.GetTitle())
	mybook1.PrintInfo()
	fmt.Println("###############################")

	//aca estamos crerando un nueva instancia, donde el objeto newTextBook, tiene como objeto includio la estructura de Book mas 2 atributos mas
	myTextBook := book.NewTextBook("comunicación", "jaime gammara", 345, "Santillana Chile", "primaria")
	myTextBook.PrintInfo()
	fmt.Println("###############################")

	// aca vamos crear 2 instancias distintas y vamos a utilizar la interfaces
	// y donde ambas instancias ocuparan el print para gatillar el PrintInfo de cada instanacia.
	mybook2 := book.NewBook("Moby Dick 3", "Herman Malville R.", 150)
	myTextBook2 := book.NewTextBook("Inteligencia Emocional", "Juan daniel", 500, "Santillana Chile", "Liderazgo")

	book.Print(mybook2)
	book.Print(myTextBook2)
	fmt.Println("###############################")

	// vamos a crear 2 instancias de la clase Animal y su clasificacion
	miPerro := animal.Perro{Nombre: "Florencia"}
	miGato := animal.Gato{Nombre: "Akiva"}

	animal.HcerSonido(&miPerro)
	animal.HcerSonido(&miGato)
	fmt.Println("###############################")
	//vamos a crear una slice, arrglo o lista, donde vamos a crear varios objetos de tipo animal.
	animales := []animal.Animal{
		&animal.Perro{Nombre: "Nala"},
		&animal.Gato{Nombre: "Alis"},
		&animal.Perro{Nombre: "Florencia"},
		&animal.Gato{Nombre: "Akiva"},
		&animal.Gato{Nombre: "Olaf"},
	}
	for _, animal := range animales {
		animal.Sonido()
	}
	fmt.Println("###############################")

}
