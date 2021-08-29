
###

```bash
imam@imam-mv:~$ go build -o ectl
imam@imam-mv:~$ ./ectl aws instance --action=start --region=ap-southeast-1 i-xxx i-xxx
imam@imam-mv:~$ ./ectl gcp instance --action=start --region=sgp i-xxx i-xxx
imam@imam-mv:~$ ./ectl digitalocean instance --action=

```



export GOOGLE_APPLICATION_CREDENTIALS=/home/imam/learning-terraform/gcp-belajarterraform/instance-simple/ardent-gearbox-editor-learning-terraform.json