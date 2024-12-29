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

[^1]: why my you may ask? Because gotcha builds with my interface name by default (see [here](#overrides) on how to override) 🤡

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

### all other distributions

Get the [latest release](https://github.com/MrSom3body/gotcha/releases) or compile it yourself if you want an useful output for your local IP. You can do that really easy by installing go and running the following commands:

```bash
git clone github.com/MrSom3body/gotcha
cd gotcha
go build -ldflags="-s -w"
```

### overrides

There is no real configuration, but you can override some values to change some things. The process of this is imo easier for nix but if you compile it yourself because you use a ~inferior~ different distro you need to do so with some flags.

| Key       | Description                                     |
| --------- | ----------------------------------------------- |
| ifaceName | The interface name from which to display the ip |

#### nix overrides

If you want to change the interface name for the IP you can override the package like this:

```nix
(inputs.gotcha.packages.${pkgs.system}.default.override {
  <key> = "<value>";
})
```

#### compiling

For every override you want to add you must add this `-X 'github.com/MrSom3body/gotcha/cmd.<key>=<value>'` to the `-ldflags` like so:

```bash
go build -ldflags="-s -w -X 'github.com/MrSom3body/gotcha/cmd.<key>=<value>'"
```

## why this name?

go + fetch → gofetch → gotch → gotcha

I know gotcha makes absolutely no sense for a fetcher but idc :)
