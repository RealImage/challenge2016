package cli

import (
	"errors"
	"flag"
	"fmt"
	"net/rpc/jsonrpc"
	"strings"
)

const (
	createCmd = "create"
	getCmd    = "get"
	hasCmd    = "has"
	updateCmd = "update"
	removeCmd = "remove"

	commandFlag = "command"
	regionFlag  = "region"
	includeFlag = "include"
	excludeFlag = "exclude"
	newDistFlag = "newDist"
	distFlag    = "dist"
	helpFlag    = "help"
)

var (
	command = flag.String(commandFlag, "", fmt.Sprintf(
		"Name of a command that will be called: \n"+
			fmt.Sprintf("\t%s - Creates new permissions. Requires: %s, %s. Optional: %s, %s \n", createCmd, newDistFlag, includeFlag, distFlag, excludeFlag)+
			fmt.Sprintf("\t%s - Gets permissions by distributor. Requires: %s \n", getCmd, distFlag)+
			fmt.Sprintf("\t%s - Checks if a distributor has permissions to a region. Requires: %s, %s\n", hasCmd, distFlag, regionFlag)+
			fmt.Sprintf("\t%s - Updates permissions by distributor. Requires: %s, %s or/both %s \n", updateCmd, distFlag, includeFlag, excludeFlag)+
			fmt.Sprintf("\t%s - Removes permissions by distributor. Requires: %s", removeCmd, distFlag)),
	)
	region       = flag.String(regionFlag, "", "Name of a region")
	include      = flag.String(includeFlag, "", "List of regions` deltas separated by ','. If the region exists, it will be deleted, if it does not exist, it will be created. Example: -include=CHENNAI-TAMILNADU-INDIA,UNITEDSTATES")
	exclude      = flag.String(excludeFlag, "", "List of regions` deltas separated by ','. If the region exists, it will be deleted, if it does not exist, it will be created. Example: -exclude=CHENNAI-TAMILNADU-INDIA,UNITEDSTATES")
	newDist      = flag.String(newDistFlag, "", "Name for new child distributor")
	existingDist = flag.String(distFlag, "", "Name of existing distributor")
	help         = flag.Bool(helpFlag, false, "Show info about commands")

	ErrCommandDoesntExist = errors.New("command does not exist")
)

type PermissionClient struct{}

func NewPermissionClient() *PermissionClient {
	pc := &PermissionClient{}

	return pc
}

func (pc *PermissionClient) Handle() error {
	client, err := jsonrpc.Dial("tcp", defaultAddress)
	if err != nil {
		return err
	}

	switch *command {
	case createCmd:
		err := checkArguments(newDistFlag, includeFlag)
		if err != nil {
			return err
		}

		args := &CreatePermsArgs{
			*existingDist,
			*newDist,
			strings.Split(*include, ","),
			strings.Split(*exclude, ","),
		}
		resp := ""
		err = client.Call(createPermissionsFunc, args, &resp)
		if err != nil {
			return err
		}

		fmt.Printf("Permissions for '%s' created:\n%s", *newDist, resp)
	case getCmd:
		err := checkArguments(distFlag)
		if err != nil {
			return err
		}

		resp := ""
		err = client.Call(getPermissionsFunc, *existingDist, &resp)
		if err != nil {
			return err
		}

		fmt.Printf("Permissions of '%s':\n%s", *existingDist, resp)
	case hasCmd:
		err := checkArguments(distFlag, regionFlag)
		if err != nil {
			return err
		}

		req := &HasPermsArgs{
			Distributor: *existingDist,
			Region:      *region,
		}
		resp := false
		err = client.Call(hasPermissionsFunc, req, &resp)
		if err != nil {
			return err
		}

		result := "DOES NOT HAVE"
		if resp {
			result = "HAS"
		}

		fmt.Printf("Distributor '%s' %s permissions to '%s' region\n", *existingDist, result, *region)
	case updateCmd:
		err := checkArguments(distFlag)
		if err != nil {
			return err
		}

		if *include == "" && *exclude == "" {
			return errors.New("at least 'excludes' or 'includes' should be provided")
		}

		args := &UpdatePermsArgs{
			*existingDist,
			strings.Split(*include, ","),
			strings.Split(*exclude, ","),
		}
		resp := ""
		err = client.Call(updatePermissionsFunc, args, &resp)
		if err != nil {
			return err
		}

		fmt.Printf("Permissions for '%s' updated:\n%s", *existingDist, resp)
	case removeCmd:
		err := checkArguments(distFlag)
		if err != nil {
			return err
		}

		err = client.Call(removePermissionsFunc, *existingDist, nil)
		if err != nil {
			return err
		}

		fmt.Printf("Successfully removed permissions for '%s'\n", *existingDist)
	default:
		return errors.New(ErrCommandDoesntExist.Error() + " " + *command)
	}

	return nil
}

func checkArguments(argsNames ...string) error {
	for _, arg := range argsNames {
		f := flag.Lookup(arg)
		if f.Value.String() == "" {
			return errors.New("flag " + arg + " is not set")
		}
	}

	return nil
}
