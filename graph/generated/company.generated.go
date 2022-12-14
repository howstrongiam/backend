// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"strconv"

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

func (ec *executionContext) _Company__id(ctx context.Context, field graphql.CollectedField, obj *model.Company) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Company__id(ctx, field)
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

func (ec *executionContext) fieldContext_Company__id(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Company",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Company_url(ctx context.Context, field graphql.CollectedField, obj *model.Company) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Company_url(ctx, field)
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

func (ec *executionContext) fieldContext_Company_url(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Company",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Company_description(ctx context.Context, field graphql.CollectedField, obj *model.Company) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Company_description(ctx, field)
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

func (ec *executionContext) fieldContext_Company_description(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Company",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Company_user(ctx context.Context, field graphql.CollectedField, obj *model.Company) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Company_user(ctx, field)
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
		return obj.User, nil
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
	res := resTmp.(*model.User)
	fc.Result = res
	return ec.marshalNUser2???github???com???howstrongiam???backend???graph???model???User(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Company_user(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Company",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "_id":
				return ec.fieldContext_User__id(ctx, field)
			case "firstName":
				return ec.fieldContext_User_firstName(ctx, field)
			case "lastName":
				return ec.fieldContext_User_lastName(ctx, field)
			case "email":
				return ec.fieldContext_User_email(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type User", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _Company_isVerified(ctx context.Context, field graphql.CollectedField, obj *model.Company) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Company_isVerified(ctx, field)
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
		return obj.IsVerified, nil
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
	res := resTmp.(bool)
	fc.Result = res
	return ec.marshalNBoolean2bool(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Company_isVerified(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Company",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Boolean does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Company_image(ctx context.Context, field graphql.CollectedField, obj *model.Company) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Company_image(ctx, field)
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
	return ec.marshalNImage2???github???com???howstrongiam???backend???graph???model???Image(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Company_image(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Company",
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

func (ec *executionContext) _Company_certification(ctx context.Context, field graphql.CollectedField, obj *model.Company) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Company_certification(ctx, field)
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
	return ec.marshalNCertification2???github???com???howstrongiam???backend???graph???model???Certification(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Company_certification(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Company",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "_id":
				return ec.fieldContext_Certification__id(ctx, field)
			case "Name":
				return ec.fieldContext_Certification_Name(ctx, field)
			case "LogoLink":
				return ec.fieldContext_Certification_LogoLink(ctx, field)
			case "Industry":
				return ec.fieldContext_Certification_Industry(ctx, field)
			case "ProvidingCompany":
				return ec.fieldContext_Certification_ProvidingCompany(ctx, field)
			case "Certifies":
				return ec.fieldContext_Certification_Certifies(ctx, field)
			case "Type":
				return ec.fieldContext_Certification_Type(ctx, field)
			case "Audited":
				return ec.fieldContext_Certification_Audited(ctx, field)
			case "Auditor":
				return ec.fieldContext_Certification_Auditor(ctx, field)
			case "ProvidingCompanyWebsite":
				return ec.fieldContext_Certification_ProvidingCompanyWebsite(ctx, field)
			case "FoundWhere":
				return ec.fieldContext_Certification_FoundWhere(ctx, field)
			case "HowToGetIt":
				return ec.fieldContext_Certification_HowToGetIt(ctx, field)
			case "Notes":
				return ec.fieldContext_Certification_Notes(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Certification", field.Name)
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputAddCompanyRequest(ctx context.Context, obj interface{}) (model.AddCompanyRequest, error) {
	var it model.AddCompanyRequest
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"url", "description", "userId", "isVerified", "imageLocation", "certification"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "url":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("url"))
			it.URL, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "description":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("description"))
			it.Description, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "userId":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("userId"))
			it.UserID, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "isVerified":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("isVerified"))
			it.IsVerified, err = ec.unmarshalNBoolean2bool(ctx, v)
			if err != nil {
				return it, err
			}
		case "imageLocation":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("imageLocation"))
			it.ImageLocation, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "certification":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("certification"))
			it.Certification, err = ec.unmarshalNAddCertificationRequest2???github???com???howstrongiam???backend???graph???model???AddCertificationRequest(ctx, v)
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

var companyImplementors = []string{"Company"}

func (ec *executionContext) _Company(ctx context.Context, sel ast.SelectionSet, obj *model.Company) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, companyImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Company")
		case "_id":

			out.Values[i] = ec._Company__id(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "url":

			out.Values[i] = ec._Company_url(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "description":

			out.Values[i] = ec._Company_description(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "user":

			out.Values[i] = ec._Company_user(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "isVerified":

			out.Values[i] = ec._Company_isVerified(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "image":

			out.Values[i] = ec._Company_image(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "certification":

			out.Values[i] = ec._Company_certification(ctx, field, obj)

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

func (ec *executionContext) unmarshalNAddCompanyRequest2github???com???howstrongiam???backend???graph???model???AddCompanyRequest(ctx context.Context, v interface{}) (model.AddCompanyRequest, error) {
	res, err := ec.unmarshalInputAddCompanyRequest(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNCompany2github???com???howstrongiam???backend???graph???model???Company(ctx context.Context, sel ast.SelectionSet, v model.Company) graphql.Marshaler {
	return ec._Company(ctx, sel, &v)
}

func (ec *executionContext) marshalNCompany2???github???com???howstrongiam???backend???graph???model???Company(ctx context.Context, sel ast.SelectionSet, v *model.Company) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._Company(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************
