
```r
require(data.table)              # print?
DT = data.table(x=1:3, y=4:6)    # no
DT                               # yes
```

```
##    x y
## 1: 1 4
## 2: 2 5
## 3: 3 6
```

```r
DT[, z := 7:9]                   # no
print(DT[, z := 10:12])          # yes
```

```
##    x y  z
## 1: 1 4 10
## 2: 2 5 11
## 3: 3 6 12
```

```r
if (1 < 2) DT[, a := 1L]         # no
DT                               # yes
```

```
##    x y  z a
## 1: 1 4 10 1
## 2: 2 5 11 1
## 3: 3 6 12 1
```
Some text.

