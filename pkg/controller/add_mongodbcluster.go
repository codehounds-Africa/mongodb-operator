package controller

import (
	"github.com/codehounds/mongodb-operator/pkg/controller/mongodbcluster"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, mongodbcluster.Add)
}
