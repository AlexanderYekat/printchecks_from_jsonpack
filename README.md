нужно папку ./jsons/works/закутнуть файлы json заданий
в этой же папке в файле (создать этот файл, ксли не создан) connection.txt укзазать
номер порта соединения с кассовым аппаратом
в файле printed.txt одним стобиком можно указать номера чеков, по которым не надо
бить чеки. Номер чека берётся из названия файла json задания 7281440500582818_2792.json,
где число после прочерка и есть номер чека (по которому мы хотим длеать исправление)

запустить программу. Она будет пробивать все json задания в папке ./jsons/works/ пропуская чеки из файла printed.txt.
при каждом успешном напечатавании задания в файл printed.txt будет добавляться число - номер чека (json задания), который
берётся из названия json задания 7281440500582818_2792.json (число после прочекрка). И при повтороном запуске программы
чек повторно пробиваться поэтому не будет.
