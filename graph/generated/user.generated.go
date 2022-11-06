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

func (ec *executionContext) _User__id(ctx context.Context, field graphql.CollectedField, obj *model.User) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_User__id(ctx, field)
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

func (ec *executionContext) fieldContext_User__id(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "User",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _User_firstName(ctx context.Context, field graphql.CollectedField, obj *model.User) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_User_firstName(ctx, field)
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
		return obj.FirstName, nil
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

func (ec *executionContext) fieldContext_User_firstName(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "User",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _User_lastName(ctx context.Context, field graphql.CollectedField, obj *model.User) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_User_lastName(ctx, field)
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
		return obj.LastName, nil
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

func (ec *executionContext) fieldContext_User_lastName(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "User",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _User_favourites(ctx context.Context, field graphql.CollectedField, obj *model.User) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_User_favourites(ctx, field)
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
		return obj.Favourites, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]*model.Favourite)
	fc.Result = res
	return ec.marshalOFavourite2ᚕᚖgithubᚗcomᚋhowstrongiamᚋbackendᚋgraphᚋmodelᚐFavourite(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_User_favourites(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "User",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "_id":
				return ec.fieldContext_Favourite__id(ctx, field)
			case "product":
				return ec.fieldContext_Favourite_product(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Favourite", field.Name)
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputNewUser(ctx context.Context, obj interface{}) (model.NewUser, error) {
	var it model.NewUser
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"firstName", "lastName", "email", "password"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "firstName":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("firstName"))
			it.FirstName, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "lastName":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("lastName"))
			it.LastName, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "email":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("email"))
			it.Email, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "password":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("password"))
			it.Password, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputUpdateUser(ctx context.Context, obj interface{}) (model.UpdateUser, error) {
	var it model.UpdateUser
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"userId", "firstName", "lastName"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "userId":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("userId"))
			it.UserID, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "firstName":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("firstName"))
			it.FirstName, err = ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
		case "lastName":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("lastName"))
			it.LastName, err = ec.unmarshalOString2ᚖstring(ctx, v)
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

var userImplementors = []string{"User"}

func (ec *executionContext) _User(ctx context.Context, sel ast.SelectionSet, obj *model.User) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, userImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("User")
		case "_id":

			out.Values[i] = ec._User__id(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "firstName":

			out.Values[i] = ec._User_firstName(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "lastName":

			out.Values[i] = ec._User_lastName(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "favourites":

			out.Values[i] = ec._User_favourites(ctx, field, obj)

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

func (ec *executionContext) unmarshalNNewUser2githubᚗcomᚋhowstrongiamᚋbackendᚋgraphᚋmodelᚐNewUser(ctx context.Context, v interface{}) (model.NewUser, error) {
	res, err := ec.unmarshalInputNewUser(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalNUpdateUser2githubᚗcomᚋhowstrongiamᚋbackendᚋgraphᚋmodelᚐUpdateUser(ctx context.Context, v interface{}) (model.UpdateUser, error) {
	res, err := ec.unmarshalInputUpdateUser(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNUser2githubᚗcomᚋhowstrongiamᚋbackendᚋgraphᚋmodelᚐUser(ctx context.Context, sel ast.SelectionSet, v model.User) graphql.Marshaler {
	return ec._User(ctx, sel, &v)
}

func (ec *executionContext) marshalNUser2ᚕᚖgithubᚗcomᚋhowstrongiamᚋbackendᚋgraphᚋmodelᚐUser(ctx context.Context, sel ast.SelectionSet, v []*model.User) graphql.Marshaler {
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
			ret[i] = ec.marshalOUser2ᚖgithubᚗcomᚋhowstrongiamᚋbackendᚋgraphᚋmodelᚐUser(ctx, sel, v[i])
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

func (ec *executionContext) marshalNUser2ᚖgithubᚗcomᚋhowstrongiamᚋbackendᚋgraphᚋmodelᚐUser(ctx context.Context, sel ast.SelectionSet, v *model.User) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._User(ctx, sel, v)
}

func (ec *executionContext) marshalOUser2ᚖgithubᚗcomᚋhowstrongiamᚋbackendᚋgraphᚋmodelᚐUser(ctx context.Context, sel ast.SelectionSet, v *model.User) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._User(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************
