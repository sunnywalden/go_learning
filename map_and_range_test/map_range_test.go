package map_and_range_test

import "testing"

func TestMapRange(t *testing.T) {
	m1 := map [string]string{"henry": "Henry Zhang", "walden": "SunnyWalden"}
	for k, v := range m1 {
		t.Log(k, v)
		if v,ok := m1["Walden"]; ok {
			t.Logf("%s in map", v)
		} else {
			t.Logf("%s not exist in map", "Walden")
		}
	}

}
