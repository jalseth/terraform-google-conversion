// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const CloudFunctionsCloudFunctionIAMAssetType string = "cloudfunctions.googleapis.com/CloudFunction"

func resourceConverterCloudFunctionsCloudFunctionIamPolicy() ResourceConverter {
	return ResourceConverter{
		AssetType:         CloudFunctionsCloudFunctionIAMAssetType,
		Convert:           GetCloudFunctionsCloudFunctionIamPolicyCaiObject,
		MergeCreateUpdate: MergeCloudFunctionsCloudFunctionIamPolicy,
	}
}

func resourceConverterCloudFunctionsCloudFunctionIamBinding() ResourceConverter {
	return ResourceConverter{
		AssetType:         CloudFunctionsCloudFunctionIAMAssetType,
		Convert:           GetCloudFunctionsCloudFunctionIamBindingCaiObject,
		FetchFullResource: FetchCloudFunctionsCloudFunctionIamPolicy,
		MergeCreateUpdate: MergeCloudFunctionsCloudFunctionIamBinding,
		MergeDelete:       MergeCloudFunctionsCloudFunctionIamBindingDelete,
	}
}

func resourceConverterCloudFunctionsCloudFunctionIamMember() ResourceConverter {
	return ResourceConverter{
		AssetType:         CloudFunctionsCloudFunctionIAMAssetType,
		Convert:           GetCloudFunctionsCloudFunctionIamMemberCaiObject,
		FetchFullResource: FetchCloudFunctionsCloudFunctionIamPolicy,
		MergeCreateUpdate: MergeCloudFunctionsCloudFunctionIamMember,
		MergeDelete:       MergeCloudFunctionsCloudFunctionIamMemberDelete,
	}
}

func GetCloudFunctionsCloudFunctionIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	return newCloudFunctionsCloudFunctionIamAsset(d, config, expandIamPolicyBindings)
}

func GetCloudFunctionsCloudFunctionIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	return newCloudFunctionsCloudFunctionIamAsset(d, config, expandIamRoleBindings)
}

func GetCloudFunctionsCloudFunctionIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	return newCloudFunctionsCloudFunctionIamAsset(d, config, expandIamMemberBindings)
}

func MergeCloudFunctionsCloudFunctionIamPolicy(existing, incoming Asset) Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeCloudFunctionsCloudFunctionIamBinding(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAuthoritativeBindings)
}

func MergeCloudFunctionsCloudFunctionIamBindingDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAuthoritativeBindings)
}

func MergeCloudFunctionsCloudFunctionIamMember(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAdditiveBindings)
}

func MergeCloudFunctionsCloudFunctionIamMemberDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAdditiveBindings)
}

func newCloudFunctionsCloudFunctionIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]IAMBinding, error),
) ([]Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := assetName(d, config, "//cloudfunctions.googleapis.com/projects/{{project}}/locations/{{region}}/functions/{{cloud_function}}")
	if err != nil {
		return []Asset{}, err
	}

	return []Asset{{
		Name: name,
		Type: CloudFunctionsCloudFunctionIAMAssetType,
		IAMPolicy: &IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchCloudFunctionsCloudFunctionIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("region"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("cloud_function"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}

	return fetchIamPolicy(
		CloudFunctionsCloudFunctionIamUpdaterProducer,
		d,
		config,
		"//cloudfunctions.googleapis.com/projects/{{project}}/locations/{{region}}/functions/{{cloud_function}}",
		CloudFunctionsCloudFunctionIAMAssetType,
	)
}
