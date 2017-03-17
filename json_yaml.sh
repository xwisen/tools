# SHELL 
# ***********************************************
# 
#       Filename: json-yaml.sh
# 
#         Author: xwisen 1031649164@qq.com
#    Description: ---
#         Create: 2017-03-17 17:15:13
#  Last Modified: 2017-03-17 17:15:13
# ***********************************************


#eg: yamltojson test.yaml test.json
function yamltojson () {
	python -c '
import sys, yaml, json;
json.dump(yaml.load(sys.stdin), sys.stdout, indent=4)
'< $1 > $2
	if [[ $? -ne 0  ]];then
		echo ">>>>>>convert failed !"
	else
		echo ">>>>>>convert succeed ! input is : $1, output is : $2"
	fi
}

#eg: jsontoyaml test.json test.yaml
function jsontoyaml () {
	python -c '
import sys, yaml, json;
yaml.safe_dump(json.load(sys.stdin), sys.stdout, default_flow_style=False)
'< $1 > $2
	if [[ $? -ne 0  ]];then
		echo ">>>>>>convert failed !"
	else
		echo ">>>>>>convert succeed ! input is : $1, output is : $2"
	fi
}

