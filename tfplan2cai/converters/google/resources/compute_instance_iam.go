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
const ComputeInstanceIAMAssetType string = "compute.googleapis.com/Instance"

func resourceConverterComputeInstanceIamPolicy() ResourceConverter {
	return ResourceConverter{
		AssetType:         ComputeInstanceIAMAssetType,
		Convert:           GetComputeInstanceIamPolicyCaiObject,
		MergeCreateUpdate: MergeComputeInstanceIamPolicy,
	}
}

func resourceConverterComputeInstanceIamBinding() ResourceConverter {
	return ResourceConverter{
		AssetType:         ComputeInstanceIAMAssetType,
		Convert:           GetComputeInstanceIamBindingCaiObject,
		FetchFullResource: FetchComputeInstanceIamPolicy,
		MergeCreateUpdate: MergeComputeInstanceIamBinding,
		MergeDelete:       MergeComputeInstanceIamBindingDelete,
	}
}

func resourceConverterComputeInstanceIamMember() ResourceConverter {
	return ResourceConverter{
		AssetType:         ComputeInstanceIAMAssetType,
		Convert:           GetComputeInstanceIamMemberCaiObject,
		FetchFullResource: FetchComputeInstanceIamPolicy,
		MergeCreateUpdate: MergeComputeInstanceIamMember,
		MergeDelete:       MergeComputeInstanceIamMemberDelete,
	}
}

func GetComputeInstanceIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	return newComputeInstanceIamAsset(d, config, expandIamPolicyBindings)
}

func GetComputeInstanceIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	return newComputeInstanceIamAsset(d, config, expandIamRoleBindings)
}

func GetComputeInstanceIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	return newComputeInstanceIamAsset(d, config, expandIamMemberBindings)
}

func MergeComputeInstanceIamPolicy(existing, incoming Asset) Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeComputeInstanceIamBinding(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAuthoritativeBindings)
}

func MergeComputeInstanceIamBindingDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAuthoritativeBindings)
}

func MergeComputeInstanceIamMember(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAdditiveBindings)
}

func MergeComputeInstanceIamMemberDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAdditiveBindings)
}

func newComputeInstanceIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]IAMBinding, error),
) ([]Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := assetName(d, config, "//compute.googleapis.com/projects/{{project}}/zones/{{zone}}/instances/{{instance_name}}")
	if err != nil {
		return []Asset{}, err
	}

	return []Asset{{
		Name: name,
		Type: ComputeInstanceIAMAssetType,
		IAMPolicy: &IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchComputeInstanceIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("zone"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("instance_name"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}

	return fetchIamPolicy(
		ComputeInstanceIamUpdaterProducer,
		d,
		config,
		"//compute.googleapis.com/projects/{{project}}/zones/{{zone}}/instances/{{instance_name}}",
		ComputeInstanceIAMAssetType,
	)
}
