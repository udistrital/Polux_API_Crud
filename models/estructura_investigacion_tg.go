package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type EstructuraInvestigacionTg struct {
	Id                            int           `orm:"column(id);pk;auto"`
	IdTrabajoGrado                *TrabajoGrado `orm:"column(id_trabajo_grado);rel(fk)"`
	CodigoEstructuraInvestigacion float64       `orm:"column(codigo_estructura_investigacion)"`
}

func (t *EstructuraInvestigacionTg) TableName() string {
	return "estructura_investigacion_tg"
}

func init() {
	orm.RegisterModel(new(EstructuraInvestigacionTg))
}

// AddEstructuraInvestigacionTg insert a new EstructuraInvestigacionTg into database and returns
// last inserted Id on success.
func AddEstructuraInvestigacionTg(m *EstructuraInvestigacionTg) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetEstructuraInvestigacionTgById retrieves EstructuraInvestigacionTg by Id. Returns error if
// Id doesn't exist
func GetEstructuraInvestigacionTgById(id int) (v *EstructuraInvestigacionTg, err error) {
	o := orm.NewOrm()
	v = &EstructuraInvestigacionTg{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllEstructuraInvestigacionTg retrieves all EstructuraInvestigacionTg matches certain condition. Returns empty list if
// no records exist
func GetAllEstructuraInvestigacionTg(query map[string]string, fields []string, sortby []string, order []string, related []interface{},
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(EstructuraInvestigacionTg))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []EstructuraInvestigacionTg
	if len(related) > 0 {
		qs = qs.OrderBy(sortFields...).RelatedSel(related...)
	} else {
		qs = qs.OrderBy(sortFields...)
	}
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateEstructuraInvestigacionTg updates EstructuraInvestigacionTg by Id and returns error if
// the record to be updated doesn't exist
func UpdateEstructuraInvestigacionTgById(m *EstructuraInvestigacionTg) (err error) {
	o := orm.NewOrm()
	v := EstructuraInvestigacionTg{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteEstructuraInvestigacionTg deletes EstructuraInvestigacionTg by Id and returns error if
// the record to be deleted doesn't exist
func DeleteEstructuraInvestigacionTg(id int) (err error) {
	o := orm.NewOrm()
	v := EstructuraInvestigacionTg{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&EstructuraInvestigacionTg{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
