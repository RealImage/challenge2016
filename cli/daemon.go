package cli

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"challenge2016/permissions"
)

const (
	defaultAddress = "0.0.0.0:4444"

	createPermissionsFunc = "Listener.CreatePermissions"
	getPermissionsFunc    = "Listener.GetPermissions"
	hasPermissionsFunc    = "Listener.HasPermissions"
	updatePermissionsFunc = "Listener.UpdatePermissions"
	removePermissionsFunc = "Listener.RemovePermissions"
)

type (
	CreatePermsArgs struct {
		ParentDistributor, NewDistributor string
		Includes, Excludes                []string
	}

	HasPermsArgs struct {
		Distributor string
		Region      string
	}

	UpdatePermsArgs struct {
		Distributor        string
		Includes, Excludes []string
	}

	Listener struct {
		perms *permissions.Service
	}

	PermissionsDaemon struct {
		Listener *Listener
	}
)

func NewPermissionsDaemon() *PermissionsDaemon {
	return &PermissionsDaemon{
		Listener: &Listener{
			permissions.NewService(),
		},
	}
}

func (pd *PermissionsDaemon) Listen() error {
	addy, err := net.ResolveTCPAddr("tcp", defaultAddress)
	if err != nil {
		log.Fatal(err)
	}

	inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
		log.Fatal(err)
	}

	err = rpc.Register(pd.Listener)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := inbound.Accept()
		if err != nil {
			continue
		}
		jsonrpc.ServeConn(conn)
	}
}

func (l *Listener) CreatePermissions(req CreatePermsArgs, permsResp *string) error {
	perms, err := l.perms.CreatePermissions(req.ParentDistributor, req.NewDistributor, req.Includes, req.Excludes)
	if err != nil {
		return err
	}

	*permsResp = perms.String()
	return nil
}

func (l *Listener) GetPermissions(distributor string, permsResp *string) error {
	perms, err := l.perms.GetPermissions(distributor)
	if err != nil {
		return err
	}

	*permsResp = perms.String()
	return nil
}

func (l *Listener) HasPermissions(hasPerms *HasPermsArgs, hasResp *bool) error {
	has, err := l.perms.HasPermissions(hasPerms.Distributor, hasPerms.Region)
	if err != nil {
		return err
	}

	*hasResp = has
	return nil
}

func (l *Listener) UpdatePermissions(updatePerms *UpdatePermsArgs, updateResp *string) error {
	perms, err := l.perms.GetPermissions(updatePerms.Distributor)
	if err != nil {
		return err
	}

	for _, in := range updatePerms.Includes {
		err = perms.Include(in)
		if err != nil {
			return err
		}
	}

	for _, ex := range updatePerms.Excludes {
		err = perms.Exclude(ex)
		if err != nil {
			return err
		}
	}

	err = l.perms.UpdatePermissions(updatePerms.Distributor, perms)
	if err != nil {
		return err
	}

	*updateResp = perms.String()
	return nil
}

func (l *Listener) RemovePermissions(distributor string, _ *string) error {
	return l.perms.RemovePermissions(distributor)
}
