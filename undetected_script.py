import undetected_chromedriver as uc
from selenium.webdriver.chrome.options import Options
import time

# Настройки Chrome
options = Options()
options.binary_location = '/usr/bin/google-chrome'

# Указание пути к ChromeDriver вручную
driver_path = '/home/gaz358/mygo/sobes_pars_new/chromedriver'

try:
    driver = uc.Chrome(driver_executable_path=driver_path, options=options)
    driver.get('https://www.galco.com/manufacturer')
    print("Title:", driver.title)
    print("Current URL:", driver.current_url)
    
    time.sleep(60)  # Увеличьте время ожидания, если нужно
    
    # Получаем HTML-код страницы
    page_source = driver.page_source
    
    # Запись HTML-кода в файл
    with open('page_source.html', 'w') as file:
        file.write(page_source)
    print("HTML code written to page_source.html")
  
finally:
    if 'driver' in locals():
        driver.quit()
