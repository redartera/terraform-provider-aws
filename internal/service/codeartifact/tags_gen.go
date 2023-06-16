// Code generated by internal/generate/tags/main.go; DO NOT EDIT.
package codeartifact

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/codeartifact"
	"github.com/aws/aws-sdk-go/service/codeartifact/codeartifactiface"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// listTags lists codeartifact service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func listTags(ctx context.Context, conn codeartifactiface.CodeArtifactAPI, identifier string) (tftags.KeyValueTags, error) {
	input := &codeartifact.ListTagsForResourceInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTagsForResourceWithContext(ctx, input)

	if err != nil {
		return tftags.New(ctx, nil), err
	}

	return KeyValueTags(ctx, output.Tags), nil
}

// ListTags lists codeartifact service tags and set them in Context.
// It is called from outside this package.
func (p *servicePackage) ListTags(ctx context.Context, meta any, identifier string) error {
	tags, err := listTags(ctx, meta.(*conns.AWSClient).CodeArtifactConn(ctx), identifier)

	if err != nil {
		return err
	}

	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = types.Some(tags)
	}

	return nil
}

// []*SERVICE.Tag handling

// Tags returns codeartifact service tags.
func Tags(tags tftags.KeyValueTags) []*codeartifact.Tag {
	result := make([]*codeartifact.Tag, 0, len(tags))

	for k, v := range tags.Map() {
		tag := &codeartifact.Tag{
			Key:   aws.String(k),
			Value: aws.String(v),
		}

		result = append(result, tag)
	}

	return result
}

// KeyValueTags creates tftags.KeyValueTags from codeartifact service tags.
func KeyValueTags(ctx context.Context, tags []*codeartifact.Tag) tftags.KeyValueTags {
	m := make(map[string]*string, len(tags))

	for _, tag := range tags {
		m[aws.StringValue(tag.Key)] = tag.Value
	}

	return tftags.New(ctx, m)
}

// getTagsIn returns codeartifact service tags from Context.
// nil is returned if there are no input tags.
func getTagsIn(ctx context.Context) []*codeartifact.Tag {
	if inContext, ok := tftags.FromContext(ctx); ok {
		if tags := Tags(inContext.TagsIn.UnwrapOrDefault()); len(tags) > 0 {
			return tags
		}
	}

	return nil
}

// setTagsOut sets codeartifact service tags in Context.
func setTagsOut(ctx context.Context, tags []*codeartifact.Tag) {
	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = types.Some(KeyValueTags(ctx, tags))
	}
}

// updateTags updates codeartifact service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func updateTags(ctx context.Context, conn codeartifactiface.CodeArtifactAPI, identifier string, oldTagsMap, newTagsMap any) error {
	oldTags := tftags.New(ctx, oldTagsMap)
	newTags := tftags.New(ctx, newTagsMap)

	removedTags := oldTags.Removed(newTags)
	removedTags = removedTags.IgnoreSystem(names.CodeArtifact)
	if len(removedTags) > 0 {
		input := &codeartifact.UntagResourceInput{
			ResourceArn: aws.String(identifier),
			TagKeys:     aws.StringSlice(removedTags.Keys()),
		}

		_, err := conn.UntagResourceWithContext(ctx, input)

		if err != nil {
			return fmt.Errorf("untagging resource (%s): %w", identifier, err)
		}
	}

	updatedTags := oldTags.Updated(newTags)
	updatedTags = updatedTags.IgnoreSystem(names.CodeArtifact)
	if len(updatedTags) > 0 {
		input := &codeartifact.TagResourceInput{
			ResourceArn: aws.String(identifier),
			Tags:        Tags(updatedTags),
		}

		_, err := conn.TagResourceWithContext(ctx, input)

		if err != nil {
			return fmt.Errorf("tagging resource (%s): %w", identifier, err)
		}
	}

	return nil
}

// UpdateTags updates codeartifact service tags.
// It is called from outside this package.
func (p *servicePackage) UpdateTags(ctx context.Context, meta any, identifier string, oldTags, newTags any) error {
	return updateTags(ctx, meta.(*conns.AWSClient).CodeArtifactConn(ctx), identifier, oldTags, newTags)
}
