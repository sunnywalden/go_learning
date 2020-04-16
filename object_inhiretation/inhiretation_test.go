package object_inhiretation

import (
	"fmt"
	"testing"
)

type Traffic interface {
	Run()
}

type Car struct {
	Traffic
	Brand string
	Source string
}

func (c *Car) Run() {
	fmt.Print("Dou Dou!\n")
}

func (c *Car) IsElectricCar() bool {
	if c.Source == "electricity" {
		return true
	} else {
		return false
	}
}

func (c *Car) InitBrand(s string) {
	c.Brand = s
}

func (c *Car) InitSource() {
	c.Source = "gas"
}

func (c *Car) CarBrand() string {
	return c.Brand
}


type NewPowerCar struct {
	Car
	Capacity int
}

func (n *NewPowerCar) Run() {
	fmt.Print("Di Di!\n")
}

func (n *NewPowerCar) InitSource(num int) {
	n.Source = "electricity"
	n.Capacity = num
}

type Horse struct {
	Traffic
	Animal
	Name string
	Source string
}

func (h *Horse) Run() {
	fmt.Printf("Run %s!\n", h.Name)
}

func (h *Horse) TypePrint() {
	fmt.Print("Horse\n")
}

func (h *Horse) Bugger() string {
	h.TypePrint()
	return "pu pu!\n"
}

func TestExtension(t *testing.T) {
	MyHorse := new(Horse)
	MyHorse.Name = "Hans"
	MyHorse.Run()
	t.Logf("Your horse %s comes to you!", MyHorse.Name)

	NormalCar := new(Car)
	NormalCar.InitSource()
	NormalCar.InitBrand("Chevrolet")
	NormalCar.Run()
	t.Log("Your ", NormalCar.Brand, " using ", NormalCar.Source, " as power. It's electricity car or not:", NormalCar.IsElectricCar(), ".")

	NewCar := new(NewPowerCar)
	NewCar.InitSource(5000)
	NewCar.InitBrand("Tesla")
	NewCar.Run()
	t.Log("Your ", NewCar.Brand, " using ", NewCar.Capacity, " kwh of ", NewCar.Source, " as power. It's electricity car or not:", NewCar.IsElectricCar(), ".")
}

type Animal struct {
	//Traffic
	Name string
}

func (a *Animal) TypePrint() {
	fmt.Print("Animal\n")
}

func (a *Animal) Bugger() string {
	a.TypePrint()
	return "...\n"
}

type Dog struct {
	Animal
}

func (d *Dog) TypePrint() {
	fmt.Print("Dog\n")
}

func (d *Dog) Bugger() string {
	d.TypePrint()
	return "Wang!Wang!\n"
}

func TestDog(t *testing.T) {
	a := new(Animal)
	a.Name = "Piggey"
	t.Log("Animal name is",a.Name)
	res1 := a.Bugger()
	t.Log(res1)

	h := new(Horse)
	h.Name = "Hans"
	t.Log("Horse name is", h.Name)
	res2 := h.Bugger()
	t.Log(res2)

	d := new(Dog)
	d.Name = "Snoppy"
	t.Log("Dog name is",d.Name)
	res := d.Bugger()
	t.Log(res)


}
