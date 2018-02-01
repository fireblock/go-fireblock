#!/usr/bin/env python
# coding: utf-8
"""Script to build all supported platforms"""
import sys
import os
import shutil
from subprocess import call

PLATFORMS = [ "windows/amd64", "windows/386", "darwin/amd64", "linux/amd64", "linux/386" ]
VERSION = "0.1.1"
if __name__ == "__main__":
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
        archive = "fbk-%s-%s-%s.tar.gz" % (eos, arch, VERSION)
        cmdBuild = "GOOS=%s GOARCH=%s go build -o build/%s github.com/fireblock/go-fireblock" % (eos, arch, "fbk")
        cmdArch = "tar zcvf dist/%s build/%s build/LICENSE" % (archive, "fbk")
        call(cmdBuild, shell=True)
        call(cmdArch, shell=True)

