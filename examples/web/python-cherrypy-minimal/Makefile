# SPDX-License-Identifier: MIT
# Copyright (c) 2019 Gennady Trafimenkov

default:
	true

build:
	docker build . -t demos:cherrypy

test:
	echo "Endpoint: http://localhost:9090/"
	docker run --rm -it -p 9090:9090 demos:cherrypy python3 /server.py

publish:
	docker tag democherrypy:latest quay.io/gtrafimenkov/demos:cherrypy
	docker push quay.io/gtrafimenkov/demos:cherrypy
