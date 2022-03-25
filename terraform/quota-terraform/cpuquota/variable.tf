variable "region" {
  description = "GCP region name"
  type        = string
}
variable "projectid" {
  description = "GCP projectid"
  type        = string
}
variable "override" {
  type = number
}

# variable "overridedisk" {
#   type = number
# }

variable "cpu_names" {
  description = "CPUs to be restricted"
  type        = list
  default = []
}