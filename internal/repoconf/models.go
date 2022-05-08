package repoconf

type Memo struct {
	Tags        []string `json:"tags"`
	Content     string   `json:"content"`
	Description string   `json:"description"`
	Name        string   `json:"name"`
}

type Config struct {
	Tags  []string        `json:"tags"`
	Items map[string]Memo `json:"items"`
}

func NewMemo(content string, name string) Memo {
	return Memo{Content: content, Name: name}
}

func NewConfig() Config {
	return Config{}
}
