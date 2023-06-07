# kaken-search

## Build/Run
```sh
go build kaken-search.go
# or
go run kaken-search.go
```


## Commands
```sh
$ ./kaken-search --help
```
```
NAME:
   kaken-search - a CLI tool to search for researchers on KAKEN database

USAGE:
   kaken-search [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --id value                CiNii App ID
   --input value, -i value   Path for input CSV
   --name value, -n value    Column number containing researcher's name (start counting from 0) (default: "0")
   --affl value, -a value    Column number containing researcher's affiliation (start counting from 0) (default: "1")
   --output value, -o value  Path for output TSV (default: "output.tsv")
   --help, -h                show help
```

## Example
### Check the input data
```sh
$ head data.csv
```
```
通番号,審査区分,氏名,所  属  等
1,医歯薬学,田中 花子,学術大学 学術学部 教授
2,医歯薬学,佐藤 太郎,学術大学 学術学部 教授
```

### Run
```sh
$ ./kaken-search --id YOUR_CINII_APP_ID --input data.csv --output out.tsv --name 2 --affl 3
```

### Check the output
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