package controller

import (
	"github.com/dos1701/nsprovisioner-operator/pkg/controller/namespaceprovisioner"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, namespaceprovisioner.Add)
}
