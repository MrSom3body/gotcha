# 🐹 gotcha, my small fetch tool

![preview](.github/assets/preview.png)

gotcha is a small fetcher written in go. It has absolutely no customization (at
least for now) and is primarily intended for my system.

## ✨ features

- speedy boi (on my machine [<2ms](#benchmarks))
- no dependencies besides linux and the binary itself (and go for building duh)
- can update itself with `gotcha update`
- displays:
  - distribution
  - kernel version
  - uptime (in days too 😈)
  - shell
  - desktop environment/window manager
  - memory usage
  - ip (of an automagically[^1] selected interface)

[^1]:
    automagically in the sense of that it _should_ select the appropriate
    interface

## 🛠️ benchmarks

With my `AMD Ryzen 9 6900HX` CPU I get the following results from some fetches I
know (and mine ofc):

| Command     |    Mean [ms] | Min [ms] | Max [ms] |       Relative |
| :---------- | -----------: | -------: | -------: | -------------: |
| `gotcha`    |    1.6 ± 0.2 |      1.2 |      2.1 |           1.00 |
| `macchina`  |   11.6 ± 0.7 |     10.3 |     15.4 |    7.06 ± 0.83 |
| `neofetch`  | 527.8 ± 11.7 |    507.7 |    547.8 | 320.03 ± 33.31 |
| `fastfetch` |   62.5 ± 1.2 |     60.6 |     67.1 |   37.91 ± 3.93 |
| `pfetch`    |  130.8 ± 1.3 |    127.9 |    134.0 |   79.33 ± 8.11 |

<sub>created with [hyperfine](https://github.com/sharkdp/hyperfine)</sub>

## 📦 installation

### ❄️ nix (with flakes)

Add this thingy to your inputs:

```nix
inputs = {
  # ...
  gotcha = {
    url = "github:MrSom3body/gotcha";
    inputs = {
      nixpkgs.follows = "nixpkgs";
    };
  };
  # ...
};
```

Andddd add this the package to home-manager or your system wide nix config:

```nix
inputs.gotcha.packages.${pkgs.system}.default
```

### 🐧 all other distributions

Get the [latest release](https://github.com/MrSom3body/gotcha/releases) or just
compile it yourself. Compiling it is extremely easy if you have `just` installed
(or else copy the command with `go build` in it from the `justfile`):

```bash
just build
```

After installing it you can update it by simply using `gotcha update`.

### 🧰 overrides

There is no real configuration, but you can override some values to change some
things. The process of this is imo easier for nix but if you compile it yourself
because you use a ~inferior~ different distro you need to do so with some flags.

| Key   | Default Value | Description                                                           |
| ----- | ------------- | --------------------------------------------------------------------- |
| color | `6`           | The integer value of the color (0-7[^2]) you want to use for the keys |

[^2]:
    Use one of these commands to display the 8 colors:  
    bash: `for i in $(seq 0 7); do tput setaf $i && echo $i; done`  
    fish: `for i in (seq 0 7); tput setaf $i && echo $i; end`

#### ❄️ nix overrides

If you want to change the interface name for the IP you can override the package
like this:

```nix
(inputs.gotcha.packages.${pkgs.system}.default.override {
  <key> = "<value>";
})
```

#### 🔨 compiling

For every override you want to add you must add this `-X
'github.com/MrSom3body/gotcha/cmd.<key>=<value>'` to the `-ldflags` like so:

```bash
go build -ldflags="-s -w -X 'github.com/MrSom3body/gotcha/cmd.<key>=<value>'"
```

## 🤔 why this name?

go + fetch → gofetch → gotch → gotcha

I know gotcha makes absolutely no sense for a fetcher but idc :)

## ⭐ Stargraph

<picture>
  <source media="(prefers-color-scheme: dark)"
    srcset="https://api.star-history.com/svg?repos=MrSom3body/gotcha&type=Date&theme=dark"/>
  <source media="(prefers-color-scheme: light)"
    srcset="https://api.star-history.com/svg?repos=MrSom3body/gotcha&type=Date"/>
  <img alt="Star History Chart"
    src="https://api.star-history.com/svg?repos=MrSom3body/gotcha&type=Date"/>
</picture>
