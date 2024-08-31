Определить версию undetected

pip show undetected-chromedriver

ВЕРСИИ undetected

(newenv) gaz358@Home-PC:~/mygo/sobes_pars_new$ pip install undetected-chromedriver==2.9.0
ERROR: Ignored the following yanked versions: 3.1.5
ERROR: Could not find a version that satisfies the requirement undetected-chromedriver==2.9.0 (from versions: 1.3.5, 1.3.6, 1.3.7, 1.4.0, 1.4.2, 1.5.0, 1.5.1, 1.5.2, 2.0b0, 2.0.0, 2.0.1, 2.0.2, 2.1.0, 2.1.1, 2.1.2, 2.2.0, 2.2.1, 2.2.7, 3.0.0, 3.0.1, 3.0.2, 3.0.3, 3.0.4, 3.0.5, 3.0.6, 3.1.0rc1, 3.1.0, 3.1.1, 3.1.2, 3.1.3, 3.1.5.post1, 3.1.5.post2, 3.1.5.post3, 3.1.5.post4, 3.1.6, 3.1.7, 3.2.0a0, 3.2.0, 3.2.1, 3.4, 3.4.1, 3.4.2, 3.4.4, 3.4.5, 3.4.6, 3.4.7, 3.5.0, 3.5.1a0, 3.5.1, 3.5.2, 3.5.3, 3.5.4, 3.5.5)

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




