package controllers

import (
	"app/models"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

//  WatchingStatusController operations for WatchingStatus
type WatchingStatusController struct {
	beego.Controller
}

// URLMapping ...
func (c *WatchingStatusController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create WatchingStatus
// @Param	body		body 	models.WatchingStatus	true		"body for WatchingStatus content"
// @Success 201 {int} models.WatchingStatus
// @Failure 403 body is empty
// @router / [post]
func (c *WatchingStatusController) Post() {
	var v models.WatchingStatus
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if _, err := models.AddWatchingStatus(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = v
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
	w, _ := models.GetTvProgramById(v.TvProgramId)
	if v.Watched {
		w.CountWatched++
	} else {
		w.CountWantToWatch++
	}
	if err := models.UpdateTvProgramById(w); err != nil {
		fmt.Println(err.Error())
	}
}

// GetOne ...
// @Title Get One
// @Description get WatchingStatus by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.WatchingStatus
// @Failure 403 :id is empty
// @router /:id [get]
func (c *WatchingStatusController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetWatchingStatusById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get WatchingStatus
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.WatchingStatus
// @Failure 403
// @router / [get]
func (c *WatchingStatusController) GetAll() {
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

	l, err := models.GetAllWatchingStatus(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the WatchingStatus
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.WatchingStatus	true		"body for WatchingStatus content"
// @Success 200 {object} models.WatchingStatus
// @Failure 403 :id is not int
// @router /:id [put]
func (c *WatchingStatusController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	u, _ := models.GetWatchingStatusById(id)
	v := models.WatchingStatus{Id: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateWatchingStatusById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
	// fmt.Println(v.TvProgramId)
	w, _ := models.GetTvProgramById(v.TvProgramId)
	if v.Watched != u.Watched {
		if v.Watched {
			w.CountWatched++
		} else {
			w.CountWatched--
		}
	} else {
		if v.WantToWatch {
			w.CountWantToWatch++
		} else {
			w.CountWantToWatch--
		}
	}
	if err := models.UpdateTvProgramById(w); err != nil {
		fmt.Println(err.Error())
	}
}

// Delete ...
// @Title Delete
// @Description delete the WatchingStatus
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *WatchingStatusController) Delete() {
	// idStr := c.Ctx.Input.Param(":id")
	// id, _ := strconv.ParseInt(idStr, 0, 64)
	// if err := models.DeleteWatchingStatus(id); err == nil {
	// 	c.Data["json"] = "OK"
	// } else {
	// 	c.Data["json"] = err.Error()
	// }
	// c.ServeJSON()
}
