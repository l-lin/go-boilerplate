package conf

import (
	"testing"
)

// cannot unit test prompts with -race
// see https://github.com/manifoldco/promptui/pull/169 or https://github.com/manifoldco/promptui/issues/148
//type mockReadWriteCloser struct {
//	io.Reader
//}
//
//func (m mockReadWriteCloser) Close() error {
//	return nil
//}
//
//func TestCreator_Create(t *testing.T) {
//	// GIVEN
//	expected := &Conf{
//		SomeProperty: "test some property",
//		Email:        "test@test.com",
//	}
//	c := creator{
//		Conf:   NewConf(),
//		stdout: os.Stdout,
//	}
//
//	// WHEN
//	c.stdin = &mockReadWriteCloser{strings.NewReader(fmt.Sprintf("%s\n", expected.SomeProperty))}
//	c.askSomeProperty()
//	c.stdin = &mockReadWriteCloser{strings.NewReader(fmt.Sprintf("%s\n", expected.Email))}
//	c.askEmail()
//	actual, err := c.create()
//
//	// THEN
//	if err != nil {
//		t.Errorf("no error expected, got %v", err)
//		t.Fail()
//	}
//	if actual.SomeProperty != expected.SomeProperty || actual.Email != expected.Email {
//		t.Errorf("expected: %v, actual: %v", expected, actual)
//	}
//}

func TestValidateNotEmpty(t *testing.T) {
	var tests = map[string]struct {
		given            string
		expectedHasError bool
	}{
		"valid content": {
			given:            "foobar",
			expectedHasError: false,
		},
		"only spaces": {
			given:            "  ",
			expectedHasError: true,
		},
		"only tabs": {
			given:            "\t",
			expectedHasError: true,
		},
		"empty string": {
			given:            "",
			expectedHasError: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual := validateNotEmpty(tt.given)
			if tt.expectedHasError && actual == nil {
				t.Errorf("expected error")
			}
			if !tt.expectedHasError && actual != nil {
				t.Errorf("expected valid content")
			}
		})
	}
}
func TestValidateEmail(t *testing.T) {
	var tests = map[string]struct {
		given            string
		expectedHasError bool
	}{
		"valid email": {
			given:            "louis.lin@bioserenity.com",
			expectedHasError: false,
		},
		"invalid email": {
			given:            "louis.lin@",
			expectedHasError: true,
		},
		"not enough characters": {
			given:            "po",
			expectedHasError: true,
		},
		"too many characters": {
			given:            `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam efficitur finibus turpis id facilisis. Donec eu sem mattis est euismod bibendum. Curabitur interdum purus rhoncus tortor vestibulum dictum. Donec non sollicitudin turpis. Donec mollis eget sem a rhoncus. Duis dictum consectetur enim, sed commodo metus lacinia vitae. Praesent non imperdiet mi, eget fermentum turpis. Phasellus convallis elit ut dolor molestie sodales. Sed aliquam fermentum posuere. Donec nec aliquet felis. Morbi varius eros ligula, et tincidunt tellus laoreet ut. `,
			expectedHasError: true,
		},
		"empty string": {
			given:            "",
			expectedHasError: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual := validateEmail(tt.given)
			if tt.expectedHasError && actual == nil {
				t.Errorf("expected error")
			}
			if !tt.expectedHasError && actual != nil {
				t.Errorf("expected valid email")
			}
		})
	}

}
