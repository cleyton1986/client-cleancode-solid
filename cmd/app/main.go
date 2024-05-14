package main

import routes "github.com/cleyton1986/client-cleancode-solid/api/routers"

func main() {
	router := routes.SetupRouter()
	router.Run(":3333") // Inicia o servidor na porta 3333
}
