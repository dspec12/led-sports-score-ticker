#!/bin/bash

echo "Updating Ticker Repo"
git -C /home/pi/led-sports-score-ticker fetch origin && git -C /home/pi/led-sports-score-ticker reset --hard origin/master

exit