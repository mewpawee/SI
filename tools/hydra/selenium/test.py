from selenium import webdriver
from selenium.webdriver.common.by import By
import time

URL = "https://www.facebook.com"

name = URL.split('.')[1]
chrome_path = '/bin/chromedriver'
chrome_options = webdriver.ChromeOptions()
# chrome_options.add_argument("--headless")
# chrome_options.add_argument("--disable-gpu")
# chrome_options.add_argument('--lang=es-US')
# prefs = {'intl.accept_languages': 'es-US'}
# chrome_options.add_experimental_option("prefs", prefs)
chrome_options.add_argument("--window-size=1024x768")
chrome_options.add_argument("--headless")
with webdriver.Chrome(executable_path=chrome_path, options=chrome_options) as driver:
    driver.maximize_window()  # For maximizing window
    driver.get(URL)
    driver.save_screenshot(f"./{name}.png")
    username = driver.find_element_by_name("email")
    password = driver.find_element_by_name("pass")
    username.send_keys('pawee_tanti@hotmail.com')
    password.send_keys('LFWSEp6BXa5R49')
    button = driver.find_element_by_name("login")
    # button = driver.find_element_by_xpath(
    #     '//button[normalize-space()="เข้าสู่ระบบ"]')
    button.click()
    time.sleep(3)
    driver.save_screenshot(f"./{name}.png")
