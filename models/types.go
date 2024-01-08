package models

import "encoding/json"

type CreateRequest struct {
	OriginalURL   string         `json:"original_url"`
	URLParams     []URLParam     `json:"url_params"`
	HeaderParams  []HeaderParam  `json:"header_params"`
	AutoGenParams []AutoGenParam `json:"auto_gen_params"`
}

type URLParam struct {
	SourceParam string `json:"source_param"`
	TargetParam string `json:"target_param"`
	IsMandatory bool   `json:"is_mandatory"`
}

type HeaderParam struct {
	SourceParam string `json:"source_param"`
	TargetParam string `json:"target_param"`
	IsMandatory bool   `json:"is_mandatory"`
}

type AutoGenParam struct {
	Type      string `json:"type"`
	TargetKey string `json:"target_key"`
}

// required by redis for marshalling the structs
func (i URLParam) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}

// required by redis for marshalling the structs
func (i HeaderParam) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}

// required by redis for marshalling the structs
func (i AutoGenParam) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}
