# ext-go

**ext-go** provides an extension for the [mus-go](https://github.com/mus-format/mus-go)
serializer, enabling additional functionality for the MUS format.

This package includes:

- `MarshallerMUS` — an interface for types that can marshal themselves into the
  MUS format.
- `MarshalMUS` — a generic function to marshal values implementing
  `MarshallerMUS`.
- `MarshallerTypedMUS` — an interface for types that support typed MUS
  serialization, designed for use with [DTS](https://github.com/mus-format/dts-go).
- `MarshalTypedMUS` — a generic function to marshal values implementing
  `MarshallerTypedMUS`.

