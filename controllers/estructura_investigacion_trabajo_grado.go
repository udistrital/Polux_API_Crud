package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/udistrital/Polux_API_Crud/models"

	"github.com/astaxie/beego"
)

// EstructuraInvestigacionTrabajoGradoController operations for EstructuraInvestigacionTrabajoGrado
type EstructuraInvestigacionTrabajoGradoController struct {
	beego.Controller
}

// URLMapping ...
func (c *EstructuraInvestigacionTrabajoGradoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create EstructuraInvestigacionTrabajoGrado
// @Param	body		body 	models.EstructuraInvestigacionTrabajoGrado	true		"body for EstructuraInvestigacionTrabajoGrado content"
// @Success 201 {int} models.EstructuraInvestigacionTrabajoGrado
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *EstructuraInvestigacionTrabajoGradoController) Post() {
	var v models.EstructuraInvestigacionTrabajoGrado
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddEstructuraInvestigacionTrabajoGrado(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			beego.Error(err)
			c.Abort("400")
		}
	} else {
			beego.Error(err)
			c.Abort("400")
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get EstructuraInvestigacionTrabajoGrado by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.EstructuraInvestigacionTrabajoGrado
// @Failure 404 not found resource
// @router /:id [get]
func (c *EstructuraInvestigacionTrabajoGradoController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetEstructuraInvestigacionTrabajoGradoById(id)
	if err != nil {
		beego.Error(err)
		c.Abort("404")
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get EstructuraInvestigacionTrabajoGrado
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.EstructuraInvestigacionTrabajoGrado
// @Failure 404 not found resource
// @router / [get]
func (c *EstructuraInvestigacionTrabajoGradoController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllEstructuraInvestigacionTrabajoGrado(query, fields, sortby, order, offset, limit)
	if err != nil {
		beego.Error(err)
		c.Abort("404")
	} else {
		if l == nil {
			l = append(l, map[string]interface{}{})
		}
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the EstructuraInvestigacionTrabajoGrado
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.EstructuraInvestigacionTrabajoGrado	true		"body for EstructuraInvestigacionTrabajoGrado content"
// @Success 200 {object} models.EstructuraInvestigacionTrabajoGrado
// @Failure 400 the request contains incorrect syntax
// @router /:id [put]
func (c *EstructuraInvestigacionTrabajoGradoController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.EstructuraInvestigacionTrabajoGrado{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateEstructuraInvestigacionTrabajoGradoById(&v); err == nil {
			c.Data["json"] = v
		} else {
			beego.Error(err)
			c.Abort("400")
		}
	} else {
			beego.Error(err)
			c.Abort("400")
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the EstructuraInvestigacionTrabajoGrado
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 404 not found resource
// @router /:id [delete]
func (c *EstructuraInvestigacionTrabajoGradoController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteEstructuraInvestigacionTrabajoGrado(id); err == nil {
		c.Data["json"] = map[string]interface{}{"Id": id}
	} else {
		beego.Error(err)
		c.Abort("404")
	}
	c.ServeJSON()
}
