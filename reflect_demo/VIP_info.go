package reflect_demo

import "time"
import "go_learning/unit_test"

type VIP struct {
	Id string
	Name         string
	Birth          time.Time
}

func (vip *VIP) GetVIPName() string {
	return unit.Capitalize(vip.Name)
}

func (vip *VIP) GetVIPBirth() time.Time {
	return vip.Birth
}

func (vip *VIP) SetVIPBirth(birth time.Time) {
	vip.Birth = birth
}

func (vip *VIP) SetVIPName(name string) {
	vip.Name = name
}
