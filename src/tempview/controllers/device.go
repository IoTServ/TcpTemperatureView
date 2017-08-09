package controllers

import (
	"tempview/models"
	"strconv"

	"github.com/astaxie/beego"
	"fmt"
)

// DEVICEController operations for DEVICE
type DEVICEController struct {
	beego.Controller
}

// URLMapping ...
func (c *DEVICEController) URLMapping() {
	c.Mapping("Get", c.GetOne)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
}

// GetOne ...
// @Title Get One
// @Description get DEVICE by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.DEVICE
// @Failure 403 :id is empty
// @router /:id [get]
func (c *DEVICEController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetDEVICEById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	fmt.Printf("getone")
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get DEVICE
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.DEVICE
// @Failure 403
// @router / [get]
func (c *DEVICEController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64
	var offset int64

	l, err := models.GetAllDEVICE(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	fmt.Printf("getall")
	c.ServeJSON()
}