import "strconv"

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */

func deserialize(s string) *NestedInteger {
	var res NestedInteger

	if string(s[0]) != "[" {
		num, _ := strconv.Atoi(s)
		res.SetInteger(num)
		return &res
	}
	var niList []NestedInteger
	var numStr string
	var level int
	for i := 0; i < len(s); i++ {
		switch string(s[i]) {
		case "[":
			numStr = ""
			niList = append(niList, NestedInteger{})
			level++
		case ",":
			if numStr != "" {
				num, _ := strconv.Atoi(numStr)
				var n NestedInteger
				n.SetInteger(num)
				niList[level-1].Add(n)
			}
			numStr = ""
		case "]":
			if numStr != "" {
				num, _ := strconv.Atoi(numStr)
				var n NestedInteger
				n.SetInteger(num)
				niList[level-1].Add(n)
			}
			level--
			if level >= 1 {
				niList[level-1].Add(niList[level])
				niList[level] = NestedInteger{}
			}
			numStr = ""
		default:
			numStr += string(s[i])
		}
	}
	return &niList[0]
}