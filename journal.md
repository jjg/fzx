# Journal

## 10162024
Probably should have created this when I started this repository but anyway...

Continued work on `Inode` after getting it to sucessfully load existing JSFS inodes earlier.  Next step is to store new/modified inodes, and see if JSFS can read them.

Looks like `time.UnixMilli()` and `time.UnixMilli(ms)` can be used to correctly convert golang `time` types to and from the integer JSFS uses to store the `Created` timestamp.  I'm not sure if I want to parse this into a go time type when loading an inode or not, since doing so would allow for more utility in working with the time data but it also means extra steps while serializing/deserializing inodes.

I'm also trying to decide how "intelligent" the `Inode` struct should be.  In JSFS, the inode object configures defaults and such but I'm not sure this is how it should work.

Something not directly related that occurs as I write these tests is that the "storage location" (currently just a string) could be an actual struct that contains the `read()`/`write()` methods of various storage drivers.  This might be a good way to implement the "pluggable" block storage stuff.
