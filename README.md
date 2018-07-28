# redis-health

Can be used to check the health of your Redis instance.

We use it for our Redis-in-Docker-on-Nomad environment to make sure rolling deployment of our Redis clusters are safe and without data loss

## Current checks

- Can connect to Redis
- Can run `info` command
- `loading` is `0` - Redis can't serve traffic while loading data from Disk
- `master_sync_in_progress` is `0` - Redis slaves can't serve traffic while getting initial data from its Master node
- `master_link_status` is `up` - Redis instances that have lost connection to their Redis master should not serve data

## Config

- `REDIS_ADDR` Redis IP+Port to connect to - example: `127.0.0.1:6379`
- `REDIS_PASS` Redis password to connect with - example: `SoSecure`
