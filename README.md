# SSA - Метод "Гусеница"

![gofer](https://i.ibb.co/72yJnQ5/gofer2.png)

## Запуск программы

После корректной установки Вам будет доступен терминал(Вид>Терминал).

Пропишите в терминале команду запуска программы:

```
$ make run
```

Во время выполнения программы выводится прогресс, который содержит дату, время, выполненный рассчёт.

### Отображение графиков в MatLab

Программа поддерживает экспорт данных в Matlab, посредством файлов формата xlsx.

В случае наличия на Вашем локальном компьютере MatLab, Вы можете отобразить выходные графики, запустив в MatLab файл plotting.m, который находится в папке "File_For_MatLab".

### Выходные данные

На данный момент выводятся следующие данные из следующих папок:

#### File_For_Matlab

Данная папка содержит директории со следующими данными:

1. Covariance matrix
2. Eigenvalues
3. Original time series and reconstruction sET12
4. Original time series and reconstruction sET34
5. Визуализация АКФ сингулярных троек для сегментов pw
6. Огибающие АКФ сингулярных троек sET12 сегментов pw
7. Нормированные АКФ сингулярных троек sET12 сегментов pw

#### files

| Название .xlsx | Описание                  | К-во значений |
| ---------------------- | --------------------------------- | ------------------------ |
| AcfNrm_sET12           | Нормированные АКФ | 102x200                  |
|                        |                                   |                          |

---

## Установка окружения

Для *корректной* работы данной программы необходимо установить:

* golang 1.16+ - Язык программирования Golang версии 1.16+;
* git - Система контроля версии;
* VS Code - Редактор кода.

### Windows

#### Golang

Для установки Golang под ОС Windows перейдите на [оффициальный сайт Golang](https://go.dev/dl/) и скачайте, нажав на кнопку "Microsoft Windows", далее установите в соответствии с установочным файлом.

#### git

Для установки системы контроля версии git под ОС Windows перейдите на [оффициальный сайт git](https://git-scm.com/download/win) и скачайте файл из раздела "Standalone Installer" в соответсвии с разрядностью Вашей системы. Далее проведите установку в соотвествии со скаченным установленным файлом.

#### VS Code

Для установки текстового редактора VS Code под ОС Windows перейдите на [оффициальный сайт VS Code](https://code.visualstudio.com/downloadhttps://code.visualstudio.com/download) и скачайте файл, нажав на кнопку "Windows". Далее проведите установку в соотвествии со скаченным установленным файлом. Допускается использование иного текстового редактора.

---

### Linux(Ubuntu)

#### Golang

Сначала обновите систему до самой новой версии:

```
$ sudo apt update
$ sudo apt upgrade
```

Затем для установки достаточно выполнить такую команду:

```
$ sudo apt install golang
```

На этом установка завершена. Для просмотра версии выполните:

```
$ go version
```

#### git

Для получения последней версии git выполните:

```
$ apt-get install gitapt-get install git
```

#### VS Code

Для установки текстового редактора VS Code можно поспользоваться встроенным центр приложений "Ubuntu Software" по запросу "Visuad Studio Code" или возпользоваться менеджером пакетов snap, прописав команду:

```
$ sudo snap install --classic codesudo snap install --classic code
```

---

## Установка проекта

Для установки проекта достаточно скачать и распоковать данный проект на локальный компьютер, однако стоит отдельно рассмотреть установку для разных операционных систем.

### Windows

Для уставноки рекомендуется использовать встроенную утилиту "Windows PowerShell", котрую можно найти в поиске программ на Вашем устройстве. Открываем утилиту "Windows PowerShell" и прописываем команду:

```
# git clone https://github.com/RB-PRO/ssa.git
```

После этого проект будет установлен на Ваш локальный комьютер. Пожалуйста, откройте его с помощью текстового редактора VS Code. Для этого откройте VS Code, на верхнем меню выберите "Файл>Открыть папку" и укажите путь к скаченной папке проекта. После этого для работы с проектом Вам будет необходим терминал, который можно открыть, выбрав на вернем меню "Вид>Терминал"

### Linux(Ubuntu)

Для установки проекта откройте терминал и перейдите в нужную для Вас директорию с помощью команд "cd" и "ls". Пропишите команду для копирования проекта:s

```
$ git clone https://github.com/RB-PRO/ssa.git
```

Откойте проект в VS Code:

```
$ code ssa/
```

---