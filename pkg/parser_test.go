package dotenv_parser

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseString(t *testing.T) {
	t.Run("test parse string with no error", func(t *testing.T){
		data := "key1=value1\nkey2=value2"
		got, err := ParseString2(data)
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

	t.Run("test multi-line value" , func(t *testing.T) {
		data :=fmt.Sprintf("KEY3=\"MULTI\nLINE\nVALUE3 and with # inside\"")
		got, err := ParseString2(data)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		expected := map[string]string{
			"KEY3": fmt.Sprintf("\"MULTI\nLINE\nVALUE3 and with # inside\""),
		}
		if !reflect.DeepEqual(got, expected) {
			t.Fatalf("expected %v, got %v", expected, got)
		}
	})

	t.Run("test empty value", func(t *testing.T) {
		data := "KEY3"
		got, err := ParseString2(data)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		expected := map[string]string{
			"KEY3": "",
		}
		if !reflect.DeepEqual(got, expected) {
			t.Fatalf("expected %v, got %v", expected, got)
		}
	})

	t.Run("test comments", func(t *testing.T) {
		data := "# comment\nKEY=VALUE"
		got, err := ParseString2(data)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		expected := map[string]string{
			"KEY": "VALUE",
		}
		if !reflect.DeepEqual(got, expected) {
			t.Fatalf("expected %v, got %v", expected, got)
		}
	})

	t.Run("test inline comments", func(t *testing.T) {
		data := "KEY=VALUE # comment"
		got, err := ParseString2(data)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		expected := map[string]string{
			"KEY": "VALUE",
		}
		if !reflect.DeepEqual(got, expected) {
			t.Fatalf("expected %v, got %v", expected, got)
		}
	})

	t.Run("test missing quotes", func(t *testing.T) {
		data := fmt.Sprintf("\"KEY=\"VALUE")
		_, err := ParseString2(data)
		if err == nil {
			t.Fatalf("expected error missing closing quote, got %v", err)
		}
	})

}