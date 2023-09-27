// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a single custom entity store assignment by name for the provided subscription and resource group.
func LookupCustomEntityStoreAssignment(ctx *pulumi.Context, args *LookupCustomEntityStoreAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupCustomEntityStoreAssignmentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupCustomEntityStoreAssignmentResult
	err := ctx.Invoke("azure-native:security/v20210701preview:getCustomEntityStoreAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCustomEntityStoreAssignmentArgs struct {
	// Name of the custom entity store assignment. Generated name is GUID.
	CustomEntityStoreAssignmentName string `pulumi:"customEntityStoreAssignmentName"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Custom entity store assignment
type LookupCustomEntityStoreAssignmentResult struct {
	// The link to entity store database.
	EntityStoreDatabaseLink *string `pulumi:"entityStoreDatabaseLink"`
	// Resource Id
	Id string `pulumi:"id"`
	// Resource name
	Name string `pulumi:"name"`
	// The principal assigned with entity store. Format of principal is: [AAD type]=[PrincipalObjectId];[TenantId]
	Principal *string `pulumi:"principal"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource type
	Type string `pulumi:"type"`
}

func LookupCustomEntityStoreAssignmentOutput(ctx *pulumi.Context, args LookupCustomEntityStoreAssignmentOutputArgs, opts ...pulumi.InvokeOption) LookupCustomEntityStoreAssignmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCustomEntityStoreAssignmentResult, error) {
			args := v.(LookupCustomEntityStoreAssignmentArgs)
			r, err := LookupCustomEntityStoreAssignment(ctx, &args, opts...)
			var s LookupCustomEntityStoreAssignmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCustomEntityStoreAssignmentResultOutput)
}

type LookupCustomEntityStoreAssignmentOutputArgs struct {
	// Name of the custom entity store assignment. Generated name is GUID.
	CustomEntityStoreAssignmentName pulumi.StringInput `pulumi:"customEntityStoreAssignmentName"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCustomEntityStoreAssignmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomEntityStoreAssignmentArgs)(nil)).Elem()
}

// Custom entity store assignment
type LookupCustomEntityStoreAssignmentResultOutput struct{ *pulumi.OutputState }

func (LookupCustomEntityStoreAssignmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomEntityStoreAssignmentResult)(nil)).Elem()
}

func (o LookupCustomEntityStoreAssignmentResultOutput) ToLookupCustomEntityStoreAssignmentResultOutput() LookupCustomEntityStoreAssignmentResultOutput {
	return o
}

func (o LookupCustomEntityStoreAssignmentResultOutput) ToLookupCustomEntityStoreAssignmentResultOutputWithContext(ctx context.Context) LookupCustomEntityStoreAssignmentResultOutput {
	return o
}

func (o LookupCustomEntityStoreAssignmentResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupCustomEntityStoreAssignmentResult] {
	return pulumix.Output[LookupCustomEntityStoreAssignmentResult]{
		OutputState: o.OutputState,
	}
}

// The link to entity store database.
func (o LookupCustomEntityStoreAssignmentResultOutput) EntityStoreDatabaseLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomEntityStoreAssignmentResult) *string { return v.EntityStoreDatabaseLink }).(pulumi.StringPtrOutput)
}

// Resource Id
func (o LookupCustomEntityStoreAssignmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomEntityStoreAssignmentResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name
func (o LookupCustomEntityStoreAssignmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomEntityStoreAssignmentResult) string { return v.Name }).(pulumi.StringOutput)
}

// The principal assigned with entity store. Format of principal is: [AAD type]=[PrincipalObjectId];[TenantId]
func (o LookupCustomEntityStoreAssignmentResultOutput) Principal() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomEntityStoreAssignmentResult) *string { return v.Principal }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupCustomEntityStoreAssignmentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCustomEntityStoreAssignmentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type
func (o LookupCustomEntityStoreAssignmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomEntityStoreAssignmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCustomEntityStoreAssignmentResultOutput{})
}
