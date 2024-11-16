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
