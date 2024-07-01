terraform {}

resource "aws_ssm_parameter" "parameter" {
  name  = "tfx-test-sensitive-value"
  type  = "StringList"
  value = "string 1, string 2, string 3"
}

data "aws_ssm_parameter" "data_parameter" {
  name = aws_ssm_parameter.parameter.name
}
