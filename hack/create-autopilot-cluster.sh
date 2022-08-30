#!/bin/bash

cluster_name=autopilot-cluster
region=us-central1

gcloud container clusters create-auto $cluster_name \
  --region=$region
