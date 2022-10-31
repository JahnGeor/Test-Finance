package utils

import "time"

func TryConnection(check func() error, attempts int, delay time.Duration) (err error) {
	for attempts > 0 {
		if err = check(); err != nil {
			time.Sleep(delay)
			attempts--
			continue
		}
		return nil
	}
	return
}
