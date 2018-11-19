#!/usr/bin/env python
# coding: utf-8
"""Script to build all supported platforms"""
from subprocess import call

if __name__ == "__main__":
    cmdTest = "go test github.com/fireblock/go-fireblock/fireblocklib"
    call(cmdTest, shell=True)
