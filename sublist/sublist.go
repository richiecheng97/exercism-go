package sublist

// Relation type is defined in relations.go file.

func Sublist(list1, list2 []int) Relation {
	if listsEqual(list1, list2) {
		return RelationEqual
	}

	if isSublist(list1, list2) {
		return RelationSublist
	}

	if isSublist(list2, list1) {
		return RelationSuperlist
	}

	return RelationUnequal
}

// listsEqual 判断两个列表是否相等
func listsEqual(list1, list2 []int) bool {
	if len(list1) != len(list2) {
		return false
	}

	for i := 0; i < len(list1); i++ {
		if list1[i] != list2[i] {
			return false
		}
	}

	return true
}

// isSublist 判断list1是否是list2的子列表（list2包含list1的连续子序列）
func isSublist(list1, list2 []int) bool {
	// 空列表是任何列表的子列表
	if len(list1) == 0 {
		return true
	}

	// 如果list1比list2长，不可能是子列表
	if len(list1) > len(list2) {
		return false
	}

	// 检查list2中是否存在连续的子序列等于list1
	for i := 0; i <= len(list2)-len(list1); i++ {
		match := true
		for j := 0; j < len(list1); j++ {
			if list2[i+j] != list1[j] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}

	return false
}
