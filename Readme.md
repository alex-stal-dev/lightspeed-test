IPv4 Counter

1. The main idea is to represent an IPv4 string as a 4-byte integer as it has 4 sections of 0-255 (1 byte) values. And for tracking uniqueness I used bitset (a slice of bool type) of size 2^32 which is the maximum number of different IPv4 addresses). The bitset uses 2^32 bits == 512 MB of memory.
2. The filename is hardcoded for simplicity and the file is read sequentially. It most likely can be optimized by reading file chunks in parallel for better performance.