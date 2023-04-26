package v1

import "go.mongodb.org/mongo-driver/bson/primitive"

type InstallResource struct {
	// meta specs
	ID                    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ObjID                 string             `bson:"objid,omitempty" json:"objid"`
	ServiceName           string             `bson:"servicename,omitempty" json:"servicename"`
	Environment           string             `bson:"environment,omitempty" json:"environment"`
	Delete                bool               `bson:"delete,omitempty" json:"delete"`
	AdditionalInformation struct {
		RelatedObjects []RelatedObject `bson:"relatedobjects,omitempty" json:"relatedobjects"`
		GitPointers    []GitPointer    `bson:"gitpointers,omitempty" json:"gitpointers"`
	} `bson:"additionalinfo,omitempty" json:"additionalinfo"`
}

type RelatedObject struct {
	ObjectType string `bson:"objecttype" json:"objecttype"`
	ApiVersion string `bson:"apiversion" json:"apiversion"`
	Namespace  string `bson:"namespace" json:"namespace"`
	Name       string `bson:"name" json:"name"`
}

type GitPointer struct {
	GitlabAPIV4 string `bson:"gitlabapiv4" json:"gitlabapiv4"`
	ProjectID   int    `bson:"projectid" json:"projectid"`
	AccessToken string `bson:"accesstoken" json:"accesstoken"`
	Author      string `bson:"author" json:"author"`
	Mail        string `bson:"mail" json:"mail"`
	Branch      string `bson:"branch" json:"branch"`
	File        string `bson:"file" json:"file"`
}
