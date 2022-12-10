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
