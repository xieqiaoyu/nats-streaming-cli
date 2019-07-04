# Nats-streaming Cli

[![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/apeirophobia/nats-streaming-cli.svg)](https://hub.docker.com/r/apeirophobia/nats-streaming-cli)


An interactive command line tool for nats-streaming server

Make it easy to monitor and publish message to a nats-streaming server in development

For security reasons ,**it's not recommend to use this cli in production!**

Support linux and macOS (windows may work, not sure)

## Installation

### Build from source
1. Clone this repo
1. cd to repo path
1. Run `./build`
1. Put `nats-streaming-cli` in `pkg` dir to your `$PATH`  

### Docker
If you prefer to use docker:
1. Clone this repo
1. cd to repo path
1. Run `docker build -t nats-streaming-cli .`


## Usage

```
nats-streaming-cli [options]
Options:
      -h   --host <string>           set nats-streaming server host (default: localhost)
      -p   --port <int>              set nats-streaming server port (default: 4222)
      -m   --http_port <int>         set http monitoring port (default: 8222)
      -cid --cluster_id <string>     set the server cluster ID, if not set, we will try to get cluster id from server monitor endpoint
           --client_id  <string>     specific client id cli use ,if not set, we will use a random client id
      -v   --version                 show version
```
**Notice:** cluster id can be get from  nats-streaming server's [monitor http server](https://github.com/nats-io/nats-streaming-server#monitoring)

If the monitor http server is disabled ,you should supply the cluster id

## Commands

```
show
  ├─ channel CHANNEL                show a specific channel info
  ├─ channels                       show all channel info
  ├─ server                         show server info
  ├─ store                          show store info
  └─ clients                        show clients info

pub CHANNEL MESSAGE                 publish MESSAGE to CHANNEL

list CHANNEL [START [LIMIT]]        list LIMIT(unlimit if not specific) messages in CHANNEL start at START (0 if not specific)

help                                print this help message

exit                                exit
```
**Notice:** `show` commands rely on the [monitor http server](https://github.com/nats-io/nats-streaming-server#monitoring) 
