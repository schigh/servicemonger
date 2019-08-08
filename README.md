This is a widget to build BitBar plugins from a YAML file
It is still very much WIP.  Please don't use unless you like using WIP stuff.

Also, in retrospect, `servicemonger` is a weird name.

# Features (current)
- Build menus from YAML files
  - HTTP links
  - Bash scripts
- Built-in base64 image converter

# Features (planned)
- Desktop notifications
- Status Bar updates
- etc

# Usage

## Build
- grab all dependencies (currently only YAML parser)
  - `go mod download`
- build
  - `go build -o servicemonger ./cmd/main.go`
  
## Use
- move the `servicemonger` binary somewhere in your `PATH` or anywhere else you can reference it
- create a bash script and place it into your `~/BitBarPlugins` folder:
```shell script
#!/bin/bash

/path/to/servicemonger -settings /path/to/service-menu.yml
```
- name the bash script with the refresh period before the file extension, e.g. `foo.1m.sh` (refresh every minute), `foo.1h.sh` (refresh every hour), etc


### Base64 images
To convert an image to base64, run servicemonger as follows:
```shell script
> servicemonger -b64 -file /path/to/myimage/png
```
The caveat with the image functionality is that it currently does not support retina resolution, so YMMV
