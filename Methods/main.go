package main

import "fmt"

func main() {

	fmt.Println("welcome to structs")

	Aghul := Ras{"Aghul", "Aghulras@gmail.com", true, 23}

	fmt.Println(Aghul)
	fmt.Printf("My details are: %+v\n", Aghul)
	Aghul.GetStatus()
	Aghul.NewMail()
}

//defining struct in golang
//notice structs Ras has first letter capitalized
type Ras struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (r Ras) GetStatus() {
	fmt.Println("Is user active:", r.Status)

}

//it does not change the values of Ras, passes only the copy....hence importance of pointers
func (r Ras) NewMail() {
	r.Email = "archer@dev.go"
	fmt.Println("New mail is:", r.Email)
}
