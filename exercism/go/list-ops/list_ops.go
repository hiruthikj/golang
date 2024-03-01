package listops

type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	acc := initial

	for _, item := range s {
		acc = fn(acc, item)
	}
	return acc
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	acc := initial

	for i := s.Length() - 1; i >= 0; i-- {
		acc = fn(s[i], acc)
	}
	return acc
}

func (s IntList) Filter(fn func(int) bool) IntList {
	filteredList := make(IntList, 0, s.Length())
	i := 0

	for _, item := range s {
		if fn(item) {
			filteredList = append(filteredList, item)
			i++
		}
	}

	return filteredList
}

func (s IntList) Length() int {
	count := 0
	for range s {
		count++
	}
	return count
}

func (s IntList) Map(fn func(int) int) IntList {
	mappedList := make(IntList, 0, s.Length())
	i := 0

	for _, item := range s {
		mappedList = append(mappedList, fn(item))
		i++
	}

	return mappedList
}

func (s IntList) Reverse() IntList {
	reversedList := make(IntList, 0, s.Length())

	for i := s.Length() - 1; i >= 0; i-- {
		reversedList = append(reversedList, s[i])
	}

	return reversedList
}

func (s IntList) Append(lst IntList) IntList {
	// appendedList := make(IntList, 0, s.Length()+lst.Length())

	// for i, item := range s {
	// 	appendedList[i] = item
	// }

	// for i, item := range lst {
	// 	appendedList[s.Length()+i] = item
	// }

	// return appendedList
	return append(s, lst...)
}

func (s IntList) Concat(lists []IntList) IntList {
	concatenatedList := make(IntList, 0)
	concatenatedList = append(concatenatedList, s...)

	for _, ls := range lists {
		concatenatedList = concatenatedList.Append(ls)
	}

	return concatenatedList
}
