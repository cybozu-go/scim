package resource

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (v *ListResponse) UnmarshalJSON(data []byte) error {
	v.itemsPerPage = nil
	v.resources = nil
	v.schemas = nil
	v.startIndex = nil
	v.totalResults = nil
	v.extra = nil
	dec := json.NewDecoder(bytes.NewReader(data))
	{ // first token
		tok, err := dec.Token()
		if err != nil {
			return fmt.Errorf("failed to read next token: %s", err)
		}
		tok, ok := tok.(json.Delim)
		if !ok {
			return fmt.Errorf("expected first token to be '{', got %c", tok)
		}
	}
	var extra map[string]interface{}

LOOP:
	for {
		tok, err := dec.Token()
		if err != nil {
			return fmt.Errorf("failed to read next token: %s", err)
		}
		switch tok := tok.(type) {
		case json.Delim:
			if tok == '}' {
				break LOOP
			} else {
				return fmt.Errorf("unexpected token %c found", tok)
			}
		case string:
			switch tok {
			case ListResponseItemsPerPageKey:
				var x int
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "itemsPerPage": %w`, err)
				}
				v.itemsPerPage = &x
			case ListResponseResourcesKey:
				var rawlist []json.RawMessage
				if err := dec.Decode(&rawlist); err != nil {
					return fmt.Errorf(`failed to decode value for key "resources": %w`, err)
				}

				list := make([]interface{}, len(rawlist))
			OUTER:
				for i, raw := range rawlist {
					var x struct {
						Schemas []string `json:"schemas"`
					}
					if err := json.Unmarshal(raw, &x); err != nil {
						return fmt.Errorf(`failed to decode hint for resource %d: %w`, i, err)
					}

					for _, schema := range x.Schemas {
						switch schema {
						case UserSchemaURI:
							var u User
							if err := json.Unmarshal(raw, &u); err != nil {
								return fmt.Errorf(`failed to decode value %d for key "resources" as User resource: %w`, i, err)
							}
							list[i] = &u
							continue OUTER
						case GroupSchemaURI:
							var g Group
							if err := json.Unmarshal(raw, &g); err != nil {
								return fmt.Errorf(`failed to decode value %d for key "resources" as Group resource: %w`, i, err)
							}
							list[i] = &g
							continue OUTER
						}
						return fmt.Errorf(`failed to decode value %d for key "resources": unhandled schema %#v`, i, x.Schemas)
					}
				}
				v.resources = list
			case ListResponseSchemasKey:
				var x schemas
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "schemas": %w`, err)
				}
				v.schemas = &x
			case ListResponseStartIndexKey:
				var x int
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "startIndex": %w`, err)
				}
				v.startIndex = &x
			case ListResponseTotalResultsKey:
				var x int
				if err := dec.Decode(&x); err != nil {
					return fmt.Errorf(`failed to decode value for key "totalResults": %w`, err)
				}
				v.totalResults = &x
			default:
				var x interface{}
				if rx, ok := registry.Get(tok); ok {
					x = rx
					if err := dec.Decode(x); err != nil {
						return fmt.Errorf(`failed to decode value for key %q: %w`, tok, err)
					}
				} else {
					if err := dec.Decode(&x); err != nil {
						return fmt.Errorf(`failed to decode value for key %q: %w`, tok, err)
					}
				}
				if extra == nil {
					extra = make(map[string]interface{})
				}
				extra[tok] = x
			}
		}
	}
	if extra != nil {
		v.extra = extra
	}
	return nil
}
