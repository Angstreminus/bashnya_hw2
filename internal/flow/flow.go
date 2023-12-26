package flow

import (
	"strconv"
	"strings"
)

func GetUniqStrings(inputPrep []string) []string {
	uniqPrep := make(map[string]struct{}, 0)
	var resStr []string
	for _, elem := range inputPrep {
		if _, ok := uniqPrep[elem]; !ok {
			resStr = append(resStr, elem)
			uniqPrep[elem] = struct{}{}
		}
	}
	return resStr
}

func GetPrepCount(inputPrep []string) map[string]int {
	strCnt := make(map[string]int, 0)
	for _, elem := range inputPrep {
		strCnt[elem]++
	}
	return strCnt
}

func GetPrepInput(input, prep []string) map[string]string {
	mapPrepInput := make(map[string]string, 0)
	for idx, elem := range prep {
		if _, ok := mapPrepInput[elem]; !ok {
			mapPrepInput[elem] = input[idx]
		}
	}
	return mapPrepInput
}

func ParamF(input *[]string, num int) []string {
	var resStr []string
	for _, elem := range *input {
		words := strings.Split(elem, " ")
		if len(words) != 1 {
			elem = strings.Join(words[num:], " ")
		}
		resStr = append(resStr, elem)
	}
	return resStr
}

func ParamI(input *[]string) []string {
	var resStr []string
	for _, elem := range *input {
		resStr = append(resStr, strings.ToLower(elem))
	}
	return resStr
}

func ParamS(input *[]string, num int) []string {
	var resStr []string
	for _, elem := range *input {
		if len(elem) != 0 {
			elem = strings.TrimLeft(elem, elem[0:num])
		} else {
			elem = ""
		}
		resStr = append(resStr, elem)
	}
	return resStr
}

func ParamC(prepUniqStr []string, prepCnt map[string]int, prepInput map[string]string) []string {
	var resStr []string
	for _, elem := range prepUniqStr {
		if _, ok := prepInput[elem]; !ok {
			continue
		}
		var sb strings.Builder
		sb.WriteString(strconv.Itoa(prepCnt[elem]))
		sb.WriteString(" ")
		sb.WriteString(prepInput[elem])
		resStr = append(resStr, sb.String())
	}
	return resStr
}

func ParamD(prepUniqStr []string, prepCnt map[string]int, prepInput map[string]string) []string {
	var resStr []string
	for _, elem := range prepUniqStr {
		if _, ok := prepInput[elem]; ok && (prepCnt[elem] > 1) {
			resStr = append(resStr, prepInput[elem])
		}
	}
	return resStr
}

func ParamU(prepUniqStr []string, prepCnt map[string]int, prepInput map[string]string) []string {
	var resStr []string
	for _, elem := range prepUniqStr {
		if _, ok := prepInput[elem]; ok && prepCnt[elem] == 1 {
			resStr = append(resStr, prepInput[elem])
		}
	}
	return resStr
}

func Unify(input, prep []string) []string {
	var resStr []string
	uniqStrings := make(map[string]string, 0)
	for idx, elem := range prep {
		if _, ok := uniqStrings[elem]; !ok {
			resStr = append(resStr, input[idx])
			uniqStrings[elem] = input[idx]
		}
	}
	return resStr
}

func UnifyFSI(f, s int, i bool, input []string) []string {
	prepStr := make([]string, len(input))
	copy(prepStr, input)
	if i {
		prepStr = ParamI(&prepStr)
	}
	if s != 0 {
		prepStr = ParamS(&prepStr, s)
	}
	if f != 0 {
		prepStr = ParamF(&prepStr, f)
	}
	return prepStr
}

func UnifyCUD(c, u, d bool, prepStr, input []string) []string {
	uniqPrepString := GetUniqStrings(prepStr)
	stringToCount := GetPrepCount(prepStr)
	prepToInput := GetPrepInput(input, prepStr)
	if c {
		return ParamC(uniqPrepString, stringToCount, prepToInput)
	}
	if d {
		return ParamD(uniqPrepString, stringToCount, prepToInput)
	}
	if u {
		return ParamU(uniqPrepString, stringToCount, prepToInput)
	}
	return prepStr
}
