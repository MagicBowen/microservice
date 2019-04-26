package main

func main() {
	(&userRPCServer{}).startUp(":8899", createRepo("mongodb://mongodb:27017"))
}
