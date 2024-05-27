package resource

import "github.com/jdbann/thestacks/util/keys"

type KeyBindings struct {
	PanCamera   keys.Binding
	ResetCamera keys.Binding
}

func DefaultKeyBindings() *KeyBindings {
	return &KeyBindings{
		PanCamera:   keys.NewBinding(keys.WithKey(keys.KeySpace)),
		ResetCamera: keys.NewBinding(keys.WithKey(keys.KeyR)),
	}
}
