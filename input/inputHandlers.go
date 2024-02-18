package input

import (
	"challenge2016/dto" // Importing DTO package for data transfer objects
	"fmt"
	"log"
	"strings"

	"github.com/manifoldco/promptui"
)

// The `PromptMenu` function in Go displays a menu for selecting different choices and returns the
// selected option.
func PromptMenu() string {
	prompt := promptui.Select{
		Label: "Select one of the below choices",
		Items: []string{
			"Create a new distributor",
			"Create a sub distributor",
			"Check permission for a distributor",
			"View Distributors information",
			"Exit",
		},
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	return result
}

// The function `PromptDistributorData` in Go prompts the user to enter distributor data including
// name, regions to include and exclude, and optionally the parent distributor name.
func PromptDistributorData(subDistributor bool) dto.Distributor {
	var distributor dto.Distributor

	promptName := promptui.Prompt{
		Label:       "Enter distributor name:",
		HideEntered: true,
	}
	name, _ := promptName.Run()
	distributor.Name = name
	fmt.Println(promptName.Label, name)

	promptInclude := promptui.Prompt{
		Label:       "Enter the regions you want to include for this distributor (comma separated): ",
		HideEntered: true,
	}
	includeInput, _ := promptInclude.Run()
	distributor.Include = strings.Split(includeInput, ",")
	fmt.Println(promptInclude.Label, includeInput)

	promptExclude := promptui.Prompt{
		Label:       "Enter the regions you want to exclude for this distributor (comma separated): ",
		HideEntered: true,
	}
	excludeInput, _ := promptExclude.Run()
	distributor.Exclude = strings.Split(excludeInput, ",")
	fmt.Println(promptExclude.Label, excludeInput)

	if subDistributor {
		promptParent := promptui.Prompt{
			Label:       "Enter the name of the parent distributor: ",
			HideEntered: true,
		}
		parent, _ := promptParent.Run()
		distributor.Parent = parent
		fmt.Println(promptParent.Label, parent)
	}

	return distributor
}

// The `PromptCheckPermissionData` function in Go prompts the user to enter a distributor name and
// regions for permission checking.
func PromptCheckPermissionData() dto.CheckPermissionData {
	var data dto.CheckPermissionData

	promptName := promptui.Prompt{
		Label:       "Enter distributor name that needs to be checked:",
		HideEntered: true,
	}
	data.DistributorName, _ = promptName.Run()
	fmt.Println(promptName.Label, data.DistributorName)

	promptRegions := promptui.Prompt{
		Label:       "Enter distributor name that needs to be checked (comma separated):",
		HideEntered: true,
	}
	regionsInput, _ := promptRegions.Run()
	data.Regions = strings.Split(regionsInput, ",")
	fmt.Println(promptRegions.Label, data.Regions)
	return data
}
