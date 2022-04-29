# Parameters to configure #
PROJECT_ID="lims001-saas-mon"

# Cloud Function Parameters #
CF_REGION="us-east1"
TIMEOUT=540 # INT - In seconds max=540
MEMORY=1024 # INT - In MB max=8192MB
PAGE_SIZE=500 # INT

# Cloud Scheduler Parameters #
EXPORT_NAME="vm-cpu-metric" # Keep this name unique for each metric export, this is the scheduler name as well as the table name in BigQuery
TIME_ZONE="UTC"
SCHEDULE="0/15 * * * *" # The export job will be triggered by this expression, for more information please look at https://cloud.google.com/scheduler/docs/configuring/cron-job-schedules.
WEEKS=0 # INT
DAYS=0 # INT
HOURS=1 # INT
FILTER='metric.type = "compute.googleapis.com/instance/cpu/usage_time"','metric.type = "compute.googleapis.com/instance/cpu/utilization"','metric.type = "compute.googleapis.com/instance/cpu/reserved_cores"','metric.type = "compute.googleapis.com/instance/uptime_total"' # Change to your metric filter
# BigQuery Parameters - Configure only at First deployment #
BQ_DATASET="brcm_gcp_usage_metrics" #
BQ_LOCATION="US"


# -------------------------------------------------------------------------------------------------------------------- #

# System parameters - Don't change #

# Cloud Function Parameters #

CF_NAME="metric_exporter_cf"
CF_SA="metric-exporter-cf-sa@"$(PROJECT_ID)".iam.gserviceaccount.com"
RUNTIME="python37"
SOURCE_PATH="./cloud_function" # Source file path for the cloud function
ENTRY_POINT="export"

# Cloud Scheduler Parameters #

SCHEDULER_SA="metric-exporter-scheduler-sa@"$(PROJECT_ID)".iam.gserviceaccount.com" # Cloud Function Invoker
HEADERS="Content-Type=application/json,User-Agent=Google-Cloud-Scheduler"

# BigQuery Parameters #
BQ_TABLE=$(EXPORT_NAME)

# GCS Bucket Parameters#
BUCKET_NAME=$(PROJECT_ID)"-vm-metric-exporter"

TOPIC_NAME=$(EXPORT_NAME)

SUBSCRIPTION_NAME=$(TOPIC_NAME)

PUSH_ENDPOINT="/projects/"$(PROJECT_ID)"/topic/"$(EXPORT_NAME)
# System Parameters - Don't change #

MSG_TMP_DIR="./msg_tmp"
MSG_BODY_FILE_NAME="msg.json"

# -------------------------------------------------------------------------------------------------------------------- #

deploy_topic:
	gcloud pubsub topics create $(TOPIC_NAME) --project=$(PROJECT_ID)

# subscribe_topic:
# 	gcloud pubsub subscriptions create $(SUBSCRIPTION_NAME) --topic=$(TOPIC_NAME) --topic-project=$(PROJECT_ID) \
# 	--ack-deadline=600 --expiration-period=never --message-retention-duration=2d --project=$(PROJECT_ID) 

deploy_cloud_function:
	gcloud functions deploy $(CF_NAME) --region=$(CF_REGION) --runtime=$(RUNTIME) --trigger-topic=$(TOPIC_NAME) --source=$(SOURCE_PATH) \
	--ingress-settings="internal-only" --entry-point=$(ENTRY_POINT) --project=$(PROJECT_ID) --service-account=$(CF_SA) \
	--memory=$(MEMORY) --timeout=$(TIMEOUT)

deploy_scheduler: build_json_msg
	gcloud scheduler jobs create pubsub $(EXPORT_NAME) --project=$(PROJECT_ID) --schedule=$(SCHEDULE) \
	--topic=$(TOPIC_NAME) --message-body-from-file=$(MSG_TMP_DIR)"/"$(MSG_BODY_FILE_NAME) \
	--time-zone=$(TIME_ZONE) --location=$(CF_REGION)

test_filter_api:
	python validate_filter.py --project=$(PROJECT_ID) --filter=$(FILTER)

build_json_msg:
	python build_message_body.py --project=$(PROJECT_ID) --MSG_TMP_DIR=$(MSG_TMP_DIR) --MSG_BODY_FILE_NAME=$(MSG_BODY_FILE_NAME)

clean:
	rm $(MSG_TMP_DIR)"/"$(MSG_BODY_FILE_NAME)

delete_cloud_function:
	gcloud functions delete $(CF_NAME) --region=$(CF_REGION) --project=$(PROJECT_ID)

delete_scheduler:
	gcloud scheduler jobs delete $(EXPORT_NAME) --project=$(PROJECT_ID) --location=$(CF_REGION)

delete_topic:
	gcloud pubsub topics delete $(TOPIC_NAME) --project=$(PROJECT_ID)

# delete_subscribe_topic:
# 	gcloud pubsub subscriptions create $(SUBSCRIPTION_NAME) --project=$(PROJECT_ID)

create_bq_dataset:
	bq --location=$(BQ_LOCATION) mk --dataset $(PROJECT_ID):$(BQ_DATASET)

get_cf_sa_name:
	@echo $(CF_SA)

get_scheduler_sa_name:
	@echo $(SCHEDULER_SA)

schedule_metric_export: deploy_scheduler clean

deploy_pub_sub: deploy_topic

full_deploy: deploy_pub_sub deploy_cloud_function schedule_metric_export