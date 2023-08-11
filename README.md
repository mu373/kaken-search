# kaken-search

CLI tool to search for researchers on [KAKEN database](https://kaken.nii.ac.jp/ja/). CiNii App ID is required. Registration can be done from [this page](https://support.nii.ac.jp/ja/cinii/api/developer).

Following informations are available upon search. Search results can be exported as TSV file.
- Researcher's name
- Research keywords
- Affiliation (institution, department, title)
- KAKEN researcher page URL (e.g., https://nrid.nii.ac.jp/ja/nrid/0000000000000/)

## Build/Run
```sh
go build kaken-search.go
# or
go run kaken-search.go
```

## Commands

### Getting started
```sh
$ ./kaken-search --help
```
```
NAME:
   kaken-search - a CLI tool to search for researchers on KAKEN database

USAGE:
   kaken-search [global options] command [command options] [arguments...]

COMMANDS:
   single, s  Search a single researcher
   bulk, b    Search multiple researchers from CSV
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

### Single search
```sh
$ ./kaken-search single --help
```
```
NAME:
   kaken-search single - Search a single researcher

USAGE:
   kaken-search single [command options] [arguments...]

OPTIONS:
   --id value                CiNii App ID
   --name value, -n value    Researcher's name
   --affl value, -a value    Researcher's affiliation (institution)
   --format value, -f value  Output format [json]
   --help, -h                show help
```

### Bulk search
```sh
$ ./kaken-search bulk --help
```
```
NAME:
   kaken-search bulk - Search multiple researchers from CSV

USAGE:
   kaken-search bulk [command options] [arguments...]

OPTIONS:
   --id value                CiNii App ID
   --input value, -i value   Path for input CSV
   --name value, -n value    Column number containing researcher's name (start counting from 0) (default: 0)
   --affl value, -a value    Column number containing researcher's affiliation (start counting from 0) (default: 1)
   --output value, -o value  Path for output TSV (default: "output.tsv")
   --help, -h                show help
```

## Examples

### Case 1: Searching a single researcher

```sh
$ ./kaken-search single --id YOUR_CINII_APP_ID --name 田中太郎 --affl 学術大学 --format json
```


### Case 2: Searching multiple researchers from CSV

Check the input data
```sh
$ head data.csv
```
```
通番号,審査区分,氏名,所  属  等
1,医歯薬学,田中 花子,学術大学 学術学部 教授
2,医歯薬学,佐藤 太郎,学術大学 学術学部 教授
```

Run
```sh
$ ./kaken-search bulk --id YOUR_CINII_APP_ID --input data.csv --output out.tsv --name 2 --affl 3
```

Check the output
```sh
$ head out.tsv
```
```
OriginalName    OriginalRole    Name    Keywords        AffiliationInstitution  AffiliationDepartment   AffiliationTitle        KakenUrl
田中 花子       学術大学 学術学部 教授    田中 花子       老化,認知症,フレイル,Sirt1,エストロゲン     学術大学 学術学部  教授    https://nrid.nii.ac.jp/ja/nrid/1000000999999/
```

## Resources
- [Kaken API Documentation](https://bitbucket.org/niijp/kaken_definition/src/master/)
- [Registration for CiNii WebAPI](https://support.nii.ac.jp/ja/cinii/api/developer)