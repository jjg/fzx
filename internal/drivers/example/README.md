# Storage Driver Example

Storage drivers allow `fzx` to store blocks and inodes anywhere data can be written to or read from.  By default fzx uses the `local` driver to write to a local disk (this driver is used an example here), but by adapting the `read`, `write` and `exists` primatives to interface with any other storage system fzx will seamlessly utilize that storage as well.
