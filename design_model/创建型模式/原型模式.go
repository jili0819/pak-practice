package main

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
)

// 原型是一种创建型设计模式，使你能够复制对象，甚至是复杂对象，而又无需使代码依赖它们所属的类。
// 所有的原型类都必须有一个通用的接口， 使得即使在对象所属的具体类未知的情况下也能复制对象。
// 原型对象可以生成自身的完整副本， 因为相同类的对象可以相互访问对方的私有成员变量。

// 纸质文件可以通过复印机轻松拷贝出多份，设置Paper接口，包含读取文件内容和克隆文件两个方法。
// 同时声明两个类报纸（Newspaper）和简历（Resume）实现了Paper接口，通过复印机（Copier）复印出两类文件的副本，并读取文件副本内容。

// Paper 纸张，包含读取内容的方法，拷贝纸张的方法，作为原型模式接口
type Paper interface {
	io.Reader
	Clone() Paper
}

// Newspaper 报纸 实现原型接口
type Newspaper struct {
	headline string
	content  string
}

func NewNewspaper(headline string, content string) *Newspaper {
	return &Newspaper{
		headline: headline,
		content:  content,
	}
}

func (np *Newspaper) Read(p []byte) (n int, err error) {
	buf := bytes.NewBufferString(fmt.Sprintf("headline:%s,content:%s", np.headline, np.content))
	return buf.Read(p)
}

func (np *Newspaper) Clone() Paper {
	return &Newspaper{
		headline: np.headline + "_copied",
		content:  np.content,
	}
}

// Resume 简历 实现原型接口
type Resume struct {
	name       string
	age        int
	experience string
}

func NewResume(name string, age int, experience string) *Resume {
	return &Resume{
		name:       name,
		age:        age,
		experience: experience,
	}
}

func (r *Resume) Read(p []byte) (n int, err error) {
	buf := bytes.NewBufferString(fmt.Sprintf("name:%s,age:%d,experience:%s", r.name, r.age, r.experience))
	return buf.Read(p)
}

func (r *Resume) Clone() Paper {
	return &Resume{
		name:       r.name + "_copied",
		age:        r.age,
		experience: r.experience,
	}
}

// Copier 复印机
type Copier struct {
	name string
}

func NewCopier(name string) *Copier {
	return &Copier{name: name}
}

func (c *Copier) copy(paper Paper) Paper {
	fmt.Printf("copier name:%v is copying:%v ", c.name, reflect.TypeOf(paper).String())
	fmt.Println()
	return paper.Clone()
}

func main() {
	copier := NewCopier("云打印机")
	oneNewspaper := NewNewspaper("Go是最好的编程语言", "Go语言十大优势")
	oneResume := NewResume("小明", 29, "5年码农")

	otherNewspaper := copier.copy(oneNewspaper)
	copyNewspaperMsg := make([]byte, 100)
	byteSize, _ := otherNewspaper.Read(copyNewspaperMsg)
	fmt.Println("copyNewspaperMsg:" + string(copyNewspaperMsg[:byteSize]))

	otherResume := copier.copy(oneResume)
	copyResumeMsg := make([]byte, 100)
	byteSize, _ = otherResume.Read(copyResumeMsg)
	fmt.Println("copyResumeMsg:" + string(copyResumeMsg[:byteSize]))
}
