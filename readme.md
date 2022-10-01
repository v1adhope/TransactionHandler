# Overview

The application was created for educational purposes, recording transactions in a NoSQL database using gRPC.

# Launch preparation

[CouchDB](https://couchdb.apache.org/) and [Go](https://go.dev/) must be installed on the device.
Default username and password for couch are admin, admin.


# Client side

Usage:
```
main.go <command> [arguments]
```
The commands are:
```
send        make a transaction
getlast     print latest transactions
```

## Make a transaction
```
main.go send [address from] [address to] [amount]
```
For example:
```
main.go send mr1NAUrG3TXFZaF4sUGvRRwA6T1x5BPkF7 mqkGkAUvZMG6jckqNXhENceDfPp5qQbnbm 44.780
```
## Print latest transactions
```
main.go getlast [count]
```
If more transactions are requested than exist, all available transactions are outputted.

# Result

[client side](https://share.cleanshot.com/qV99AlUtcO2IIhe1H0vj)

[server side](https://share.cleanshot.com/T9NL2mWLOSf0d8zaDxsT)

[base side](https://share.cleanshot.com/4cnDcc9MyZqNIX7Ju6x0)
