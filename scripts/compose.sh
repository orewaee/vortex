#!/bin/bash
docker compose \
  -f ./deploy/compose.yaml \
  -p vortex_app \
  --env-file ./config/postgres.env \
  --env-file ./config/pgadmin.env \
  --env-file ./config/redis.env \
  up -d
