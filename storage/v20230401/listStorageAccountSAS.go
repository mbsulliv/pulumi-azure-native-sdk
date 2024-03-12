// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// List SAS credentials of a storage account.
func ListStorageAccountSAS(ctx *pulumi.Context, args *ListStorageAccountSASArgs, opts ...pulumi.InvokeOption) (*ListStorageAccountSASResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListStorageAccountSASResult
	err := ctx.Invoke("azure-native:storage/v20230401:listStorageAccountSAS", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListStorageAccountSASArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName string `pulumi:"accountName"`
	// An IP address or a range of IP addresses from which to accept requests.
	IPAddressOrRange *string `pulumi:"iPAddressOrRange"`
	// The key to sign the account SAS token with.
	KeyToSign *string `pulumi:"keyToSign"`
	// The signed permissions for the account SAS. Possible values include: Read (r), Write (w), Delete (d), List (l), Add (a), Create (c), Update (u) and Process (p).
	Permissions string `pulumi:"permissions"`
	// The protocol permitted for a request made with the account SAS.
	Protocols *HttpProtocol `pulumi:"protocols"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The signed resource types that are accessible with the account SAS. Service (s): Access to service-level APIs; Container (c): Access to container-level APIs; Object (o): Access to object-level APIs for blobs, queue messages, table entities, and files.
	ResourceTypes string `pulumi:"resourceTypes"`
	// The signed services accessible with the account SAS. Possible values include: Blob (b), Queue (q), Table (t), File (f).
	Services string `pulumi:"services"`
	// The time at which the shared access signature becomes invalid.
	SharedAccessExpiryTime string `pulumi:"sharedAccessExpiryTime"`
	// The time at which the SAS becomes valid.
	SharedAccessStartTime *string `pulumi:"sharedAccessStartTime"`
}

// The List SAS credentials operation response.
type ListStorageAccountSASResult struct {
	// List SAS credentials of storage account.
	AccountSasToken string `pulumi:"accountSasToken"`
}

func ListStorageAccountSASOutput(ctx *pulumi.Context, args ListStorageAccountSASOutputArgs, opts ...pulumi.InvokeOption) ListStorageAccountSASResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListStorageAccountSASResult, error) {
			args := v.(ListStorageAccountSASArgs)
			r, err := ListStorageAccountSAS(ctx, &args, opts...)
			var s ListStorageAccountSASResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListStorageAccountSASResultOutput)
}

type ListStorageAccountSASOutputArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// An IP address or a range of IP addresses from which to accept requests.
	IPAddressOrRange pulumi.StringPtrInput `pulumi:"iPAddressOrRange"`
	// The key to sign the account SAS token with.
	KeyToSign pulumi.StringPtrInput `pulumi:"keyToSign"`
	// The signed permissions for the account SAS. Possible values include: Read (r), Write (w), Delete (d), List (l), Add (a), Create (c), Update (u) and Process (p).
	Permissions pulumi.StringInput `pulumi:"permissions"`
	// The protocol permitted for a request made with the account SAS.
	Protocols HttpProtocolPtrInput `pulumi:"protocols"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The signed resource types that are accessible with the account SAS. Service (s): Access to service-level APIs; Container (c): Access to container-level APIs; Object (o): Access to object-level APIs for blobs, queue messages, table entities, and files.
	ResourceTypes pulumi.StringInput `pulumi:"resourceTypes"`
	// The signed services accessible with the account SAS. Possible values include: Blob (b), Queue (q), Table (t), File (f).
	Services pulumi.StringInput `pulumi:"services"`
	// The time at which the shared access signature becomes invalid.
	SharedAccessExpiryTime pulumi.StringInput `pulumi:"sharedAccessExpiryTime"`
	// The time at which the SAS becomes valid.
	SharedAccessStartTime pulumi.StringPtrInput `pulumi:"sharedAccessStartTime"`
}

func (ListStorageAccountSASOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStorageAccountSASArgs)(nil)).Elem()
}

// The List SAS credentials operation response.
type ListStorageAccountSASResultOutput struct{ *pulumi.OutputState }

func (ListStorageAccountSASResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStorageAccountSASResult)(nil)).Elem()
}

func (o ListStorageAccountSASResultOutput) ToListStorageAccountSASResultOutput() ListStorageAccountSASResultOutput {
	return o
}

func (o ListStorageAccountSASResultOutput) ToListStorageAccountSASResultOutputWithContext(ctx context.Context) ListStorageAccountSASResultOutput {
	return o
}

// List SAS credentials of storage account.
func (o ListStorageAccountSASResultOutput) AccountSasToken() pulumi.StringOutput {
	return o.ApplyT(func(v ListStorageAccountSASResult) string { return v.AccountSasToken }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListStorageAccountSASResultOutput{})
}
