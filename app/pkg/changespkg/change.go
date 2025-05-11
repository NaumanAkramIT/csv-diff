package changespkg

type Change struct {
	ChangeType string   `json:"type"`
	Key        string   `json:"key"`
	OldValue   []string `json:"old,omitempty"`
	NewValue   []string `json:"new,omitempty"`
	ChangeNote string   `json:"note,omitempty"`
}

func New(changeType string, key string, oldValue []string, newValue []string, changeNote string) *Change {
	return &Change{
		ChangeType: changeType,
		Key:        key,
		OldValue:   oldValue,
		NewValue:   newValue,
		ChangeNote: changeNote,
	}
}
