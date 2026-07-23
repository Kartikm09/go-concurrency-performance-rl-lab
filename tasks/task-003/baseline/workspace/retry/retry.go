package retry

import (
	"context"
	"time"
)

func Do(ctx context.Context, attempts int, operation func(context.Context) error, retryable func(error) bool, sleep func(time.Duration)) error {
	var err error
	for attempt := 1; attempt <= attempts; attempt++ {
		err = operation(ctx)
		if err == nil || !retryable(err) {
			return err
		}
		if attempt < attempts {
			sleep(time.Duration(attempt) * time.Millisecond)
		}
	}
	return err
}
