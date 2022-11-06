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

func (ec *executionContext) _Product__id(ctx context.Context, field graphql.CollectedField, obj *model.Product) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Product__id(ctx, field)
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

func (ec *executionContext) fieldContext_Product__id(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Product",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Product_Title(ctx context.Context, field graphql.CollectedField, obj *model.Product) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Product_Title(ctx, field)
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

func (ec *executionContext) fieldContext_Product_Title(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Product",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Product_Url(ctx context.Context, field graphql.CollectedField, obj *model.Product) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Product_Url(ctx, field)
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
		return obj.URL, nil
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

func (ec *executionContext) fieldContext_Product_Url(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Product",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Product_Description(ctx context.Context, field graphql.CollectedField, obj *model.Product) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Product_Description(ctx, field)
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
		return obj.Description, nil
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

func (ec *executionContext) fieldContext_Product_Description(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Product",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Product_UserId(ctx context.Context, field graphql.CollectedField, obj *model.Product) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Product_UserId(ctx, field)
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
		return obj.UserID, nil
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

func (ec *executionContext) fieldContext_Product_UserId(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Product",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Product_Image(ctx context.Context, field graphql.CollectedField, obj *model.Product) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Product_Image(ctx, field)
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
		return obj.Image, nil
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
	res := resTmp.(*model.Image)
	fc.Result = res
	return ec.marshalNImage2ᚖgithubᚗcomᚋhowstrongiamᚋbackendᚋgraphᚋmodelᚐImage(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Product_Image(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Product",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "_id":
				return ec.fieldContext_Image__id(ctx, field)
			case "Location":
				return ec.fieldContext_Image_Location(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Image", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _Product_Certification(ctx context.Context, field graphql.CollectedField, obj *model.Product) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Product_Certification(ctx, field)
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
		return obj.Certification, nil
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
	res := resTmp.(*model.Certification)
	fc.Result = res
	return ec.marshalNCertification2ᚖgithubᚗcomᚋhowstrongiamᚋbackendᚋgraphᚋmodelᚐCertification(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Product_Certification(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Product",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "_id":
				return ec.fieldContext_Certification__id(ctx, field)
			case "CertifyingCompany":
				return ec.fieldContext_Certification_CertifyingCompany(ctx, field)
			case "CertName":
				return ec.fieldContext_Certification_CertName(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Certification", field.Name)
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputAddProductRequest(ctx context.Context, obj interface{}) (model.AddProductRequest, error) {
	var it model.AddProductRequest
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"Title", "Url", "Description", "UserId", "ImageLocation", "Certification", "StyleId"}
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
		case "Url":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("Url"))
			it.URL, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "Description":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("Description"))
			it.Description, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "UserId":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("UserId"))
			it.UserID, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "ImageLocation":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("ImageLocation"))
			it.ImageLocation, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "Certification":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("Certification"))
			it.Certification, err = ec.unmarshalNAddCertificationRequest2ᚖgithubᚗcomᚋhowstrongiamᚋbackendᚋgraphᚋmodelᚐAddCertificationRequest(ctx, v)
			if err != nil {
				return it, err
			}
		case "StyleId":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("StyleId"))
			it.StyleID, err = ec.unmarshalNString2string(ctx, v)
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

var productImplementors = []string{"Product"}

func (ec *executionContext) _Product(ctx context.Context, sel ast.SelectionSet, obj *model.Product) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, productImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Product")
		case "_id":

			out.Values[i] = ec._Product__id(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "Title":

			out.Values[i] = ec._Product_Title(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "Url":

			out.Values[i] = ec._Product_Url(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "Description":

			out.Values[i] = ec._Product_Description(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "UserId":

			out.Values[i] = ec._Product_UserId(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "Image":

			out.Values[i] = ec._Product_Image(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "Certification":

			out.Values[i] = ec._Product_Certification(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
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

func (ec *executionContext) unmarshalNAddProductRequest2githubᚗcomᚋhowstrongiamᚋbackendᚋgraphᚋmodelᚐAddProductRequest(ctx context.Context, v interface{}) (model.AddProductRequest, error) {
	res, err := ec.unmarshalInputAddProductRequest(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNProduct2githubᚗcomᚋhowstrongiamᚋbackendᚋgraphᚋmodelᚐProduct(ctx context.Context, sel ast.SelectionSet, v model.Product) graphql.Marshaler {
	return ec._Product(ctx, sel, &v)
}

func (ec *executionContext) marshalNProduct2ᚖgithubᚗcomᚋhowstrongiamᚋbackendᚋgraphᚋmodelᚐProduct(ctx context.Context, sel ast.SelectionSet, v *model.Product) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._Product(ctx, sel, v)
}

func (ec *executionContext) marshalOProduct2ᚕᚖgithubᚗcomᚋhowstrongiamᚋbackendᚋgraphᚋmodelᚐProduct(ctx context.Context, sel ast.SelectionSet, v []*model.Product) graphql.Marshaler {
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
			ret[i] = ec.marshalOProduct2ᚖgithubᚗcomᚋhowstrongiamᚋbackendᚋgraphᚋmodelᚐProduct(ctx, sel, v[i])
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

func (ec *executionContext) marshalOProduct2ᚖgithubᚗcomᚋhowstrongiamᚋbackendᚋgraphᚋmodelᚐProduct(ctx context.Context, sel ast.SelectionSet, v *model.Product) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._Product(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************
