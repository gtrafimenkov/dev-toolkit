# SPDX-License-Identifier: MIT
# Copyright (c) 2019 Gennady Trafimenkov

import cherrypy

class HelloWorld(object):
    @cherrypy.expose
    def index(self):
        return "Hello World!"

cherrypy.config.update({'server.socket_host': '0.0.0.0', 'server.socket_port': 9090})
cherrypy.quickstart(HelloWorld())
