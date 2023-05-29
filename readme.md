[![Audit](https://github.com/gluck1986/test_work_xml/actions/workflows/audit.yml/badge.svg)](https://github.com/gluck1986/test_work_xml/actions/workflows/audit.yml)

# How to run

```shell
make up
```

# Task

## A) Реализовать приложение на golang
#### Методы

1. localhost:8080/update
   - ```shell
        curl -X POST -H "Content-Type: application/json" -d '' http://localhost:8080/update
     ```
   - Импорт / обновление необходимых данных из https://www.treasury.gov/ofac/downloads/sdn.xml в
   локальную базу PostgreSQL 14. в базу должны попадать записи с sdnType=Individual
   - Результат: success:
   ```json
      {"result": true, "info": "", "code": 200}
   ```
   - fail:
   ```json
      {"result": false, "info": "service unavailable", "code": 503}
   ```
2. localhost:8080/state
    - ```shell
        curl -X GET -H "Content-Type: application/json" http://localhost:8080/state
      ```
    - Получение текущего состояния данных нет данных:
      ```json
        {"result": false, "info": "empty"}
      ```
    - в процессе обновления:
     ```json
       {"result": false, "info": "updating"}
     ```
    - данные готовы к использованию:
        ```json
         {"result": true, "info": "ok"}
        ```

3. localhost:8080/get_names?name={SOME_VALUE}&type={strong|weak}

      Получение списка всех возможных имён человека из локальной базы данных с указанием
      основного uid в виде JSON.
      Если параметр type не указан / указан ошибочно, то выдаём список состоящий из всех типов.
      Параметр type независим от регистра. strong - это точное совпадение имени и фамилии, weak -
      должно найти любое совпадение в имени либо фамилии
   
4. Запрос: localhost:8080/get_names?name=MUZONZINI&type=strong
      ```shell
   curl -X GET -H "Content-Type: application/json" http://localhost:8080/get_names?name=MUZONZINI&type=strong
      ```
   Результат:
   ```json
      [{"uid":7535, "first_name":"Elisha", "last_name":"Muzonzini"}]
   ```
   
5. Запрос: localhost:8080/get_names?name=Elisha Muzonzini
   ```shell
   curl -X GET -H "Content-Type: application/json" http://localhost:8080/get_names?name=Elisha%20Muzonzini
   ```
   Результат:
   ```json
      [{"uid":7535, "first_name":"Elisha", "last_name":"Muzonzini"}]
   ```
   
6. Запрос: localhost:8080/get_names?name=Mohammed Musa&type=weak
      ```shell
   curl -X GET -H "Content-Type: application/json" http://localhost:8080/get_names?name=Mohammed%20Musa&type=weak
      ```
   Результат:
   ```json
   [{"uid":15582, "first_name":"Musa", "last_name":"Kalim"}, {"uid":15582, "first_name":"Barich", "last_name":"Musa Kalim"},
    {"uid":15582, "first_name":"Mohammed Musa", "last_name":"Kalim"}, {"uid":15582, "first_name":"Musa Khalim",
    "last_name":"Alizari"}, {"uid":15582, "first_name":"Qualem", "last_name":"Musa"}, {"uid":15582,
   "first_name":"Qualim", "last_name":"Musa"}, {"uid":15582, "first_name":"Khaleem", "last_name":"Musa"},
   {"uid":15582, "first_name":"Kaleem", "last_name":"Musa"}]
   ```

 
## B) Написать инструкции Docker Compose для разворачивания
   реализованного приложения на порту 8080 с использованием Postgresql14.
## C) Описать алгоритм для более эффективного обновления данных при
   повторном вызове метода localhost:8080/update ( можно реализовать, но не
   обязательно )
   
