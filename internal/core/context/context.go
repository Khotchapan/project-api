package context

import (
	"github.com/labstack/echo/v4"
)

type CustomContext struct {
	echo.Context
	Parameters interface{}
}

func SetCustomContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &CustomContext{Context: c}
		return next(cc)
	}
}

//const pathKey = "path"

// BindAndValidate bind and validate form
func (c *CustomContext) BindAndValidate(i interface{}) error {
	if err := c.Bind(i); err != nil {
		return err
	}
	//c.parsePathParams(i)
	if err := c.Validate(i); err != nil {
		return err
	}
	//c.Parameters = i
	return nil
}

// func (c *CustomContext) parsePathParams(form interface{}) {
// 	formValue := reflect.ValueOf(form)
// 	if formValue.Kind() == reflect.Ptr {
// 		formValue = formValue.Elem()
// 	}
// 	t := reflect.TypeOf(formValue.Interface())
// 	for i := 0; i < t.NumField(); i++ {
// 		tag := t.Field(i).Tag.Get(pathKey)
// 		if tag != "" {
// 			fieldName := t.Field(i).Name
// 			paramValue := formValue.FieldByName(fieldName)
// 			if paramValue.IsValid() {
// 				paramValue.Set(reflect.ValueOf(c.Param(tag)))
// 			}
// 		}
// 	}
// }
