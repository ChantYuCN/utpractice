# Unit test tutorial

## Abstract 

Unit test (UT) test the alogrithm and logic per module. In this tutorial, use gomock and mockgen to do UT.


### Download the source code

```bash
git clone https://github.com/ChantYuCN/utpractice.git utp && cd utp
```
### Download the relative dependence

```bash
make dependence
```

### Create the interface to the funtions we want to test

```bash
vi pkg/hello.go
```

### Download testing tools "gomock and mockge". Using tools to do "mocking"

```bash
make mockprepare
```

### code a test case

```bash
vi pkg/hello_test.go
```

### test

```bash
make test
```

### After practice, don't forget to clean build environemnt

```bash
make clean
```

