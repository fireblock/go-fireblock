FIO
===

[![Gitter chat](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/fireblock/Lobby)

[![Go Report Card](https://goreportcard.com/badge/github.com/fireblock/go-fireblock)](https://goreportcard.com/report/github.com/fireblock/go-fireblock)

[![Build Status](https://travis-ci.org/fireblock/go-fireblock.svg?branch=master)](https://travis-ci.org/fireblock/go-fireblock)

Check integrity and ownership of a file registered on [Fireblock.io](https://fireblock.io)

Check by user id:

- ```fio -u 0xf104f872ef3608c99c55a1f1d65a029db103e825712c4c6f0d0ba752f3923c8e ~/Downloads/HelloWorld.zip```

Check by card id:

- ```fio -c 0xd8d94a286b1012b7612c2f2f49c843c913ae3c18ec7265fd58814a3197ab76c7 ~/Downloads/HelloWorld.zip```

To add more information use the -j flag to have an detailled output in JSON

- ```fio -u 0xf104f872ef3608c99c55a1f1d65a029db103e825712c4c6f0d0ba752f3923c8e -j ~/Downloads/HelloWorld.zip```
