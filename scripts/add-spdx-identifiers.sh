#!/bin/bash

SD=$(dirname $0)

for f in $(find $SD/../examples -type f -print); do
    if ! grep -q "# SPDX-License-Identifier" $f; then
        sed -i '1s!^!# SPDX-License-Identifier: MIT\n# Copyright (c) 2019 Gennady Trafimenkov\n\n!' $f
    fi
done
