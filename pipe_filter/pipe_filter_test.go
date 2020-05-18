package pipe_filter

import (
	"errors"
	"fmt"
	"github.com/sunnywalden/go_learning/unit"
	"strings"
	"testing"
	"time"
)

type GetEmployeeNames struct {
	*unit.Employee
}

func (ge *GetEmployeeNames) Process() ([]string, error) {
	epName := ge.Name
	nameList := strings.Split(epName, string(' '))
	if len(nameList) >= 1 {
		fmt.Printf("Names  %s.\n", nameList)
		return nameList,nil
	}  else {
		fmt.Printf("Name is null\n")
	}
	return []string{""},errors.New("name is null\n")
}


type GetEmployeeCapNames struct {
	*unit.Employee
	*GetEmployeeNames
}

func (ge *GetEmployeeCapNames) Process() ([]string, error) {
	epNames,err := ge.GetEmployeeNames.Process()
	var namesCap []string
	if err == nil {
		fmt.Printf("Names  %s.\n", epNames)
		for _,name := range epNames {
			namesCap = append(namesCap, unit.Capitalize(name))
		}
		return namesCap, nil
	}
	return []string{""},errors.New("name is null\n")
}

type GetEmployeeBirth struct {
	*unit.Employee
}

func (gp *GetEmployeeBirth) Process() (int, error) {
	epYears := gp.Age
	currentYear := time.Now().Year()
	birthYear := currentYear - epYears
	if birthYear != 0 {
		fmt.Printf("Names  %d.\n", birthYear)
		return birthYear,nil
	}  else {
		fmt.Printf("birthYear is null\n")
	}
	return 0,errors.New("birthYear is null\n")
}


type MyFilter interface {
	Process() ([]string, error)
}

type GetEmployeeName struct {
	*unit.Employee
	Filters []func()([]string, error)
}

func (gen *GetEmployeeName) Process() ([]string, error) {
	var names []string
	for _,filter := range gen.Filters {
		res, err := filter()

		if err == nil {
			names = append(names, res...)
			fmt.Print(res)
		} else {
			fmt.Printf("%s\n", err)
			break
		}
	}
	return names, nil
}


func TestPipeFilter(t *testing.T) {
	employeeDemo := new(unit.Employee)
	employeeDemo.SetEmployeeName("sunny walden")
	employeeDemo.SetEmployeeAge(30)

	en := &GetEmployeeNames{employeeDemo}
	eCn := GetEmployeeCapNames{employeeDemo, en}
	Filters := []func()([]string, error){en.Process,eCn.Process}
	getEmName := &GetEmployeeName{employeeDemo, Filters}

	names, err := getEmName.Process()
	if err == nil {
		t.Logf("Cap names: %s\n",names)
	} else {
		t.Error(err)
	}

}



