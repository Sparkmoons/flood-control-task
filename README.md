Когда завершите задачу, в этом README опишите свой ход мыслей: как вы пришли к решению, какие были варианты и почему выбрали именно этот. 

# Что нужно сделать

Реализовать интерфейс с методом для проверки правил флуд-контроля. Если за последние N секунд вызовов метода Check будет больше K, значит, проверка на флуд-контроль не пройдена.

- Интерфейс FloodControl располагается в файле main.go.

- Флуд-контроль может быть запущен на нескольких экземплярах приложения одновременно, поэтому нужно предусмотреть общее хранилище данных. Допустимо использовать любое на ваше усмотрение. 

# Необязательно, но было бы круто

Хорошо, если добавите поддержку конфигурации итоговой реализации. Параметры — на ваше усмотрение.


# Решение

Выбрал Redis из-за быстрого чтения и записи в хранилище (ну и я других, честно говоря, не знаю. Загуглил и начал разбираться с ним).

Реализовал структуру floodControl, которая реализует интерфейс FloodControl с полями клиента, интервала времени и порога запросов. В функции NewFloodControl реализовано создание нового нового экзампляра структуры по заданным параметрам. 

Метод Check осуществляет проверку прохождения пользователя на флуд. Для пользователя создается ключ, получается количество вызовов за временной интервал и проверка, был превышен порог или нет. Если порог не был превышен - текущий вызов добавляется в хранилище  с временной меткой и устанавливается время жизни ключа. При успешном прохождении контроля вовзращаем true, в противном случае - false. 

В main добавил пример проверки для пользователя с идентификатором 159. Создается новый экземпляр floodControl с адресом и паролем для соединения с базой Redis. Интервал времени 5 секунд, порог запросов равен 3. Между каждой проверкой программа ожидает одну секунду. 
