package gogetset

import (
	"fmt"
	"reflect"
)

const (
	setPtrFieldStr  = "setPtrField"
	setMapIndexStr  = "setMapIndex"
	forLoop         = "forLoop"
	AfterStepChange = "afterStepChange"
)

func Set(path string, input interface{}, data interface{}) (err error) {
	var (
		nextStepIsArr  bool
		idx            int
		arrKey         string
		lastStep       bool
		stepReflectVal reflect.Value
		//prevStepObj    interface{}
		//success bool
	)
	stepReflectVal = reflect.ValueOf(input)

	if stepReflectVal.Kind() == reflect.Struct {
		err = structNotAcceptedError
		return
	}

	if input == nil {
		err = nilNotAccepted
		return
	}
	var stepObj interface{}
	//prevStepObj := input
	steps := stepsFromPath(path)
	for stepIndex, step := range steps {
		if stepReflectVal.IsValid() {
			stepObj = stepReflectVal.Interface()
			fmt.Printf("%s %s: Step= %s, Obj: %v, Kind: %s\n", forLoop, Begin, step, stepObj, stepReflectVal.Kind().String())
		} else {
			err = unknownError
		}
		if !nextStepIsArr {
			arrKey, idx, nextStepIsArr, err = nextStepIsArray(step)
			if err != nil {
				return
			}
			if nextStepIsArr {
				step = arrKey
			}
		}
		lastStep = isLastStep(stepIndex, steps)

		// stepObj of kind Map, should have a key `step` with the right value set
		stepReflectVal, err = setMapIndex(&stepReflectVal, step, data, nextStepIsArr, lastStep, idx)
		if stepReflectVal.IsValid() {
			stepObj = stepReflectVal.Interface()
			fmt.Printf("%s %s: Step= %s, Obj: %v, Kind: %s\n", setMapIndexStr, End, step, stepObj, stepReflectVal.Kind().String())
		} else {
			err = unknownError
		}
		// stepObj of kind Ptr, should have a field `step` with the right value set
		stepReflectVal, err = setPtrField(&stepReflectVal, step, data, nextStepIsArr, lastStep, idx)
		if stepReflectVal.IsValid() {
			stepObj = stepReflectVal.Interface()
			fmt.Printf("%s %s: Step= %s, Obj: %v, Kind: %s\n", forLoop, End, step, stepObj, stepReflectVal.Kind().String())
		} else {
			err = unknownError
		}
		//prevStepObj = stepObj
	}

	return
}

func setMapIndex(refValPtr *reflect.Value, index string, data interface{}, dataIsSliceType, isLastStep bool, sliceIdx int) (refVal reflect.Value, err error) {
	var obj interface{}
	refVal = *refValPtr
	if refVal.IsValid() {
		obj = refVal.Interface()
		fmt.Printf("%s %s: Index= %s, Obj: %v, Kind: %s\n", setMapIndexStr, Begin, index, obj, refVal.Kind().String())
	}
	var ok bool
	if obj, ok = obj.(map[string]interface{}); !ok {
		return
	}
	refVal = reflect.ValueOf(obj)
	if refVal.Kind() != reflect.Map {
		return
	}
	if isLastStep {
		if dataIsSliceType {
			sliceVal := reflect.MakeSlice(reflect.TypeOf([]interface{}{}), sliceIdx+1, sliceIdx+10)
			sliceVal.Index(sliceIdx).Set(reflect.ValueOf(data))
			refVal.SetMapIndex(reflect.ValueOf(index), sliceVal)
			obj = refVal.Interface()
		} else {
			if refVal.IsZero() {
				refVal = reflect.MakeMap(reflect.MapOf(reflect.TypeOf(index), reflect.TypeOf(data)))
			}
			refVal.SetMapIndex(reflect.ValueOf(index), reflect.ValueOf(data))
		}
	} else {
		if dataIsSliceType {

		} else {
			refVal.SetMapIndex(reflect.ValueOf(index), reflect.MakeMap(reflect.TypeOf(map[string]interface{}{})))
			obj = refVal.Interface()
			obj = obj.(map[string]interface{})[index]
			refVal = reflect.ValueOf(obj)
		}
	}
	if refVal.IsValid() {
		obj = (refVal).Interface()
		fmt.Printf("%s %s: FieldName= %s, Obj: %v, Kind: %s\n", setMapIndexStr, End, index, obj, refVal.Kind().String())
	}
	return
}

func setPtrField(refValPtr *reflect.Value, fieldName string, data interface{}, dataIsSliceType, isLastStep bool, sliceIdx int) (fieldVal reflect.Value, err error) {
	var obj interface{}
	refVal := *refValPtr
	if refVal.Kind() != reflect.Ptr {
		fieldVal = refVal
		return
	}
	if refVal.IsValid() {
		obj = refVal.Interface()
		fmt.Printf("%s %s: FieldName= %s, Obj: %v, Kind: %s\n", setPtrFieldStr, Begin, fieldName, obj, refVal.Kind().String())
	}
	refValElem := refVal.Elem()

	fieldVal = refValElem.FieldByName(fieldName)
	if !fieldVal.IsValid() {
		err = invalidPathError
		return
	}
	if fieldVal.IsValid() {
		obj = refVal.Interface()
		fmt.Printf("%s %s: FieldName= %s, Obj: %v, Kind: %s\n", setMapIndexStr, AfterStepChange, fieldName, obj, refVal.Kind().String())
	}
	if isLastStep {
		if fieldVal.CanSet() {
			if dataIsSliceType {
				fieldVal = setSliceValueAtIndex(fieldVal, sliceIdx, data)
			} else {
				fieldVal.Set(reflect.ValueOf(data))
			}
		} else {
			err = cantSetError
			return
		}
	} else {
		if fieldVal.CanSet() {
			if dataIsSliceType {
				fieldVal = setSliceValueAtIndex(fieldVal, sliceIdx, make(map[string]interface{}))
				fieldVal = fieldVal.Index(sliceIdx)
			} else {
				fieldVal.Set(reflect.ValueOf(map[string]interface{}{}))
			}
		}
	}
	if fieldVal.IsValid() {
		obj = fieldVal.Interface()
		fmt.Printf("%s %s: FieldName= %s, Obj: %v, Kind: %s\n", setMapIndexStr, End, fieldName, obj, fieldVal.Kind().String())
	}
	return
}

func setSliceValueAtIndex(sliceVal reflect.Value, sliceIdx int, data interface{}) reflect.Value {
	if sliceIdx >= sliceVal.Cap() {
		nCap := 2 * sliceIdx
		if nCap < 4 {
			nCap = 4
		}
		nF := reflect.MakeSlice(sliceVal.Type(), sliceIdx, nCap)
		reflect.Copy(nF, sliceVal)
		sliceVal.Set(nF)
		sliceVal.SetLen(sliceIdx + 1)
	}
	sliceVal.Index(sliceIdx).Set(reflect.ValueOf(data))
	return sliceVal
}
