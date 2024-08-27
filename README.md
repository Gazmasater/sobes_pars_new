

ГУГЛ ХРОМ

https://mirror.cs.uchicago.edu/google-chrome/pool/main/g/google-chrome-stable/

Чтобы установить undetected-chromedriver, выполните следующие шаги:

1. Установка виртуального окружения (необязательно, но рекомендуется)
Создайте и активируйте виртуальное окружение для изоляции зависимостей вашего проекта:

# Создание виртуального окружения
python3 -m venv myenv

# Активация виртуального окружения
source myenv/bin/activate  # На Linux или macOS
myenv\Scripts\activate  # На Windows
2. Установка пакета undetected-chromedriver
После активации виртуального окружения установите undetected-chromedriver через pip:

pip install undetected-chromedriver
3. Установка необходимых зависимостей
undetected-chromedriver требует selenium для работы. Эта библиотека обычно устанавливается автоматически вместе с undetected-chromedriver. Если по какой-то причине selenium не установился, вы можете сделать это вручную:

pip install selenium
Также для автоматического управления ChromeDriver рекомендуется установить webdriver-manager:

pip install webdriver-manager




