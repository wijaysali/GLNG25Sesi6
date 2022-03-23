package main

import "ginproject/routers"

func main(){
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}

