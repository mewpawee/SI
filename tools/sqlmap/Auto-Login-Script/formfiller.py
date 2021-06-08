from selenium import webdriver
from selenium.webdriver.common.by import By
from browsermobproxy import Server
import time
import requests
from bs4 import BeautifulSoup as bs

#def is_vulnerable(response):
#    """A simple boolean function that determines whether a page 
#    is SQL Injection vulnerable from its `response`"""
#    errors = {
#        # MySQL
#        "you have an error in your sql syntax;",
#        "warning: mysql",
#        # SQL Server
#        "unclosed quotation mark after the character string",
#        # Oracle
#        "quoted string not properly terminated",
#    }
#    for error in errors:
#        # if you find one of these errors, return True
#        if error in response.content.decode().lower():
#            return True
#    # no error detected
#    return False

#def get_all_forms(url):
#    """Given a `url`, it returns all forms from the HTML content"""
#    soup = bs(s.get(url).content, "html.parser")
#    return soup.find_all("form")


#def get_form_details(form):
#    """
#    This function extracts all possible useful information about an HTML `form`
#    """
#    details = {}
#    # get the form action (target url)
#    try:
#        action = form.attrs.get("action").lower()
#    except:
#        action = None
#    # get the form method (POST, GET, etc.)
#    method = form.attrs.get("method", "get").lower()
#    # get all the input details such as type and name
#    inputs = []
#    for input_tag in form.find_all("input"):
#        input_type = input_tag.attrs.get("type", "text")
#        input_name = input_tag.attrs.get("name")
#        input_value = input_tag.attrs.get("value", "")
#        inputs.append({"type": input_type, "name": input_name, "value": input_value})
#    # put everything to the resulting dictionary
#    details["action"] = action
#    details["method"] = method
#    details["inputs"] = inputs
#    return details

#URL = "https://www.facebook.com"
URL = "http://128.199.166.185/DVWA/login.php"
settingURL = "http://128.199.166.185/DVWA/security.php"
SQLiURL = "http://128.199.166.185/DVWA/vulnerabilities/sqli/"
name = URL.split('.')[1]
chrome_path = './chromedriver'
gecko_path = './geckodriver'
browsermob_path = 'D:/capstone/Auto-Login-Script/browsermob-proxy-2.1.4/bin/browsermob-proxy'
server = Server(browsermob_path)
server.start()
proxy = server.create_proxy()
#chrome_options = webdriver.ChromeOptions()
# chrome_options.add_argument("--headless")
# chrome_options.add_argument("--disable-gpu")
# chrome_options.add_argument('--lang=es-US')
# prefs = {'intl.accept_languages': 'es-US'}
# chrome_options.add_experimental_option("prefs", prefs)
#chrome_options.add_argument("--window-size=1024x768")
#chrome_options.add_argument("--headless")
#chrome_options.add_argument('--proxy-server=%s' % proxy.proxy)
#chrome_options.add_argument('--ignore-certificate-errors')

profile  = webdriver.FirefoxProfile()
profile.set_proxy(proxy.selenium_proxy())
driver = webdriver.Firefox(firefox_profile=profile)
#proxy.new_har("", options={'captureHeaders': True, 'captureContent': True})
#driver.get("your_url")
#proxy.wait_for_traffic_to_stop(1, 60)
#for ent in proxy.har['log']['entries']:
#    print(ent)

#server.stop()
#driver.quit()


with webdriver.Firefox(executable_path=gecko_path, firefox_profile=profile) as driver:
    #print("------------------------------------------------------------------------------------------------------------------------------------------------")
    driver.maximize_window()  # For maximizing window
    #proxy.new_har("", options={'captureHeaders': True, 'captureContent': True})
    driver.get(URL)
    driver.save_screenshot(f"./{name}.png")
    #username = driver.find_element_by_name("email")
    userID = driver.find_element_by_name("username")
    password = driver.find_element_by_name("password")
    #username.send_keys('<!mmmmmmm!>')
    userID.send_keys('admin')
    password.send_keys('password')
    button = driver.find_element_by_name("Login") #("login")
    # button = driver.find_element_by_xpath(
    #     '//button[normalize-space()="เข้าสู่ระบบ"]')
    button.click()
    time.sleep(3)
    #driver.save_screenshot(f"./{name}.png")
    driver.get(settingURL)
    proxy.wait_for_traffic_to_stop(1, 60)
    driver.find_element_by_xpath("//select[security='low']").click()
    driver.get(settingURL)
    #for ent in proxy.har['log']['entries']:
    #   print(ent)
    #server.stop()
    #driver.quit()
   # print("..........................................................................")
    
#    #result = proxy.har
#    #print(result)
#    #for entry in result['log']['entries']:
#    #    print("******************************************************************************************************************************")
#    #    print(entry['request']['url'])
#    for i in range(len(proxy.har['log']['entries'])):
#        print("******************************************************************************************************************************")
#        print(proxy.har['log']['entries'][i]['request']['headers'])
#        print(proxy.har['log']['entries'][i]['request']['content'])
#    server.stop()
#    driver.quit()
    #=======================================================================================================================
    #server = Server('path_to_executable')
    #server.start()
    #proxy = server.create_proxy()
    #profile  = webdriver.FirefoxProfile()
    #profile.set_proxy(proxy.selenium_proxy())
    #driver = webdriver.Firefox(firefox_profile=profile)
    #proxy.new_har("file_name", options={'captureHeaders': True, 'captureContent': True})
    #driver.get("your_url")
    #proxy.wait_for_traffic_to_stop(1, 60)
    #for ent in proxy.har['log']['entries']:
    #    print(ent)

    #server.stop()
    #driver.quit()



    #---------------------------for chrome-------------------------------------------------------------------
    #reff : https://stackoverflow.com/questions/52394408/how-to-use-chrome-profile-in-selenium-webdriver-python-3
    #from selenium import webdriver
    #from selenium.webdriver.chrome.options import Options

    #options = webdriver.ChromeOptions()
    #options.add_argument("user-data-dir=C:\\Users\\AtechM_03\\AppData\\Local\\Google\\Chrome\\User Data\\Default")
    #driver = webdriver.Chrome(executable_path=r'C:\path\to\chromedriver.exe', chrome_options=options)
    #driver.get("https://www.google.co.in")

    #------------------------custom chrome profile------------------------------------------------------------

    #from selenium import webdriver
    #from selenium.webdriver.chrome.options import Options

    #options = Options()
    #options.add_argument("user-data-dir=C:\\Users\\AtechM_03\\AppData\\Local\\Google\\Chrome\\User Data\\Profile 2")
    #driver = webdriver.Chrome(executable_path=r'C:\path\to\chromedriver.exe', chrome_options=options)
    #driver.get("https://www.google.co.in")
    