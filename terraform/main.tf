variable "bucket_name" {
  default = "zukaishi.shiritori"
}

data "aws_iam_policy_document" "s3_policy" {
  statement {
    actions = ["s3:GetObject"]
    effect  = "Allow"
    principals {
      type        = "AWS"
      identifiers = ["*"]
    }
    resources = ["arn:aws:s3:::${var.bucket_name}/*"]
    sid       = "PublicReadGetObject"
  }
}

resource "aws_s3_bucket" "b" {
  bucket = var.bucket_name
  acl    = "private"
  policy = data.aws_iam_policy_document.s3_policy.json
  website {
    index_document = "index.html"
    error_document = "error.html"
  }
}

output "url" {
  value = aws_s3_bucket.b.website_endpoint
}