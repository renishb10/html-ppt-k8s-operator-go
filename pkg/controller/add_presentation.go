package controller

import (
	"github.com/renishb10/html-ppt-k8s-operator-go/pkg/controller/presentation"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, presentation.Add)
}
