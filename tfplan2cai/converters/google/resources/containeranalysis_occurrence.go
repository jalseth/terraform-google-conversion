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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const ContainerAnalysisOccurrenceAssetType string = "containeranalysis.googleapis.com/Occurrence"

func resourceConverterContainerAnalysisOccurrence() ResourceConverter {
	return ResourceConverter{
		AssetType: ContainerAnalysisOccurrenceAssetType,
		Convert:   GetContainerAnalysisOccurrenceCaiObject,
	}
}

func GetContainerAnalysisOccurrenceCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	name, err := assetName(d, config, "//containeranalysis.googleapis.com/projects/{{project}}/occurrences/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetContainerAnalysisOccurrenceApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: ContainerAnalysisOccurrenceAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/containeranalysis/v1/rest",
				DiscoveryName:        "Occurrence",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetContainerAnalysisOccurrenceApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	resourceUriProp, err := expandContainerAnalysisOccurrenceResourceUri(d.Get("resource_uri"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("resource_uri"); !tpgresource.IsEmptyValue(reflect.ValueOf(resourceUriProp)) && (ok || !reflect.DeepEqual(v, resourceUriProp)) {
		obj["resourceUri"] = resourceUriProp
	}
	noteNameProp, err := expandContainerAnalysisOccurrenceNoteName(d.Get("note_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("note_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(noteNameProp)) && (ok || !reflect.DeepEqual(v, noteNameProp)) {
		obj["noteName"] = noteNameProp
	}
	remediationProp, err := expandContainerAnalysisOccurrenceRemediation(d.Get("remediation"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("remediation"); !tpgresource.IsEmptyValue(reflect.ValueOf(remediationProp)) && (ok || !reflect.DeepEqual(v, remediationProp)) {
		obj["remediation"] = remediationProp
	}
	attestationProp, err := expandContainerAnalysisOccurrenceAttestation(d.Get("attestation"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("attestation"); !tpgresource.IsEmptyValue(reflect.ValueOf(attestationProp)) && (ok || !reflect.DeepEqual(v, attestationProp)) {
		obj["attestation"] = attestationProp
	}

	return resourceContainerAnalysisOccurrenceEncoder(d, config, obj)
}

func resourceContainerAnalysisOccurrenceEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// encoder logic only in non-GA versions

	return obj, nil
}

func expandContainerAnalysisOccurrenceResourceUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandContainerAnalysisOccurrenceNoteName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandContainerAnalysisOccurrenceRemediation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandContainerAnalysisOccurrenceAttestation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSerializedPayload, err := expandContainerAnalysisOccurrenceAttestationSerializedPayload(original["serialized_payload"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSerializedPayload); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["serializedPayload"] = transformedSerializedPayload
	}

	transformedSignatures, err := expandContainerAnalysisOccurrenceAttestationSignatures(original["signatures"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSignatures); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["signatures"] = transformedSignatures
	}

	return transformed, nil
}

func expandContainerAnalysisOccurrenceAttestationSerializedPayload(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandContainerAnalysisOccurrenceAttestationSignatures(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedSignature, err := expandContainerAnalysisOccurrenceAttestationSignaturesSignature(original["signature"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSignature); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["signature"] = transformedSignature
		}

		transformedPublicKeyId, err := expandContainerAnalysisOccurrenceAttestationSignaturesPublicKeyId(original["public_key_id"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPublicKeyId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["publicKeyId"] = transformedPublicKeyId
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandContainerAnalysisOccurrenceAttestationSignaturesSignature(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandContainerAnalysisOccurrenceAttestationSignaturesPublicKeyId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
