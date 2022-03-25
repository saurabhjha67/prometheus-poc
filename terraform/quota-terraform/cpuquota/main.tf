provider "google" {
  project = var.projectid
  region = var.region
}

resource "google_service_usage_consumer_quota_override" "cpuquota" {
  provider = google-beta
  project = var.projectid
  service = "compute.googleapis.com"
  metric =  urlencode("compute.googleapis.com/n2d_cpus")
  limit = urlencode("/project/region")
  override_value = var.override
  dimensions = {
    region = var.region
  }
  force = true
}

# resource "google_service_usage_consumer_quota_override" "diskquota" {
#   provider = google-beta
#   project = var.projectid
#   service = "compute.googleapis.com"
#   metric =  urlencode("compute.googleapis.com/disks_total_storage")
#   limit = urlencode("/project/region")
#   override_value = var.overridedisk
#   dimensions = {
#     region = var.region
#   }
#   force = true
# }

resource "google_service_usage_consumer_quota_override" "othercpu" {
  count = length(var.cpu_names)
  provider = google-beta
  project = var.projectid
  service = "compute.googleapis.com"
  metric =  urlencode(var.cpu_names[count.index])
  limit = urlencode("/project/region")
  override_value = 0
  dimensions = {
    region = var.region
  }
  force = true
}