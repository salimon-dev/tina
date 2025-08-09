#!/bin/bash

set -e

ENV="$1"

if [[ "$ENV" != "dev" && "$ENV" != "prod" ]]; then
  echo "Usage: $0 dev|prod"
  exit 1
fi


echo "Deploying to $ENV environment..."

if [ "$ENV" == "dev" ]; then
  rm ./layers/env/.env
  cp ./layers/.env.development ./layers/env/.env

  sam build
  sam deploy --no-confirm-changeset --config-file ./development.toml
else
  rm ./layers/env/.env
  cp ./layers/.env.production ./layers/env/.env
  
  sam build
  sam deploy --no-confirm-changeset --config-file ./production.toml
fi

echo "Deployment to $ENV completed successfully."
