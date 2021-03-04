# gogallery
Photo gallery written in Go

## Development
 * install KinD
 * create a KinD cluster
 ```bash
kind create cluster --name gogallery --config kind.config # for M1 add: --image rossgeorgiev/kind-node-arm64:v1.20.0
 ```
 * setup ingress
 ```bash
kubectl apply --filename https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/provider/kind/deploy.yaml

kubectl wait --namespace ingress-nginx \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/component=controller \
  --timeout=90s
 ```
 * start skaffold
 ```bash
skaffold dev
 ```

 * remove the KinD cluster
 ```bash
kind delete cluster --name gogallery
 ```

## Todo
 * finalise tests
 * recreate thumbs.bin when dir has changed / is newer
 * stream data/info to user when creating thumb
 * make sure the tree doesn't get squashed