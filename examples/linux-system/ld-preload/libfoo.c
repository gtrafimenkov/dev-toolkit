// SPDX-License-Identifier: MIT
// Copyright (c) 2019 Gennady Trafimenkov

#define _GNU_SOURCE
#include <dlfcn.h>
#include <stdio.h>
#include <sys/types.h>

static size_t bytes_written;

static ssize_t (*original_write)(int fd, const void *buf, size_t count);

// GCC trick to make __libfoo_init run before main and __libfoo_done after
static void __libfoo_init(void) __attribute__((constructor));
static void __libfoo_done(void) __attribute__((destructor));

static void __libfoo_init(void)
{
    printf("== libfoo init\n");
    original_write = dlsym(RTLD_NEXT, "write");
}

static void __libfoo_done(void)
{
    printf("== libfoo done\n");
    printf("== bytes written: %ld\n", bytes_written);
}

// This is a wrapper around the standard write
ssize_t write(int fd, const void *buf, size_t count)
{
    ssize_t written = (*original_write)(fd, buf, count);
    if(written > 0)
    {
        bytes_written += written;
    }
    return written;
}
