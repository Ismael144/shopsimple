.PHONY: buf build installation setup

buf: 
		buf generate buf.build/shopsimple/shopsimple
build: 
		docker compose build currencyservice 
		docker compose up currencyservice -d 
installation: 
		apt install python3.12-venv
		sudo apt install python3-pip
setup: installation 
		python3 -m venv venv
		source ./venv/bin/activate
		pip install -r requirements.txt