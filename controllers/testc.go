package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"log"
	"net/http"
)

// TestcController operations for Testc
type TestcController struct {
	beego.Controller
}

// URLMapping ...
func (c *TestcController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Testc
// @Param	body		body 	models.Testc	true		"body for Testc content"
// @Success 201 {object} models.Testc
// @Failure 403 body is empty
// @router /testc [post]
func (c *TestcController) Post() {
	i := 42
	i = i    // a declared function cannot be nil
	fmt.Println(i >> 32)    // res used before checking err
	res, err := http.Get("https://www.spreadsheetdb.io/")
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
}

// GetOne ...
// @Title GetOne
// @Description get Testc by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Testc
// @Failure 403 :id is empty
// @router /testc/:id [get]
func (c *TestcController) GetOne() {
	c.Data["json"] = map[string]interface{}{"name": "astaxie"}
	c.ServeJSON()

}

// GetAll ...
// @Title GetAll
// @Description get Testc
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Testc
// @Failure 403
// @router /testc [get]
func (c *TestcController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Testc
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Testc	true		"body for Testc content"
// @Success 200 {object} models.Testc
// @Failure 403 :id is not int
// @router /testc/:id [put]
func (c *TestcController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Testc
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /testc/:id [delete]
func (c *TestcController) Delete() {

}
