# cryptopro-cert-export
Экспорт контейнеров JCP в файлики из реестра

Экспортированная папка может быть установлена в JCP HDIMAGE хранилище

# Usage

### Аргументы 
* `-out` папка вывода, по умолчанию - '.'
* `-sid` SID Пользователя (eg S-1-5-21-2408434269-2496968396-1739759940-3298), по умолчанию - '' (Все пользователи)
* `-container` Контейнер, по умолчанию - '' (Все контейнеры)

### Пример
```
PS C:\users\user\Documents> .\main.exe -out cpcerts2
2023/08/21 21:42:07 Arg: container     = ''
2023/08/21 21:42:07 Arg: userSid       = ''
2023/08/21 21:42:07 Arg: outputFolder  = 'cpcerts2'
2023/08/21 21:42:07 Opening cryptopro root (WOW64)...
2023/08/21 21:42:07 Reading users nodes...
2023/08/21 21:42:07 Found user S-1-5-18 ...
2023/08/21 21:42:07 Found user S-1-5-21-2408434269-2496968396-1739759940-3298 ...
2023/08/21 21:42:07 Found user S-1-5-21-2408434269-2496968396-1739759940-6064 ...
2023/08/21 21:42:07 User nodes count - 3
2023/08/21 21:42:07 Reading containers...
2023/08/21 21:42:07 Cant read user Keys dir - HKLM\SOFTWARE\WOW6432Node\Crypto Pro\Settings\Users\S-1-5-18 / Keys
2023/08/21 21:42:07 Reading containers...
2023/08/21 21:42:07 Found Container 185c352d6-d22d-9a42-1f9b-52a7e5b6ea3 ...
2023/08/21 21:42:07 Found Container f5cde7a28-803b-d030-4f7e-9a0079e259f ...
2023/08/21 21:42:07 Full container name - HKLM\SOFTWARE\WOW6432Node\Crypto Pro\Settings\Users\S-1-5-21-2408434269-2496968396-1739759940-3298\Keys\185c352d6-d22d-9a42-1f9b-52a7e5b6ea3
2023/08/21 21:42:07 Saving cpcerts2\185c352d.000\name.key file
2023/08/21 21:42:07 Saving cpcerts2\185c352d.000\header.key file
2023/08/21 21:42:07 Saving cpcerts2\185c352d.000\primary.key file
2023/08/21 21:42:07 Saving cpcerts2\185c352d.000\masks.key file
2023/08/21 21:42:07 Saving cpcerts2\185c352d.000\primary2.key file
2023/08/21 21:42:07 Saving cpcerts2\185c352d.000\masks2.key file
2023/08/21 21:42:07 Full container name - HKLM\SOFTWARE\WOW6432Node\Crypto Pro\Settings\Users\S-1-5-21-2408434269-2496968396-1739759940-3298\Keys\f5cde7a28-803b-d030-4f7e-9a0079e259f
2023/08/21 21:42:07 Saving cpcerts2\f5cde7a2.000\name.key file
2023/08/21 21:42:07 Saving cpcerts2\f5cde7a2.000\header.key file
2023/08/21 21:42:07 Saving cpcerts2\f5cde7a2.000\primary.key file
2023/08/21 21:42:07 Saving cpcerts2\f5cde7a2.000\masks.key file
2023/08/21 21:42:07 Saving cpcerts2\f5cde7a2.000\primary2.key file
2023/08/21 21:42:07 Saving cpcerts2\f5cde7a2.000\masks2.key file
2023/08/21 21:42:07 Reading containers...
2023/08/21 21:42:07 Found Container rnd-2-7F25-A318-B78A-3677-EA1B-127C-F478 ...
2023/08/21 21:42:07 Found Container rnd-D-6F82-4EC4-8762-24B2-7D11-BD49-F9A8 ...
2023/08/21 21:42:07 Found Container rnd-F-9F1B-78B0-2B35-22CD-CBCC-8FA7-519B ...
2023/08/21 21:42:07 Full container name - HKLM\SOFTWARE\WOW6432Node\Crypto Pro\Settings\Users\S-1-5-21-2408434269-2496968396-1739759940-6064\Keys\rnd-2-7F25-A318-B78A-3677-EA1B-127C-F478
2023/08/21 21:42:07 Saving cpcerts2\rnd-2-7F.000\name.key file
2023/08/21 21:42:07 Saving cpcerts2\rnd-2-7F.000\header.key file
2023/08/21 21:42:07 Saving cpcerts2\rnd-2-7F.000\primary.key file
2023/08/21 21:42:07 Saving cpcerts2\rnd-2-7F.000\masks.key file
2023/08/21 21:42:07 Saving cpcerts2\rnd-2-7F.000\primary2.key file
2023/08/21 21:42:07 Saving cpcerts2\rnd-2-7F.000\masks2.key file
2023/08/21 21:42:07 Full container name - HKLM\SOFTWARE\WOW6432Node\Crypto Pro\Settings\Users\S-1-5-21-2408434269-2496968396-1739759940-6064\Keys\rnd-D-6F82-4EC4-8762-24B2-7D11-BD49-F9A8
2023/08/21 21:42:07 Saving cpcerts2\rnd-D-6F.000\name.key file
2023/08/21 21:42:07 Saving cpcerts2\rnd-D-6F.000\header.key file
2023/08/21 21:42:07 Saving cpcerts2\rnd-D-6F.000\primary.key file
2023/08/21 21:42:07 Saving cpcerts2\rnd-D-6F.000\masks.key file
2023/08/21 21:42:07 Saving cpcerts2\rnd-D-6F.000\primary2.key file
2023/08/21 21:42:07 Saving cpcerts2\rnd-D-6F.000\masks2.key file
2023/08/21 21:42:07 Full container name - HKLM\SOFTWARE\WOW6432Node\Crypto Pro\Settings\Users\S-1-5-21-2408434269-2496968396-1739759940-6064\Keys\rnd-F-9F1B-78B0-2B35-22CD-CBCC-8FA7-519B
2023/08/21 21:42:07 Saving cpcerts2\rnd-F-9F.000\name.key file
2023/08/21 21:42:07 Saving cpcerts2\rnd-F-9F.000\header.key file
2023/08/21 21:42:07 Saving cpcerts2\rnd-F-9F.000\primary.key file
2023/08/21 21:42:07 Saving cpcerts2\rnd-F-9F.000\masks.key file
2023/08/21 21:42:07 Saving cpcerts2\rnd-F-9F.000\primary2.key file
2023/08/21 21:42:07 Saving cpcerts2\rnd-F-9F.000\masks2.key file
PS C:\users\user\Documents> dir cpcerts2


    Каталог: C:\users\user\Documents\cpcerts2


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----        21.08.2023     21:42                185c352d.000
d-----        21.08.2023     21:42                f5cde7a2.000
d-----        21.08.2023     21:42                rnd-2-7F.000
d-----        21.08.2023     21:42                rnd-D-6F.000
d-----        21.08.2023     21:42                rnd-F-9F.000


PS C:\users\user\Documents> dir .\cpcerts2\rnd-2-7F.000\


    Каталог: C:\users\user\Documents\cpcerts2\rnd-2-7F.000


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
-a----        21.08.2023     21:42           1726 header.key
-a----        21.08.2023     21:42             56 masks.key
-a----        21.08.2023     21:42             56 masks2.key
-a----        21.08.2023     21:42             44 name.key
-a----        21.08.2023     21:42             36 primary.key
-a----        21.08.2023     21:42             36 primary2.key


PS C:\users\user\Documents>
```

Что бы поставить в JCP, копируем полученную папку с сертификатом (называется примерно "`rnd-2-7F.000`") в JCP HDIMAGE хранилще (Обычно лежит в `${user.home}\Local Settings\Application Data\Crypto Pro\`)