package permissions

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPermissions(t *testing.T) {
	includes := []string{
		"ITEM0-ONE",
		"ITEN2-TWO",
		"ITEN3-FREE",
	}
	excludes := []string{
		"EX-ITEM0-ONE",
	}

	perms, err := NewPermissions(nil, includes, excludes)
	require.NoError(t, err, err)
	require.NotNil(t, perms)

	t.Run("IsAllowed", func(t *testing.T) {
		for _, include := range includes {
			v := perms.IsAllowed(include)
			assert.True(t, v)
		}

		for _, exclude := range excludes {
			assert.False(t, perms.IsAllowed(exclude))
		}
	})

	t.Run("Include", func(t *testing.T) {
		newIncl := "TEST-INCLUDE"
		// create include
		err := perms.Include(newIncl)
		assert.NoError(t, err, err)
		assert.True(t, perms.IsAllowed(newIncl))

		t.Run("ExcludeIncluded", func(t *testing.T) {
			err := perms.Exclude(newIncl)
			assert.Error(t, err, ErrCannotInclude)
		})

		//delete include
		err = perms.Include(newIncl)
		assert.NoError(t, err, err)
		assert.False(t, perms.IsAllowed(newIncl))
	})

	t.Run("Exclude", func(t *testing.T) {
		newIncl := "TEST-EXCLUDE"
		// create exclude
		err := perms.Include(newIncl)
		assert.NoError(t, err, err)
		assert.True(t, perms.IsAllowed(newIncl))

		newExcl := "FAKE-" + newIncl
		assert.True(t, perms.IsAllowed(newExcl))

		err = perms.Exclude(newExcl)
		assert.NoError(t, err, err)
		assert.False(t, perms.IsAllowed(newExcl))

		t.Run("IncludeExcluded", func(t *testing.T) {
			err := perms.Include(newExcl)
			assert.Error(t, err, ErrCannotExclude)
		})

		//delete exclude
		err = perms.Exclude(newExcl)
		assert.NoError(t, err, err)
		assert.True(t, perms.IsAllowed(newExcl))
	})
}
