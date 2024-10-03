# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

# tflint-ignore: terraform_unused_declarations
data "aws_batch_scheduling_policy" "test" {
  arn = aws_batch_scheduling_policy.test.arn
}

resource "aws_batch_scheduling_policy" "test" {
  name = var.rName

  fair_share_policy {
    compute_reservation = 0
    share_decay_seconds = 0
  }

  tags = var.resource_tags
}

variable "rName" {
  description = "Name for resource"
  type        = string
  nullable    = false
}

variable "resource_tags" {
  description = "Tags to set on resource. To specify no tags, set to `null`"
  # Not setting a default, so that this must explicitly be set to `null` to specify no tags
  type     = map(string)
  nullable = true
}
