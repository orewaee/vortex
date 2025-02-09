#!/bin/bash
ENV_FILE_PATH=config/.env
FILE_PATH=deploy/compose.yaml
PROJECT_NAME=vortex

docker compose \
  --env-file config/.env \
  --file $FILE_PATH \
  --project-name $PROJECT_NAME \
  up --detach
