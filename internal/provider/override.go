// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
	"github.com/hashicorp/terraform-provider-aws/names"
)

var (
	overrideProviderSchema = &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"default_tags": { // TODO: Only for aws_s3_object?
					Type:     schema.TypeList,
					Optional: true,
					MaxItems: 1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							names.AttrTags: {
								Type:        schema.TypeMap,
								Optional:    true,
								Elem:        &schema.Schema{Type: schema.TypeString},
								Description: `Resource tags.`,
							},
						},
					},
					Description: `Per-resource default tags override.`,
				},
				names.AttrRegion: {
					Type:         schema.TypeString,
					Optional:     true,
					Computed:     true,
					ForceNew:     true,
					ValidateFunc: verify.ValidRegionName, // TODO: Valid for partition
					Description:  `Per-resource Region override.`,
				},
			},
		},
		Description: `Configuration block to override the provider configuration.`,
	}
)

type overrideProviderModel struct {
	DefaultTagsConfig *tftags.DefaultConfig
	Region            *string
}

func expandOverrideProviderModel(ctx context.Context, tfMap map[string]interface{}) *overrideProviderModel {
	if tfMap == nil {
		return nil
	}

	data := &overrideProviderModel{}

	if v, ok := tfMap["default_tags"].([]interface{}); ok && len(v) > 0 {
		if v[0] != nil {
			data.DefaultTagsConfig = expandOverrideProviderDefaultTags(ctx, v[0].(map[string]interface{}))
		} else {
			data.DefaultTagsConfig = expandOverrideProviderDefaultTags(ctx, map[string]interface{}{})
		}
	}

	return data
}

func expandOverrideProviderDefaultTags(ctx context.Context, tfMap map[string]interface{}) *tftags.DefaultConfig {
	if tfMap == nil {
		return nil
	}

	data := &tftags.DefaultConfig{}

	if v, ok := tfMap[names.AttrTags].(map[string]interface{}); ok {
		data.Tags = tftags.New(ctx, v)
	}

	return data
}

func providerRegionCustomizeDiffFunc(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
	return d.SetNew("override_provider", []interface{}{map[string]interface{}{names.AttrRegion: "us-east-2"}})
}
