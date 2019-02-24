# SPDX-License-Identifier: MIT
# Copyright (c) 2019 Gennady Trafimenkov

build:
	gcc libfoo.c -shared -fPIC -ldl -o libfoo.so
	gcc testwrite.c -o testwrite

test:
	LD_PRELOAD=./libfoo.so ./testwrite

test-debug:
	LD_DEBUG=all LD_PRELOAD=./libfoo.so ./testwrite
