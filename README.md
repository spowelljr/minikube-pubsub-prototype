# minikube-pubsub-prototype

## How to run on minikube
```
$ minikube start --driver docker
$ minikube ssh
$ curl -LO https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-cli-392.0.0-linux-x86_64.tar.gz
$ tar -xf google-cloud-cli-392.0.0-linux-x86_64.tar.gz
$ ./google-cloud-sdk/install.sh
$ . .bashrc
$ gcloud components install beta
$ sudo mkdir -p /usr/share/man/man1
$ sudo apt-get update
$ sudo apt-get install default-jre
$ gcloud components install pubsub-emulator
$ gcloud beta emulators pubsub start --host-port=0.0.0.0:8085 --project=test

# in another terminal run the following
$ git clone https://github.com/spowelljr/minikube-pubsub-prototype.git
$ cd minikube-pubsub-prototype
$ eval $(minikube docker-env)
$ docker build . -t pubsub
$ kubectl apply -f pubsub.yml
```
