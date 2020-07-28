all: autoscaler ingress externaldns

autoscaler:
	helm install cluster-autoscaler --namespace kube-system stable/cluster-autoscaler --values=cluster-autoscaler-chart-values.yaml

uninstall_autoscaler:
	helm uninstall cluster-autoscaler --namespace kube-system stable/cluster-autoscaler


ingress:
	kubectl apply -f rbac-role.yaml
	kubectl apply -f alb-ingress-controller.yaml

uninstall_ingress:
	kubectl delete -f alb-ingress-controller.yaml
	kubectl delete -f rbac-role.yaml


externaldns:
	kubectl apply -f external-dns.yaml

uninstall_externaldns:
	kubectl delete -f external-dns.yaml


prometheus:
	kubectl create namespace prometheus
	helm install -f prometheus.yaml prometheus --namespace prometheus stable/prometheus

uninstall_prometheus:
	helm uninstall prometheus --namespace prometheus
	kubectl delete ns prometheus

grafana:
	kubectl create namespace grafana
	helm install grafana stable/grafana \
		--namespace grafana \
		--set persistence.storageClassName="gp2" \
		--set datasources."datasources\.yaml".apiVersion=1 \
		--set datasources."datasources\.yaml".datasources[0].name=Prometheus \
		--set datasources."datasources\.yaml".datasources[0].type=prometheus \
		--set datasources."datasources\.yaml".datasources[0].url=http://prometheus-server \
		--set datasources."datasources\.yaml".datasources[0].access=proxy \
		--set datasources."datasources\.yaml".datasources[0].isDefault=true \
		--set service.type=LoadBalancer
	kubectl get secret --namespace grafana grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo
	export ELB=$(kubectl get svc -n grafana grafana -o jsonpath='{.status.loadBalancer.ingress[0].hostname}')
	echo "http://$ELB"

uinstall_grafana:
	helm uninstall grafana --namespace grafana
	kubectl delete ns grafana

termination_handler:
	kubectl apply -f https://github.com/aws/aws-node-termination-handler/releases/download/v1.4.0/all-resources.yaml

# argo:
# 	kubectl create namespace argocd
# 	kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
