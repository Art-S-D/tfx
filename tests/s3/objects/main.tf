variable "bucket_name" {
  type = string
}
variable "object_count" {
  type    = number
  default = 10
}

resource "aws_s3_object" "object" {
  count   = var.object_count
  bucket  = var.bucket_name
  key     = "object-${count.index}"
  content = "this is the ${count.index}th test object"
}
