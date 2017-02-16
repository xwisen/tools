curl -i -v -X GET -H Content-Type:application/json http://20.26.2.52:8080/apis/extensions/v1beta1/namespaces/default/deployments
curl -i -v -X POST -d@tm.json -H Content-Type:application/json http://20.26.2.52:8080/apis/extensions/v1beta1/namespaces/default/deployments
curl -i -v -X DELETE -H Content-Type:application/json http://20.26.2.52:8080/apis/extensions/v1beta1/namespaces/default/deployments/{name}
