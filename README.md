#### Упражнение 4.13

Веб-служба Open Movie Database https://omdbapi.com/ на базе JSON позволяет выполнять поиск фильма по названию и загружать 
его афишу. Напишите инструмент poster, который загружает афишу фильма по переданному в командной строке названию.

Пример использования:

Загруженные постеры находяться в папке /tmp/lego07102024_115310
```bash
$ go test -v --search lego --apikey <API_KEY> 
=== RUN   TestSearch
    pstr_test.go:38: ---------------------------------------------------------------------------------------------------------------------------------
    pstr_test.go:39: poster: https://m.media-amazon.com/images/M/MV5BMTg4MDk1ODExN15BMl5BanBnXkFtZTgwNzIyNjg3MDE@._V1_SX300.jpg
    pstr_test.go:40: title: The Lego Movie
    pstr_test.go:41: type: movie
    pstr_test.go:42: year: 2014
    pstr_test.go:43: imdbID: 
    pstr_test.go:38: ---------------------------------------------------------------------------------------------------------------------------------
    pstr_test.go:39: poster: https://m.media-amazon.com/images/M/MV5BMTcyNTEyOTY0M15BMl5BanBnXkFtZTgwOTAyNzU3MDI@._V1_SX300.jpg
    pstr_test.go:40: title: The Lego Batman Movie
    pstr_test.go:41: type: movie
    pstr_test.go:42: year: 2017
    pstr_test.go:43: imdbID: 
    pstr_test.go:38: ---------------------------------------------------------------------------------------------------------------------------------
    pstr_test.go:39: poster: https://m.media-amazon.com/images/M/MV5BMTkyOTkwNDc1N15BMl5BanBnXkFtZTgwNzkyMzk3NjM@._V1_SX300.jpg
    pstr_test.go:40: title: The Lego Movie 2: The Second Part
    pstr_test.go:41: type: movie
    pstr_test.go:42: year: 2019
    pstr_test.go:43: imdbID: 
    pstr_test.go:38: ---------------------------------------------------------------------------------------------------------------------------------
    pstr_test.go:39: poster: https://m.media-amazon.com/images/M/MV5BNWQ3Zjk0ZGEtOTU4ZS00OWY1LTlkM2YtNTllZjBhNTk3NDAxXkEyXkFqcGc@._V1_SX300.jpg
    pstr_test.go:40: title: The Lego Ninjago Movie
    pstr_test.go:41: type: movie
    pstr_test.go:42: year: 2017
    pstr_test.go:43: imdbID: 
    pstr_test.go:38: ---------------------------------------------------------------------------------------------------------------------------------
    pstr_test.go:39: poster: https://m.media-amazon.com/images/M/MV5BMDVlZGE3NDgtMTZjMi00YjlkLWFhYmYtODc2NmZlMzcxNWYzXkEyXkFqcGdeQXVyMTI0NTA1MDI3._V1_SX300.jpg
    pstr_test.go:40: title: The Lego Star Wars Holiday Special
    pstr_test.go:41: type: movie
    pstr_test.go:42: year: 2020
    pstr_test.go:43: imdbID: 
    pstr_test.go:38: ---------------------------------------------------------------------------------------------------------------------------------
    pstr_test.go:39: poster: https://m.media-amazon.com/images/M/MV5BNjRjOGU2NzUtMjAwOC00MDI3LThmNmUtNTVkZTQ0MDEyYTc1XkEyXkFqcGdeQXVyMTA5NzUzODM4._V1_SX300.jpg
    pstr_test.go:40: title: Lego Batman: The Movie - DC Super Heroes Unite
    pstr_test.go:41: type: movie
    pstr_test.go:42: year: 2013
    pstr_test.go:43: imdbID: 
    pstr_test.go:38: ---------------------------------------------------------------------------------------------------------------------------------
    pstr_test.go:39: poster: https://m.media-amazon.com/images/M/MV5BZDEwYTYwMjEtZjMzYy00YWE1LTg4ODEtNzEwMjcwZjZmNGQyXkEyXkFqcGc@._V1_SX300.jpg
    pstr_test.go:40: title: Lego Masters
    pstr_test.go:41: type: series
    pstr_test.go:42: year: 2020–
    pstr_test.go:43: imdbID: 
    pstr_test.go:38: ---------------------------------------------------------------------------------------------------------------------------------
    pstr_test.go:39: poster: https://m.media-amazon.com/images/M/MV5BYWJiMWM2YTktZjhjOS00YmQyLTgyYzYtMDcwMDQ1NWRjYTI2XkEyXkFqcGdeQXVyNzcxMjQzMDU@._V1_SX300.jpg
    pstr_test.go:40: title: Lego Star Wars: The Complete Saga
    pstr_test.go:41: type: game
    pstr_test.go:42: year: 2007
    pstr_test.go:43: imdbID: 
    pstr_test.go:38: ---------------------------------------------------------------------------------------------------------------------------------
    pstr_test.go:39: poster: https://m.media-amazon.com/images/M/MV5BOGRiYWEzZjctMTMzMi00M2VhLWEzZDItNDIyOTVkZjBkNTI0XkEyXkFqcGc@._V1_SX300.jpg
    pstr_test.go:40: title: Lego Star Wars Terrifying Tales
    pstr_test.go:41: type: movie
    pstr_test.go:42: year: 2021
    pstr_test.go:43: imdbID: 
    pstr_test.go:38: ---------------------------------------------------------------------------------------------------------------------------------
    pstr_test.go:39: poster: https://m.media-amazon.com/images/M/MV5BZjViNzNkNjAtMGZmZi00MDM2LTk1OGUtYjVlNzhjNjIxNDgwXkEyXkFqcGdeQXVyMTAyNzc0MDkz._V1_SX300.jpg
    pstr_test.go:40: title: Lego Batman: The Videogame
    pstr_test.go:41: type: game
    pstr_test.go:42: year: 2008
    pstr_test.go:43: imdbID: 
    pstr_test.go:45: all posters successfully uploaded
--- PASS: TestSearch (3.73s)
PASS
ok      github.com/unixlinuxgeek/pstr   3.733s
```

![logo](https://raw.githubusercontent.com/unixlinuxgeek/logos/refs/heads/main/gopl/ex_4_13.png)