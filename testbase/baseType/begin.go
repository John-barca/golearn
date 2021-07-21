package main

import "fmt"

// test slice
/*
func main () {
	var sli_1 [] int	// nil 切片
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(sli_1), cap(sli_1), sli_1);

	var sli_2 = [] int {} // 空切片
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(sli_2), cap(sli_2), sli_2);

	var sli_3 = [] int {1, 2, 3, 4, 5}
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(sli_3), cap(sli_3), sli_3);

	sli_4 := [] int {1, 2, 3, 4, 5}
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(sli_4), cap(sli_4), sli_4);

	var sli_5 [] int = make([] int, 5, 8)
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(sli_5), cap(sli_5), sli_5);

	sli_6 := make([] int, 5, 9)
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(sli_6), cap(sli_6), sli_6);
}
*/

// 截取切片
/*
func main() {
	sli := [] int {1, 2, 3, 4, 5, 6}
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(sli), cap(sli), sli);

	fmt.Println("sli[1] ==", sli[1])
	fmt.Println("sli[:] ==", sli[:])
	fmt.Println("sli[1:] ==", sli[1:])
	fmt.Println("sli[:4] ==", sli[:4])

	fmt.Println("sli[0:3] ==", sli[0:3])
	fmt.Printf("len=%d, cap=%d, slice[0:3]=%v\n", len(sli[0:3]), cap(sli[0:3]), sli[0:3])

	fmt.Println("sli[0:3:4] ==", sli[0:3:4])
	fmt.Printf("len=%d, cap=%d, slice[0:3:4]=%v\n", len(sli[0:3:4]), cap(sli[0:3:4]), sli[0:3:4])
}
*/

/*
// struct test
type Person struct {
	Name string
	Age int
}

func main () {
	var p1 Person
	p1.Name = "aaa"
	p1.Age = 20
	fmt.Println("p1 =", p1)

	var p2 = Person{Name: "Burke", Age: 31}
	fmt.Println("p2 =", p2)

	p3 := Person{Name: "Aaron", Age: 32}
	fmt.Println("p3 =", p3)

	// 匿名
	p4 := struct {
		Name string
		Age int
	} {Name: "niming", Age: 33}
	fmt.Println("p4 =", p4)
}
*/

/*
// JSON 生成
type Result struct {
	Code int `json:"code"`
	Message string `json:"msg"`
}

func main() {
	var res Result
	res.Code = 200
	res.Message = "success"

	// 序列化
	jsons, errs := json.Marshal(res)
	if errs != nil {
		fmt.Println("json marshal error:", errs)
	}
	fmt.Println("json data :", string(jsons))
	fmt.Printf("jsons type is %s\n", reflect.TypeOf(jsons));

	// 反序列化
	var res2 Result
	errs = json.Unmarshal(jsons, &res2)
	if errs != nil {
		fmt.Println("json unmarshal error:", errs)
	}
	fmt.Println("res2 :", res2)
	fmt.Printf("res2 type is %s\n", reflect.TypeOf(res2));
}
*/

/*
// update json data
type Result struct {
	Code int `json:"code"`
	Message string `json:"msg"`
}

func setData (res *Result) {
	res.Code = 500
	res.Message = "fail"
}

func toJson (res *Result) {
	jsons, errs := json.Marshal(res)
	if errs != nil {
		fmt.Println("json marshal error:", errs)
	}
	fmt.Println("json data :", string(jsons))
}

func main () {
	var res Result
	res.Code = 200
	res.Message = "success"
	toJson(&res)

	setData(&res)
	toJson(&res)
}
*/

// Map test
// no sort key-value
/*
func main() {
	var p1 map[int]string
	p1 = make(map[int]string)
	p1[1] = "Tom"
	fmt.Println("p1 :", p1)

	var p2 map[int]string = map[int]string{}
	p2[1] = "Tom"
	fmt.Println("p2 :", p2)

	var p3 map[int]string = make(map[int]string)
	p3[1] = "Tom"
	fmt.Println("p3 :", p3)

	p4 := map[int]string{}
	p4[1] = "Tom"
	fmt.Println("p4 :", p4)

	p5 := make(map[int]string)
	p5[1] = "Tom"
	fmt.Println("p5 :", p5)

	p6 := map[int]string{
		1 : "Tom",
	}
	fmt.Println("p6 :", p6)
}
*/

/*
// Map 生成 JSON
func main() {
	res := make(map[string]interface{})
	res["code"] = 200
	res["msg"] = "success"
	res["data"] = map[string]interface{}{
		"username": "Tom",
		"age"			: "10",
		"hobby"		: []string{"walk", "play"},
	}
	fmt.Println("map data :", res)

	// 序列化
	jsons, errs := json.Marshal(res)
	if errs != nil {
		fmt.Println("json marshal error:", errs)
	}
	fmt.Println("")
	fmt.Println("--- map to json ---")
	fmt.Println("json data :", string(jsons))

	// 反序列化
	res2 := make(map[string]interface{})
	errs = json.Unmarshal([]byte(jsons), &res2)
	if errs != nil {
		fmt.Println("json marshal error:", errs)
	}
	fmt.Println("")
	fmt.Println("--- json to map ---")
	fmt.Println("map data :", res2)
}
*/

// editor or delete Map
/*
func main() {
	person := map[int]string{
		1: "Tom",
		2: "Aaron",
		3: "John",
	}
	fmt.Println("data :", person)

	delete(person, 2)
	fmt.Println("data :", person)

	person[2] = "Jack"
	person[3] = "Kevin"
	fmt.Println("data :", person)
}
*/

// always control
func main() {
	person := [3] string {"Tom", "Aaron", "John"}
	fmt.Printf("len=%d, cap=%d, array=%v\n", len(person), cap(person), person)
	fmt.Println("")
	for k, v := range person {
		fmt.Printf("person[%d]: %s\n", k, v)
	}
	fmt.Println("")
	for i := range person {
		fmt.Printf("person[%d]: %s\n", i, person[i])
	}
	fmt.Println("")
	for i := 0; i < len(person); i++ {
		fmt.Printf("person[%d]: %s\n", i, person[i])
	}
	fmt.Println("")
	// 使用空白符
	for _, name := range person {
		fmt.Println("name :", name)
	}
}