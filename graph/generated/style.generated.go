// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/99designs/gqlgen/graphql"
	"github.com/howstrongiam/backend/graph/model"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _Style__id(ctx context.Context, field graphql.CollectedField, obj *model.Style) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Style__id(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Style__id(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Style",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Style_Title(ctx context.Context, field graphql.CollectedField, obj *model.Style) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Style_Title(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Title, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Style_Title(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Style",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Style_Products(ctx context.Context, field graphql.CollectedField, obj *model.Style) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Style_Products(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Products, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]*model.Product)
	fc.Result = res
	return ec.marshalOProduct2??????github???com???howstrongiam???backend???graph???model???Product(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Style_Products(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Style",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "_id":
				return ec.fieldContext_Product__id(ctx, field)
			case "Title":
				return ec.fieldContext_Product_Title(ctx, field)
			case "Description":
				return ec.fieldContext_Product_Description(ctx, field)
			case "Certification":
				return ec.fieldContext_Product_Certification(ctx, field)
			case "ProductCertifications":
				return ec.fieldContext_Product_ProductCertifications(ctx, field)
			case "CompanyCertifications":
				return ec.fieldContext_Product_CompanyCertifications(ctx, field)
			case "MaterialsAndIngredients":
				return ec.fieldContext_Product_MaterialsAndIngredients(ctx, field)
			case "GiveBackPrograms":
				return ec.fieldContext_Product_GiveBackPrograms(ctx, field)
			case "OwnersAndFounders":
				return ec.fieldContext_Product_OwnersAndFounders(ctx, field)
			case "Section":
				return ec.fieldContext_Product_Section(ctx, field)
			case "Department":
				return ec.fieldContext_Product_Department(ctx, field)
			case "Category":
				return ec.fieldContext_Product_Category(ctx, field)
			case "Type":
				return ec.fieldContext_Product_Type(ctx, field)
			case "Style":
				return ec.fieldContext_Product_Style(ctx, field)
			case "ImageLinks":
				return ec.fieldContext_Product_ImageLinks(ctx, field)
			case "PurchaseInfo":
				return ec.fieldContext_Product_PurchaseInfo(ctx, field)
			case "Verified":
				return ec.fieldContext_Product_Verified(ctx, field)
			case "VerifiedBy":
				return ec.fieldContext_Product_VerifiedBy(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Product", field.Name)
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputAddStyleRequest(ctx context.Context, obj interface{}) (model.AddStyleRequest, error) {
	var it model.AddStyleRequest
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"Title", "TypeId"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "Title":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("Title"))
			it.Title, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "TypeId":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("TypeId"))
			it.TypeID, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var styleImplementors = []string{"Style"}

func (ec *executionContext) _Style(ctx context.Context, sel ast.SelectionSet, obj *model.Style) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, styleImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Style")
		case "_id":

			out.Values[i] = ec._Style__id(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "Title":

			out.Values[i] = ec._Style_Title(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "Products":

			out.Values[i] = ec._Style_Products(ctx, field, obj)

		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) unmarshalNAddStyleRequest2github???com???howstrongiam???backend???graph???model???AddStyleRequest(ctx context.Context, v interface{}) (model.AddStyleRequest, error) {
	res, err := ec.unmarshalInputAddStyleRequest(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNStyle2github???com???howstrongiam???backend???graph???model???Style(ctx context.Context, sel ast.SelectionSet, v model.Style) graphql.Marshaler {
	return ec._Style(ctx, sel, &v)
}

func (ec *executionContext) marshalNStyle2???github???com???howstrongiam???backend???graph???model???Style(ctx context.Context, sel ast.SelectionSet, v *model.Style) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._Style(ctx, sel, v)
}

func (ec *executionContext) marshalOStyle2??????github???com???howstrongiam???backend???graph???model???Style(ctx context.Context, sel ast.SelectionSet, v []*model.Style) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalOStyle2???github???com???howstrongiam???backend???graph???model???Style(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	return ret
}

func (ec *executionContext) marshalOStyle2???github???com???howstrongiam???backend???graph???model???Style(ctx context.Context, sel ast.SelectionSet, v *model.Style) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._Style(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************
