package repoconf

type Item struct {
	Tags        []string `json:"tags"`
	Content     string   `json:"content"`
	Description string   `json:"description"`
	Name        string   `json:"name"`
}

type Config struct {
	Tags  []string        `json:"tags"`
	Items map[string]Item `json:"items"`
}

func NewItem(content string, name string) Item {
	return Item{Content: content, Name: name}
}

func NewConfig() Config {
	return Config{}
}
