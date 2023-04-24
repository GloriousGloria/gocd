package v1

import "go.mongodb.org/mongo-driver/bson/primitive"

type InstallResource struct {
	// meta specs
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ObjID       string             `bson:"objid" json:"objid"`
	ServiceName string             `bson:"servicename" json:"servicename"`
	Environment string             `bson:"environment" json:"environment"`
	Delete      bool               `bson:"delete" json:"delete"`
}
