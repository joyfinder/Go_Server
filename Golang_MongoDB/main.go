package main

func main() {

	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
}
