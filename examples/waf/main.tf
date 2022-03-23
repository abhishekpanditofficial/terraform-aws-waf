terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}

# Configure the AWS Provider
provider "aws" {
  region     = "us-east-1"
  access_key = ""
  secret_key = ""
}

module "aws_waf" {
  source= "../../"
  acl_name= var.acl_name
  acl_description= var.acl_description
  metric_name= var.metric_name
  project= var.project
}


