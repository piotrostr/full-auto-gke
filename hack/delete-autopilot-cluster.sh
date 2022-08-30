#!/bin/bash

cluster_name=autopilot-cluster
region=us-central1

gcloud container clusters delete $cluster_name \
  --region $region \
  --verbosity=info \
  --quiet
