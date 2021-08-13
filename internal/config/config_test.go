package config

import (
	"errors"
	"testing"
)

const configPath = "./../../configs/"

var updateConfigTestCases = []struct {
	caseName   string
	configPath string
	expected   error
}{
	{
		"success update",
		configPath + "config.json",
		nil,
	},
	{
		"error update",
		configPath + "config_wrong.json",
		errors.New("open ./../../configs/config_wrong.json: no such file or directory"),
	},
}

func ConfigUpdaterMock(configPath string) error {
	for i := 0; i < 5; i++ {
		err := readConfig(configPath)
		if err != nil {
			return err
		}
	}

	return nil
}

func TestUpdateConfig(t *testing.T) {
	for _, testCase := range updateConfigTestCases {
		err := UpdateConfig(testCase.configPath, ConfigUpdaterMock)
		if err != nil {
			if testCase.expected == nil {
				t.Fatalf(`%v. Expected err: %v. Actual: %v`, testCase.caseName, testCase.expected, err)
			}
			if err.Error() != testCase.expected.Error() {
				t.Fatalf(`%v. Expected err: %v. Actual: %v`, testCase.caseName, testCase.expected, err)
			}
		}
	}
}
