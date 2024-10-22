package utils

import (
	"github.com/JorgeO3/flowcast/internal/catalog/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func genID() string {
	return primitive.NewObjectID().Hex()
}

// GenerateIDs generates the IDs for the act and its assets
func GenerateIDs(act *entity.Act) {
	if !IsStructEmpty(act) && act.ID == "" {
		act.ID = genID()
	}

	for i := range act.Albums {
		if !IsStructEmpty(act.Albums[i]) && act.Albums[i].ID == "" {
			act.Albums[i].ID = genID()
		}

		for j := range act.Albums[i].Songs {
			if !IsStructEmpty(act.Albums[i].Songs[j]) && act.Albums[i].Songs[j].ID == "" {
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

// IsActEmpty checks if the act is empty
func IsActEmpty(act entity.Act) bool {
	return act.ID == "" && act.
}