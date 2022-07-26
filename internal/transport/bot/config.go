package bot

import "time"

type Config struct {
	Token             string
	LongPollerTimeout time.Duration
}
