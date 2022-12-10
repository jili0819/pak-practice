package main

import "fmt"

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
