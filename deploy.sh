#!/usr/bin/env bash

export PROJECT_ID=catbot-391520

export ENV_VARS_FILE=~/tokens/env.yaml

gcloud run deploy catbot \
  --source . \
  --env-vars-file ${ENV_VARS_FILE} \
  --allow-unauthenticated \
  --project ${PROJECT_ID} \
  --region europe-west1