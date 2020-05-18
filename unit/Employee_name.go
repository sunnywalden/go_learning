package unit

import "fmt"

type Employee struct {
	Name         string
	Age          int
	SelfDescribe string `info:"self describe"`
}

func Capitalize(str string) string {
	var upperStr string
	vv := []rune(str)   // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 {  // 后文有介绍
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				fmt.Println("Not begins with lowercase letter,")
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}

func Capitalized(str string) string {
	StrHead := str[0]
	if StrHead >= 97 && StrHead <= 122 { // 后文有介绍
		StrHead -= 32
	}
	StrTail := str[1:]
	return string(StrHead) + StrTail
}

func (ep *Employee) GetEmployeeName() string {
	return Capitalize(ep.Name)
}

func (ep *Employee) GetEmployeeAge() int {
	return ep.Age
}

func (ep *Employee) SetEmployeeAge(age int) {
	ep.Age = age
}

func (ep *Employee) SetEmployeeName(name string) {
	ep.Name = name
}
