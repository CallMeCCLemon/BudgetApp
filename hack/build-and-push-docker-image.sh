TAG="thomas-latest"
ENDPOINT="100.69.236.43:32000"

docker build -t $ENDPOINT/budget-app:$TAG .
docker push $ENDPOINT/budget-app:$TAG
