echo "\nBuilding 'companies-microservice' Docker image...\n"
docker build --tag companies-microservice .

echo "\nStarting 'companies-microservice' and 'postgres' containers...\n"
docker compose up -d
