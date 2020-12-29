# LED-Sports-Score-Ticker

Project consists of the tickers that run on a Raspi. These tickers consume text from either the Philly Scores API or the Scores Scraper script.


* Philly Scores API - API app written in Go. Returns formatted strings for major philly sports teams. 

*  Scores Scraper - Python script that scrapes ESPN APIs for the latest scores then writes the formated string to a text file in S3.

* Ticker - This folder contains different tickers (python scripts) that run on a Raspberry Pi. It grabs strings from a http endpoint and writes to LED panels using the [Henner Zeller RPi RGB LED Matrix](https://github.com/hzeller/rpi-rgb-led-matrix/) library.

<br></br>
### Demo:
[![Demo](https://i9.ytimg.com/vi/Lv7JURRsOHY/mq2.jpg?sqp=COyXrv8F&rs=AOn4CLCsGIur0KjiWUuSvLwIqfxeAueSFA)](https://youtu.be/Lv7JURRsOHY "Demo")
