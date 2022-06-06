package permissions

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegionFlow(t *testing.T) {
	r := newRegions()
	regions := []string{
		"Bolondron-Matanzas-Cuba",
		"Batabano-Mayabeque-Cuba",
		"Alacranes-Matanzas-Cuba",
		"Batabano-Mayabeque-Cuba",
	}

	for _, region := range regions {
		r.Add(region)
	}

	output := strings.Split(r.String(), "\n")
	for _, region := range regions {
		assert.Contains(t, output, region)
	}

	assert.True(t, r.Contains(regions[0], true))
	assert.True(t, r.Contains(regions[0], false))

	// check that strict and non-strict CUBA match
	assert.True(t, r.Contains("Cuba", false))
	assert.False(t, r.Contains("Cuba", true))

	assert.True(t, r.Contains("Mayabeque-Cuba", false))
	assert.False(t, r.Contains("Mayabeque-Cuba", true))

	// check fake element
	assert.False(t, r.Contains("FAKE", false))
	assert.False(t, r.Contains("FAKE", true))

	r.Remove(regions[1])
	assert.False(t, r.Contains(regions[1], true))

	r.Remove("Cuba")
	assert.True(t, len(r.table) == 0)
}
