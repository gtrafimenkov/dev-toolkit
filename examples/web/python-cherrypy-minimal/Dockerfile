# SPDX-License-Identifier: MIT
# Copyright (c) 2019 Gennady Trafimenkov

FROM ubuntu:18.04

RUN echo "Install dependencies" \
  && apt-get update \
  && apt-get upgrade -y \
  && apt-get install -y python3-pip \
  && pip3 install -U cherrypy \
  && apt-get clean && rm -rf /var/lib/apt/lists/* \
  && echo "== done"

ADD server.py /
