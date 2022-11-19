package marker

import (
	"bytes"
	"fmt"
)

type AttributeBuilder func() (string, string)

func Attributes(builders ...AttributeBuilder) AttributeStorage {
	storage := make(AttributeStorage)

	for _, builder := range builders {
		key, value := builder()
		storage[key] = value
	}

	return storage
}

// Атрибуты узла
type AttributeStorage map[string]string

func (a AttributeStorage) String() string {
	buffer := new(bytes.Buffer)

	for key, value := range a {
		if value == "" {
			buffer.Write([]byte(fmt.Sprintf(" %s", key)))
		} else {
			buffer.Write([]byte(fmt.Sprintf(` %s="%s"`, key, value)))
		}
	}

	return buffer.String()
}

func Set(name, value string) AttributeBuilder {
	return func() (string, string) {
		return name, value
	}
}

func SetClass(class string) AttributeBuilder {
	return func() (string, string) {
		return "class", class
	}
}

func SetID(id string) AttributeBuilder {
	return func() (string, string) {
		return "id", id
	}
}

func SetStyle(style string) AttributeBuilder {
	return func() (string, string) {
		return "style", style
	}
}
