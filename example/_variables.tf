variable "notion_integration_token" {
  type        = string
  description = "Notion Integration Token"
  sensitive   = true
}

variable "notion_api_version" {
  type        = string
  description = "Notion API Version"
  default     = "2022-06-28"
}

variable "notion_parent_id" {
  type        = string
  description = "Notion Parent ID"
}
