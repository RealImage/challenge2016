package permissions

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestService_DistributorFlow(t *testing.T) {
	includes := []string{
		"ITEM0-ONE",
		"ITEN2-TWO",
		"ITEN3-FREE",
	}
	excludes := []string{
		"EX-ITEM0-ONE",
	}
	dist := "one"

	s := NewService()

	// create permission
	parentPerms, err := s.CreatePermissions("", dist, includes, excludes)
	assert.NoError(t, err, err)
	assert.NotNil(t, parentPerms)

	t.Run("Distributors", func(t *testing.T) {
		contains := false
		for _, d := range s.Distributors() {
			if d.Name == dist {
				contains = true
				break
			}
		}
		assert.True(t, contains)
	})

	t.Run("GetPermissions", func(t *testing.T) {
		t.Run("BadDistributor", func(t *testing.T) {
			perms, err := s.GetPermissions("fake")
			assert.Error(t, err, ErrDistributorNotFound)
			assert.Nil(t, perms)
		})

		t.Run("ExistingDistributor", func(t *testing.T) {
			perms, err := s.GetPermissions(dist)
			assert.NoError(t, err, err)
			assert.EqualValues(t, parentPerms, perms)
		})
	})

	t.Run("HasPermissions", func(t *testing.T) {
		t.Run("Has", func(t *testing.T) {
			for _, include := range includes {
				isAllowed, err := s.HasPermissions(dist, include)
				assert.NoError(t, err, err)
				assert.True(t, isAllowed)
			}
		})

		t.Run("DoesntHave", func(t *testing.T) {
			for _, exclude := range excludes {
				isAllowed, err := s.HasPermissions(dist, exclude)
				assert.NoError(t, err, err)
				assert.False(t, isAllowed)
			}
		})

		t.Run("BadDistributor", func(t *testing.T) {
			isAllowed, err := s.HasPermissions("fake", includes[0])
			assert.Error(t, err, ErrDistributorNotFound)
			assert.False(t, isAllowed)
		})
	})

	t.Run("UpdatePermissions", func(t *testing.T) {
		updatePerms, err := NewPermissions(nil, []string{"NEW-ONE"}, []string{})
		require.NoError(t, err, err)
		require.NotNil(t, updatePerms)

		err = s.UpdatePermissions(dist, updatePerms)
		assert.NoError(t, err, err)

		perms, err := s.GetPermissions(dist)
		assert.NoError(t, err, err)
		assert.EqualValues(t, updatePerms, perms)

		// set back old permissions for other tests
		err = s.UpdatePermissions(dist, parentPerms)
		assert.NoError(t, err, err)
	})

	t.Run("SubDistributor", func(t *testing.T) {
		// allows only the last region
		childIncludes := []string{
			includes[len(includes)-1],
		}
		childExcludes := []string{
			"EXCLUDE-" + includes[len(includes)-1],
		}
		childDist := "child"

		// create sub permission
		childPerms, err := s.CreatePermissions(dist, childDist, childIncludes, childExcludes)
		require.NoError(t, err, err)
		require.NotNil(t, childPerms)

		t.Run("HasPermissions", func(t *testing.T) {
			t.Run("DoesntHave", func(t *testing.T) {
				// doesn't have access except the last region
				for i := 0; i < len(includes)-1; i++ {
					isAllowed, err := s.HasPermissions(childDist, includes[i])
					assert.NoError(t, err, err)
					assert.False(t, isAllowed)
				}

				for _, exclude := range excludes {
					isAllowed, err := s.HasPermissions(childDist, exclude)
					assert.NoError(t, err, err)
					assert.False(t, isAllowed)
				}

				for _, exclude := range childExcludes {
					isAllowed, err := s.HasPermissions(childDist, exclude)
					assert.NoError(t, err, err)
					assert.False(t, isAllowed)
				}
			})

			t.Run("Has", func(t *testing.T) {
				isAllowed, err := s.HasPermissions(childDist, includes[len(includes)-1])
				assert.NoError(t, err, err)
				assert.True(t, isAllowed)

				isAllowed, err = s.HasPermissions(childDist, "TEST-"+includes[len(includes)-1])
				assert.NoError(t, err, err)
				assert.True(t, isAllowed)
			})
		})

		t.Run("UpdateParentPermission", func(t *testing.T) {
			updatePerms, err := NewPermissions(nil, []string{"NEW-ONE"}, []string{})
			require.NoError(t, err, err)
			require.NotNil(t, updatePerms)

			err = s.UpdatePermissions(dist, updatePerms)
			assert.NoError(t, err, err)

			perms, err := s.GetPermissions(childDist)
			assert.NoError(t, err, err)
			assert.EqualValues(t, updatePerms, perms.parentPerms)
			assert.EqualValues(t, newRegions(), perms.includes)
			assert.EqualValues(t, newRegions(), perms.excludes)

			isAllowed, err := s.HasPermissions(childDist, includes[len(includes)-1])
			assert.NoError(t, err, err)
			assert.False(t, isAllowed)

			isAllowed, err = s.HasPermissions(childDist, "TEST"+includes[len(includes)-1])
			assert.NoError(t, err, err)
			assert.False(t, isAllowed)
		})
	})

	t.Run("RemovePermissions", func(t *testing.T) {
		t.Run("BadDistributor", func(t *testing.T) {
			err := s.RemovePermissions("fake")
			assert.Error(t, err, ErrDistributorNotFound)
		})

		t.Run("ExistingDistributor", func(t *testing.T) {
			err := s.RemovePermissions(dist)
			assert.NoError(t, err, err)

			contains := false
			for _, distributor := range s.Distributors() {
				if distributor.Name == dist {
					contains = true
					break
				}
			}
			assert.False(t, contains)
		})
	})
}
