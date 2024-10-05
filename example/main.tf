resource "notion_page" "page1" {
  name      = "Example Notion Resource"
  parent_id = var.notion_parent_id
}
