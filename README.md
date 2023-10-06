#Go-urlmining
Mendapatkan data teks dari pencarian url


### Install
```
go install github.com/1amkaizen/Go-urlmining@latest
```

### Penggunaan

```
Usage of ./gomining:
-h    Show help
-p int
    Number of pages (default 1)
-q string
    Search Query
```


```
Go-urlmining -q jokowi -p 1 > output.txt
```
Atau
```
Go-urlmining -q "masalah masyarakat" -p 1 > output.txt
```

Selanjutnya gunakan tools [BashTextMiner](https://github.com/1amkaizen/BashTextMiner)  untuk mendapatkan hasil akhirnya.        
```
./BashTextminer -s combine-stopwords.txt hasil.txt > output.txt
```
