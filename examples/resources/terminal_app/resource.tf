resource "terminal_app" "example_app" {
  id = "cli_XXXXXXXXXXXXXXXXXXXXXXXXX"
  name = "Example App"
  redirect_uri = "https://example.com/callback"
  secret = "sec_******XXXX"
}
