package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/udistrital/Polux_API_Crud/utilidades"
)

type FormatoEvaluacionCarrera struct {
	Id             int        `orm:"column(id);pk;auto"`
	IdFormato      *Formato   `orm:"column(id_formato);rel(fk)"`
	IdModalidad    *Modalidad `orm:"column(id_modalidad);rel(fk);null"`
	Activo         bool       `orm:"column(activo)"`
	CodigoProyecto float64    `orm:"column(codigo_proyecto)"`
	FechaInicio    time.Time  `orm:"column(fecha_inicio);type(date)"`
	FechaFin       time.Time  `orm:"column(fecha_fin);type(date);null"`
}

func (t *FormatoEvaluacionCarrera) TableName() string {
	return "formato_evaluacion_carrera"
}

func init() {
	orm.RegisterModel(new(FormatoEvaluacionCarrera))
}

//Transaccion formato_evaluacion_carrera
func TrFormatoEvaluacionCarrera(m map[string]interface{}) (aceptado string, err error) {
	fmt.Println("modelo")
	formato_evaluacion := []FormatoEvaluacionCarrera{}

	err = utilidades.FillStruct(m["formato_facultad"], &formato_evaluacion)
	fmt.Println("PPPPPPPPPPPPPPPPPPPPP")
	o := orm.NewOrm()
	o.Begin()
	if err == nil {
		fmt.Println("AAAAAAAA")
		for _, data := range formato_evaluacion {
			fmt.Println("CCCCCCCCCCCCCCCCCCC")
			fmt.Println("formato: ", data)
			_, err = o.Insert(&data)
		}
		fmt.Println("BBBBBBBBBBB")

		if err == nil {
			o.Commit()
			aceptado = "OK"
			return
		} else {

			fmt.Println(err.Error())
			o.Rollback()
			return
		}
	} else {
		fmt.Println("111111111111111111111")
		fmt.Println(err.Error())
		o.Rollback()
		return
	}
}

// AddFormatoEvaluacionCarrera insert a new FormatoEvaluacionCarrera into database and returns
// last inserted Id on success.
func AddFormatoEvaluacionCarrera(m *FormatoEvaluacionCarrera) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFormatoEvaluacionCarreraById retrieves FormatoEvaluacionCarrera by Id. Returns error if
// Id doesn't exist
func GetFormatoEvaluacionCarreraById(id int) (v *FormatoEvaluacionCarrera, err error) {
	o := orm.NewOrm()
	v = &FormatoEvaluacionCarrera{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFormatoEvaluacionCarrera retrieves all FormatoEvaluacionCarrera matches certain condition. Returns empty list if
// no records exist
func GetAllFormatoEvaluacionCarrera(query map[string]string, fields []string, sortby []string, order []string, related []interface{},
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FormatoEvaluacionCarrera))
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

	var l []FormatoEvaluacionCarrera
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

// UpdateFormatoEvaluacionCarrera updates FormatoEvaluacionCarrera by Id and returns error if
// the record to be updated doesn't exist
func UpdateFormatoEvaluacionCarreraById(m *FormatoEvaluacionCarrera) (err error) {
	o := orm.NewOrm()
	v := FormatoEvaluacionCarrera{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFormatoEvaluacionCarrera deletes FormatoEvaluacionCarrera by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFormatoEvaluacionCarrera(id int) (err error) {
	o := orm.NewOrm()
	v := FormatoEvaluacionCarrera{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FormatoEvaluacionCarrera{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
