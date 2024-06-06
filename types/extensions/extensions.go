package extensions

import (
	"errors"

	"github.com/go-viper/mapstructure/v2"
)

type Extensions interface {
	Get(name string, target interface{}) (bool, error)
	GetMap() map[string]interface{}
	DeepCopyExtensions() Extensions
}

type Map map[string]interface{}

func (m Map) Get(name string, target interface{}) (bool, error) {
	if v, ok := m[name]; ok {
		err := mapstructure.Decode(v, target)
		return true, err
	}
	return false, nil
}

func (m Map) GetMap() map[string]interface{} {
	return m
}

func (m Map) DeepCopyExtensions() Extensions {
	result := make(Map)
	for k, v := range m {
		val, ok := v.(Extensions)
		if !ok {
			panic(errors.New("value not deepcopyable"))
		}
		result[k] = val.DeepCopyExtensions()
	}
	return result
}
