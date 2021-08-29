
###

## GCP ##
imam@imam-mv:~$ export GOOGLE_APPLICATION_CREDENTIALS=


## AWS ##
imam@imam-mv:~$ aws configure


## Digitalocean ##



## How to Use ##
```bash
imam@imam-mv:~$ go build -o ectl
imam@imam-mv:~$ ./ectl aws instance --action=start --region=ap-southeast-1 i-xxx i-xxx
imam@imam-mv:~$ ./ectl gcp instance --action=start --region=sgp i-xxx i-xxx
imam@imam-mv:~$ ./ectl digitalocean instance --action=

```
