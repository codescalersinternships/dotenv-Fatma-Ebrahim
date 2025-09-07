package dotenv_parser

import (
	"fmt"
	"strings"
)

func ParseString(s string) (map[string]string, error) {
	parsed := make(map[string]string)
	key_counter := 0
	for len(s) > 0 {
		fmt.Println("Remaining string to parse:", s)
		key_counter++
		separator_idx := strings.Index(s, "=")
		if separator_idx == -1 {
			return nil, fmt.Errorf("invalid format in key %d", key_counter)
		}
		key := strings.TrimSpace(s[:separator_idx])
		fmt.Println("Parsed key:", key)
		value := strings.TrimSpace(s[separator_idx+1:])
		fmt.Println("Parsed value before processing:", value)
	
		start_idx := separator_idx + 1
		if value[0] == '"' {
			// handle quoted values
			value = value[1:]
			end_idx := strings.Index(value, "\"")
			if end_idx == -1 {
				return nil, fmt.Errorf("missing closing quote in key %d", key_counter)
			}

			value = strings.TrimSpace(value[:end_idx])
			pair_end_idx := strings.Index(s, value)
			if pair_end_idx >= len(s) {
				pair_end_idx = len(s)-1
			}
			s = strings.TrimSpace(s[pair_end_idx+len(value)+1:])

		} else {
			// handle non-quoted values
			end_idx := strings.Index(s, "\n")
			if end_idx == -1 {
				end_idx = len(s)
			}
			if start_idx > end_idx {
				return nil, fmt.Errorf("invalid format in key %d", key_counter)
			}
			value = strings.TrimSpace(s[start_idx:end_idx])
			if end_idx >= len(s){
				end_idx = len(s)-1
			}
			s = strings.TrimSpace(s[end_idx+1:])
		}
		parsed[key] = value
	}
	return parsed, nil
	// support comments, multilines and quoted values, export keyword
}

func ParseFile(path string) (map[string]string, error) {
	return nil, fmt.Errorf("not implemented")
}
