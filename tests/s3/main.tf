terraform {}

variable "bucket_name" {
  type = string
}
variable "bucket_count" {
  type    = number
  default = 10
}

resource "aws_s3_bucket" "example" {
  count  = var.bucket_count
  bucket = "${var.bucket_name}-${count.index}"
}

module "objects" {
  count       = var.bucket_count
  source      = "./objects"
  bucket_name = aws_s3_bucket.example[count.index].bucket
}
