package testing

import (
	"github.com/RealImage/challenge2016/service"
	"testing"
)

func TestCheckPermissionsScenario1(t *testing.T) {
	//Example permissions for DISTRIBUTOR1
	permission := service.Permissions{
		Include: []string{"INDIA", "UNITEDSTATES"},
		Exclude: []string{"KARNATAKA-INDIA", "CHENNAI-TAMILNADU-INDIA"},
	}

	permission1 := service.Permissions{
		Include: []string{"GUJARAT-INDIA"},
	}
	permission2 := service.Permissions{
		Include: []string{"CHENNAI-TAMILNADU-INDIA"},
	}
	permission3 := service.Permissions{
		Include: []string{"BANGALORE-KARNATAKA-INDIA"},
	}

	//The first parameter is a subset check and second is superset
	check := service.CheckPermissions(permission1, permission)
	if check != true {
		t.Errorf("got %v; want %v", check, true)
	}
	check = service.CheckPermissions(permission2, permission)
	if check != false {
		t.Errorf("got %v; want %v", check, false)
	}
	check = service.CheckPermissions(permission3, permission)
	if check != false {
		t.Errorf("got %v; want %v", check, false)
	}
}

func TestCheckPermissionsScenario2(t *testing.T) {
	//Example permissions for DISTRIBUTOR2
	permission := service.Permissions{
		Include: []string{"INDIA"},
		Exclude: []string{"GUJARAT-INDIA", "KARNATAKA-INDIA", "CHENNAI-TAMILNADU-INDIA"},
	}

	permission1 := service.Permissions{
		Include: []string{"CHINA"},
	}
	permission2 := service.Permissions{
		Include: []string{"INDIA"},
	}
	permission3 := service.Permissions{
		Include: []string{"TAMILNADU-INDIA"},
	}

	//The first parameter is a subset check and second is superset
	check := service.CheckPermissions(permission1, permission)
	if check != false {
		t.Errorf("got %v; want %v", check, false)
	}
	check = service.CheckPermissions(permission2, permission)
	if check != true {
		t.Errorf("got %v; want %v", check, true)
	}
	check = service.CheckPermissions(permission3, permission)
	if check != true {
		t.Errorf("got %v; want %v", check, true)
	}
}

func TestCheckPermissionsScenario3(t *testing.T) {
	//Example permissions for DISTRIBUTOR1
	permission := service.Permissions{
		Include: []string{"HUBLI-KARNATAKA-INDIA"},
	}
	permission1 := service.Permissions{
		Include: []string{"KARNATAKA-INDIA"},
	}
	permission2 := service.Permissions{
		Include: []string{"HUBLI-KARNATAKA-INDIA"},
	}

	//The first parameter is a subset check and second is superset
	check := service.CheckPermissions(permission1, permission)
	if check != false {
		t.Errorf("got %v; want %v", check, false)
	}
	check = service.CheckPermissions(permission2, permission)
	if check != true {
		t.Errorf("got %v; want %v", check, true)
	}
}
