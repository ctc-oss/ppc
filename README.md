ppc
==

### server
_ private cloud

#### server config

| Field                           | Description                                      | Default Value
| ------------------------------- | ------------------------------------------------ | --------------- |
| BROKER_HOSTNAME                 | Hostname of the event broker                     | localhost       |

### cli

```
$ particle 
 NAME:
    particle - A new cli application
 
 USAGE:
    particle [global options] command [command options] [arguments...]
 
 COMMANDS:
    call     Call a particular function on a device
    help, h  Shows a list of commands or help for one command
 
 GLOBAL OPTIONS:
    --help, -h  show help (default: false)
```

#### cli config

| Field                           | Description                                      | Default Value
| ------------------------------- | ------------------------------------------------ | --------------- |
| CLOUD_API                       | URI of cloud server                              | localhost:9000  |


#### cli config



### build

`make`