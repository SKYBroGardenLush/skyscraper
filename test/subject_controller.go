package test

import (
	"fmt"
	"github.com/SKYBroGardenLush/skyscraper/framework"
	"skyscraper/provider/demo"
)

func SubjectAddController(c *framework.Context) error {
	c.SetOKStatus().Json("ok, SubjectAddController")
	return nil
}

// 对应路由 /subject/list/all
func SubjectListController(c *framework.Context) error {
	// 获取demo服务实例
	demoService := c.MustMake(demo.Key).(demo.Service)

	// 调用服务实例的方法
	foo := demoService.GetFoo()

	// 输出结果
	c.SetOKStatus().Json(foo)
	return nil
}

func SubjectDelController(c *framework.Context) error {
	c.SetOKStatus().Json("ok, SubjectDelController")
	return nil
}

func SubjectUpdateController(c *framework.Context) error {
	c.SetOKStatus().Json("ok, SubjectUpdateController")
	return nil
}

func SubjectGetController(c *framework.Context) error {
	subjectId, _ := c.ParamInt("id", 0)
	c.SetOKStatus().Json("ok, SubjectGetController:" + fmt.Sprint(subjectId))
	return nil

}

func SubjectNameController(c *framework.Context) error {
	c.SetOKStatus().Json("ok, SubjectNameController")
	return nil
}
