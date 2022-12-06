package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jili/pkg-practice/api"
	"strconv"
)

type TypeName struct {
	Bs []string `json:"bs"`
}
type A struct {
	Name     string   `json:"name" xml:"name"`
	Password string   `json:"password" xml:"password"`
	T        TypeName `json:"t" xml:"t"`
}

func main() {
	g := gin.Default()
	g.GET("/json", func(c *gin.Context) {
		var ss []A
		for i := 0; i < 10; i++ {
			ss = append(ss, A{
				Name:     "name_" + strconv.Itoa(i),
				Password: "pass_" + strconv.Itoa(i),
				T: TypeName{
					Bs: []string{"1", "fw"},
				},
			})
		}
		api.WriteSuccessResponseJSON(c.Writer, ss)
	})
	g.GET("/xml", func(c *gin.Context) {
		api.WriteSuccessResponseXML(c.Writer, A{
			Name:     "fw",
			Password: "fw",
			T: TypeName{
				Bs: []string{"1", "fw"},
			},
		})
	})
	g.Run()
}
