// SPDX-License-Identifier: MIT
// Copyright (c) 2019 Gennady Trafimenkov

#include <unistd.h>
#include <string.h>

int main()
{
    const char *message = "hello\n";
    write(1, message, strlen(message));
    return 0;
}
