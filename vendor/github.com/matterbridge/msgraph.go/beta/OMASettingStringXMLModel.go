// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// OMASettingStringXML undocumented
type OMASettingStringXML struct {
	// OMASetting is the base model of OMASettingStringXML
	OMASetting
	// FileName File name associated with the Value property (*.xml).
	FileName *string `json:"fileName,omitempty"`
	// Value Value. (UTF8 encoded byte array)
	Value *Binary `json:"value,omitempty"`
}