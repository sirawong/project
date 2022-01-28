docker pull siraphopdocker/user:latest
docker pull siraphopdocker/media:latest
docker pull siraphopdocker/showtime:latest
docker pull siraphopdocker/cinema:latest
docker pull siraphopdocker/movie:latest
docker pull siraphopdocker/reservation:latest
docker pull siraphopdocker/movie-booking:latest

docker tag siraphopdocker/user:latest asia.gcr.io/$PROJECT_ID/movie-user:dev
docker tag siraphopdocker/user:latest asia.gcr.io/$PROJECT_ID/movie-user:prd
docker tag siraphopdocker/media:latest asia.gcr.io/$PROJECT_ID/movie-media:dev
docker tag siraphopdocker/media:latest asia.gcr.io/$PROJECT_ID/movie-media:prd
docker tag siraphopdocker/showtime:latest asia.gcr.io/$PROJECT_ID/movie-showtime:dev
docker tag siraphopdocker/showtime:latest asia.gcr.io/$PROJECT_ID/movie-showtime:prd
docker tag siraphopdocker/cinema:latest asia.gcr.io/$PROJECT_ID/movie-cinema:dev
docker tag siraphopdocker/cinema:latest asia.gcr.io/$PROJECT_ID/movie-cinema:prd
docker tag siraphopdocker/movie:latest asia.gcr.io/$PROJECT_ID/movie-movie:dev
docker tag siraphopdocker/movie:latest asia.gcr.io/$PROJECT_ID/movie-movie:prd
docker tag siraphopdocker/reservation:latest asia.gcr.io/$PROJECT_ID/movie-reservation:dev
docker tag siraphopdocker/reservation:latest asia.gcr.io/$PROJECT_ID/movie-reservation:prd
docker tag siraphopdocker/movie-booking:latest asia.gcr.io/$PROJECT_ID/movie-booking:dev
docker tag siraphopdocker/movie-booking:latest asia.gcr.io/$PROJECT_ID/movie-booking:prd

docker push asia.gcr.io/$PROJECT_ID/movie-user:dev
docker push asia.gcr.io/$PROJECT_ID/movie-user:prd
docker push asia.gcr.io/$PROJECT_ID/movie-media:dev
docker push asia.gcr.io/$PROJECT_ID/movie-media:prd
docker push asia.gcr.io/$PROJECT_ID/movie-showtime:dev
docker push asia.gcr.io/$PROJECT_ID/movie-showtime:prd
docker push asia.gcr.io/$PROJECT_ID/movie-cinema:dev
docker push asia.gcr.io/$PROJECT_ID/movie-cinema:prd
docker push asia.gcr.io/$PROJECT_ID/movie-movie:dev
docker push asia.gcr.io/$PROJECT_ID/movie-movie:prd
docker push asia.gcr.io/$PROJECT_ID/movie-reservation:dev
docker push asia.gcr.io/$PROJECT_ID/movie-reservation:prd
docker push asia.gcr.io/$PROJECT_ID/movie-booking:dev
docker push asia.gcr.io/$PROJECT_ID/movie-booking:prd

gcloud container --project "$PROJECT_ID" clusters create "$K8S_NAME" --zone "$K8S_ZONE" \
  --cluster-version "1.20.12-gke.1500" --release-channel "stable" --machine-type "e2-medium" \
  --enable-ip-alias --image-type "COS_CONTAINERD" --disk-size "100" --num-nodes "2" \
  --network "default" --subnetwork "default" --preemptible

kubectl create namespace movie-dev
kubectl create namespace movie-prd

kubectl config set-context $(kubectl config current-context) --namespace=movie-dev

kubectl create secret generic registry-movie \
  --from-file=.dockerconfigjson=$HOME/.docker/config.json \
  --type=kubernetes.io/dockerconfigjson

kubectl apply -f ~/movie-secret/movie-dev-mongodb-secret.yaml

echo $(kubectl get secret movie-dev-mongodb-secret \
  -o jsonpath="{.data.mongodb-root-password}" | base64 --decode)


echo -n "movie-dev-root" | base64

kubectl create configmap movie-dev-mongodb-initdb \
  --from-file=databases/users_data.json \
  --from-file=databases/script.sh --dry-run=client -o yaml | kubectl apply -f -

helm install -f ~/k8s/helm-values/values-movie-dev-mongodb.yaml \
  movie-dev-mongodb bitnami/mongodb --version v10.10.1

export MONGODB_ROOT_PASSWORD=$(kubectl get secret movie-dev-mongodb-secret \
  -o jsonpath="{.data.mongodb-root-password}" | base64 --decode)

kubectl run mongodb-client --rm --tty -i --restart='Never' --image bitnami/mongodb:4.4.4-debian-10-r27 \
  --command -- mongo admin --host movie-dev-mongodb --authenticationDatabase admin \
  -u root -p $MONGODB_ROOT_PASSWORD