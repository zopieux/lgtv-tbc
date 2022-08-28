# LG TV: Take Back Control

A replacement HTTP server for LG webOS bootstrap mechanism that prevents unwanted forced updates, 
disables bloatware and forbids phoning home with unique device IDs.

## Setup

You'll need one or multiple machines running: 

* a DNS server to poison LG-controlled domains to your own HTTP server
* an HTTP server reverse-proxying `lgtv-tbc`, to terminate SSL 
* and of course `lgtv-tbc` itself.  

The instructions below are not completely copy-pastable, you'll need some fair amount of understanding to adapt them to your setup.

1. Have the TV use a DNS server you control. It might be possible to do so using the TV network settings, but personally 
   I have my own home DHCP server. I make it announce my DNS server as authoritative.
   To achieve this with `dnsmasq`, use a config similar to:

   ```ini
   # No DNS, only DHCP.
   port=0
   except-interface=lo
   dhcp-authoritative
   # Gateway IP.
   dhcp-option=3,<gateway IP here>
   # DNS server is ourselves.
   # Change 0.0.0.0 to the machine running the DNS server, if different from dnsmasq's.
   dhcp-option=6,0.0.0.0
   ```
   
1. Configure your DNS server to poison the following domains to resolve to the machine running the HTTP server:
   
   `*.lgtvsdp.com`, `*.lgappstv.com`, `*.lge.com`, `*.lgsmartad.com`
   
   but *only* for requests emanating from the LG TV, or add a pass-thru exception for the machine running `lgtv-tbc`, so 
   that `lgtv-tbc` itself can correctly resolve these domains when forwarding some of the non-problematic requests.

   To achieve this using a [Response Policy Zone](https://en.wikipedia.org/wiki/Response_policy_zone) on BIND,
   use something similar to:

   ```zonefile
   ; (SOA header here)
   ; This passthru setup is not necessary if lgtv-tbc resolves through a different, non-poisoned DNS server.
   ; If lgtv-tbc runs on the same machine as this DNS server, allow localhost/16 to passthru.
   16.0.0.0.127.rpz-client-ip CNAME rpz-passthru.
   ; If lgtv-tbc runs on another machine, say 192.168.1.20/32, passthru that instead:
   ; 32.20.1.168.192.rpz-client-ip CNAME rpz-passthru.
   *.lgtvsdp.com     IN      A       <IPv4 where HTTP server runs>
   *.lgappstv.com    IN      A       <IPv4 where HTTP server runs>
   *.lge.com         IN      A       <IPv4 where HTTP server runs>
   *.lgsmartad.com   IN      A       <IPv4 where HTTP server runs>
   *.lgtvsdp.com     IN      AAAA    <IPv6 where HTTP server runs>
   *.lgappstv.com    IN      AAAA    <IPv6 where HTTP server runs>
   *.lge.com         IN      AAAA    <IPv6 where HTTP server runs>
   *.lgsmartad.com   IN      AAAA    <IPv6 where HTTP server runs>
   ```

1. Run the `lgtv-tbc` binary, using the `-addr` flag to specify the listen port.
   Reuse it in your reverse-proxy, as described below.

   ```ini
   # Stub systemd service definition.
   [Service]
   ExecStart=/path/to/lgtv-tbc -addr :8765
   ```

1. Create a self-signed certificate for the poisoned domains. It will be invalid, since you obviously cannot emit 
   genuine certificates for LG-controlled domains. If *you* can, please ask your management to stop creating privacy-invading 
   walled gardens and allow people who buy LG TVs to use them as they see fit.
   We're lucky here as webOS *does not validate the certificate* for these particular start-up requests, allowing us 
   to intercept and replace the replies.

   ```shell
   $ openssl req -x509 -newkey rsa:4096 -sha256 -days 3560 -nodes -keyout lgtv.key -out lgtv.crt \
       -subj '/CN=lgtvsdp.com' -extensions san \
       -config <(printf '[req]\ndistinguished_name=req\n[san]\nsubjectAltName=DNS:*.lgtvsdp.com,DNS:*.lge.com,DNS:*.lgsmartad.com,DNS:*.lgappstv.com\n')
   ```
   
1. Configure your HTTP server with virtual hosts for the poisoned domains *on both HTTP and HTTPS*.
   To achieve this using nginx, use a config similar to:

   ```nginx
   server {
     listen 0.0.0.0:443 ssl;
     listen [::]:443 ssl;
     listen 0.0.0.0:80;
     listen [::]:80;
     server_name *.lgtvsdp.com *.lgappstv.com *.lge.com *.lgsmartad.com;
     ssl_certificate /path/to/lgtv.crt;
     ssl_certificate_key /path/to/lgtv.key;
     location / {
       proxy_pass http://<ip where lgtv-tbc runs>:<lgtv-tbc port eg. 8765>;
     }
   }
   ```

If all goes well, `lgtv-tbc` should start receiving requests when you power on your TV. 
You might want to add the `-notify` flag for debugging, which creates a TV notification when `lgtv-tbc` successfully
intercepts the start-up request.

## License

[GNU General Public License v3.0 or later](https://spdx.org/licenses/GPL-3.0-or-later.html).
