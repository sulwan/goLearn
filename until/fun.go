package until

import "errors"

func Add(a int, b int) (ret int, err error) {
	if a < 0 || b < 0 {
		err = errors.New("Should be nono-megative numbers!")
		return
	}
	return a + b, nil
}
