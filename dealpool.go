package gogetset

import (
	"errors"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	dotDelim string = "."
	End      string = "End"
	Begin    string = "Begin"
)

var (
	invalidPathError       = errors.New("invalid path")
	sliceExpectedNotFound  = errors.New("slice expected, not found")
	structNotAcceptedError = errors.New("struct not accepted")
	nilNotAccepted         = errors.New("nil not accepted, pass an empty map instead")
	unknownError           = errors.New("unknown error")
	cantSetError           = errors.New("cant set field for struct")
)

func stepsFromPath(path string) []string {
	return strings.Split(path, dotDelim)
}

func isLastStep(stepIndex int, steps []string) bool {
	return stepIndex == len(steps)-1
}

func nextStepIsArray(step string) (arrayKey string, idx int, isArray bool, err error) {
	if !strings.Contains(step, "[") {
		return "", -1, false, nil
	}
	alphaReg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		log.Fatal(err)
	}
	arrayKey = alphaReg.ReplaceAllString(step, "")
	numReg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	idx, err = strconv.Atoi(numReg.ReplaceAllString(step, ""))
	if err != nil {
		return "", -1, false, err
	}
	return arrayKey, idx, true, nil
}
