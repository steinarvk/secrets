Secrets
=======

This is a tiny Go library and associated tools for handling
files that contain secrets. Certain standards are enforced
for secret files; namely currently that they are not
world- or group-readable and that they are stored in files
with names that clearly mark them as secret (".secret.yaml").

This is a useful-seeming utility that was split off from
the github.com/steinarvk/watcher project for use in other
projects.

Tools
=====

*pg\_fromyaml*: a tool for creating a Postgres connection
string from a Postgres secret file.

Legal stuff
===========

I (@steinarvk) hold the copyright on this code. It is not
associated with any employer of mine, past or present.

The code is made available for use under the MIT license;
see the LICENSE file for details.
