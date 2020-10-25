package _25__长按键入

func isLongPressedName(name string, typed string) bool {
	i := 0
	j := 0
	for j < len(typed) {
		switch {
		case i < len(name) && name[i] == typed[j]:
			// name到最后一个字符时，name[len(name)]会溢出，需要跳过
			i++
			j++
		case j > 0 && name[i-1] == typed[j]:
			// 可能首个字母就不相同，此时i-1会溢出，需要跳过
			j++
		default:
			return false
		}
	}
	return i == len(name)
}
