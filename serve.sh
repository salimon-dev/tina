rm ./layers/env/.env
cp ./layers/.env.development ./layers/env/.env
  
sam build
sam local start-api --port 8081