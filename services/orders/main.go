package main

func main() {
	httpServer := NewHttpServer(":8000")
	go httpServer.Run()

	gRPCServer := newGRPCServer(":9000")
	gRPCServer.Run()
}
