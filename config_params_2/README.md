# Serving a struct
this design pattern uses protobuf definition as the starting point to
define structs and services on those structs. Then a non-grpc service
can be defined in a similar manner and a client can be written on top
of that.