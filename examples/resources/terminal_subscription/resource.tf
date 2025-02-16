resource "terminal_subscription" "example_subscription" {
  id = "sub_XXXXXXXXXXXXXXXXXXXXXXXXX"
  address_id = "shp_XXXXXXXXXXXXXXXXXXXXXXXXX"
  card_id = "crd_XXXXXXXXXXXXXXXXXXXXXXXXX"
  frequency = "fixed"
  product_variant_id = "var_XXXXXXXXXXXXXXXXXXXXXXXXX"
  quantity = 1
  next = "2025-02-01T19:36:19.000Z"
  schedule = {
    interval = 3
    type = "weekly"
  }
}
