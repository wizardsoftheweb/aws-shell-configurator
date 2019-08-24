package main

type AwsSetting struct {
	EnvironmentVariable string   `json:"env"`
	Value               string   `json:"value"`
	AllowedValues       []string `json:"allowed"`
}

func (s *AwsSetting) Set(newValue string) {
	if 0 < len(s.AllowedValues) {
		for _, allowedValue := range s.AllowedValues {
			if allowedValue == newValue {
				s.Value = newValue
			}
		}
	} else {
		s.Value = newValue
	}
}
