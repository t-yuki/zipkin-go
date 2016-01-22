package models

import "github.com/go-swagger/go-swagger/strfmt"

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*DependencyLink DependencyLink

swagger:model DependencyLink
*/
type DependencyLink struct {

	/* CallCount call count
	 */
	CallCount *int64 `json:"callCount,omitempty"`

	/* Child child
	 */
	Child *string `json:"child,omitempty"`

	/* Parent parent
	 */
	Parent *string `json:"parent,omitempty"`
}

// Validate validates this dependency link
func (m *DependencyLink) Validate(formats strfmt.Registry) error {
	return nil
}
