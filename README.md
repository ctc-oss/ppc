ppc
==


#### server config

| Field                           | Description                                      | Default Value
| ------------------------------- | ------------------------------------------------ | --------------- |
| BROKER_URI                      | URI of the event broker                          | localhost:1883  |
| APP_PREFIX                      | App topic prefix                                 | xr              |
| EVENT_CHANNEL                   | Event Channel                                    | E               |
| FUNCTION_CHANNEL                | Function Channel                                 | F               |

### cli

```
$ polyform 
 NAME:
    polyform - A new cli application
 
 USAGE:
    polyform [global options] command [command options] [arguments...]
 
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


### build

`make`