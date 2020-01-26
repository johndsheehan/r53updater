## Update AWS Route53 record

go program to

- check external public ip every 600 seconds
- update `route53` entry if external ip has changed

(uses <https://api.ipify.org> to determine external ip)

## Requirements

- aws account and cli credentials (better to create a specific `IAM` role for this)
- docker

## Quickstart

- populate `~/.aws/config`, `~/.aws/credentials`
- `cd cmd/r53updater; ./r53updater  --fqdn <VALUE>  --zone <VALUE>`
- alternatively, fqdn and zone can be stored in enviroment variables; `AWSFQDN`, and `AWSZONE`


## Docker

- `make docker`
- `docker run  --rm  --init  --read-only  --name='r53updater'  --net='host'  --env AWSFQDN=<VALUE>  --env AWSZONE=<VALUE>  -v /full/path/to/.aws:/home/user/.aws  r53updater:<VALUE>`
- note: permissions on `~/.aws/config` and `~/.aws/credentials` should at least be `444`. This is because when the directory is mounted in the container, its files will (probably) have a different `uid` and `gid` to the user running in the container.
