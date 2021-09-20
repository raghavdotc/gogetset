package gogetset

import (
	"errors"
	"reflect"
)

func Get(path string, input interface{}) (interface{}, error) {
	var (
		nextStepIsArr bool
		idx           int
		arrKey        string
		err           error
		temp          interface{}
	)
	temp = input
	steps := stepsFromPath(path)
	for stepIndex, step := range steps {
		if !nextStepIsArr {
			arrKey, idx, nextStepIsArr, err = nextStepIsArray(step)
			if err != nil {
				return nil, err
			}
			if nextStepIsArr {
				step = arrKey
			}
		}
		tempVal := reflect.ValueOf(temp)

		// this block of code executes for the `temp` whose value is of kind `reflect.Map`
		if tempVal.Kind() == reflect.Map {
			valVal := tempVal.MapIndex(reflect.ValueOf(step))

			if valVal.IsValid() && !nextStepIsArr && isLastStep(stepIndex, steps) {
				return valVal.Interface(), nil
			} else if !valVal.IsValid() {
				return nil, invalidPathError
			}
			temp = valVal.Interface()

			// this block of code handles the occurrence of a slice
			if nextStepIsArr {
				newTemp, replaceTemp, returnTemp, err := getFromSlice(temp, idx, isLastStep(stepIndex, steps))
				if returnTemp {
					return newTemp, err
				}
				if replaceTemp {
					temp = newTemp
				}
			}

			// reaches here only it is not the last step and keyExistsInMap
			// hence continue to iterate over the next steps
			continue
		}

		if tempVal.Kind() == reflect.Struct || tempVal.Kind() == reflect.Ptr {

			if tempVal.Kind() == reflect.Ptr {
				tempVal = tempVal.Elem()
			}

			// 'dereference' with Elem() and get the field by name
			fieldVal := tempVal.FieldByName(step)

			// check if the field exists in the interface
			if !fieldVal.IsValid() {
				return nil, invalidPathError
			}
			temp = fieldVal.Interface()
			if nextStepIsArr {
				newTemp, replaceTemp, returnTemp, err := getFromSlice(temp, idx, isLastStep(stepIndex, steps))
				if returnTemp {
					return newTemp, err
				}
				if replaceTemp {
					temp = newTemp
				}
			}
			if isLastStep(stepIndex, steps) {
				return fieldVal.Interface(), nil
			}
		}

	}
	return nil, errors.New("path not found")
}

func getFromSlice(temp interface{}, idx int, isLastStep bool) (val interface{}, replaceTemp bool, returnTemp bool, err error) {
	tempVal := reflect.ValueOf(temp)
	if tempVal.Kind() == reflect.Slice {
		arrEle := tempVal.Index(idx)
		if arrEle.IsValid() && isLastStep {
			val = arrEle.Interface()
			returnTemp = true
			return
		} else if !arrEle.IsValid() && isLastStep {
			err = invalidPathError
			returnTemp = true
			return
		}
		replaceTemp = true
		val = arrEle.Interface()
		return
	}
	err = sliceExpectedNotFound
	replaceTemp = true
	return
}
