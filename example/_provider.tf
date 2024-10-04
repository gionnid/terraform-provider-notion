terraform {
  required_providers {
    notion = {
      source  = "terraform.local/local/notion"
      version = "0.0.1"
    }
  }
}

provider "notion" {
  notion_integration_token = var.notion_integration_token
  notion_api_version       = var.notion_api_version
}

variable "notion_integration_token" {
  type        = string
  description = "Notion Integration Token"
  default     = "xxxxx"
  sensitive   = true
}

variable "notion_api_version" {
  type        = string
  description = "Notion API Version"
  default     = "2022-06-28"
}