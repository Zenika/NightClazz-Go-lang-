% title: IPFS
% subtitle: The Permanent Web
% author: Emile Vauge
% author: Vincent Demeester
% thankyou: Thanks everyone!
% contact: <a href="https://twitter.com/vdemeest">@vdemeest</a>
% contact: <a href="https://twitter.com/emilevauge">@emilevauge</a>

---
title: HTTP(s) is broken
build_lists: true

- Centralized
- Censorship issue
- Unsecured, even with SSL (CVE-2016-0701, CVE-2015-1793, CVE-2015-0291...)
- Inefficient
- No history ⛔ *404 error*

---
title: Enters IPVS
subtitle: Distributed Internet
class: img-top-center

<img height=195 src=https://upload.wikimedia.org/wikipedia/en/b/ba/Centralised-decentralised-distributed.png />

> With HTTP, you search for **locations**. With IPFS, you search for **content**.

---
title: How it works?

- DHT *you know Bittorrent?*
- Git *hashes everywhere!*
<a href="https://neocities.org/img/neocitieslogo.svg">https://neocities.org/img/neocitieslogo.svg</a> > <a href="https://ipfs.io/ipfs/QmXGTaGWTT1uUtfSb2sBAvArMEVLK4rQEcQg5bv7wwdzwU">QmXGTaGWTT1uUtfSb2sBAvArMEVLK4rQEcQg5bv7wwdzwU</a>
- Efficient *Only 20 hops for a network of 10,000,000*

---
title: Bootstrap

<pre class="prettyprint" data-lang="sh">
$ ipfs init
initializing ipfs node at /Users/jbenet/.go-ipfs
generating 2048-bit RSA keypair...done
peer identity: Qmcpo2iLBikrdf1d6QU6vXuNb6P7hwrbNPW9kLAH8eG67z
to get started, enter:

 ipfs cat /ipfs/QmPXME1oRtoT627YKaDPDQ3PwA8tdP9rWuAAweLzqSwAWT/readme
</pre>

<pre class="prettyprint" data-lang="sh">
$ ipfs cat /ipfs/QmPXME1oRtoT627YKaDPDQ3PwA8tdP9rWuAAweLzqSwAWT/readme
 Hello and Welcome to IPFS!

██╗██████╗ ███████╗███████╗
██║██╔══██╗██╔════╝██╔════╝
██║██████╔╝█████╗  ███████╗
██║██╔═══╝ ██╔══╝  ╚════██║
██║██║     ██║     ███████║
╚═╝╚═╝     ╚═╝     ╚══════╝
</pre>

---
title: Bootstrap

<pre class="prettyprint" data-lang="sh">
$ ipfs daemon
Initializing daemon...
API server listening on /ip4/127.0.0.1/tcp/5001
Gateway server listening on /ip4/127.0.0.1/tcp/8080
</pre>

<pre class="prettyprint" data-lang="sh">
$ echo "I <3 IPFS -$(whoami)" | ipfs add -q
QmZnEKRTo2bXuFdTqnK45wBe9j6sffceBtuvcKRNbZxM4e
$ curl "http://127.0.0.1:8080/ipfs/QmZnEKRTo2bXuFdTqnK45wBe9j6sffceBtuvcKRNbZxM4e"
</pre>
