#!/usr/bin/env python
# coding: utf-8
"""Script to build all supported platforms"""
import os
import sys
import shutil
from subprocess import call

import argparse

# read version file
VERSION = "0.0.0"
with open('version.txt') as f:
    VERSION = f.read().rstrip('\r\n')
    f.close()
# create version.go
with open('cmd/fio/version.go', 'w') as f2:
    code = '''package main

const (
	// Version - fbk version
	Version = "$#$VERSION$#$"
	// Author fireblock
	Author = "Laurent Mallet laurent.mallet at gmail dot com"
)
'''
    code = code.replace('$#$VERSION$#$', VERSION)
    f2.write(code)
    f2.close()

PLATFORMS = ["windows/amd64", "windows/386", "darwin/amd64", "linux/amd64", "linux/386"]

if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument('--platform', help='select a platform')
    args = parser.parse_args()
    if args.platform:
        PLATFORMS = [ args.platform ]
    if not os.path.exists("build"):
        os.makedirs("build")
    if not os.path.exists("dist"):
        os.makedirs("dist")
    for p in PLATFORMS:
        # clean build dir
        for f in os.listdir('build'):
            os.remove('build/' + f)
        # copy LICENSE
        shutil.copy2('LICENSE', 'build/')
        els = p.split('/')
        eos = els[0]
        arch = els[1]
        if  p.find('windows') != -1:
            archive = "fio-%s-%s-%s.zip" % (eos, arch, VERSION)
            cmdBuild = "GOOS=%s GOARCH=%s go build -o build/%s github.com/fireblock/go-fireblock/cmd/fio" % (eos, arch, "fio")
            cmdArch = "zip -r dist/%s build/*" % (archive)
            call(cmdBuild, shell=True)
            call(cmdArch, shell=True)
        else:
            archive = "fio-%s-%s-%s.tar.gz" % (eos, arch, VERSION)
            cmdBuild = "GOOS=%s GOARCH=%s go build -o build/%s github.com/fireblock/go-fireblock/cmd/fio" % (eos, arch, "fio")
            cmdArch = "tar zcvf dist/%s build/%s build/LICENSE" % (archive, "fio")
            call(cmdBuild, shell=True)
            call(cmdArch, shell=True)

