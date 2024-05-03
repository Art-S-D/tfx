package tfstate

import "fmt"

func resourceIndexToStr(index any) string {
	switch index.(type) {
	case int:
		return fmt.Sprintf("[%d]", index)
	case string:
		return fmt.Sprintf("[\"%s\"]", index)
	}
	return ""
}
