package v1

import "go.mongodb.org/mongo-driver/bson/primitive"

type ServiceDeclaration struct {
	// meta specs
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name,omitempty" json:"name"`
	Description string             `bson:"description,omitempty" json:"description"`
	ApiVersion  string             `bson:"apiversion,omitempty" json:"apiversion"`

	// deployment specs
	Maintainer         string   `bson:"maintainer,omitempty" json:"maintainer"`
	Environment        string   `bson:"environment,omitempty" json:"environment"`
	ImageName          string   `bson:"imagename,omitempty" json:"imagename"`
	ImageTag           string   `bson:"imagetag,omitempty" json:"imagetag"`
	Replicas           int      `bson:"replicas,omitempty" json:"replicas"`
	Subdomains         []string `bson:"subdomains,omitempty" json:"subdomains"`
	Subpath            string   `bson:"subpath,omitempty" json:"subpath"`
	AppSettingsContent string   `bson:"appsettingscontent,omitempty" json:"appsettingscontent"`
	Envs               []struct {
		Name  string `bson:"name" json:"name"`
		Value string `bson:"value" json:"value"`
	} `bson:"envs,omitempty" json:"envs"`
}
