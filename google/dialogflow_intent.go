// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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

import "reflect"

func GetDialogflowIntentCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//dialogflow.googleapis.com/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetDialogflowIntentApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: "dialogflow.googleapis.com/Intent",
			Resource: &AssetResource{
				Version:              "v2",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/dialogflow/v2/rest",
				DiscoveryName:        "Intent",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetDialogflowIntentApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandDialogflowIntentDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	webhookStateProp, err := expandDialogflowIntentWebhookState(d.Get("webhook_state"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("webhook_state"); !isEmptyValue(reflect.ValueOf(webhookStateProp)) && (ok || !reflect.DeepEqual(v, webhookStateProp)) {
		obj["webhookState"] = webhookStateProp
	}
	priorityProp, err := expandDialogflowIntentPriority(d.Get("priority"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("priority"); !isEmptyValue(reflect.ValueOf(priorityProp)) && (ok || !reflect.DeepEqual(v, priorityProp)) {
		obj["priority"] = priorityProp
	}
	isFallbackProp, err := expandDialogflowIntentIsFallback(d.Get("is_fallback"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("is_fallback"); !isEmptyValue(reflect.ValueOf(isFallbackProp)) && (ok || !reflect.DeepEqual(v, isFallbackProp)) {
		obj["isFallback"] = isFallbackProp
	}
	mlDisabledProp, err := expandDialogflowIntentMlDisabled(d.Get("ml_disabled"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ml_disabled"); !isEmptyValue(reflect.ValueOf(mlDisabledProp)) && (ok || !reflect.DeepEqual(v, mlDisabledProp)) {
		obj["mlDisabled"] = mlDisabledProp
	}
	inputContextNamesProp, err := expandDialogflowIntentInputContextNames(d.Get("input_context_names"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("input_context_names"); !isEmptyValue(reflect.ValueOf(inputContextNamesProp)) && (ok || !reflect.DeepEqual(v, inputContextNamesProp)) {
		obj["inputContextNames"] = inputContextNamesProp
	}
	eventsProp, err := expandDialogflowIntentEvents(d.Get("events"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("events"); !isEmptyValue(reflect.ValueOf(eventsProp)) && (ok || !reflect.DeepEqual(v, eventsProp)) {
		obj["events"] = eventsProp
	}
	actionProp, err := expandDialogflowIntentAction(d.Get("action"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("action"); !isEmptyValue(reflect.ValueOf(actionProp)) && (ok || !reflect.DeepEqual(v, actionProp)) {
		obj["action"] = actionProp
	}
	resetContextsProp, err := expandDialogflowIntentResetContexts(d.Get("reset_contexts"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("reset_contexts"); !isEmptyValue(reflect.ValueOf(resetContextsProp)) && (ok || !reflect.DeepEqual(v, resetContextsProp)) {
		obj["resetContexts"] = resetContextsProp
	}
	defaultResponsePlatformsProp, err := expandDialogflowIntentDefaultResponsePlatforms(d.Get("default_response_platforms"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("default_response_platforms"); !isEmptyValue(reflect.ValueOf(defaultResponsePlatformsProp)) && (ok || !reflect.DeepEqual(v, defaultResponsePlatformsProp)) {
		obj["defaultResponsePlatforms"] = defaultResponsePlatformsProp
	}
	parentFollowupIntentNameProp, err := expandDialogflowIntentParentFollowupIntentName(d.Get("parent_followup_intent_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("parent_followup_intent_name"); !isEmptyValue(reflect.ValueOf(parentFollowupIntentNameProp)) && (ok || !reflect.DeepEqual(v, parentFollowupIntentNameProp)) {
		obj["parentFollowupIntentName"] = parentFollowupIntentNameProp
	}

	return obj, nil
}

func expandDialogflowIntentDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowIntentWebhookState(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowIntentPriority(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowIntentIsFallback(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowIntentMlDisabled(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowIntentInputContextNames(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowIntentEvents(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowIntentAction(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowIntentResetContexts(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowIntentDefaultResponsePlatforms(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowIntentParentFollowupIntentName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
