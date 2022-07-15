package resource

import (
	"encoding/json"
	"fmt"
)

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
}
