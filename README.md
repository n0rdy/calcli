# Calcli - a CLI calculator

## Description
A CLI calculator app.

Features:
- basic arithmetic operations: `+`, `-`, `*`, `/`, `%`, `^`, `!`, as well as parentheses
- integer and floating-point numbers (`.` as a decimal separator)
- a set of predefined constants: `pi`, `e`
- math functions:
  - `abs(x)` - the absolute value of `x`
  - `acos(x)` - the arccosine of `x`
  - `asin(x)` - the arcsine of `x`
  - `atan(x)` - the arctangent of `x`
  - `ceil(x)` - the smallest integer value greater than or equal to `x`
  - `cos(x)` - the cosine of `x`
  - `exp(x)` - the value of `e^x`
  - `exp2(x)` - the value of `2^x`
  - `floor(x)` - the largest integer value less than or equal to `x`
  - `ln(x)` - the natural logarithm of `x`
  - `log(x, base)` - the logarithm of `x` to the specified `base`
  - `log2(x)` - the base-2 logarithm of `x`
  - `log10(x)` - the base-10 logarithm of `x`
  - `mod(x, y)` - the remainder of the division of `x` by `y`
  - `nrt(x, degree)` - the root of `x` of the specified `degree`
  - `percent(x, y)` - the percentage of `x` from `y`
  - `round(x)` - the value of `x` rounded to the nearest integer
  - `sin(x)` - the sine of `x`
  - `sqrt(x)` - the square root of `x`
  - `tan(x)` - the tangent of `x`
- the result of the previous calculation is stored in the variable `$pr` and can be used in further calculations.

## Installation
### Prerequisites
- [Go](https://golang.org/doc/install) (version 1.22 or higher) if you want to build the app from the source code.

### Manual
Download the latest release for your OS from [GitHub](https://github.com/n0rdy/calcli/releases).

### MacOS
- via Homebrew:
```shell
brew tap n0rdy/n0rdy
brew install calcli
```

### Linux
- via APT:
#### Prerequisites
To enable, add the following file /etc/apt/sources.list.d/fury.list:
```text
deb [trusted=yes] https://apt.fury.io/n0rdy/ /
```
You can do this either manually or by running the following command:
```shell
echo "deb [trusted=yes] https://apt.fury.io/n0rdy/ /" > /etc/apt/sources.list.d/fury.list
```
If you experienced this error:
```text
bash: /etc/apt/sources.list.d/fury.list: Permission denied
```
try to do the following:
```shell
sudo -i
# enter your password
echo "deb [trusted=yes] https://apt.fury.io/n0rdy/ /" > /etc/apt/sources.list.d/fury.list
# click Ctrl+D to exit
```

#### Installation
```shell
sudo apt update && sudo apt install calcli
```

- via YUM:
#### Prerequisites
To enable, add the following file /etc/yum.repos.d/fury.repo:
```text
[fury]
name=Gemfury n0rdy Private Repo
baseurl=https://yum.fury.io/n0rdy/
enabled=1
gpgcheck=0
```

#### Installation
```shell
sudo yum install calcli
```

## Usage
```shell
calcli
```

This will start the app in the interactive mode.
Type your expression and press Enter to get the result.