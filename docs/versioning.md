# Tagging and versioning

When pushing a driver it can (must) be tagged, e.g. for a driver named 'bar' by author 'foo', when doing

    drivers push -t test
	
which would result in a driver identified by `foo/bar:test`.
Tags can be alphanumeric strings, e.g. git commit hashes, branch names, or semantic verions.
Multiple tags can be added at the same time:

    drivers push -t test -t temporary


## Semantic Versioning

Semantic versioning (https://semver.org/) is an informal standard on how software version numbers
change across different releases of a software product.

Version numbers follow the semantic versioning consist of a major version, a minor version,
and a patch version, separated by dots: major.minor.patch, e.g. `1.0.2`.
Changes to the major version indicate backwards-incompatible changes,
changes to the minor version number indicate backwards-compatible feature additions,
and change to the patch level indicate backwards-compatible bug fixes.

We strongly recommend use of semantic versioning for drivers.
If you decide to do so, we consider it best to adhere the following rules:

Increment the major version number when you:
* use a different schema version
* rename a field in the schema
* change the type of a field in the schema
* change the physical unit, e.g. from degree Celsius to Fahrenheit

Increment the minor version number when you:
* add new fields to the schema

Increment the patch level when you:
* fix incorrect calculations of data
* fix crashes


## Fixed and Floating Tags

For semantic versioning in driver development,
it is recommended to use the pattern of floating and fixed tags,
as commonly used in the docker community.
For this, the driver should be tagged with all version prefixes of the real version number:

    drivers push -t 1.0.0 -t 1.0 -t 1

This will create multiple tags for the same driver and a device can be configured
to use any of these:

    foo/bar:1
    foo/bar:1.0
    foo/bar:1.0.0

Right after this driver was pushed, there's no behavioral difference in using any of those tags.
However, when a new versions of the driver becomes available at a later point of time,
and the driver author applies the pattern again, the tags will be updated differently.
Imagine the driver requires patching and as a consequence, a new version 1.0.1 is released.
Still following the pattern, the driver's author, will now tag this new version like this:

    drivers push -t 1.0.1 -t 1.0 -t 1

Note that this will overwrite the tags `1` and `1.0`, but create a new tag `1.0.1`
and leave the tag `1.0.0` unchanged.
If a device was set to use version `1.0.0` it will now still use the same driver.
However, if the driver was configured to use the tags `1` or `1.0`,
it will automatically use the newer version.

If a new feature is added to the driver, a new minor version number would be released.
In this case that would be version 1.1.0, so the author would push the driver with

    drivers push -t 1.1.0 -t 1.1 -t 1
	
Notice that now only the tag `1` is overwritten, while tags `1.0` and `1.0.1` are unchanged
and tags `1.1` and `1.1.0` are created.

This means a device set to use tag `1` updates automatically now,
while a device set to use `1.0`, `1.0.0` or `1.0.1` will use the same driver as before.

Note that all of this happens as convention and is not hardcoded.
Any driver developer is free to not use semantic versioning,
so he can just as well follow a device's firmware versions,
even when those do not follow the semantic versioning standard.



