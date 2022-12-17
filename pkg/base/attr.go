package base

// хранилище для атрибутов html тега
type Attributes map[string]string

// Функция строит строку атрибутов для html тега
func (a Attributes) build() string {
	attributes := ""

	for key, value := range a {
		if value == "" {
			attributes += key + " "
		} else {
			attributes += key + "=\"" + value + "\" "
		}
	}

	return attributes
}

// Интерфейс, которому должны удовлетворять структуры для хранения тегов
type AttributeEditor interface {
	Get(string) string
	Set(string, string)
	Append(string, string)
	Delete(string)
}
