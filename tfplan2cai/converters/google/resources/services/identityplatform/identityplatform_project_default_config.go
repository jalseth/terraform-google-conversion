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

package identityplatform

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const IdentityPlatformProjectDefaultConfigAssetType string = "identitytoolkit.googleapis.com/ProjectDefaultConfig"

func ResourceConverterIdentityPlatformProjectDefaultConfig() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: IdentityPlatformProjectDefaultConfigAssetType,
		Convert:   GetIdentityPlatformProjectDefaultConfigCaiObject,
	}
}

func GetIdentityPlatformProjectDefaultConfigCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//identitytoolkit.googleapis.com/projects/{{project}}/config")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetIdentityPlatformProjectDefaultConfigApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: IdentityPlatformProjectDefaultConfigAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "v2",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/identitytoolkit/v2/rest",
				DiscoveryName:        "ProjectDefaultConfig",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetIdentityPlatformProjectDefaultConfigApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	signInProp, err := expandIdentityPlatformProjectDefaultConfigSignIn(d.Get("sign_in"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("sign_in"); !tpgresource.IsEmptyValue(reflect.ValueOf(signInProp)) && (ok || !reflect.DeepEqual(v, signInProp)) {
		obj["signIn"] = signInProp
	}

	return obj, nil
}

func expandIdentityPlatformProjectDefaultConfigSignIn(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEmail, err := expandIdentityPlatformProjectDefaultConfigSignInEmail(original["email"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEmail); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["email"] = transformedEmail
	}

	transformedPhoneNumber, err := expandIdentityPlatformProjectDefaultConfigSignInPhoneNumber(original["phone_number"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPhoneNumber); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["phoneNumber"] = transformedPhoneNumber
	}

	transformedAnonymous, err := expandIdentityPlatformProjectDefaultConfigSignInAnonymous(original["anonymous"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAnonymous); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["anonymous"] = transformedAnonymous
	}

	transformedAllowDuplicateEmails, err := expandIdentityPlatformProjectDefaultConfigSignInAllowDuplicateEmails(original["allow_duplicate_emails"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowDuplicateEmails); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["allowDuplicateEmails"] = transformedAllowDuplicateEmails
	}

	transformedHashConfig, err := expandIdentityPlatformProjectDefaultConfigSignInHashConfig(original["hash_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHashConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["hashConfig"] = transformedHashConfig
	}

	return transformed, nil
}

func expandIdentityPlatformProjectDefaultConfigSignInEmail(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnabled, err := expandIdentityPlatformProjectDefaultConfigSignInEmailEnabled(original["enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enabled"] = transformedEnabled
	}

	transformedPasswordRequired, err := expandIdentityPlatformProjectDefaultConfigSignInEmailPasswordRequired(original["password_required"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPasswordRequired); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["passwordRequired"] = transformedPasswordRequired
	}

	return transformed, nil
}

func expandIdentityPlatformProjectDefaultConfigSignInEmailEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformProjectDefaultConfigSignInEmailPasswordRequired(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformProjectDefaultConfigSignInPhoneNumber(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnabled, err := expandIdentityPlatformProjectDefaultConfigSignInPhoneNumberEnabled(original["enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enabled"] = transformedEnabled
	}

	transformedTestPhoneNumbers, err := expandIdentityPlatformProjectDefaultConfigSignInPhoneNumberTestPhoneNumbers(original["test_phone_numbers"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTestPhoneNumbers); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["testPhoneNumbers"] = transformedTestPhoneNumbers
	}

	return transformed, nil
}

func expandIdentityPlatformProjectDefaultConfigSignInPhoneNumberEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformProjectDefaultConfigSignInPhoneNumberTestPhoneNumbers(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandIdentityPlatformProjectDefaultConfigSignInAnonymous(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnabled, err := expandIdentityPlatformProjectDefaultConfigSignInAnonymousEnabled(original["enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enabled"] = transformedEnabled
	}

	return transformed, nil
}

func expandIdentityPlatformProjectDefaultConfigSignInAnonymousEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformProjectDefaultConfigSignInAllowDuplicateEmails(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformProjectDefaultConfigSignInHashConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAlgorithm, err := expandIdentityPlatformProjectDefaultConfigSignInHashConfigAlgorithm(original["algorithm"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAlgorithm); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["algorithm"] = transformedAlgorithm
	}

	transformedSignerKey, err := expandIdentityPlatformProjectDefaultConfigSignInHashConfigSignerKey(original["signer_key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSignerKey); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["signerKey"] = transformedSignerKey
	}

	transformedSaltSeparator, err := expandIdentityPlatformProjectDefaultConfigSignInHashConfigSaltSeparator(original["salt_separator"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSaltSeparator); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["saltSeparator"] = transformedSaltSeparator
	}

	transformedRounds, err := expandIdentityPlatformProjectDefaultConfigSignInHashConfigRounds(original["rounds"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRounds); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["rounds"] = transformedRounds
	}

	transformedMemoryCost, err := expandIdentityPlatformProjectDefaultConfigSignInHashConfigMemoryCost(original["memory_cost"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMemoryCost); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["memoryCost"] = transformedMemoryCost
	}

	return transformed, nil
}

func expandIdentityPlatformProjectDefaultConfigSignInHashConfigAlgorithm(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformProjectDefaultConfigSignInHashConfigSignerKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformProjectDefaultConfigSignInHashConfigSaltSeparator(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformProjectDefaultConfigSignInHashConfigRounds(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformProjectDefaultConfigSignInHashConfigMemoryCost(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}