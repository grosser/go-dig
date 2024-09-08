package dig

import (
	"fmt"
)

// Dig extracts the nested value specified by the keys from v
func Dig(v interface{}, keys ...interface{}) (interface{}, error) {
	n := len(keys)
	for i, key := range keys {
		// try lookup with string key
		if stringKey, ok := key.(string); ok {
			raw, ok := v.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("%v isn't a map", v)
			}
			found, ok := raw[stringKey]
			if !ok {
				return nil, fmt.Errorf("key %v not found in %v", stringKey, v)
			}
			v = found
			if i == n-1 {
				return v, nil
			}
			continue
		}

		// try lookup with int key
		if intKey, ok := key.(int); ok {
			raw, ok := v.([]interface{})
			if !ok {
				return nil, fmt.Errorf("%v isn't a slice", v)
			}
			if intKey < 0 || intKey >= len(raw) {
				return nil, fmt.Errorf("index out of range [%v]: %v", intKey, raw)
			}
			v = raw[intKey]
			if i == n-1 {
				return v, nil
			}
			continue
		}

		// not a supported key format
		return nil, fmt.Errorf("unsupported key type: %v", key)
	}

	return nil, fmt.Errorf("no key given")
}
