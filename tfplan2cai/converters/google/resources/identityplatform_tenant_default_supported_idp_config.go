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
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const IdentityPlatformTenantDefaultSupportedIdpConfigAssetType string = "identitytoolkit.googleapis.com/TenantDefaultSupportedIdpConfig"

func resourceConverterIdentityPlatformTenantDefaultSupportedIdpConfig() ResourceConverter {
	return ResourceConverter{
		AssetType: IdentityPlatformTenantDefaultSupportedIdpConfigAssetType,
		Convert:   GetIdentityPlatformTenantDefaultSupportedIdpConfigCaiObject,
	}
}

func GetIdentityPlatformTenantDefaultSupportedIdpConfigCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	name, err := assetName(d, config, "//identitytoolkit.googleapis.com/projects/{{project}}/tenants/{{tenant}}/defaultSupportedIdpConfigs/{{idp_id}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetIdentityPlatformTenantDefaultSupportedIdpConfigApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: IdentityPlatformTenantDefaultSupportedIdpConfigAssetType,
			Resource: &AssetResource{
				Version:              "v2",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/identitytoolkit/v2/rest",
				DiscoveryName:        "TenantDefaultSupportedIdpConfig",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetIdentityPlatformTenantDefaultSupportedIdpConfigApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	clientIdProp, err := expandIdentityPlatformTenantDefaultSupportedIdpConfigClientId(d.Get("client_id"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("client_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(clientIdProp)) && (ok || !reflect.DeepEqual(v, clientIdProp)) {
		obj["clientId"] = clientIdProp
	}
	clientSecretProp, err := expandIdentityPlatformTenantDefaultSupportedIdpConfigClientSecret(d.Get("client_secret"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("client_secret"); !tpgresource.IsEmptyValue(reflect.ValueOf(clientSecretProp)) && (ok || !reflect.DeepEqual(v, clientSecretProp)) {
		obj["clientSecret"] = clientSecretProp
	}
	enabledProp, err := expandIdentityPlatformTenantDefaultSupportedIdpConfigEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(enabledProp)) && (ok || !reflect.DeepEqual(v, enabledProp)) {
		obj["enabled"] = enabledProp
	}

	return obj, nil
}

func expandIdentityPlatformTenantDefaultSupportedIdpConfigClientId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformTenantDefaultSupportedIdpConfigClientSecret(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformTenantDefaultSupportedIdpConfigEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
