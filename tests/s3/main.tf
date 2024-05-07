terraform {}

variable "bucket_name" {
  type = string
}

resource "aws_s3_bucket" "example" {
  bucket = var.bucket_name
}

resource "aws_s3_object" "object" {
  bucket  = aws_s3_bucket.example.id
  key     = "object"
  content = "this is a test object"
}
