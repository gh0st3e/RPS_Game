# RPC_Game

This is famous game (rock-paper-scissors) and I've done it using next technologies:
* gRPC
* Kubernetes (k8s)

To run you need to install locale k8s cluster - minikube

``minikube start``

The next step is enable ingress addons

``minikube addons enable ingress``

You next step will be

``eval $(minikube -p minikube docker-env)``

It allows you to switch on docker inside the minikube

Next you need to pull images

``docker pull ghostfed/user_service``

``docker pull ghostfed/computer_service``

``docker pull ghostfed/result_service``

After this to deploy all in minikube you need to switch dir to ``k8s`` and run bash script

``./full_deploy.sh``

And at the finish you need to add hostname and ip to ``hosts``

At first run next command

``kubectl get ing``

You will find host and address there

After this you need to open terminal and write next commands

``sudo su``

``sudo gedit /etc/hosts``

Add address and host at the end of file like this

``192.168.49.2 rsp.game``

After this using endpoint ``rsp.game/start/?value`` you will be able to play