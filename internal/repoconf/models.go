package repoconf

type Memo struct {
	Tags        []string `json:"tags"`
	Content     string   `json:"content"`
	Description string   `json:"description"`
}

type Config struct {
	Tags  []string        `json:"tags"`
	Items map[string]Memo `json:"items,omitempty"`
}

func NewMemo(content string) Memo {
	return Memo{Content: content}
}

func NewConfig() Config {
	return Config{}
}
