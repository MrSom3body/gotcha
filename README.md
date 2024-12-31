# 🐹 gotcha, my small fetch tool

![preview](.github/assets/preview.png)

gotcha is a small fetcher written in go. It has absolutely no customization (at least for now) and is primarily intended for my system.

## ✨ features

- speedy boi (on my machine [<2ms](#benchmarks))
- no dependencies besides linux and the binary itself (and go for building duh)
- displays:
  - distribution
  - kernel version
  - uptime (in days too 😈)
  - shell
  - desktop environment/window manager
  - memory usage
  - (by default my[^1]) local ip

[^1]: why my you may ask? Because gotcha builds with my interface name by default (see [here](#overrides) on how to override) 🤡

## 🛠️ benchmarks

With my `AMD Ryzen 9 6900HX` CPU I get the following results from some fetches I know (and mine ofc):

| Command     |   Mean [ms] | Min [ms] | Max [ms] |       Relative |
| :---------- | ----------: | -------: | -------: | -------------: |
| `gotcha`    |   1.7 ± 0.2 |      1.3 |      2.3 |           1.00 |
| `macchina`  |   2.4 ± 0.3 |      1.7 |      3.4 |    1.35 ± 0.19 |
| `neofetch`  | 529.6 ± 7.4 |    521.0 |    544.1 | 302.87 ± 28.97 |
| `fastfetch` |  63.1 ± 3.7 |     59.4 |     78.8 |   36.11 ± 4.01 |
| `pfetch`    | 122.8 ± 3.7 |    118.8 |    131.6 |   70.25 ± 6.98 |

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

Get the [latest release](https://github.com/MrSom3body/gotcha/releases) or compile it yourself if you want an useful output for your local IP. You can do that really easy by installing go and running the following commands:

```bash
git clone github.com/MrSom3body/gotcha
cd gotcha
go build -ldflags="-s -w"
```

### 🧰 overrides

There is no real configuration, but you can override some values to change some things. The process of this is imo easier for nix but if you compile it yourself because you use a ~inferior~ different distro you need to do so with some flags.

| Key       | Description                                     |
| --------- | ----------------------------------------------- |
| ifaceName | The interface name from which to display the ip |

#### ❄️ nix overrides

If you want to change the interface name for the IP you can override the package like this:

```nix
(inputs.gotcha.packages.${pkgs.system}.default.override {
  <key> = "<value>";
})
```

#### 🔨 compiling

For every override you want to add you must add this `-X 'github.com/MrSom3body/gotcha/cmd.<key>=<value>'` to the `-ldflags` like so:

```bash
go build -ldflags="-s -w -X 'github.com/MrSom3body/gotcha/cmd.<key>=<value>'"
```

## 🤔 why this name?

go + fetch → gofetch → gotch → gotcha

I know gotcha makes absolutely no sense for a fetcher but idc :)
