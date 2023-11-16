// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the properties of the specified user.
func LookupUser(ctx *pulumi.Context, args *LookupUserArgs, opts ...pulumi.InvokeOption) (*LookupUserResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupUserResult
	err := ctx.Invoke("azure-native:databoxedge/v20220301:getUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUserArgs struct {
	// The device name.
	DeviceName string `pulumi:"deviceName"`
	// The user name.
	Name string `pulumi:"name"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Represents a user who has access to one or more shares on the Data Box Edge/Gateway device.
type LookupUserResult struct {
	// The password details.
	EncryptedPassword *AsymmetricEncryptedSecretResponse `pulumi:"encryptedPassword"`
	// The path ID that uniquely identifies the object.
	Id string `pulumi:"id"`
	// The object name.
	Name string `pulumi:"name"`
	// List of shares that the user has rights on. This field should not be specified during user creation.
	ShareAccessRights []ShareAccessRightResponse `pulumi:"shareAccessRights"`
	// Metadata pertaining to creation and last modification of User
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The hierarchical type of the object.
	Type string `pulumi:"type"`
	// Type of the user.
	UserType string `pulumi:"userType"`
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
	// The device name.
	DeviceName pulumi.StringInput `pulumi:"deviceName"`
	// The user name.
	Name pulumi.StringInput `pulumi:"name"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserArgs)(nil)).Elem()
}

// Represents a user who has access to one or more shares on the Data Box Edge/Gateway device.
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

// The password details.
func (o LookupUserResultOutput) EncryptedPassword() AsymmetricEncryptedSecretResponsePtrOutput {
	return o.ApplyT(func(v LookupUserResult) *AsymmetricEncryptedSecretResponse { return v.EncryptedPassword }).(AsymmetricEncryptedSecretResponsePtrOutput)
}

// The path ID that uniquely identifies the object.
func (o LookupUserResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Id }).(pulumi.StringOutput)
}

// The object name.
func (o LookupUserResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Name }).(pulumi.StringOutput)
}

// List of shares that the user has rights on. This field should not be specified during user creation.
func (o LookupUserResultOutput) ShareAccessRights() ShareAccessRightResponseArrayOutput {
	return o.ApplyT(func(v LookupUserResult) []ShareAccessRightResponse { return v.ShareAccessRights }).(ShareAccessRightResponseArrayOutput)
}

// Metadata pertaining to creation and last modification of User
func (o LookupUserResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupUserResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The hierarchical type of the object.
func (o LookupUserResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Type }).(pulumi.StringOutput)
}

// Type of the user.
func (o LookupUserResultOutput) UserType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.UserType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserResultOutput{})
}
