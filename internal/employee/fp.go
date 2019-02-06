// Code generated by 'gofp'. DO NOT EDIT.
package employee 
import "sync" 

func Map(f func(Employee) Employee, list []Employee) []Employee {
	if f == nil {
		return []Employee{}
	}
	newList := make([]Employee, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}
func Filter(f func(Employee) bool, list []Employee) []Employee {
	if f == nil {
		return []Employee{}
	}
	var newList []Employee
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}
func Remove(f func(Employee) bool, list []Employee) []Employee {
	if f == nil {
		return []Employee{}
	}
	var newList []Employee
	for _, v := range list {
		if !f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}
func Some(f func(Employee) bool, list []Employee) bool {
	if f == nil {
		return false
	}
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}
func Every(f func(Employee) bool, list []Employee) bool {
	if f == nil || len(list) == 0 {
		return false
	}
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}
func DropWhile(f func(Employee) bool, list []Employee) []Employee {
	if f == nil {
		return []Employee{}
	}
	var newList []Employee
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]Employee, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}
func TakeWhile(f func(Employee) bool, list []Employee) []Employee {
	if f == nil {
		return []Employee{}
	}
	var newList []Employee
	for _, v := range list {
		if !f(v) {
			return newList
		}
		newList = append(newList, v)
	}
	return newList
}
func Pmap(f func(Employee) Employee, list []Employee) []Employee {
	if f == nil {
		return []Employee{}
	}

	ch := make(chan map[int]Employee)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]Employee, i int, v Employee) {
			defer wg.Done()
			ch <- map[int]Employee{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]Employee, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}
func FilterMap(fFilter func(Employee) bool, fMap func(Employee) Employee, list []Employee) []Employee {
	if fFilter == nil || fMap == nil {
		return []Employee{}
	}
	var newList []Employee
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}
func Rest(l []Employee) []Employee {
	if l == nil {
		return []Employee{}
	}

	len := len(l)
	if len == 0 || len == 1 {
		return []Employee{}
	}

	newList := make([]Employee, len-1)

	for i, v := range l[1:] {
		newList[i] = v
	}

	return newList
}
func Reduce(f func(Employee, Employee) Employee, list []Employee, initializer ...Employee) Employee {
	var init Employee 
	lenList := len(list)

	if initializer != nil {
		init = initializer[0]
	} else if lenList > 0 {
		init = list[0]
		if lenList == 1 {
			return list[0]
		}
		if lenList >= 2 {
			list = list[1:]
		}
	}
	
	if lenList == 0 {
		return init
	}
	r := f(init, list[0])
	return Reduce(f, list[1:], r)
}
func MapTeacher(f func(Teacher) Teacher, list []Teacher) []Teacher {
	if f == nil {
		return []Teacher{}
	}
	newList := make([]Teacher, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}
func FilterTeacher(f func(Teacher) bool, list []Teacher) []Teacher {
	if f == nil {
		return []Teacher{}
	}
	var newList []Teacher
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}
func RemoveTeacher(f func(Teacher) bool, list []Teacher) []Teacher {
	if f == nil {
		return []Teacher{}
	}
	var newList []Teacher
	for _, v := range list {
		if !f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}
func SomeTeacher(f func(Teacher) bool, list []Teacher) bool {
	if f == nil {
		return false
	}
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}
func EveryTeacher(f func(Teacher) bool, list []Teacher) bool {
	if f == nil || len(list) == 0 {
		return false
	}
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}
func DropWhileTeacher(f func(Teacher) bool, list []Teacher) []Teacher {
	if f == nil {
		return []Teacher{}
	}
	var newList []Teacher
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]Teacher, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}
func TakeWhileTeacher(f func(Teacher) bool, list []Teacher) []Teacher {
	if f == nil {
		return []Teacher{}
	}
	var newList []Teacher
	for _, v := range list {
		if !f(v) {
			return newList
		}
		newList = append(newList, v)
	}
	return newList
}
func PmapTeacher(f func(Teacher) Teacher, list []Teacher) []Teacher {
	if f == nil {
		return []Teacher{}
	}

	ch := make(chan map[int]Teacher)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]Teacher, i int, v Teacher) {
			defer wg.Done()
			ch <- map[int]Teacher{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]Teacher, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}
func FilterMapTeacher(fFilter func(Teacher) bool, fMap func(Teacher) Teacher, list []Teacher) []Teacher {
	if fFilter == nil || fMap == nil {
		return []Teacher{}
	}
	var newList []Teacher
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}
func RestTeacher(l []Teacher) []Teacher {
	if l == nil {
		return []Teacher{}
	}

	len := len(l)
	if len == 0 || len == 1 {
		return []Teacher{}
	}

	newList := make([]Teacher, len-1)

	for i, v := range l[1:] {
		newList[i] = v
	}

	return newList
}
func ReduceTeacher(f func(Teacher, Teacher) Teacher, list []Teacher, initializer ...Teacher) Teacher {
	var init Teacher 
	lenList := len(list)

	if initializer != nil {
		init = initializer[0]
	} else if lenList > 0 {
		init = list[0]
		if lenList == 1 {
			return list[0]
		}
		if lenList >= 2 {
			list = list[1:]
		}
	}
	
	if lenList == 0 {
		return init
	}
	r := f(init, list[0])
	return ReduceTeacher(f, list[1:], r)
}