// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// List all entities (Management Groups, Subscriptions, etc.) for the authenticated user.
func GetEntity(ctx *pulumi.Context, args *GetEntityArgs, opts ...pulumi.InvokeOption) (*GetEntityResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetEntityResult
	err := ctx.Invoke("azure-native:management/v20200501:getEntity", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetEntityArgs struct {
	// The filter parameter allows you to filter on the the name or display name fields. You can check for equality on the name field (e.g. name eq '{entityName}')  and you can check for substrings on either the name or display name fields(e.g. contains(name, '{substringToSearch}'), contains(displayName, '{substringToSearch')). Note that the '{entityName}' and '{substringToSearch}' fields are checked case insensitively.
	Filter *string `pulumi:"filter"`
	// A filter which allows the get entities call to focus on a particular group (i.e. "$filter=name eq 'groupName'")
	GroupName *string `pulumi:"groupName"`
	// The $search parameter is used in conjunction with the $filter parameter to return three different outputs depending on the parameter passed in.
	// With $search=AllowedParents the API will return the entity info of all groups that the requested entity will be able to reparent to as determined by the user's permissions.
	// With $search=AllowedChildren the API will return the entity info of all entities that can be added as children of the requested entity.
	// With $search=ParentAndFirstLevelChildren the API will return the parent and  first level of children that the user has either direct access to or indirect access via one of their descendants.
	// With $search=ParentOnly the API will return only the group if the user has access to at least one of the descendants of the group.
	// With $search=ChildrenOnly the API will return only the first level of children of the group entity info specified in $filter.  The user must have direct access to the children entities or one of it's descendants for it to show up in the results.
	Search *string `pulumi:"search"`
	// This parameter specifies the fields to include in the response. Can include any combination of Name,DisplayName,Type,ParentDisplayNameChain,ParentChain, e.g. '$select=Name,DisplayName,Type,ParentDisplayNameChain,ParentNameChain'. When specified the $select parameter can override select in $skipToken.
	Select *string `pulumi:"select"`
	// Number of entities to skip over when retrieving results. Passing this in will override $skipToken.
	Skip *int `pulumi:"skip"`
	// Page continuation token is only used if a previous operation returned a partial result.
	// If a previous response contains a nextLink element, the value of the nextLink element will include a token parameter that specifies a starting point to use for subsequent calls.
	Skiptoken *string `pulumi:"skiptoken"`
	// Number of elements to return when retrieving results. Passing this in will override $skipToken.
	Top *int `pulumi:"top"`
	// The view parameter allows clients to filter the type of data that is returned by the getEntities call.
	View *string `pulumi:"view"`
}

// Describes the result of the request to view entities.
type GetEntityResult struct {
	// Total count of records that match the filter
	Count int `pulumi:"count"`
	// The URL to use for getting the next set of results.
	NextLink string `pulumi:"nextLink"`
	// The list of entities.
	Value []EntityInfoResponse `pulumi:"value"`
}

func GetEntityOutput(ctx *pulumi.Context, args GetEntityOutputArgs, opts ...pulumi.InvokeOption) GetEntityResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetEntityResult, error) {
			args := v.(GetEntityArgs)
			r, err := GetEntity(ctx, &args, opts...)
			var s GetEntityResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetEntityResultOutput)
}

type GetEntityOutputArgs struct {
	// The filter parameter allows you to filter on the the name or display name fields. You can check for equality on the name field (e.g. name eq '{entityName}')  and you can check for substrings on either the name or display name fields(e.g. contains(name, '{substringToSearch}'), contains(displayName, '{substringToSearch')). Note that the '{entityName}' and '{substringToSearch}' fields are checked case insensitively.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// A filter which allows the get entities call to focus on a particular group (i.e. "$filter=name eq 'groupName'")
	GroupName pulumi.StringPtrInput `pulumi:"groupName"`
	// The $search parameter is used in conjunction with the $filter parameter to return three different outputs depending on the parameter passed in.
	// With $search=AllowedParents the API will return the entity info of all groups that the requested entity will be able to reparent to as determined by the user's permissions.
	// With $search=AllowedChildren the API will return the entity info of all entities that can be added as children of the requested entity.
	// With $search=ParentAndFirstLevelChildren the API will return the parent and  first level of children that the user has either direct access to or indirect access via one of their descendants.
	// With $search=ParentOnly the API will return only the group if the user has access to at least one of the descendants of the group.
	// With $search=ChildrenOnly the API will return only the first level of children of the group entity info specified in $filter.  The user must have direct access to the children entities or one of it's descendants for it to show up in the results.
	Search pulumi.StringPtrInput `pulumi:"search"`
	// This parameter specifies the fields to include in the response. Can include any combination of Name,DisplayName,Type,ParentDisplayNameChain,ParentChain, e.g. '$select=Name,DisplayName,Type,ParentDisplayNameChain,ParentNameChain'. When specified the $select parameter can override select in $skipToken.
	Select pulumi.StringPtrInput `pulumi:"select"`
	// Number of entities to skip over when retrieving results. Passing this in will override $skipToken.
	Skip pulumi.IntPtrInput `pulumi:"skip"`
	// Page continuation token is only used if a previous operation returned a partial result.
	// If a previous response contains a nextLink element, the value of the nextLink element will include a token parameter that specifies a starting point to use for subsequent calls.
	Skiptoken pulumi.StringPtrInput `pulumi:"skiptoken"`
	// Number of elements to return when retrieving results. Passing this in will override $skipToken.
	Top pulumi.IntPtrInput `pulumi:"top"`
	// The view parameter allows clients to filter the type of data that is returned by the getEntities call.
	View pulumi.StringPtrInput `pulumi:"view"`
}

func (GetEntityOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEntityArgs)(nil)).Elem()
}

// Describes the result of the request to view entities.
type GetEntityResultOutput struct{ *pulumi.OutputState }

func (GetEntityResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEntityResult)(nil)).Elem()
}

func (o GetEntityResultOutput) ToGetEntityResultOutput() GetEntityResultOutput {
	return o
}

func (o GetEntityResultOutput) ToGetEntityResultOutputWithContext(ctx context.Context) GetEntityResultOutput {
	return o
}

// Total count of records that match the filter
func (o GetEntityResultOutput) Count() pulumi.IntOutput {
	return o.ApplyT(func(v GetEntityResult) int { return v.Count }).(pulumi.IntOutput)
}

// The URL to use for getting the next set of results.
func (o GetEntityResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v GetEntityResult) string { return v.NextLink }).(pulumi.StringOutput)
}

// The list of entities.
func (o GetEntityResultOutput) Value() EntityInfoResponseArrayOutput {
	return o.ApplyT(func(v GetEntityResult) []EntityInfoResponse { return v.Value }).(EntityInfoResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetEntityResultOutput{})
}
