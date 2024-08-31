import undetected_chromedriver as uc
from selenium.webdriver.chrome.options import Options
import time 

# Настройки Chrome
options = Options()
options.binary_location = '/usr/bin/google-chrome'
# options.add_argument('--headless')  # Если хотите запускать браузер в фоновом режиме

# Указание пути к ChromeDriver вручную
driver_path = '/home/gaz358/mygo/sobes_pars_new/chromedriver'

# Создание экземпляра драйвера с указанием пути к драйверу
try:
    driver = uc.Chrome(driver_executable_path=driver_path, options=options)
    driver.get('https://www.galco.com')
    print(driver.title)

    time.sleep(500)

  

finally:
    if 'driver' in locals():
        driver.quit()
