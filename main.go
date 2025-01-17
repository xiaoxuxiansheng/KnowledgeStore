package main

func main() {
	if err := GetServer().Run(); err != nil {
		panic(err)
	}
}
