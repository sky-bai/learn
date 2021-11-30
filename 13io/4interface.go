package main

type A interface {
	M()
	Read() (int, error)
}

type B interface {
	N()
}
type c interface {
	Read() (int, error)
}

func M() {

}

func N() {

}

func copy(c c) {

}
func main() {
	copy(A)
}
