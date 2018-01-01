golang-init
====

golang microservice as init.

This project runs at http://golang-init.zifnab.net. Basic rundown:

1) Configure eth0
2) Add default route
3) Do something on :80 (in this case, an http echo server)

This project is running on a Linode at the url listed above. There is no actual linux distribution installed, the (statically linked) binary created by ./build.sh is on disk at /bin/bash, and init=/bin/bash is set on the Linode's profile.

License
===

I have no idea what the legal implications are of statically linking to glibc. All code here is provided under Unlicense (https://unlicense.org/), use as you see fit.
