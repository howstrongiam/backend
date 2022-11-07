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

func (ec *executionContext) _Type__id(ctx context.Context, field graphql.CollectedField, obj *model.Type) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Type__id(ctx, field)
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

func (ec *executionContext) fieldContext_Type__id(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Type",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Type_Title(ctx context.Context, field graphql.CollectedField, obj *model.Type) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Type_Title(ctx, field)
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

func (ec *executionContext) fieldContext_Type_Title(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Type",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Type_Styles(ctx context.Context, field graphql.CollectedField, obj *model.Type) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Type_Styles(ctx, field)
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
		return obj.Styles, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]*model.Style)
	fc.Result = res
	return ec.marshalOStyle2ᚕᚖgithubᚗcomᚋhowstrongiamᚋbackendᚋgraphᚋmodelᚐStyle(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Type_Styles(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Type",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "_id":
				return ec.fieldContext_Style__id(ctx, field)
			case "Title":
				return ec.fieldContext_Style_Title(ctx, field)
			case "Products":
				return ec.fieldContext_Style_Products(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Style", field.Name)
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputAddTypeRequest(ctx context.Context, obj interface{}) (model.AddTypeRequest, error) {
	var it model.AddTypeRequest
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"Title", "CategoryId"}
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
		case "CategoryId":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("CategoryId"))
			it.CategoryID, err = ec.unmarshalNString2string(ctx, v)
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

var typeImplementors = []string{"Type"}

func (ec *executionContext) _Type(ctx context.Context, sel ast.SelectionSet, obj *model.Type) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, typeImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Type")
		case "_id":

			out.Values[i] = ec._Type__id(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "Title":

			out.Values[i] = ec._Type_Title(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "Styles":

			out.Values[i] = ec._Type_Styles(ctx, field, obj)

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

func (ec *executionContext) unmarshalNAddTypeRequest2githubᚗcomᚋhowstrongiamᚋbackendᚋgraphᚋmodelᚐAddTypeRequest(ctx context.Context, v interface{}) (model.AddTypeRequest, error) {
	res, err := ec.unmarshalInputAddTypeRequest(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNType2githubᚗcomᚋhowstrongiamᚋbackendᚋgraphᚋmodelᚐType(ctx context.Context, sel ast.SelectionSet, v model.Type) graphql.Marshaler {
	return ec._Type(ctx, sel, &v)
}

func (ec *executionContext) marshalNType2ᚖgithubᚗcomᚋhowstrongiamᚋbackendᚋgraphᚋmodelᚐType(ctx context.Context, sel ast.SelectionSet, v *model.Type) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._Type(ctx, sel, v)
}

func (ec *executionContext) marshalOType2ᚕᚖgithubᚗcomᚋhowstrongiamᚋbackendᚋgraphᚋmodelᚐType(ctx context.Context, sel ast.SelectionSet, v []*model.Type) graphql.Marshaler {
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
			ret[i] = ec.marshalOType2ᚖgithubᚗcomᚋhowstrongiamᚋbackendᚋgraphᚋmodelᚐType(ctx, sel, v[i])
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

func (ec *executionContext) marshalOType2ᚖgithubᚗcomᚋhowstrongiamᚋbackendᚋgraphᚋmodelᚐType(ctx context.Context, sel ast.SelectionSet, v *model.Type) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._Type(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************