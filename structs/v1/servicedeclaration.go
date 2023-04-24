package v1

import "go.mongodb.org/mongo-driver/bson/primitive"

type ServiceDeclaration struct {
	// meta specs
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	ApiVersion  string             `bson:"apiversion" json:"apiversion"`

	// deployment specs
	Maintainer         string   `bson:"maintainer" json:"maintainer"`
	Environment        string   `bson:"environment" json:"environment"`
	ImageName          string   `bson:"imagename" json:"imagename"`
	ImageTag           string   `bson:"imagetag" json:"imagetag"`
	Replicas           int      `bson:"replicas" json:"replicas"`
	Subdomains         []string `bson:"subdomains" json:"subdomains"`
	Subpath            string   `bson:"subpath" json:"subpath"`
	AppSettingsContent string   `bson:"appsettingscontent" json:"appsettingscontent"`
	Envs               []struct {
		Name  string `bson:"name" json:"name"`
		Value string `bson:"value" json:"value"`
	} `bson:"envs" json:"envs"`
}
