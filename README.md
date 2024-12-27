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
  - (my[^1]) local ip

[^1]: why my you may ask? Because I hard-coded my Interface 🤡

## installation

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
inputs.gotcha.packages.${pkgs.system}.gotcha
```

## why this name?

go + fetch -> gofetch -> gotch -> gotcha

I know gotcha makes absolutely no sense for a fetcher but idc :)
