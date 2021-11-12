// main.go

package main

func main() {
	a := App{}
	a.Initialize("root", "Kamath@123", "godb_restapi")

	a.Run(":8080")
}
