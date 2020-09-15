# LED-Sports-Score-Ticker

Project consists of two parts:

*  Score Scraper - This script will scrape ESPN APIs for the latest scores then writes the formated string to a text file in S3.

* Scores Ticker - This script runs on a Raspberry Pi. It grabs a string from a http endpoint and writes to a LED panels using [Henner Zeller RPi RGB LED Matrix.](https://github.com/hzeller/rpi-rgb-led-matrix/)