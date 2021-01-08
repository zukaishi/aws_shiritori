variable "bucket_name" {
  default = "website.shiritori.com"
}

data "aws_iam_policy_document" "s3_bucket_policy" {
  statement {
    sid       = ""
    effect  = "Allow"
    principals {
      type        = "*"
      identifiers = ["*"]
    }
    actions = ["s3:GetObject"]
    resources = [
      "arn:aws:s3:::${var.bucket_name}/",
      "arn:aws:s3:::${var.bucket_name}//*"
    ]
  }
}

resource "aws_s3_bucket" "website_shiritori_com" {
  bucket = var.bucket_name
  policy = data.aws_iam_policy_document.s3_bucket_policy.json
 
  website {
    index_document = "index.html"
    error_document = "error.html"
  }
}

resource "aws_s3_bucket_public_access_block" "website_shiritori" {
  bucket                  = var.bucket_name
  block_public_acls       = true
  block_public_policy     = false
  ignore_public_acls      = true
  restrict_public_buckets = false
}

output "url" {
  value = aws_s3_bucket.website_shiritori_com.website_endpoint
}