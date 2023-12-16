# int32linter

int32linter is a static analyzer to check usage of int32 variable.

# Install
to get int32linter execute this :

 $ go get github.com/beyzaekici/int32linter


 # Example

 func foo() {
	var x int32 = 25
	fmt.Println(x)
}

func main() {
	foo()
}

../main.go:11:8: int32 type is used, consider using int or int64

