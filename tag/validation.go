package tag

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

// 接口：方法的集合。实现接口中所有的方法，就实现了接口。
// 空接口中没有任何方法，所以可以认为所有类型都实现了空接口。也就是空接口可以接受所有类型的原因。

type Validation struct {
	Data   interface{}
	Errors []error
}

func NewValidation(structData interface{}) *Validation {
	v := new(Validation)
	v.Data = structData
	return v
}

const ValidTagName = "valid"

func (v *Validation) Validate() {
	var errs []error
	runData := reflect.ValueOf(v.Data)
	if runData.Type().Kind() != reflect.Struct {
		panic("must be a struct")
	}

	for i := 0; i < runData.NumField(); i++ {
		tag := runData.Type().Field(i).Tag.Get(ValidTagName)
		curValidator := v.GetValidator(tag)
		ok, err := curValidator.Validate(runData.Field(i).Interface())
		if !ok && err != nil {
			errs = append(v.Errors, err)
		}
	}

	v.Errors = errs
}

var (
	MaxRe = regexp.MustCompile(`\Amaxsize\(\d+\)\z`)
	MinRe = regexp.MustCompile(`\Aminsize\(\d+\)\z`)
)

func (v *Validation) GetValidator(tag string) Validator {
	tagArgs := strings.Split(strings.ToLower(tag), ";")
	switch tagArgs[0] {
	case "string":
		sv := StringValidator{}
		for _, s := range tagArgs[1:] {
			switch {
			case MaxRe.MatchString(s):
				fmt.Sscanf(s, "maxsize(%d)", &sv.Max)
			case MinRe.MatchString(s):
				fmt.Sscanf(s, "minsize(%d)", &sv.Min)
			}
		}

		if sv.Min > sv.Max {
			panic("max value must be greater then min value")
		}
		return &sv
	case "number":
		nv := NumberValidator{}
		fmt.Sscanf(tagArgs[1], "range(%d,%d)", &nv.Min, &nv.Max)

		if nv.Min > nv.Max {
			panic("max value must be greater then min value")
		}
		return &nv
	case "email":
		return &EmailValidator{}
	}

	return &DefaultValidator{}
}

type Validator interface {
	Validate(date interface{}) (bool, error)
}

// DefaultValidator 默认验证器，什么都不验证
type DefaultValidator struct {
}

func (sv *DefaultValidator) Validate(data interface{}) (bool, error) {
	return true, nil
}

// StringValidator 字符串验证器
type StringValidator struct {
	Min int
	Max int
}

func (dv *StringValidator) Validate(data interface{}) (bool, error) {
	str := data.(string)
	l := len(str)
	if l < dv.Min {
		return false, fmt.Errorf("string length should be greater than %d", dv.Min)
	}
	if l > dv.Max {
		return false, fmt.Errorf("string length should be less than %d", dv.Max)
	}

	return true, nil
}

type NumberValidator struct {
	Min int
	Max int
}

func (nv *NumberValidator) Validate(data interface{}) (bool, error) {
	n := data.(int)
	if n < nv.Min {
		return false, fmt.Errorf("number should be greater then %d", nv.Min)
	}
	if n > nv.Max {
		return false, fmt.Errorf("number should be less than %d", nv.Max)
	}

	return true, nil
}

var emailRe = regexp.MustCompile(`\A[\w+\d\-.]+@[a-z\d\-]+(\.[a-z\d\-]+)*\.[a-z]+\z`)

type EmailValidator struct {
}

func (ev *EmailValidator) Validate(data interface{}) (bool, error) {
	email := data.(string)
	if !emailRe.MatchString(email) {
		return false, fmt.Errorf("email format error")
	}

	return true, nil
}
