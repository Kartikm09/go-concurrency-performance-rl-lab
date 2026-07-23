// Package retry separates retry classification, timing, and transport work.
package retry

import "context"

type Classifier interface{ Retryable(error) bool }
type Backoff interface {
	Wait(context.Context, int) error
}
type Operation func(context.Context) error

func Do(ctx context.Context, attempts int, classifier Classifier, backoff Backoff, operation Operation) error {
	var err error
	for attempt := 1; attempt <= attempts; attempt++ {
		if err = operation(ctx); err == nil || !classifier.Retryable(err) {
			return err
		}
		if attempt < attempts {
			if waitErr := backoff.Wait(ctx, attempt); waitErr != nil {
				return waitErr
			}
		}
	}
	return err
}
