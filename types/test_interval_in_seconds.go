package types

import (
	"encoding/json"
	"fmt"
)

// Configure how often availability tests should be performed.
type TestIntervalInSeconds int

var allowedIntervals = []int{60, 300, 600, 900, 1800, 3600, 7200, 14400}

func (t *TestIntervalInSeconds) UnmarshalJSON(data []byte) error {
	var value int
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	for _, allowedInterval := range allowedIntervals {
		if value == allowedInterval {
			*t = TestIntervalInSeconds(value)
			return nil
		}
	}

	return fmt.Errorf("invalid TestIntervalInSeconds: must be one of %v", allowedIntervals)
}

func (t TestIntervalInSeconds) MarshalJSON() ([]byte, error) {
	return json.Marshal(int(t))
}
