#!/usr/bin/env python3
import sys
import os
import time
import requests
from rgbmatrix import RGBMatrix, RGBMatrixOptions, graphics


# Configuration for the matrix
options = RGBMatrixOptions()
options.scan_mode = 0
options.pwm_lsb_nanoseconds = 130
options.pwm_bits = 11
options.show_refresh_rate = 0
options.gpio_slowdown = 2
options.rows = 16
options.chain_length = 4
options.parallel = 1
options.hardware_mapping = "adafruit-hat-pwm"
font_filename = "9x15B.bdf"
text_color = 4, 106, 56
ticker_speed = 0.03


def grab_scores():
    url = "https://led-sports-score-ticker.s3.amazonaws.com/scores.txt"
    try:
        scores = requests.get(url)
        return scores.text
    except requests.exceptions.ConnectionError as e:
        print("Could not connect to endpoint:")
        print(e)
    except requests.exceptions.HTTPError as e:
        print("Http error:")
        print(e)
    except Exception as e:
        print("Unknown error:")
        print(type(e))
        print(e)


def led_scroll_text():
    matrix = RGBMatrix(options=options)
    offscreen_canvas = matrix.CreateFrameCanvas()
    cwd = os.path.dirname(__file__)
    font_path = os.path.join(cwd, font_filename)
    font = graphics.Font()
    font.LoadFont(font_path)
    textColor = graphics.Color(text_color[0], text_color[1], text_color[2])
    pos = offscreen_canvas.width
    scroll_text = grab_scores()
    count = 0

    while True:
        offscreen_canvas.Clear()
        len = graphics.DrawText(offscreen_canvas, font, pos, 13, textColor, scroll_text)
        pos -= 1
        if pos + len < 0:
            pos = offscreen_canvas.width
            count += 1

        if count >= 1:
            count = 0
            print("Refreshing scores...")
            scroll_text = grab_scores()

        time.sleep(ticker_speed)
        offscreen_canvas = matrix.SwapOnVSync(offscreen_canvas)


if __name__ == "__main__":
    print("Starting...")
    led_scroll_text()
