package dotenv_parser

import (
	"testing"
	"reflect"
)

func TestParseString(t *testing.T) {
	t.Run("test parse string with no error", func(t *testing.T){
		data := "key1=value1\nkey2=value2"
		got, err := ParseString(data)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		expected := map[string]string{
			"key1": "value1",
			"key2": "value2",
		}
		if !reflect.DeepEqual(got, expected) {
			t.Fatalf("expected %v, got %v", expected, got)
		}
	})
}