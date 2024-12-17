// Package utils provides utility functions for the catalog service
package utils

import (
	"github.com/JorgeO3/flowcast/internal/catalogv2/domain/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func genID() string {
	return primitive.NewObjectID().Hex()
}

// GenerateIDs generates the IDs for the act and its assets
func GenerateIDs(act *entity.Act) {
	if !act.IsEmpty() && act.ID == "" {
		act.ID = genID()
	}

	for i := range act.Albums {
		if !act.Albums[i].IsEmpty() && act.Albums[i].ID == "" {
			act.Albums[i].ID = genID()
		}

		for j := range act.Albums[i].Songs {
			if !act.Albums[i].Songs[j].IsEmpty() && act.Albums[i].Songs[j].ID == "" {
				act.Albums[i].Songs[j].ID = genID()
			}
		}
	}
}

// GenerateIDsFromActs generates the IDs for the acts and their assets
func GenerateIDsFromActs(acts []entity.Act) {
	for i := range acts {
		GenerateIDs(&acts[i])
	}
}
