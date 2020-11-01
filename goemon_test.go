package dotenv

import (
	"os"
	"reflect"
	"testing"
)

func TestLoadEnv(t *testing.T) {
	os.Clearenv()

	dir, _ := os.Getwd()
	filename := dir + "/fixtures/env/.foo_env"

	if err := LoadEnv(filename); err != nil {
		t.Errorf("%v", err)
	}

	expectedEnvValues := map[string]string{
		"FOO_1": "foo_1",
		"FOO_2": "foo_2",
	}

	for k := range expectedEnvValues {
		actual := os.Getenv(k)
		expected := expectedEnvValues[k]

		if actual != expected {
			t.Errorf("actual:%v expected:%v", actual, expected)
		}
	}
}

func TestLoadMultipleEnv(t *testing.T) {
	os.Clearenv()

	dir, _ := os.Getwd()
	env := dir + "/fixtures/env/.foo_env"
	localEnv := dir + "/fixtures/env/.bar_env"

	if err := LoadEnv(env, localEnv); err != nil {
		t.Errorf("%v", err)
	}

	expectedEnvValues := map[string]string{
		"FOO_1": "foo_1",
		"FOO_2": "foo_2",
		"BAR_1": "bar_1",
		"BAR_2": "bar_2",
	}

	for k := range expectedEnvValues {
		actual := os.Getenv(k)
		expected := expectedEnvValues[k]

		if actual != expected {
			t.Errorf("actual:%v expected:%v", actual, expected)
		}
	}
}

func TestSetFilename(t *testing.T) {
	cases := []struct {
		actual   []string
		expected []string
	}{
		{
			actual:   setFilename([]string{".foo_env"}),
			expected: []string{".foo_env"},
		},
		{
			actual:   setFilename([]string{}),
			expected: []string{".env"},
		},
	}

	for _, c := range cases {
		if !reflect.DeepEqual(c.actual, c.expected) {
			t.Errorf("actual:%v expected:%v\n", c.actual, c.expected)
		}
	}
}

func TestReadFile(t *testing.T) {
	dir, _ := os.Getwd()
	filename := dir + "/fixtures/env/.foo_env"

	actual, err := readFile(filename)
	if err != nil {
		t.Errorf("%v", err)
	}

	expected := []string{"FOO_1=foo_1", "FOO_2=foo_2"}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual:%v expected:%v", actual, expected)
	}
}

func TestSetEnv(t *testing.T) {
	os.Clearenv()

	lines := []string{"FOO_1=foo_1", "FOO_2=foo_2"}

	setEnv(lines)

	expectedEnvValues := map[string]string{
		"FOO_1": "foo_1",
		"FOO_2": "foo_2",
	}

	for k := range expectedEnvValues {
		actual := os.Getenv(k)
		expected := expectedEnvValues[k]

		if actual != expected {
			t.Errorf("actual:%v expected:%v", actual, expected)
		}
	}
}
