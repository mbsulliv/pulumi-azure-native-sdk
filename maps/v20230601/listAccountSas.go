// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Create and list an account shared access signature token. Use this SAS token for authentication to Azure Maps REST APIs through various Azure Maps SDKs. As prerequisite to create a SAS Token.
//
// Prerequisites:
// 1. Create or have an existing User Assigned Managed Identity in the same Azure region as the account.
// 2. Create or update an Azure Map account with the same Azure region as the User Assigned Managed Identity is placed.
func ListAccountSas(ctx *pulumi.Context, args *ListAccountSasArgs, opts ...pulumi.InvokeOption) (*ListAccountSasResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListAccountSasResult
	err := ctx.Invoke("azure-native:maps/v20230601:listAccountSas", args.Defaults(), &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAccountSasArgs struct {
	// The name of the Maps Account.
	AccountName string `pulumi:"accountName"`
	// The date time offset of when the token validity expires. For example "2017-05-24T10:42:03.1567373Z". Maximum duration allowed is 24 hours between `start` and `expiry`.
	Expiry string `pulumi:"expiry"`
	// Required parameter which represents the desired maximum request per second to allowed for the given SAS token. This does not guarantee perfect accuracy in measurements but provides application safe guards of abuse with eventual enforcement.
	MaxRatePerSecond int `pulumi:"maxRatePerSecond"`
	// The principal Id also known as the object Id of a User Assigned Managed Identity currently assigned to the Map Account. To assign a Managed Identity of the account, use operation Create or Update an assign a User Assigned Identity resource Id.
	PrincipalId string `pulumi:"principalId"`
	// Optional, allows control of which region locations are permitted access to Azure Maps REST APIs with the SAS token. Example: "eastus", "westus2". Omitting this parameter will allow all region locations to be accessible.
	Regions []string `pulumi:"regions"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Map account key to use for signing. Picking `primaryKey` or `secondaryKey` will use the Map account Shared Keys, and using `managedIdentity` will use the auto-renewed private key to sign the SAS.
	SigningKey string `pulumi:"signingKey"`
	// The date time offset of when the token validity begins. For example "2017-05-24T10:42:03.1567373Z". Maximum duration allowed is 24 hours between `start` and `expiry`.
	Start string `pulumi:"start"`
}

// Defaults sets the appropriate defaults for ListAccountSasArgs
func (val *ListAccountSasArgs) Defaults() *ListAccountSasArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if utilities.IsZero(tmp.MaxRatePerSecond) {
		tmp.MaxRatePerSecond = 500
	}
	return &tmp
}

// A new Sas token which can be used to access the Maps REST APIs and is controlled by the specified Managed identity permissions on Azure (IAM) Role Based Access Control.
type ListAccountSasResult struct {
	// The shared access signature access token.
	AccountSasToken string `pulumi:"accountSasToken"`
}

func ListAccountSasOutput(ctx *pulumi.Context, args ListAccountSasOutputArgs, opts ...pulumi.InvokeOption) ListAccountSasResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListAccountSasResult, error) {
			args := v.(ListAccountSasArgs)
			r, err := ListAccountSas(ctx, &args, opts...)
			var s ListAccountSasResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListAccountSasResultOutput)
}

type ListAccountSasOutputArgs struct {
	// The name of the Maps Account.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The date time offset of when the token validity expires. For example "2017-05-24T10:42:03.1567373Z". Maximum duration allowed is 24 hours between `start` and `expiry`.
	Expiry pulumi.StringInput `pulumi:"expiry"`
	// Required parameter which represents the desired maximum request per second to allowed for the given SAS token. This does not guarantee perfect accuracy in measurements but provides application safe guards of abuse with eventual enforcement.
	MaxRatePerSecond pulumi.IntInput `pulumi:"maxRatePerSecond"`
	// The principal Id also known as the object Id of a User Assigned Managed Identity currently assigned to the Map Account. To assign a Managed Identity of the account, use operation Create or Update an assign a User Assigned Identity resource Id.
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
	// Optional, allows control of which region locations are permitted access to Azure Maps REST APIs with the SAS token. Example: "eastus", "westus2". Omitting this parameter will allow all region locations to be accessible.
	Regions pulumi.StringArrayInput `pulumi:"regions"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The Map account key to use for signing. Picking `primaryKey` or `secondaryKey` will use the Map account Shared Keys, and using `managedIdentity` will use the auto-renewed private key to sign the SAS.
	SigningKey pulumi.StringInput `pulumi:"signingKey"`
	// The date time offset of when the token validity begins. For example "2017-05-24T10:42:03.1567373Z". Maximum duration allowed is 24 hours between `start` and `expiry`.
	Start pulumi.StringInput `pulumi:"start"`
}

func (ListAccountSasOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAccountSasArgs)(nil)).Elem()
}

// A new Sas token which can be used to access the Maps REST APIs and is controlled by the specified Managed identity permissions on Azure (IAM) Role Based Access Control.
type ListAccountSasResultOutput struct{ *pulumi.OutputState }

func (ListAccountSasResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAccountSasResult)(nil)).Elem()
}

func (o ListAccountSasResultOutput) ToListAccountSasResultOutput() ListAccountSasResultOutput {
	return o
}

func (o ListAccountSasResultOutput) ToListAccountSasResultOutputWithContext(ctx context.Context) ListAccountSasResultOutput {
	return o
}

func (o ListAccountSasResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListAccountSasResult] {
	return pulumix.Output[ListAccountSasResult]{
		OutputState: o.OutputState,
	}
}

// The shared access signature access token.
func (o ListAccountSasResultOutput) AccountSasToken() pulumi.StringOutput {
	return o.ApplyT(func(v ListAccountSasResult) string { return v.AccountSasToken }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListAccountSasResultOutput{})
}
