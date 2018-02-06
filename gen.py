#!/usr/bin/env python
# coding: utf-8
"""Script to generate go code from ABI"""
import os
import shutil
from subprocess import call

if __name__ == "__main__":
    cmdStore = "abigen --abi abi/store.abi --pkg contracts --type Store --out contracts/store.go"
    cmdStore = "abigen --abi abi/fireblock.abi --pkg contracts --type Fireblock --out contracts/fireblock.go"
    call(cmdStore, shell=True)
