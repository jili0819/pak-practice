package main

import "fmt"

// 迭代器模式
// 迭代器模式是一种行为设计模式，让你能在不暴露集合底层表现形式 （列表、 栈和树等）的情况下遍历集合中所有的元素。
// 在迭代器的帮助下， 客户端可以用一个迭代器接口以相似的方式遍历不同集合中的元素。
// 这里需要注意的是有两个典型的迭代器接口需要分清楚；
// 一个是集合类实现的可以创建迭代器的工厂方法接口一般命名为Iterable，包含的方法类似CreateIterator；
// 另一个是迭代器本身的接口，命名为Iterator，有Next及hasMore两个主要方法；

// 一个班级类中包括一个老师和若干个学生，我们要对班级所有成员进行遍历，班级中老师存储在单独的结构字段中，学生存储在另外一个slice字段中.
// 通过迭代器，我们实现统一遍历处理；
type (
	// Member 成员接口
	Member interface {
		Desc() string // 输出成员描述信息
	}

	// Teacher 老师
	Teacher struct {
		name    string // 名称
		subject string // 所教课程
	}

	// Student 学生
	Student struct {
		name     string // 姓名
		sumScore int    // 考试总分数
	}
)

// NewTeacher 根据姓名、课程创建老师对象
func NewTeacher(name, subject string) *Teacher {
	return &Teacher{
		name:    name,
		subject: subject,
	}
}

func (t *Teacher) Desc() string {
	return fmt.Sprintf("%s班主任老师负责教%s", t.name, t.subject)
}

// NewStudent 创建学生对象
func NewStudent(name string, sumScore int) *Student {
	return &Student{
		name:     name,
		sumScore: sumScore,
	}
}

func (t *Student) Desc() string {
	return fmt.Sprintf("%s同学考试总分为%d", t.name, t.sumScore)
}

type (
	CommonIteratorImp interface {
		getIterator() CommonIterator
	}
	CommonIterator interface {
		hasMore() bool
		next() interface{}
	}

	Int64Slice struct {
		data []int64
	}

	// Int64SliceIterator int64slice迭代器
	Int64SliceIterator struct {
		index int64
		data  []int64
	}
)

func (s *Int64Slice) getIterator() CommonIterator {
	return &Int64SliceIterator{data: s.data}
}

func (s *Int64SliceIterator) hasMore() bool {
	return s.index < int64(len(s.data))
}

func (s *Int64SliceIterator) next() interface{} {
	next := s.data[s.index]
	s.index++
	return next
}

type (
	// Iterator 迭代器接口
	Iterator interface {
		Next() Member  // 迭代下一个成员
		HasMore() bool // 是否还有
	}

	// memberIterator 班级成员迭代器实现
	memberIterator struct {
		class *Class // 需迭代的班级
		index int    // 迭代索引
	}

	// Class 班级，包括老师和同学
	Class struct {
		name     string
		teacher  *Teacher
		students []*Student
	}
)

func (m *memberIterator) Next() Member {
	// 迭代索引为-1时，返回老师成员，否则遍历学生slice
	if m.index == -1 {
		m.index++
		return m.class.teacher
	}
	student := m.class.students[m.index]
	m.index++
	return student
}

func (m *memberIterator) HasMore() bool {
	return m.index < len(m.class.students)
}

// NewClass 根据班主任老师名称，授课创建班级
func NewClass(name, teacherName, teacherSubject string) *Class {
	return &Class{
		name:    name,
		teacher: NewTeacher(teacherName, teacherSubject),
	}
}

// CreateIterator 创建班级迭代器
func (c *Class) CreateIterator() Iterator {
	return &memberIterator{
		class: c,
		index: -1, // 迭代索引初始化为-1，从老师开始迭代
	}
}

func (c *Class) Name() string {
	return c.name
}

// AddStudent 班级添加同学
func (c *Class) AddStudent(students []*Student) {
	c.students = append(c.students, students...)
}

func main() {
	info := Int64Slice{data: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}
	iterator := info.getIterator()
	for iterator.hasMore() {
		fmt.Println(iterator.next().(int64))
	}

	class := NewClass("三年级一班", "王明", "数学课")
	a := []*Student{
		{"张三", 389},
		{"李四", 378},
		{"王五", 347},
	}
	class.AddStudent(a)

	fmt.Printf("%s成员如下:\n", class.Name())
	classIterator := class.CreateIterator()
	for classIterator.HasMore() {
		member := classIterator.Next()
		fmt.Println(member.Desc())
	}
}
