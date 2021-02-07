package util

// StringSet represents a set of string
type StringSet map[string]struct{}

// NewStringSet initializes a new StringSet from args
func NewStringSet(args ...string) StringSet {
	return NewStringSetFromSlice(args)
}

// NewStringSetFromSlice initializes new StringSet
func NewStringSetFromSlice(slice []string) StringSet {
	ss := make(StringSet)

	for _, s := range slice {
		ss[s] = struct{}{}
	}

	return ss
}

// Contains returns if s is contained in the StringSet
func (ss StringSet) Contains(s string) bool {
	_, ok := ss[s]

	return ok
}

// ContainsSet returns if all entries in s are contained in the StringSet
func (ss StringSet) ContainsSet(ss2 StringSet) bool {
	for k := range ss2 {
		if !ss.Contains(k) {
			return false
		}
	}

	return true
}

// Insert adds s as an entry of the StringSet
func (ss StringSet) Insert(s string) {
	ss[s] = struct{}{}
}

// Delete deletes s from the StringSet
func (ss StringSet) Delete(s string) {
	delete(ss, s)
}

// Slice returns the StringSet as a slice of string
func (ss StringSet) Slice() []string {
	res := make([]string, 0, len(ss))

	for k := range ss {
		res = append(res, k)
	}

	return res
}
