package set

type SetUint32 struct {
	nodeMap map[uint32]bool
}

// Create set object
func NewUint32(nums []uint32) *SetUint32 {
	s := &SetUint32{}
	for _, num := range nums {
		s.Add(num)
	}
	return s
}

// Add an item
func (s *SetUint32) Add(num uint32) *SetUint32 {
	if s.nodeMap == nil {
		s.nodeMap = make(map[uint32]bool)
	}
	_, ok := s.nodeMap[num]
	if !ok {
		s.nodeMap[num] = true
	}
	return s
}

// Make object empty
func (s *SetUint32) Clear() {
	s.nodeMap = make(map[uint32]bool)
}

// Remove an item
func (s *SetUint32) Remove(num uint32) bool {
	_, ok := s.nodeMap[num]
	if ok {
		delete(s.nodeMap, num)
	}
	return ok
}

// Check if item exists in set
func (s *SetUint32) Contains(num uint32) bool {
	_, ok := s.nodeMap[num]
	return ok
}

// Get set items
func (s *SetUint32) GetList() []uint32 {
	nums := []uint32{}
	for i := range s.nodeMap {
		nums = append(nums, i)
	}
	return nums
}

// Get size of set
func (s *SetUint32) Size() int {
	return len(s.nodeMap)
}

// Returns all the items that are in S or in S2
func (s *SetUint32) Union(s2 *SetUint32) *SetUint32 {
	s3 := SetUint32{}
	s3.nodeMap = make(map[uint32]bool)
	for i := range s.nodeMap {
		s3.nodeMap[i] = true
	}
	for i := range s2.nodeMap {
		_, ok := s3.nodeMap[i]
		if !ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

// Common items in S and S2
func (s *SetUint32) Intersection(s2 *SetUint32) *SetUint32 {
	s3 := SetUint32{}
	s3.nodeMap = make(map[uint32]bool)
	for i := range s2.nodeMap {
		_, ok := s.nodeMap[i]
		if ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

// s.Minus(s2) : all of S but not in S2
func (s *SetUint32) Minus(s2 *SetUint32) *SetUint32 {
	s3 := SetUint32{}
	s3.nodeMap = make(map[uint32]bool)
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			s3.nodeMap[i] = true
		}
	}
	return &s3
}

// Checks if S is subset of S2
func (s *SetUint32) Subset(s2 *SetUint32) bool {
	for i := range s.nodeMap {
		_, ok := s2.nodeMap[i]
		if !ok {
			return false
		}
	}
	return true
}

// Checks if S is superset of S2
func (s *SetUint32) Superset(s2 *SetUint32) bool {
	for i := range s2.nodeMap {
		_, ok := s.nodeMap[i]
		if !ok {
			return false
		}
	}
	size1 := s.Size()
	size2 := s2.Size()
	return size1 == size2 || size1 > size2
}
