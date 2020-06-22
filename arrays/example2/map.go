package example2

type Map struct{}

func (m Map) Run() string {
	programmingCharacteristics := map[string]string{
		"C":          "Super hard low level programming",
		"Go":         "Clear concise language",
		"Javascript": "Most popular language",
	}

	return programmingCharacteristics["Go"]
}
