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
options.chain_length = 1
options.parallel = 1
options.hardware_mapping = "adafruit-hat-pwm"
options.drop_privileges=False
font_filename = "9x15B.bdf"
ticker_speed = 0.03


def grab_team_info(team):
    base_url = "https://gophillyscores.herokuapp.com/"
    team_info = base_url + team

    try:
        team_info = requests.get(team_info)
        return team_info.text
    
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


def led_scroll_text(team, text_color):
    matrix = RGBMatrix(options=options)
    offscreen_canvas = matrix.CreateFrameCanvas()
    cwd = os.path.dirname(__file__)
    font_path = os.path.join(cwd, font_filename)
    font = graphics.Font()
    font.LoadFont(font_path)
    textColor = graphics.Color(*text_color)
    pos = offscreen_canvas.width
    print(f'Fetching info for {team}')
    scroll_text = grab_team_info(team)
    count = 0

    while True:
        offscreen_canvas.Clear()
        len = graphics.DrawText(offscreen_canvas, font, pos, 13, textColor, scroll_text)
        pos -= 1
        if pos + len < 0:
            pos = offscreen_canvas.width
            break
 
        time.sleep(ticker_speed)
        offscreen_canvas = matrix.SwapOnVSync(offscreen_canvas)
 
 
if __name__ == "__main__":
    eagles_color = 4, 106, 56
    flyers_color = 247,73,2
    phillies_color = 232,24,40
    psu_color = 4,30,66
    sixers_color = 237, 23, 76


    print("Starting...")
    while True:
        led_scroll_text("eagles", eagles_color)
        led_scroll_text("flyers", flyers_color)
        led_scroll_text("phillies", phillies_color)
        led_scroll_text("psu", psu_color)
        led_scroll_text("sixers", sixers_color)
