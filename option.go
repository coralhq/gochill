package gochill

import "fmt"

//Option struct used to set optional field to the log
type Option struct {
	Key   string
	Value interface{}
}

//Set used to set new key value
func (o *Option) Set(key string, value interface{}) {
	o.Key = "_" + key
	o.Value = fmt.Sprintf("%v", value)
}

//NewOption used to create new Option instance
func NewOption() *Option {
	o := Option{}
	return &o
}

//MergeOptionsToMap used to merge given options with Option type
//merge them into one map string interface
func MergeOptionsToMap(options []Option) map[string]interface{} {
	var maps map[string]interface{}
	maps = make(map[string]interface{})

	//only continue process if list is not empty
	if len(options) > 0 {
		for _, option := range options {
			maps[option.Key] = option.Value
		}
	}

	return maps
}

//O used to set optional fields
func O(key string, value interface{}) Option {
	o := NewOption()
	o.Set(key, value)

	return *o
}
