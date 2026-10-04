# git

A learning repository for git basics. It is where I practiced commits, feature branches, merging, resolving a merge conflict and pushing to two remotes, Gitea and GitHub. It also contains a Repo Health Analyzer written in Go that scores the commit history of this repository.

## Contents

| Path | What it is |
|---|---|
| `feature/feature.txt` | file created on the `feature` branch and merged into `main` |
| `docs/workflow.md` | the branching, merging and conflict workflow used here |
| `reflection.md` | what I learned about git |
| `main.go` | analyzer entry point: reads `git log` and prints the report |
| `internal/health/` | scoring rules: `ScoreMessage` and `Label`, with unit tests |

## Repo Health Analyzer

Run it from the root of the repository:

```bash
go run .
```

It reads `git log --pretty=format:"%ad|%s" --date=short` and prints the number of commits, the span in days, the commit cadence, the average message score out of 3 and a history label (`tidy`, `acceptable` or `messy`).

Example output for this repository:

```
=== Repo Health ===
Commits: 21
Span: 1 days
Cadence: 21.00 commits/day
Average message score: 3.0 / 3
History: tidy
```

Each commit message gets one point for each rule:

- it has 3 or more words;
- it starts with a type: `add`, `update`, `fix`, `delete`, `docs` or `refactor`;
- it references a file or a topic: a word with a dot (`main.go`), a word in capitals (`README`), or a project topic (`branch`, `merge`, `conflict`, `remote`, `analyzer`, `readme`, `workflow`, `feature`).

Run the tests:

```bash
go test ./...
```

---

# git (на русском)

Учебный репозиторий по основам git. Здесь я отрабатывал коммиты, работу в отдельных ветках, слияние, разрешение конфликта и отправку в два удалённых репозитория: Gitea и GitHub. Также здесь есть Repo Health Analyzer на Go, который оценивает историю коммитов этого репозитория.

## Содержимое

| Путь | Что это |
|---|---|
| `feature/feature.txt` | файл, созданный в ветке `feature` и влитый в `main` |
| `docs/workflow.md` | описание работы с ветками, слиянием и конфликтом |
| `reflection.md` | что я узнал о git |
| `main.go` | точка входа анализатора: читает `git log` и печатает отчёт |
| `internal/health/` | правила оценки: `ScoreMessage` и `Label` с unit-тестами |

## Repo Health Analyzer

Запуск из корня репозитория:

```bash
go run .
```

Программа читает `git log --pretty=format:"%ad|%s" --date=short` и выводит число коммитов, период в днях, частоту коммитов (cadence), средний балл сообщений из 3 и итоговую метку истории (`tidy`, `acceptable` или `messy`).

Пример вывода для этого репозитория:

```
=== Repo Health ===
Commits: 21
Span: 1 days
Cadence: 21.00 commits/day
Average message score: 3.0 / 3
History: tidy
```

Каждое сообщение коммита получает по одному баллу за каждое правило:

- в нём 3 слова или больше;
- оно начинается с типа: `add`, `update`, `fix`, `delete`, `docs` или `refactor`;
- в нём упомянут файл или тема: слово с точкой (`main.go`), слово заглавными буквами (`README`) или тема проекта (`branch`, `merge`, `conflict`, `remote`, `analyzer`, `readme`, `workflow`, `feature`).

Запуск тестов:

```bash
go test ./...
```
