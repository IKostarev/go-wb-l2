# WB Tech: level # 2 (Golang)

## Паттерны проектирования (pattern)
~~~
- Порождающие шаблоны (creational) — шаблоны проектирования, которые абстрагируют процесс инстанцирования. Они позволяют сделать систему независимой от способа создания, композиции и представления объектов. Шаблон, порождающий классы, использует наследование, чтобы изменять инстанцируемый класс, а шаблон, порождающий объекты, делегирует инстанцирование другому объекту. (2, 6)

- Поведенческие шаблоны (behavioral patterns) — шаблоны проектирования, определяющие алгоритмы и способы реализации взаимодействия различных объектов и классов. (3, 4, 5, 7, 8)

- Структурные шаблоны (structural) определяют различные сложные структуры, которые изменяют интерфейс уже существующих объектов или его реализацию, позволяя облегчить разработку и оптимизировать программу. (1)

1. Паттерн facade - "фасад"
2. Паттерн builder - "строитель"
3. Паттерн visitor - "посетитель"
4. Паттерн command - "команда"
5. Паттерн chain of resp - "цепочка вызовов"
6. Паттерн factory method - "фабричный метод"
7. Паттерн strategy - "стратегия"
8. Паттерн state - "состояние"
~~~

## Задачи на разработку (develop)
~~~
1. Базовая задача. 
 - Создать go module на базе библиотеке NTP для получения точного времени.
 - Доступ по ссылке - https://gitlab.com/IKostarev/go-wb-l2/-/tree/master/develop/dev01
2. Задача на распаковку.
 - Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны
3. Утилита sort.
 - Отсортировать строки в файле по аналогии с консольной утилитой sort, поддерживаемые ключи: -k, -n, -r, -u, -M, -b, -c, -h
4. Поиск анаграмм по словарю.
 - Написать функцию поиска всех множеств анаграмм по словарю.
5. Утилита grep.
 - Реализовать утилиту фильтрации по аналогии с консольной утилитой grep. Ключи : A, B, C, c, i, v, F, n.
6. Утилита cut.
 - Реализовать утилиту аналог консольной команды cut. Кдлючи: -f, -d, -s.
9. Утилита wget.
 - Реализовать утилиту wget с возможностью скачивать сайты целиком.
~~~