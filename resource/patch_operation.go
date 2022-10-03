package resource

import (
	"encoding/json"
	"fmt"
)

type PatchOperationValue json.RawMessage

func (v *PatchOperationValue) GetValue() interface{} {
	var dst interface{}
	if err := json.Unmarshal(*v, &dst); err != nil {
		return nil
	}
	return dst
}

func (v *PatchOperationValue) AcceptValue(in interface{}) error {
	serialized, err := json.Marshal(in)
	if err != nil {
		return fmt.Errorf(`failed to marshal value: %w`, err)
	}

	*v = PatchOperationValue(serialized)
	return nil
}

/*
func (b *PatchOperationBuilder) Value(v interface{}) *PatchOperationBuilder {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.once.Do(b.init)
	if b.err != nil {
		return b
	}

	serialized, err := json.Marshal(v)
	if err != nil {
		b.err = fmt.Errorf(`failed to marshal value: %w`, err)
	}

	if err := b.object.Set(`value`, json.RawMessage(serialized)); err != nil {
		b.err = err
	}
	return b
}*/
