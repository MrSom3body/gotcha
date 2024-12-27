{
  mkShell,
  go,
  gopls,
  gotools,
  go-tools,
}:
mkShell {
  buildInputs = [
    go
    gopls
    gotools
    go-tools
  ];
}
