package repository

import (
	"context"
	"log"

	"github.com/Akshaytermin/gqlbeego/graph/model"
	"github.com/astaxie/beego/orm"
)

func Create(ctx context.Context, doc model.Product) (model.Product, error) {
	var resp model.Product
	o := orm.NewOrm()

	id, err := o.Insert(&doc)

	_, err = o.QueryTable("Product").Filter("id", id).All(&resp)
	if err != nil {
		log.Fatal(err)
		return resp, err
	}

	return resp, nil
}

func Update(ctx context.Context, id int, doc model.Product) (model.Product, error) {
	var product model.Product
	o := orm.NewOrm()
	products := model.Product{
		Id:    id,
		Name:  doc.Name,
		Price: doc.Price,
	}

	_, err := o.Update(&products)

	_, err = o.QueryTable("Product").Filter("id", id).All(&product)
	if err != nil {
		log.Fatal(err)
		return product, err
	}

	return product, nil
}

func Delete(ctx context.Context, id int) (model.Product, error) {
	o := orm.NewOrm()

	var product model.Product
	//Fetch based on ID and delete

	product.Id = id
	_, err := o.Delete(&product, "id")
	if err != nil {
		log.Fatal(err)
		return product, err
	}

	return product, nil
}

func Find(ctx context.Context) ([]*model.Product, error) {

	o := orm.NewOrm()
	var products []*model.Product

	_, err := o.QueryTable("Product").All(&products)

	if err != nil {
		log.Fatal(err)
		return products, err
	}

	return products, err
}
