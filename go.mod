module github.com/todaatsushi/repo

go 1.18

require (
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/cobra v1.4.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	internal/repoconf v1.0.0
	internal/utils v1.0.0
)

replace internal/repoconf => ./internal/repoconf

replace internal/utils => ./internal/utils
