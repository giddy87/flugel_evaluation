resource "aws_s3_bucket" "s3" {
  bucket = "flugel-test-bucket"
  acl    = "private"

  tags = var.tags
}
