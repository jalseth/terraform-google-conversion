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

package dialogflow

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const DialogflowIntentAssetType string = "dialogflow.googleapis.com/Intent"

func ResourceConverterDialogflowIntent() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: DialogflowIntentAssetType,
		Convert:   GetDialogflowIntentCaiObject,
	}
}

func GetDialogflowIntentCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//dialogflow.googleapis.com/{{name}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetDialogflowIntentApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: DialogflowIntentAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "v2",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/dialogflow/v2/rest",
				DiscoveryName:        "Intent",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetDialogflowIntentApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandDialogflowIntentDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	webhookStateProp, err := expandDialogflowIntentWebhookState(d.Get("webhook_state"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("webhook_state"); !tpgresource.IsEmptyValue(reflect.ValueOf(webhookStateProp)) && (ok || !reflect.DeepEqual(v, webhookStateProp)) {
		obj["webhookState"] = webhookStateProp
	}
	priorityProp, err := expandDialogflowIntentPriority(d.Get("priority"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("priority"); !tpgresource.IsEmptyValue(reflect.ValueOf(priorityProp)) && (ok || !reflect.DeepEqual(v, priorityProp)) {
		obj["priority"] = priorityProp
	}
	isFallbackProp, err := expandDialogflowIntentIsFallback(d.Get("is_fallback"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("is_fallback"); !tpgresource.IsEmptyValue(reflect.ValueOf(isFallbackProp)) && (ok || !reflect.DeepEqual(v, isFallbackProp)) {
		obj["isFallback"] = isFallbackProp
	}
	mlDisabledProp, err := expandDialogflowIntentMlDisabled(d.Get("ml_disabled"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ml_disabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(mlDisabledProp)) && (ok || !reflect.DeepEqual(v, mlDisabledProp)) {
		obj["mlDisabled"] = mlDisabledProp
	}
	inputContextNamesProp, err := expandDialogflowIntentInputContextNames(d.Get("input_context_names"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("input_context_names"); !tpgresource.IsEmptyValue(reflect.ValueOf(inputContextNamesProp)) && (ok || !reflect.DeepEqual(v, inputContextNamesProp)) {
		obj["inputContextNames"] = inputContextNamesProp
	}
	eventsProp, err := expandDialogflowIntentEvents(d.Get("events"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("events"); !tpgresource.IsEmptyValue(reflect.ValueOf(eventsProp)) && (ok || !reflect.DeepEqual(v, eventsProp)) {
		obj["events"] = eventsProp
	}
	actionProp, err := expandDialogflowIntentAction(d.Get("action"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("action"); !tpgresource.IsEmptyValue(reflect.ValueOf(actionProp)) && (ok || !reflect.DeepEqual(v, actionProp)) {
		obj["action"] = actionProp
	}
	resetContextsProp, err := expandDialogflowIntentResetContexts(d.Get("reset_contexts"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("reset_contexts"); !tpgresource.IsEmptyValue(reflect.ValueOf(resetContextsProp)) && (ok || !reflect.DeepEqual(v, resetContextsProp)) {
		obj["resetContexts"] = resetContextsProp
	}
	defaultResponsePlatformsProp, err := expandDialogflowIntentDefaultResponsePlatforms(d.Get("default_response_platforms"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("default_response_platforms"); !tpgresource.IsEmptyValue(reflect.ValueOf(defaultResponsePlatformsProp)) && (ok || !reflect.DeepEqual(v, defaultResponsePlatformsProp)) {
		obj["defaultResponsePlatforms"] = defaultResponsePlatformsProp
	}
	parentFollowupIntentNameProp, err := expandDialogflowIntentParentFollowupIntentName(d.Get("parent_followup_intent_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("parent_followup_intent_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(parentFollowupIntentNameProp)) && (ok || !reflect.DeepEqual(v, parentFollowupIntentNameProp)) {
		obj["parentFollowupIntentName"] = parentFollowupIntentNameProp
	}

	return obj, nil
}

func expandDialogflowIntentDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowIntentWebhookState(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowIntentPriority(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowIntentIsFallback(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowIntentMlDisabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowIntentInputContextNames(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowIntentEvents(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowIntentAction(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowIntentResetContexts(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowIntentDefaultResponsePlatforms(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowIntentParentFollowupIntentName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}