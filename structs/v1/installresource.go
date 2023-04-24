package v1

import "go.mongodb.org/mongo-driver/bson/primitive"

type InstallResource struct {
	// meta specs
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ObjID       string             `bson:"objid,omitempty" json:"objid"`
	ServiceName string             `bson:"servicename,omitempty" json:"servicename"`
	Environment string             `bson:"environment,omitempty" json:"environment"`
	Delete      bool               `bson:"delete,omitempty" json:"delete"`
}
