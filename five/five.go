package five

import "errors"

//some comments
//need to export
func Five(num int) (bool, error) {
	if num < 5 {
		return true, nil
	}
	return false, errors.New("number more than 5")
}
