# homelab-webhook

## Architecture

apps -> homelab-webhook -> serverchan -> wechat

## Setup
```bash
SERVER_CHAN_SEND_KEY=<your serverchan key>
docker run -itd --name homelab-webhook \
  --restart=unless-stopped \
  -p 8080:8080 \
  -e SERVER_CHAN_SEND_KEY=${SERVER_CHAN_SEND_KEY} \
  futuretea/homelab-webhook:latest
```

## Uptime-Kuma notify setting
![setup.png](setup.png)

## Harvester post install notify setting
Refer to https://github.com/harvester/harvester-installer/pull/57

```yaml
install:
  webhooks:
    - event: SUCCEEDED
      method: POST
      url: http://<homelab-webhook>:8080/harvester-serverchan/SUCCEEDED
      headers:
        Content-Type:
           - 'application/json; charset=utf-8'
      payload: |
        {
          "hostname": "{{.Hostname}}",
          "ipv4": "{{.IPAddrV4}}",
          "ipv6": "{{.IPAddrV6}}",
          "mac": "{{.MACAddr}}"
        }
```