#!/usr/bin/env python3
import sys
import os
import time
from rgbmatrix import RGBMatrix, RGBMatrixOptions, graphics


# Configuration for the matrix
options = RGBMatrixOptions()
options.scan_mode = 0
options.pwm_lsb_nanoseconds = 130
options.pwm_bits = 11
options.show_refresh_rate = 0
options.gpio_slowdown = 2
options.rows = 16
options.chain_length = 2
options.parallel = 1
options.hardware_mapping = "adafruit-hat-pwm"
ticker_speed = 0.03


def grab_scores():
    scores = os.popen(
        "curl https://led-sports-score-ticker.s3.amazonaws.com/scores.txt"
    ).read()
    return scores


def led_scroll_text():
    matrix = RGBMatrix(options=options)
    offscreen_canvas = matrix.CreateFrameCanvas()
    font = graphics.Font()
    font.LoadFont("9x15B.bdf")
    textColor = graphics.Color(text_color)
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
