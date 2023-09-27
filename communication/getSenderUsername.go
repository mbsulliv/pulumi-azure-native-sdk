// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package communication

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get a valid sender username for a domains resource.
// Azure REST API version: 2023-03-31.
func LookupSenderUsername(ctx *pulumi.Context, args *LookupSenderUsernameArgs, opts ...pulumi.InvokeOption) (*LookupSenderUsernameResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSenderUsernameResult
	err := ctx.Invoke("azure-native:communication:getSenderUsername", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSenderUsernameArgs struct {
	// The name of the Domains resource.
	DomainName string `pulumi:"domainName"`
	// The name of the EmailService resource.
	EmailServiceName string `pulumi:"emailServiceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The valid sender Username.
	SenderUsername string `pulumi:"senderUsername"`
}

// A class representing a SenderUsername resource.
type LookupSenderUsernameResult struct {
	// The location where the SenderUsername resource data is stored at rest.
	DataLocation string `pulumi:"dataLocation"`
	// The display name for the senderUsername.
	DisplayName *string `pulumi:"displayName"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Provisioning state of the resource. Unknown is the default state for Communication Services.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// A sender senderUsername to be used when sending emails.
	Username string `pulumi:"username"`
}

func LookupSenderUsernameOutput(ctx *pulumi.Context, args LookupSenderUsernameOutputArgs, opts ...pulumi.InvokeOption) LookupSenderUsernameResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSenderUsernameResult, error) {
			args := v.(LookupSenderUsernameArgs)
			r, err := LookupSenderUsername(ctx, &args, opts...)
			var s LookupSenderUsernameResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSenderUsernameResultOutput)
}

type LookupSenderUsernameOutputArgs struct {
	// The name of the Domains resource.
	DomainName pulumi.StringInput `pulumi:"domainName"`
	// The name of the EmailService resource.
	EmailServiceName pulumi.StringInput `pulumi:"emailServiceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The valid sender Username.
	SenderUsername pulumi.StringInput `pulumi:"senderUsername"`
}

func (LookupSenderUsernameOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSenderUsernameArgs)(nil)).Elem()
}

// A class representing a SenderUsername resource.
type LookupSenderUsernameResultOutput struct{ *pulumi.OutputState }

func (LookupSenderUsernameResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSenderUsernameResult)(nil)).Elem()
}

func (o LookupSenderUsernameResultOutput) ToLookupSenderUsernameResultOutput() LookupSenderUsernameResultOutput {
	return o
}

func (o LookupSenderUsernameResultOutput) ToLookupSenderUsernameResultOutputWithContext(ctx context.Context) LookupSenderUsernameResultOutput {
	return o
}

func (o LookupSenderUsernameResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupSenderUsernameResult] {
	return pulumix.Output[LookupSenderUsernameResult]{
		OutputState: o.OutputState,
	}
}

// The location where the SenderUsername resource data is stored at rest.
func (o LookupSenderUsernameResultOutput) DataLocation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSenderUsernameResult) string { return v.DataLocation }).(pulumi.StringOutput)
}

// The display name for the senderUsername.
func (o LookupSenderUsernameResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSenderUsernameResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupSenderUsernameResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSenderUsernameResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupSenderUsernameResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSenderUsernameResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the resource. Unknown is the default state for Communication Services.
func (o LookupSenderUsernameResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSenderUsernameResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupSenderUsernameResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSenderUsernameResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupSenderUsernameResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSenderUsernameResult) string { return v.Type }).(pulumi.StringOutput)
}

// A sender senderUsername to be used when sending emails.
func (o LookupSenderUsernameResultOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSenderUsernameResult) string { return v.Username }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSenderUsernameResultOutput{})
}
