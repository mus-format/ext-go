# ext-mus-go
Provides a [mus-go](https://github.com/mus-format/mus-go) serializer extension.

Includes the `MarshallerMUS` interface, which represents a type that can marshal itself into the MUS format, along with the generic `MarshalMUS` function. Also includes the `MarshallerTypedMUS` interface and the `MarshalTypedMUS` function, intended for use with [DTS](https://github.com/mus-format/mus-dts-go).