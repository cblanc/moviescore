# Moviescore

Uses Metacritic API to lookup movie scores.

## Usage

Add binary to $PATH

```
$ moviescore the godfather
Found 6 matching movies

+----------------------------+-------+------------+--------------------------------+--------------------------------+
|           MOVIE            | SCORE | USER SCORE |             GENRE              |              CAST              |
+----------------------------+-------+------------+--------------------------------+--------------------------------+
| The Godfather              |   100 |        9.0 | Drama, Thriller, Crime         | Al Pacino, Marlon Brando       |
| The Godfather: Part II     | 80    |        9.2 | Drama, Thriller, Crime         | Al Pacino, Diane Keaton,       |
|                            |       |            |                                | Robert De Niro, Robert Duvall  |
| The Godfather: Part III    | 60    |        7.1 | Action, Drama, Thriller, Crime | Al Pacino, Andy Garcia, Diane  |
|                            |       |            |                                | Keaton                         |
| Tokyo Godfathers           | 73    |        8.3 | Adventure, Drama, Comedy,      | Aya Okamoto, Toru Emori,       |
|                            |       |            | Animation                      | Yoshiaki Umegaki               |
| The Godfather of Green Bay |       | tbd        | Comedy                         |                                |
| Square Grouper             | 55    |            | Documentary                    |                                |
+----------------------------+-------+------------+--------------------------------+--------------------------------+
```

Requires a Mashape API Key. Either:

```
exports MASHAPE_KEY="your_api_key"
```

or 

```
MASHAPE_KEY=your_api_key moviescore
```