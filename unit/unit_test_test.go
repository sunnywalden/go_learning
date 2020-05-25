package unit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnitTest(t *testing.T) {
	employees := []string{"sunnywalden", "henry", "robin", "VC"}
	employeesCap := []string{"Sunnywalden", "Henry", "robin", "VC"}
	for i,name := range employees {
		assert.Equal(t, Capitalize(name), employeesCap[i], "%s's capitalise should be %s", employeesCap[i])
	}
}
