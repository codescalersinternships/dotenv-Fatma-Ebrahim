package dotenv_parser

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func ParseString(s string) (map[string]string, error) {
	parsed := make(map[string]string)
	key_counter := 0
	for len(s) > 0 {
		fmt.Println("Remaining string to parse:", s)
		fmt.Println("Parsed so far:", parsed)
		// handle comments at the beginning of the line
		if s[0] == '#' {
			end_idx := strings.Index(s, "\n")
			if end_idx == -1 {
				break
			}
			s = s[end_idx+1:]
			continue
		}
		key_counter++
		separator_idx := strings.Index(s, "=")
		if separator_idx == -1 {
			return nil, fmt.Errorf("missing '=' for key %d", key_counter)
		}

		// handle spaces around keys and values
		key := strings.TrimSpace(s[:separator_idx])
		value := strings.TrimSpace(s[separator_idx+1:])
		s = key + "=" + value

		start_idx := len(key) + 1
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
				pair_end_idx = len(s) - 1
			}
			s = strings.TrimSpace(s[pair_end_idx+len(value)+1:])

		} else {
			// handle non-quoted values
			end_idx := strings.Index(value, "\n")
			if end_idx == -1 {
				end_idx = len(value) - 1
			}
			if start_idx > end_idx {
				return nil, fmt.Errorf("invalid format in key %d", key_counter)
			}

			value = strings.TrimSpace(value[:end_idx+1])

			// handle inline comments
			comment_idx := strings.Index(value, "#")
			if comment_idx != -1 {
				value = strings.TrimSpace(value[:comment_idx])
			}

			fmt.Println("Parsed value after processing:", value)
			s = strings.TrimSpace(s[end_idx+len(value):])
		}
		parsed[key] = value
	}
	return parsed, nil
}

func ParseString2(s string) (map[string]string, error) {
	parsed := make(map[string]string)
	lines := strings.Split(s, "\n")
	var value string
	var key string
	quote_flag := false
	for _, line := range lines {
		line = strings.TrimSpace(line)
		separator_index := strings.Index(line, "=")
		if separator_index == -1 {
			// handle comments and empty lines
			if line == "" ||line[0] == '#' {
				continue
			}
			// end and middle of quoted value
			if quote_flag {
				closing_quote_idx := strings.Index(line, "\"")
				if closing_quote_idx != -1 {
					line = strings.TrimSpace(line[:closing_quote_idx+1])
					quote_flag = false
					value += "\n" + line
					parsed[key] = value
					key = ""
					value = ""
					continue
				}
				value += "\n" + line
				continue
			}
			// empty value
			key = line
			value = ""
			parsed[key] = value
			key = ""
			value = ""
			continue
		} else if !quote_flag {
			key = strings.TrimSpace(line[:separator_index])
			value = strings.TrimSpace(line[separator_index+1:])
			if value[0] == '"' {
				// start of quoted value
				quote_flag = true
			}
			// one line quoted value
			val := value[1:]
			closing_quote_idx := strings.Index(val, "\"")
			if closing_quote_idx != -1 {
				value = value[:closing_quote_idx+2]
				quote_flag = false
				parsed[key] = value
				continue
			}
		} else {
			value += "\n" + line
			parsed[key] = value
			continue
		}

		comment_idx := strings.Index(value, "#")
		if comment_idx != -1 {
			value = strings.TrimSpace(value[:comment_idx])
		}
		parsed[key] = value

	}
	if quote_flag {
		return nil, fmt.Errorf("missing closing quote")
	}
	return parsed, nil
}



func ParseFile(path string) (map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return ParseString2(string(content))
}
