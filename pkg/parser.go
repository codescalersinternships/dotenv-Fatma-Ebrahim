package dotenv_parser

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func ParseString(s string) (map[string]string, error) {
	parsed := make(map[string]string)
	lines := strings.Split(s, "\n")
	var value string
	var key string
	quote_flag := false
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "export ") {
			line = strings.TrimSpace(strings.TrimPrefix(line, "export "))
		}
		separator_index := strings.Index(line, "=")
		if separator_index == -1 {
			// handle comments and empty lines
			if line == "" || line[0] == '#' {
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
			if value == "" {
				// empty value
				parsed[key] = value
				key = ""
				value = ""
				continue
			}
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
				key = ""
				value = ""
				continue
			}
		}

		// inline comments
		comment_idx := strings.Index(value, "#")
		if comment_idx != -1 {
			value = strings.TrimSpace(value[:comment_idx])
		}
		// non-quoted value
		if value[0] != '"' {
			parsed[key] = value
			continue
		}

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

	return ParseString(string(content))
}

func LoadEnvFile(path string) error {
	res, err := ParseFile(path)
	if err != nil {
		return err
	}
	for key, value := range res {
		err:=os.Setenv(key, value)
		if err != nil {
			return err
		}
		
	}
	return nil
}

func LoadEnvString(s string) error {
	res, err := ParseString(s)
	if err != nil {
		return err
	}
	for key, value := range res {
		err:=os.Setenv(key, value)
		if err != nil {
			return err
		}
		
	}
	return nil
}