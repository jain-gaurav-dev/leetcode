package recursion

type NestedInteger struct {
}

func (n NestedInteger) IsInteger() bool { return true }

func (n NestedInteger) GetInteger() int { return 0 }

func (n *NestedInteger) SetInteger(value int) {}

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation

// Return true if this NestedInteger holds a single integer, rather than a nested list.
// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check

// Set this NestedInteger to hold a single integer.

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (n *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The result is undefined if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (n NestedInteger) GetList() []*NestedInteger { return nil }

func depthSum(nestedList []*NestedInteger) int {
	var eval func(ni []*NestedInteger, depth int) int
	eval = func(ni []*NestedInteger, depth int) int {
		if len(ni) == 0 {
			return 0
		}

		if len(ni) == 1 && ni[0].IsInteger() {
			return ni[0].GetInteger() * depth
		}

		// ni is a list
		// descend with depth incremented.
		res := 0
		for i := 0; i < len(ni); i++ {
			if ni[i].IsInteger() {
				res += ni[i].GetInteger() * depth
			} else {
				res += eval(ni[i].GetList(), depth+1)
			}
		}

		return res
	}

	return eval(nestedList, 1)

}
