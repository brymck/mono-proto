mono-proto
==========

![](https://github.com/brymck/mono-proto/workflows/goreleaser/badge.svg)

This CLI utility orchestrates

* Cloning a monorepo containing Protobuf schemas
* Compiling specific Protobufs from that monorepo using a tool like [Prototool][prototool]

This in turn allows Protobuf schemas to be decoupled from projects.

[prototool]: https://github.com/uber/prototool
