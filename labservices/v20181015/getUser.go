// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20181015

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get user
func LookupUser(ctx *pulumi.Context, args *LookupUserArgs, opts ...pulumi.InvokeOption) (*LookupUserResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupUserResult
	err := ctx.Invoke("azure-native:labservices/v20181015:getUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUserArgs struct {
	// Specify the $expand query. Example: 'properties($select=email)'
	Expand *string `pulumi:"expand"`
	// The name of the lab Account.
	LabAccountName string `pulumi:"labAccountName"`
	// The name of the lab.
	LabName string `pulumi:"labName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the user.
	UserName string `pulumi:"userName"`
}

// The User registered to a lab
type LookupUserResult struct {
	// The user email address, as it was specified during registration.
	Email string `pulumi:"email"`
	// The user family name, as it was specified during registration.
	FamilyName string `pulumi:"familyName"`
	// The user given name, as it was specified during registration.
	GivenName string `pulumi:"givenName"`
	// The identifier of the resource.
	Id string `pulumi:"id"`
	// The details of the latest operation. ex: status, error
	LatestOperationResult LatestOperationResultResponse `pulumi:"latestOperationResult"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The provisioning status of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The user tenant ID, as it was specified during registration.
	TenantId string `pulumi:"tenantId"`
	// How long the user has used his VMs in this lab
	TotalUsage string `pulumi:"totalUsage"`
	// The type of the resource.
	Type string `pulumi:"type"`
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier *string `pulumi:"uniqueIdentifier"`
}

func LookupUserOutput(ctx *pulumi.Context, args LookupUserOutputArgs, opts ...pulumi.InvokeOption) LookupUserResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUserResult, error) {
			args := v.(LookupUserArgs)
			r, err := LookupUser(ctx, &args, opts...)
			var s LookupUserResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUserResultOutput)
}

type LookupUserOutputArgs struct {
	// Specify the $expand query. Example: 'properties($select=email)'
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the lab Account.
	LabAccountName pulumi.StringInput `pulumi:"labAccountName"`
	// The name of the lab.
	LabName pulumi.StringInput `pulumi:"labName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the user.
	UserName pulumi.StringInput `pulumi:"userName"`
}

func (LookupUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserArgs)(nil)).Elem()
}

// The User registered to a lab
type LookupUserResultOutput struct{ *pulumi.OutputState }

func (LookupUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserResult)(nil)).Elem()
}

func (o LookupUserResultOutput) ToLookupUserResultOutput() LookupUserResultOutput {
	return o
}

func (o LookupUserResultOutput) ToLookupUserResultOutputWithContext(ctx context.Context) LookupUserResultOutput {
	return o
}

// The user email address, as it was specified during registration.
func (o LookupUserResultOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Email }).(pulumi.StringOutput)
}

// The user family name, as it was specified during registration.
func (o LookupUserResultOutput) FamilyName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.FamilyName }).(pulumi.StringOutput)
}

// The user given name, as it was specified during registration.
func (o LookupUserResultOutput) GivenName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.GivenName }).(pulumi.StringOutput)
}

// The identifier of the resource.
func (o LookupUserResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Id }).(pulumi.StringOutput)
}

// The details of the latest operation. ex: status, error
func (o LookupUserResultOutput) LatestOperationResult() LatestOperationResultResponseOutput {
	return o.ApplyT(func(v LookupUserResult) LatestOperationResultResponse { return v.LatestOperationResult }).(LatestOperationResultResponseOutput)
}

// The location of the resource.
func (o LookupUserResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o LookupUserResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning status of the resource.
func (o LookupUserResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// The tags of the resource.
func (o LookupUserResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupUserResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The user tenant ID, as it was specified during registration.
func (o LookupUserResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.TenantId }).(pulumi.StringOutput)
}

// How long the user has used his VMs in this lab
func (o LookupUserResultOutput) TotalUsage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.TotalUsage }).(pulumi.StringOutput)
}

// The type of the resource.
func (o LookupUserResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Type }).(pulumi.StringOutput)
}

// The unique immutable identifier of a resource (Guid).
func (o LookupUserResultOutput) UniqueIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *string { return v.UniqueIdentifier }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserResultOutput{})
}
