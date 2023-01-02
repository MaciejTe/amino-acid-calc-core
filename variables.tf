variable "AWS_REGION" {
  default = "eu-west-1"
}

variable "tags" {
  description = "Default tags to apply to all resources."
  type        = map(any)
  default = {
    archuuid = "ab447f37-4fe5-4d0d-8802-4128753054c7"
    env      = "Development"
  }
}

