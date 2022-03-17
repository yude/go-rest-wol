# go-rest-wol
Wake on LAN over the Internet via the computer on your LAN
## Setup
1. Create `docker-compose.yml` and copy & paste the below:
  ```yaml
  version: '3'
  services:
    wol:
      image: ghcr.io/yude/go-rest-wol:master
      network_mode: host
      volumes:
        - type: bind
          source: "./hosts.csv"
          target: "/app/hosts.csv"
  ```
2. Create `hosts.csv` and copy & paste the below, and edit as you like.
  ```
  name,mac,ip
  mycomputer,64-07-2D-BB-BB-BF,192.168.0.254:9
  ```
3. Run `docker-compose up -d`. To update, run `docker-compose pull`.

## License
MIT License.