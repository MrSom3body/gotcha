# 🐹 gotcha, my small fetch tool

![preview](.github/assets/preview.png)

gotcha is a small fetcher written in go. It has absolutely no customization (at least for now) and is primarily intended for my system.

## features

- speedy boi (on my machine <4ms)
- no dependencies besides linux and the binary itself (and go for building duh)
- displays:
  - distribution
  - kernel version
  - uptime (in days too 😈)
  - shell
  - desktop environment/window manager
  - memory usage
  - (by default my[^1]) local ip

[^1]: why my you may ask? Because gotcha builds with my interface name by default (see [here](#overrideInterface) on how to override) 🤡

## installation

### nix (with flakes)

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

If you want to change the interface name for the IP you can override the package like this:

<a id="overrideInterface"></a>

```nix
(inputs.gotcha.packages.${pkgs.system}.default.override {interface = "ens33";})
```

### all other distributions

Get the [latest release](https://github.com/MrSom3body/gotcha/releases) or compile it yourself if you want an useful output for your local IP. You can do that really easy by installing go and running the following commands:

```bash
git clone github.com/MrSom3body/gotcha
cd gotcha
go build -ldflags="-s -w -X 'github.com/MrSom3body/gotcha/cmd.ifaceName=INTERFACE'"
```

## why this name?

go + fetch → gofetch → gotch → gotcha

I know gotcha makes absolutely no sense for a fetcher but idc :)
