package main

func main() {
	var a int
	a = 1
	var b int64
	b = 1
	test := map[string]interface{}{"camer": a, "pic": b}
	for k, v := range test {
		if v == 1 {
			print(k)
		}
	}
	type class struct {
		name string
	}
	type classPri struct {
		Class    class
		priority int
	}
	stuMap := map[string]classPri{}

	stuMap["aaa"] = classPri{
		Class:    class{name: "zhang"},
		priority: 0,
	}
	stuMap["bbb"] = classPri{
		Class:    class{name: "gaoyi"},
		priority: 0,
	}

	stuMap["aaa"] = classPri{
		Class:    class{name: "gaosi"},
		priority: 0,
	}

	for k, v := range stuMap {
		print(k, v.Class.name, "\n")
	}
}
