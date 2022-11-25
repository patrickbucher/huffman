# Huffman Encoding/Decoding

Demonstration of Huffman encoding/decoding (compression/decompression) for educational purposes.

Usage:

    $ go build cmd/huffman.go
    $ ./huffman abacadabra
    abacadabra -> [0 1 1 1 0 1 0 0 0 1 0 1 0 1 1 1 1 1 0 0] -> abacadabra
    encoded message 'abacadabra' (10 characters) in 20 bits

Compression:

```go
encoded, tree := huffman.Encode("abacadabra")
```

Decompression:

```go
decoded := huffman.Decode(encoded)
```
