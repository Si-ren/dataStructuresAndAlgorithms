package factory

import (
	"errors"
	"fmt"
)

type Config struct {
	IP string
}

type ConfigParse interface {
	Parse(string) (*Config, error)
}

type JsonConfigParse struct{}

func (j *JsonConfigParse) Parse(str string) (c *Config, e error) {
	// TODO: implement function
	fmt.Println("this is json parse")
	return c, nil
}

type YamlConfigParse struct{}

func (j *YamlConfigParse) Parse(str string) (c *Config, e error) {
	// TODO: implement function
	fmt.Println("this is yaml parse")
	return c, nil
}

// 通过工厂创建不同的对象
func NewConfigParse(mode string) (ConfigParse, error) {
	switch mode {
	case "json":
		return &JsonConfigParse{}, nil
	case "yaml":
		return &YamlConfigParse{}, nil
	default:
		fmt.Println("this is default mode")
		return nil, errors.New("this mode not supported")
	}
}
