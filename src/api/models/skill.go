package models

type Skills struct {
	ID        string `json:"_id" bson:"_id,omitempty"`
	IsActive  bool   `json:"is_active" bson:"is_active,omitempty"`
	SkillName string `json:"skill_name" bson:"skill_name,omitempty"`
	SkillType string `json:"skill_type" bson:"skill_type,omitempty"`
	PicPath   string `json:"picPath" bson:"picPath,omitempty"`
}
